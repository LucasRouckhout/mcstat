package main

import "testing"

func TestNewStatus(t *testing.T) {
    // Input which should translate to:
    // {Online: true, Version: 1.6.4, Motd: "The great world of Rouckhout", CurrentPlayers: 1, MaxPlayers: 0}
    input := []byte("\xFF\x00\x2F\x00\xA7\x00\x31\x00\x00\x00\x31\x00\x32\x00\x37\x00\x00\x00\x31\x00\x2E\x00\x31\x00\x36\x00\x2E\x00\x34\x00\x00\x00\x54\x00\x68\x00\x65\x00\x20\x00\x67\x00\x72\x00\x65\x00\x61\x00\x74\x00\x20\x00\x77\x00\x6F\x00\x72\x00\x6C\x00\x64\x00\x20\x00\x6F\x00\x66\x00\x20\x00\x52\x00\x6F\x00\x75\x00\x63\x00\x6B\x00\x68\x00\x6F\x00\x75\x00\x74\x00\x00\x00\x31\x00\x00\x00\x32\x00\x30\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")

    actualStatus, err := NewStatus(input)

    if err != nil {
        t.Error(err.Error())
    }

    expectedStatus := Status {
        Online: true,
        Version: "1.6.4",
        Motd: "The great world of Rouckhout",
        CurrentPlayers: "1",
        MaxPlayers: "20",
    }

    if actualStatus.Online != expectedStatus.Online {
        t.Errorf("Expected %t to be true but was %t\n", expectedStatus.Online, actualStatus.Online)
    }

    if actualStatus.Version != expectedStatus.Version {
        t.Errorf("Expected the Version to be: %x but got: %x\n", expectedStatus.Version, actualStatus.Version)
    }

    if actualStatus.Motd != expectedStatus.Motd {
        t.Errorf("Expected the Motd to be: %x, but got: %x\n", expectedStatus.Motd, actualStatus.Motd)
    }

    if actualStatus.CurrentPlayers != expectedStatus.CurrentPlayers {
        t.Errorf("Expected the CurrentPlayers to be: %x, but got: %x\n", expectedStatus.CurrentPlayers, actualStatus.CurrentPlayers)
    }

    if actualStatus.MaxPlayers != expectedStatus.CurrentPlayers {
        t.Errorf("Expected MaxPlayers to be: %x, but got: %x\n", expectedStatus.MaxPlayers, actualStatus.MaxPlayers)
    }

    if actualStatus != expectedStatus {
        t.Errorf("Expected: %+v but got: %+v\n", expectedStatus, actualStatus)
    }
}
