// mcstat -- A simple HTTP minecraft status server
// Copyright (C) 2020  Lucas Rouckhout
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
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

	"github.com/LucasRouckhout/mcstat/logger"
)

var (
    address     = flag.String("a", "127.0.0.1", "The address of the Minecraft server")
    port        = flag.Int("p", 25565, "The port the Minecraft server is running on")
    serverPort  = flag.Int("s", 8080, "The port the mcstat server will run at.")
    logLevel    = flag.String("l", "INFO", "The log level that mcstat should use. Can be (DEBUG, INFO or ERROR)")
)

var LOGGER logger.Logger

type Status struct {
    Online bool                 // online or offline?
    Version string              // server version
    Motd string                 // message of the day
    CurrentPlayers string       // current number of players online
    MaxPlayers string           // maximum player capacity
}

func main() {
    flag.Parse()

    switch(*logLevel) {
    case "INFO":
        LOGGER = logger.NewLogger(logger.INFO)

    case "DEBUG":
        LOGGER = logger.NewLogger(logger.DEBUG)

    case "ERROR":
        LOGGER = logger.NewLogger(logger.ERROR)

    default:
        LOGGER = logger.NewLogger(logger.INFO)
        LOGGER.Infof("Given logLevel %s was not recognized as one of (INFO, ERROR or DEBUG) so using INFO as default\n", *logLevel)
    }

    http.HandleFunc("/status", func(rw http.ResponseWriter, r *http.Request) {
        LOGGER.Infof("Getting status from %s:%d\n", *address, *port)
        status, err := GetStatus(*address, *port)

        if err != nil {
            LOGGER.Error(err.Error())
            rw.WriteHeader(500)

        } else {
            LOGGER.Infof("Retrieved status from server: %+v\n", status)
            rw.Header().Set("Content-Type", "application/json")
            json.NewEncoder(rw).Encode(status)
        }
    })

    serverPort := fmt.Sprint(*serverPort)
    LOGGER.Infof("Running mcstat on port %s\n", serverPort)
    log.Fatal(http.ListenAndServe(":" + serverPort, nil))
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
    portString := fmt.Sprint(port)

    // Opens a TCP connection to the given address and port.
    // Also accepts a timeout so that we fail fast if the opening the connection takes to long 
    LOGGER.Debugf("Attempting to open TCP connection to %s:%s\n", address, portString)
    conn, err := net.DialTimeout("tcp", address + ":" + portString, time.Duration(5) * time.Second)
    defer conn.Close()

    if err != nil {
        return Status{}, err
    }

    // Write the following bytes to the TCP connection:
    // FE 01 (hex)
    // This is the legacy protocol way of doing a Server List ping
    // More info https://wiki.vg/Server_List_Ping#1.6
    LOGGER.Debugf("Sending Server List Ping to %s:%d\n", address, port)
    _, err = conn.Write([]byte("\xFE\x01"))

    if err != nil {
        return Status{}, err
    }

    // Read the raw response from our Server List Ping
    r := make([]byte, 512)
    _, err = conn.Read(r)

    LOGGER.Debugf("Retrieved response from %s:%d\n%X\n", address, port, r)

    if err != nil {
        return Status{}, err
    }

    // Create a Status struct from this response
    status, err := NewStatus(r)

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
func NewStatus(b []byte) (Status, error) {
    r := bytes.Split(b, []byte("\x00\x00\x00"))

    return Status {
        Online: true,
        Version: string(cutByteSlice(r[2])),
        Motd: string(cutByteSlice(r[3])),
        CurrentPlayers: string(cutByteSlice(r[4])),
        MaxPlayers: string(cutByteSlice(r[5])),
    }, nil
}

// Helper method which cuts out all occurences of \x00\x00
// out of a byte slice
func cutByteSlice(b []byte) []byte {
    return bytes.ReplaceAll(b, []byte("\x00"), []byte(""))
}

