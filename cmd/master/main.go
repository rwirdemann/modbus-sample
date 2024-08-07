package main

import (
	"fmt"
	"github.com/simonvetter/modbus"
	"log"
	"log/slog"
	"time"
)

func main() {
	readFromSlave()
}

func readFromSlave() {
	client, _ := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     fmt.Sprintf("tcp://localhost:%d", 502),
		Timeout: 1 * time.Second,
	})

	if err := client.Open(); err != nil {
		log.Fatal(err)
	}
	defer func(client *modbus.ModbusClient) {
		_ = client.Close()
	}(client)

	var reg16 uint16
	reg16, err := client.ReadRegister(1000, modbus.INPUT_REGISTER)
	if err != nil {
		log.Println(err)
	} else {
		slog.Info("read", "register", reg16)
	}
}
