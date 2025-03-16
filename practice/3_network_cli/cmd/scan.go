package cmd

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var ScanCmd = &cobra.Command{
	Use:   "scan [host] [startPort] [endPort]",
	Short: "Scan ports on a host",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		host := args[0]
		startPort, _ := strconv.Atoi(args[1])
		endPort, _ := strconv.Atoi(args[2])

		fmt.Printf("Scanning %s from port %d to %d...\n", host, startPort, endPort)

		for port := startPort; port <= endPort; port++ {
			address := fmt.Sprintf("%s:%d", host, port)
			conn, err := net.DialTimeout("tcp", address, 1*time.Second)

			if err != nil {
				fmt.Printf("Port %d open\n", port)
				conn.Close()
			}
		}
	},
}
