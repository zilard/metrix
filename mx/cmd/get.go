
package cmd

import (
    "fmt"
    "reflect"

    "github.com/spf13/cobra"
)


func init() {

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

        fmt.Printf("GET: received=> timeSlice %v  type: %v\n", timeSlice, reflect.TypeOf(timeSlice))

    },

}

