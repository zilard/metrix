
package send

import (
    "fmt"

    s "github.com/zilard/metrix/metrix/structs"
    h "github.com/zilard/metrix/mx/cmd/http"
    u "github.com/zilard/metrix/mx/cmd/utils"

    "github.com/spf13/cobra"
)



func init() {

    SendNodeMetricsCmd.Flags().StringVarP(&nodeName, "nodename", "n", "", "Node Name")

    SendNodeMetricsCmd.Flags().Float64VarP(&cpu, "cpu", "c", 0, "cpu percentage")
    SendNodeMetricsCmd.Flags().Float64VarP(&mem, "mem", "m", 0, "memory percentage")

}



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



func SendNodeMetrics() {

    c := h.NewClient()

    nodeMetrics := s.NodeMeasurement{
                       TimeSlice: timeSlice,
                       Cpu: cpu,
                       Mem: mem,
                   }

    path := fmt.Sprintf("/v1/metrics/node/%s/", nodeName)
    req, _ := c.NewRequest("POST", path, nodeMetrics)


    var nodeMeasurement s.NodeMeasurement

    c.Do(req, &nodeMeasurement)


    fmt.Printf("NODE MEASUREMENT SENT: %s\n", u.PrettyPrint(nodeMeasurement))

}


