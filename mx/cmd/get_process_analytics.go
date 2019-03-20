
package cmd

import (
    "fmt"
    "reflect"
    "github.com/spf13/cobra"
)



var GetProcessAnalyticsCmd = &cobra.Command{

    Use: "pa",
    Run: func(cmd *cobra.Command, args []string) {

        fmt.Printf("GET PROCESS ANALYTICS: received=> timeSlice %v  type: %v\n", timeSlice, reflect.TypeOf(timeSlice))

    },

}

