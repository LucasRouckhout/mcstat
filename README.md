# mcstat

__NOTE__: This project is not finished yet. It was my first golang project used to learn the language.

A simple HTTP minecraft status server. mcstat is a simple HTTP server with a single endpoint `/status` which will provide you a quick status on your minecraft server.

# Usage

To run mcstat simple run the binary (see [installation](Installation) for installation instructions) with the address and port of your minecraft server.

```bash
mcstat -a myminecraftserver.example.com -p 25565
```

Once running you can query the endpoint with a HTTP GET call to `/status` to recieve a status of the server. The format of the response will be a JSON object with following fields

```json
{
    "Online": true, 
    "Version": "1.16.4",
    "Motd": "Message of the Day",
    "CurrentPlayers": 0,
    "MaxPlayers": 20
}
```


# Installation

Coming soon...

# Compiling

Just clone the repo, go to the root of the project and run.

```bash
go build
```

You should now find a binary in the root of the directory called `mcstat`.

# License

GNU GPLv3

```
Copyright (C) 2020 Lucas Rouckhout

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```
