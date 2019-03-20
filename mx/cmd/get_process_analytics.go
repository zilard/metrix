
package cmd

import (
    "fmt"

    s "github.com/zilard/metrix/mx/cmd/structs"
    h "github.com/zilard/metrix/mx/cmd/http"


    "github.com/spf13/cobra"
)



var GetProcessAnalyticsCmd = &cobra.Command{

    Use: "pa",
    Run: func(cmd *cobra.Command, args []string) {

        GetProcessAnalytics(timeSlice)

    },

}



func GetProcessAnalytics(timeSlice float64) {

    c := h.NewClient()

    req, _ := c.NewRequest("GET", "/v1/analytics/processes/", nil)

    var processAnalytics s.ProcessAnalytics

    c.Do(req, &processAnalytics)

    fmt.Printf("PROCESS ANALYTICS: %v\n", processAnalytics)

}



