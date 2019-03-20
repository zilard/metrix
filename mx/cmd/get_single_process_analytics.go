
package cmd

import (
    "fmt"
    "reflect"
    "github.com/spf13/cobra"
)



func init() {

    GetSingleProcessAnalyticsCmd.Flags().StringVarP(&processName, "processname", "p", "", "Process Name")

}


var GetSingleProcessAnalyticsCmd = &cobra.Command{

    Use: "spa",
    Run: func(cmd *cobra.Command, args []string) {


        if processName == "" {

             fmt.Println("Please specify a process name using the -p=<processname> flag")
             return

        }

        fmt.Printf("GET PROCESS ANALYTICS: received=> timeSlice %v  type: %v\n", timeSlice, reflect.TypeOf(timeSlice))

    },

}

