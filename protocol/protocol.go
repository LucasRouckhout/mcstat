package protocol

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

// Represents the status of a Minecraft server
type Status struct {
	Online         bool   // online or offline?
	Version        string // server version
	Motd           string // message of the day
	CurrentPlayers string // current number of players online
	MaxPlayers     string // maximum player capacity
}

// Retrieves the status of the minecraft server at given Address and Port.
// The initial inspiration for this code was gathered from
// https://github.com/ldilley/minestat/blob/master/Go/minestat/minestat.go
// I rewrote and documented most parts to make it more idiomatic.
//
// This method uses the old 1.6 minecraft protocol to get a simple status response
// from the server by sending over a set of specific bytes over a TCP socket.
// Modern servers still respond to this protocol correctly so this will also
// work with servers who run anything newer than 1.6.
//
// Anything older then 1.6 is not supported.
func GetStatus(address string, port int) (Status, error) {
	conn, err := net.DialTimeout("tcp", address+":"+fmt.Sprint(port), time.Duration(5)*time.Second)

	if err != nil {
		return Status{}, err
	}

	// defer after the return of the err otherwise
	//you will get nasty nullpointers
	defer conn.Close()

	// Write the following bytes to the TCP connection:
	// FE 01 (hex)
	// This is the legacy protocol way of doing a Server List ping
	// More info https://wiki.vg/Server_List_Ping#1.6
	_, err = conn.Write([]byte("\xFE\x01"))

	if err != nil {
		return Status{}, err
	}

	buf := make([]byte, 512)
	_, err = conn.Read(buf)

	if err != nil {
		return Status{}, err
	}

	status, err := newStatus(buf)

	if err != nil {
		return Status{}, err
	}

	return status, nil
}

// Creates a Status struct from the structured
// response retrieved from a Server List Ping of
// a minecraft server.
//
// The structure of such a response in a byte sequece
// (Big Endian) which is structured like so:
// https://wiki.vg/Server_List_Ping#1.6
func newStatus(b []byte) (Status, error) {
	r := bytes.Split(b, []byte("\x00\x00\x00"))

	return Status{
		Online:         true,
		Version:        string(cutByteSlice(r[2])),
		Motd:           string(cutByteSlice(r[3])),
		CurrentPlayers: string(cutByteSlice(r[4])),
		MaxPlayers:     string(cutByteSlice(r[5])),
	}, nil
}

// Helper method which cuts out all occurences of \x00\x00
// out of a byte slice
func cutByteSlice(b []byte) []byte {
	return bytes.ReplaceAll(b, []byte("\x00"), []byte(""))
}
