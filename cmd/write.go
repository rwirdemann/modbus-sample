package cmd

import (
	"github.com/goburrow/serial"
	"github.com/rwirdemann/modbus-sample/mb"
	"github.com/spf13/cobra"
	"log"
	"log/slog"
)

func init() {
	rootCmd.AddCommand(writeCmd)
}

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		p, err := serial.Open(&serial.Config{Address: port})
		if err != nil {
			log.Fatal(err)
		}
		defer p.Close()

		adu := mb.NewADU(1, 0x04, 57)
		bb, _ := adu.Build()
		_, err = p.Write(bb)
		if err != nil {
			log.Fatal(err)
		}
		slog.Info("write", "port", port, "value", 57)
	},
}
