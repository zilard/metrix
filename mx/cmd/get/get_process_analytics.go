
package get

import (
    "fmt"
    "net/url"

    s "github.com/zilard/metrix/structs"
    h "github.com/zilard/metrix/mx/cmd/http"
    u "github.com/zilard/metrix/mx/cmd/utils"

    "github.com/spf13/cobra"
)



var GetProcessAnalyticsCmd = &cobra.Command{

    Use: "pa",
    Run: func(cmd *cobra.Command, args []string) {

        GetProcessAnalytics()

    },

}



func GetProcessAnalytics() {

    c := h.NewClient()

    req, _ := c.NewRequest("GET", "/v1/analytics/processes", nil)

    q := url.Values{}
    q.Add("timeslice", fmt.Sprintf("%f", timeSlice))
    req.URL.RawQuery = q.Encode()

    var processAnalytics s.ProcessHistoryReport

    c.Do(req, &processAnalytics)

    fmt.Printf("PROCESS ANALYTICS: %s\n", u.PrettyPrint(processAnalytics))

}



