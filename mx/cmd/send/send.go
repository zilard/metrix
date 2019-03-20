
package send

import (
    "fmt"
    "net"
    "errors"

    "github.com/spf13/cobra"
)


const IP = "127.0.0.1"
const PORT = 8080

var ip net.IP
var port int
var timeSlice float64
var nodeName string
var cpu float64
var mem float64



// The send subcommand also has it's own sub-subcommands: 
// - send node metrics
// - send process metrics
// timeslice flag is valid throughout all these subcommands
func init() {

    SendCmd.PersistentFlags().IPVarP(&ip, "ip", "i", net.ParseIP(IP), "ip")
    SendCmd.PersistentFlags().IntVarP(&port, "port", "o", PORT, "port")
    SendCmd.PersistentFlags().Float64VarP(&timeSlice, "ts", "t", 60, "timeslice")

    SendCmd.AddCommand(SendNodeMetricsCmd)
    SendCmd.AddCommand(SendProcessMetricsCmd)

}


// implementation of send subcommand
var SendCmd = &cobra.Command{

    Use: "send",

    PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

        if port < 1024 || port > 65535 {
            return errors.New(" Port number out of range!\n" +
                              "\tPlease use a port number between 1024 and 65535\n")
        }

        return nil
    },

    Run: func(cmd *cobra.Command, args []string) {

        fmt.Printf("Use:" +
                   "\n\t1.) \"mx send nm\" to send Node Metrics" +
                   "\n\t2.) \"mx send pm\" to get Process Metrics\n")

    },

}


