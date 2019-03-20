
package get

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




// The get subcommand also has it's own sub-subcommands: 
// - get node analytics
// - get process analytics
// - get single process analytics
// timeslice flag is valid throughout all these subcommands
func init() {

    GetCmd.PersistentFlags().IPVarP(&ip, "ip", "i", net.ParseIP(IP), "ip")
    GetCmd.PersistentFlags().IntVarP(&port, "port", "o", PORT, "port")
    GetCmd.PersistentFlags().Float64VarP(&timeSlice, "ts", "t", 60, "timeslice")

    GetCmd.AddCommand(GetNodeAnalyticsCmd)
    GetCmd.AddCommand(GetProcessAnalyticsCmd)
    GetCmd.AddCommand(GetSingleProcessAnalyticsCmd)

}


// implementation of get subcommand
var GetCmd = &cobra.Command{

    Use: "get",

    PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

        if port < 1024 || port > 65535 {
            return errors.New(" Port number out of range!\n" +
                              "\tPlease use a port number between 1024 and 65535\n")
        }

        return nil
    },

    Run: func(cmd *cobra.Command, args []string) {

        fmt.Printf("Use:" +
                   "\n\t1.) \"mx get na\"  to get Node Analytics" +
                   "\n\t2.) \"mx get pa\" to get Process Analytics" +
                   "\n\t3.) \"mx get spa <process-name>\" to get Single Process Analytics for a given process\n")

    },

}

