package cmd

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

var PingCmd = &cobra.Command{
	Use:   "ping [host]",
	Short: "Ping a network host",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		fmt.Printf("Pinging %s...\n", host)

		for i := 0; i < 4; i++ {
			start := time.Now()
			conn, err := net.DialTimeout("tcp", host+":80", 2*time.Second)
			duration := time.Since(start)

			if err != nil {
				fmt.Printf("Ping %d: failed (%v)\n", i+1, err)
			} else {
				fmt.Printf("Ping %d: %v\n", i+1, duration)
				conn.Close()
			}
			time.Sleep(1 * time.Second)
		}
	},
}
