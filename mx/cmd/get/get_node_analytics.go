
package get

import (
    "fmt"
    "net/url"

    s "github.com/zilard/metrix/structs"
    h "github.com/zilard/metrix/mx/cmd/http"
    u "github.com/zilard/metrix/mx/cmd/utils"

    "github.com/spf13/cobra"
)



var GetNodeAnalyticsCmd = &cobra.Command{

    Use: "na",
    Run: func(cmd *cobra.Command, args []string) {

        GetNodeAnalytics()

    },

}


func GetNodeAnalytics() {

    c := h.NewClient()

    req, _ := c.NewRequest("GET", "/v1/analytics/nodes/average", nil)

    q := url.Values{}
    q.Add("timeslice", fmt.Sprintf("%f", timeSlice))
    req.URL.RawQuery = q.Encode()

    var nodeAnalytics s.NodeAverageReport

    c.Do(req, &nodeAnalytics)

    fmt.Printf("NODE ANALYTICS: %s\n", u.PrettyPrint(nodeAnalytics))

}


