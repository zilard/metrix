
package get

import (
    "fmt"

    "github.com/spf13/cobra"
)

var timeSlice float64

func init() {

    GetCmd.PersistentFlags().Float64VarP(&timeSlice, "ts", "t", 60, "timeslice")

    GetCmd.AddCommand(GetNodeAnalyticsCmd)
    GetCmd.AddCommand(GetProcessAnalyticsCmd)
    GetCmd.AddCommand(GetSingleProcessAnalyticsCmd)

}


var GetCmd = &cobra.Command{

    Use: "get",
    Run: func(cmd *cobra.Command, args []string) {

        fmt.Printf("Use:" +
                   "\n\t1.) \"mx get na\"  to get Node Analytics" +
                   "\n\t2.) \"mx get pa\" to get Process Analytics" +
                   "\n\t3.) \"mx get spa <process-name>\" to get Single Process Analytics for a given process\n")

    },

}

