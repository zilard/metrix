
package cmd

import (
    "fmt"

    s "github.com/zilard/metrix/mx/cmd/structs"
    h "github.com/zilard/metrix/mx/cmd/http"

    "github.com/spf13/cobra"
)



var GetNodeAnalyticsCmd = &cobra.Command{

    Use: "na",
    Run: func(cmd *cobra.Command, args []string) {

        GetNodeAnalytics(timeSlice)

    },

}


func GetNodeAnalytics(timeSlice float64) {

    c := h.NewClient()

    req, _ := c.NewRequest("GET", "/v1/analytics/nodes/average", nil)

    var nodeAnalytics s.NodeAnalytics

    c.Do(req, &nodeAnalytics)

    fmt.Printf("NODE ANALYTICS: %v\n", nodeAnalytics)

}


