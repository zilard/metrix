
package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)



var GetNodeAnalyticsCmd = &cobra.Command{

    Use: "na",
    Run: func(cmd *cobra.Command, args []string) {

        GetNodeAnalytics(timeSlice)

    },

}



type NodeAnalytics struct {
    TimeSlice float64     `json:"timeslice,omitempty"`
    Cpu       float64     `json:"cpu_used,omitempty"`
    Mem       float64     `json:"mem_used,omitempty"`
}



func GetNodeAnalytics(timeSlice float64) {

    c := NewClient()

    req, _ := c.newRequest("GET", "/v1/analytics/nodes/average", nil)

    var nodeAnalytics NodeAnalytics

    c.do(req, &nodeAnalytics)

    fmt.Printf("NODE ANALYTICS: %v\n", nodeAnalytics)

}


