
package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)


var nodeName string
var cpu float64
var mem float64



func init() {

    SendCmd.AddCommand(SendNodeMetricsCmd)
    SendCmd.AddCommand(SendProcessMetricsCmd)

}



var SendCmd = &cobra.Command{

    Use: "send",
    Run: func(cmd *cobra.Command, args []string) {

        fmt.Printf("Use:" +
                   "\n\t1.) \"mx send nm\" to send Node Metrics" +
                   "\n\t2.) \"mx send pm\" to get Process Metrics\n")

    },

}


