package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

var (
    address     = flag.String("a", "127.0.0.1", "The address of the Minecraft server")
    port        = flag.Int("p", 25565, "The port the Minecraft server is running on")
    serverPort  = flag.Int("serverPort", 8080, "The port the mcstat server will run at.")
)

type Status struct {
    Online bool             // online or offline?
    Version string          // server version
    Motd string             // message of the day
    CurrentPlayers string   // current number of players online
    MaxPlayers string       // maximum player capacity
    Latency time.Duration   // ping time to server in milliseconds
}

func main() {
    // Parse the command line flags
    flag.Parse()

    http.HandleFunc("/status", func(rw http.ResponseWriter, r *http.Request) {
        // Retrieve the status of the Minecraft server
        status, err := GetStatus(*address, *port)
        if err != nil {
            log.Fatal(err.Error())
        }
        log.Printf("Retrieved status from server: %+v\n", status)

        rw.Header().Set("Content-Type", "application/json")
        json.NewEncoder(rw).Encode(status)
    })

    sp := fmt.Sprint(*serverPort)
    log.Printf("Running mcstat on port %s\n", sp)
    log.Fatal(http.ListenAndServe(":" + sp, nil))
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
    s := time.Now()
    portString := fmt.Sprint(port)

    // Opens a TCP connection to the given address and port.
    // Also accepts a timeout so that we fail fast if the opening the connection takes to long 
    log.Printf("Attempting to open TCP connection to %s:%s\n", address, portString)
    conn, err := net.DialTimeout("tcp", address + ":" + portString, time.Duration(5) * time.Second)

    if err != nil {
        log.Printf("ERROR: %s\n", err.Error())
        return Status{}, err
    }

    // Calculate the latency with millisecond accuracy
    latency := time.Since(s)
    latency = latency.Round(time.Millisecond)

    // Write the following bytes to the TCP connection:
    // FE 01 (hex)
    // This is the legacy protocol way of doing a Server List ping
    // More info https://wiki.vg/Server_List_Ping#1.6
    // Modern servers should reply to this call correctly because of
    // the starting FE bytes.
    log.Printf("Sending Server List Ping to %s:%d\n", address, port)
    _, err = conn.Write([]byte("\xFE\x01"))


    // Read the raw response from our Server List Ping
    r := make([]byte, 512)
    _, err = conn.Read(r)
    log.Printf("Retrieved response from %s:%d\n%X\n", address, port, r)

    if err != nil {
        log.Printf("ERROR: %s\n", err.Error())
        return Status{}, err
    }
    conn.Close()

    // Split the response data by the byte pattern 00 00 00
    v := bytes.Split(r, []byte("\x00\x00\x00"))

    return Status{
        Online: true,
        Version: string(v[2][:]),
        Motd: string(v[3][:]),
        CurrentPlayers: string(v[4][:]),
        MaxPlayers: string(v[5][:]),
    }, nil

}
