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
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/LucasRouckhout/mcstat/logger"
	"github.com/LucasRouckhout/mcstat/protocol"
)

var (
	address    = flag.String("a", "127.0.0.1", "The address of the Minecraft server")
	port       = flag.Int("p", 25565, "The port the Minecraft server is running on")
	serverPort = flag.Int("s", 8080, "The port the mcstat server will run at.")
	logLevel   = flag.String("l", "INFO", "The log level that mcstat should use. Can be (DEBUG, INFO or ERROR)")
)

func main() {
	flag.Parse()

	var l logger.Logger

	switch *logLevel {
	case "INFO":
		l = logger.NewLogger(logger.INFO)

	case "DEBUG":
		l = logger.NewLogger(logger.DEBUG)

	case "ERROR":
		l = logger.NewLogger(logger.ERROR)

	default:
		l = logger.NewLogger(logger.INFO)
		l.Infof("Given logLevel %s was not recognized as one of (INFO, ERROR or DEBUG) so using INFO as default\n", *logLevel)
	}

	http.HandleFunc("/status", func(rw http.ResponseWriter, r *http.Request) {
		l.Infof("Getting status from %s:%d\n", *address, *port)
		status, err := protocol.GetStatus(*address, *port)

		if err != nil {
			l.Error(err.Error())
			rw.WriteHeader(500)

		} else {
			l.Infof("Retrieved status from server: %+v\n", status)
			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(status)
		}
	})

	serverPort := fmt.Sprint(*serverPort)
	l.Infof("Running mcstat on port %s\n", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}
