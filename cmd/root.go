package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var port string

func init() {
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "", "/dev/ttys003")
	rootCmd.MarkPersistentFlagRequired("port")
}

var rootCmd = &cobra.Command{
	Use: "mbs",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
