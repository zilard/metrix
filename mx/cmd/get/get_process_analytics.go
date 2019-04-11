
package get

import (
    "fmt"
    "net/url"

    s "github.com/zilard/metrix/metrix/structs"
    h "github.com/zilard/metrix/mx/cmd/http"

    "github.com/spf13/cobra"
)



// GetProcessAnalyticsCmd - Get Process Analytics sub-subcommand invokes the GetProcessAnalytics API Client
var GetProcessAnalyticsCmd = &cobra.Command{

    Use: "pa",
    Run: func(cmd *cobra.Command, args []string) {

        GetProcessAnalytics()

    },

}



// GetProcessAnalytics - API Client for sending GET request to path /v1/analytics/processes
// and a timeslice query parameter
// The response is loaded into a ProcessHistoryReport struct
func GetProcessAnalytics() {

    c := h.NewClient(ip, port)

    req, _ := c.NewRequest("GET", "/v1/analytics/processes", nil)

    q := url.Values{}
    q.Add("timeslice", fmt.Sprintf("%f", timeSlice))
    req.URL.RawQuery = q.Encode()

    var processAnalytics s.ProcessHistoryReport

    c.Do(req, &processAnalytics)

    //fmt.Printf("PROCESS ANALYTICS: %s\n", u.PrettyPrint(processAnalytics))

}



