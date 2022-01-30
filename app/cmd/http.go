package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	httpCmd = &cobra.Command{
		Use:   "http",
		Short: "Start Http REST API",
		Run:   initHTTP,
	}
)

func initHTTP(cmd *cobra.Command, args []string) {
	fmt.Println("http run...")
}

