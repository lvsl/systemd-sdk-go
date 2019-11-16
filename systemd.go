package sdk

import (
	"log"

	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatalf("Failed to connect to session bus: %v", err)
	}

	log.Printf("Connection: %v\n", conn)
}
