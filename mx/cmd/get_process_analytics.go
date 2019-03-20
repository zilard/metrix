
package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)



var GetProcessAnalyticsCmd = &cobra.Command{

    Use: "pa",
    Run: func(cmd *cobra.Command, args []string) {

        GetProcessAnalytics(timeSlice)

    },

}



type Process struct {
    Name   string   `json:"name,omitempty"`
    Url    string   `json:"url,omitempty"`
}


type ProcessAnalytics struct {
    TimeSlice   float64    `json:"timeslice,omitempty"`
    Processes   []Process  `json:"processes,omitempty"`
}



func GetProcessAnalytics(timeSlice float64) {

    c := NewClient()

    req, _ := c.newRequest("GET", "/v1/analytics/processes/", nil)

    var processAnalytics ProcessAnalytics

    c.do(req, &processAnalytics)

    fmt.Printf("PROCESS ANALYTICS: %v\n", processAnalytics)

}




