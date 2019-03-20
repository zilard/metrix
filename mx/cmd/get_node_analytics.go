
package cmd

import (
    "fmt"
    "reflect"
    "github.com/spf13/cobra"
)



var GetNodeAnalyticsCmd = &cobra.Command{

    Use: "na",
    Run: func(cmd *cobra.Command, args []string) {

        fmt.Printf("GET NODE ANALYTICS: received=> timeSlice %v  type: %v\n", timeSlice, reflect.TypeOf(timeSlice))

    },

}

