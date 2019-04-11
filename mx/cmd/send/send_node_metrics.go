
package send

import (
    "fmt"

    s "github.com/zilard/metrix/metrix/structs"
    h "github.com/zilard/metrix/mx/cmd/http"
    u "github.com/zilard/metrix/mx/cmd/utils"

    "github.com/spf13/cobra"
)





// Send Node Metrics sub-subcommand has some extra flags for nodename, cpu and mem
func init() {

    SendNodeMetricsCmd.Flags().StringVarP(&nodeName, "nodename", "n", "", "Node Name")

    SendNodeMetricsCmd.Flags().Float64VarP(&cpu, "cpu", "c", 0, "cpu percentage")
    SendNodeMetricsCmd.Flags().Float64VarP(&mem, "mem", "m", 0, "memory percentage")

}




// SendNodeMetricsCmd - Send Node Metrics sub-subcommand invokes the SendNodeMetrics API Client
// it also verifies whether the nodeName, cpu and mem has been provided through the corresponding flags
var SendNodeMetricsCmd = &cobra.Command{

    Use: "nm",
    Run: func(cmd *cobra.Command, args []string) {


        if nodeName == "" {

             fmt.Println("Please specify the Node Name using the -n=<nodename> flag")
             return

        }

        if cpu == 0 {

             fmt.Println("Please specify the CPU usage percentage using the -c=<used-cpu-percentage> flag")
             return

        }


        if mem == 0 {

             fmt.Println("Please specify the Memory usage percentage using the -m=<used-memory-percentage> flag")
             return

        }

        SendNodeMetrics()


    },

}




// SendNodeMetrics - API Client for sending POST request to the API server on path /v1/metrics/node/<nodename>/ 
// and a Json structure containing Node Metrics data: timeslice, cpu, mem
func SendNodeMetrics() {

    c := h.NewClient(ip, port)

    nodeMetrics := s.NodeMeasurement{
                       TimeSlice: timeSlice,
                       Cpu: cpu,
                       Mem: mem,
                   }

    path := fmt.Sprintf("/v1/metrics/node/%s/", nodeName)
    req, _ := c.NewRequest("POST", path, nodeMetrics)


    var nodeMeasurement s.NodeMeasurement

    c.Do(req, &nodeMeasurement)


    //fmt.Printf("NODE MEASUREMENT SENT: %s\n", u.PrettyPrint(nodeMeasurement))

}


