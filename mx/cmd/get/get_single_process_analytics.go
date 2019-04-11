
package get

import (
    "fmt"
    "net/url"

    s "github.com/zilard/metrix/metrix/structs"
    h "github.com/zilard/metrix/mx/cmd/http"
    u "github.com/zilard/metrix/mx/cmd/utils"

    "github.com/spf13/cobra"
)


var processName string



// Get Single Process Analytics sub-subcommand will get an extra flag for the Process Name
func init() {

    GetSingleProcessAnalyticsCmd.Flags().StringVarP(&processName, "processname", "p", "", "Process Name")

}



// GetSingleProcessAnalyticsCmd - Get Process Analytics sub-subcommand invokes the GetSingleProcessAnalytics API Client
var GetSingleProcessAnalyticsCmd = &cobra.Command{

    Use: "spa",
    Run: func(cmd *cobra.Command, args []string) {


        if processName == "" {

             fmt.Println("Please specify a process name using the -p=<processname> flag")
             return

        }

        GetSingleProcessAnalytics()

    },

}


// GetSingleProcessAnalytics - API Client for sending GET request to path /v1/analytics/processes/<processname>
// and a timeslice query parameter
// The response is loaded into a ProcessAverageReport struct
func GetSingleProcessAnalytics() {

    c := h.NewClient(ip, port)

    req, _ := c.NewRequest("GET", "/v1/analytics/processes/" + processName, nil)

    q := url.Values{}
    q.Add("timeslice", fmt.Sprintf("%f", timeSlice))
    req.URL.RawQuery = q.Encode()

    var processAnalytics s.ProcessAverageReport

    c.Do(req, &processAnalytics)


    //fmt.Printf("SINGLE PROCESS ANALYTICS: %s\n", u.PrettyPrint(processAnalytics))

}


