package ds

import (
	"testing"
)

func TestConnection(t *testing.T) {
	ds := &DSConnection{Username: "cn=Directory Manager", Password: "password", DN: "o=monitor"}

	if err := ds.Connect(); err != nil {
		t.Fatalf("Can't connect %s", err)
	}
}

func TestSearch(t *testing.T) {
	ds := &DSConnection{Username: "cn=Directory Manager", Password: "password", DN: "o=monitor"}
	if err := ds.Search(); err != nil {
		t.Fatalf("Can't connect %s", err)
	}
}
