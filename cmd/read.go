package cmd

import (
	"fmt"
	"github.com/simonvetter/modbus"
	"github.com/spf13/cobra"
	"log"
	"log/slog"
	"time"
)

func init() {
	rootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		client, _ := modbus.NewClient(&modbus.ClientConfiguration{
			URL:      fmt.Sprintf("rtu://%s", port),
			Speed:    19200,
			DataBits: 8,
			Parity:   modbus.PARITY_NONE,
			StopBits: 2,
			Timeout:  30 * time.Millisecond,
		})

		if err := client.Open(); err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		ticker := time.NewTicker(1 * time.Second)
		for {
			<-ticker.C
			var reg16 uint16
			reg16, err := client.ReadRegister(100, modbus.INPUT_REGISTER)
			if err != nil {
				log.Println(err)
			} else {
				slog.Info("read", "register", reg16)
			}
		}
	},
}
