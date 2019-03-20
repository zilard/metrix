
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var timeSlice float64
var processName string


var RootCmd = &cobra.Command{
    Use: "mx",
    Run: func(cmd *cobra.Command, args []string) {

        fmt.Printf("Use:" +
                   "\n\t1.) \"mx send <params>\" to upload metrics" +
                   "\n\t2.) \"mx get <params>\" to get analytics\n")

    },
}


func init() {

    RootCmd.PersistentFlags().Float64VarP(&timeSlice, "ts", "t", 60, "timeslice")

    RootCmd.AddCommand(GetCmd)
    RootCmd.AddCommand(SendCmd)


}


