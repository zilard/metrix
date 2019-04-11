
package send

import (
    "fmt"

    s "github.com/zilard/metrix/metrix/structs"
    h "github.com/zilard/metrix/mx/cmd/http"
    u "github.com/zilard/metrix/mx/cmd/utils"

    "github.com/spf13/cobra"
)


var processName string




// Send Process Metrics sub-subcommand has some extra flags for nodename, processname, cpu and mem
func init() {

    SendProcessMetricsCmd.Flags().StringVarP(&nodeName, "nodename", "n", "", "Node Name")
    SendProcessMetricsCmd.Flags().StringVarP(&processName, "processname", "p", "", "Process Name")

    SendProcessMetricsCmd.Flags().Float64VarP(&cpu, "cpu", "c", 0, "cpu percentage")
    SendProcessMetricsCmd.Flags().Float64VarP(&mem, "mem", "m", 0, "memory percentage")

}




// SendProcessMetricsCmd - Send Process Metrics sub-subcommand invokes the SendProcessMetrics API Client
// it also verifies whether the nodename, processname, cpu and mem has been provided through the corresponding flags
var SendProcessMetricsCmd = &cobra.Command{

    Use: "pm",
    Run: func(cmd *cobra.Command, args []string) {

        if nodeName == "" {

             fmt.Println("Please specify the Node Name using the -n=<nodename> flag")
             return

        }


        if processName == "" {

             fmt.Println("Please specify the Process Name using the -p=<processname> flag")
             return

        }


        if cpu == 0 {

             fmt.Println("Please specify the CPU usage percentage using the -c=<used-cpu-percentage> flag")
             return

        }


        if mem == 0 {

             fmt.Println("Please specify the Megabytes of allocated memory using the -m=<used-memory-megabytes> flag")
             return

        }

        SendProcessMetrics()

    },

}



// SendProcessMetrics - API Client for sending POST request to the API server on path /v1/metrics/nodes/<nodename>/process/<processname>/
// and a Json structure containing Process Metrics data: timeslice, cpu, mem
func SendProcessMetrics() {

    c := h.NewClient(ip, port)

    processMetrics := s.ProcessMeasurement{
                          TimeSlice: timeSlice,
                          CpuUsed: cpu,
                          MemUsed: mem,
                      }


    path := fmt.Sprintf("/v1/metrics/nodes/%s/process/%s/", nodeName, processName)
    req, _ := c.NewRequest("POST", path, processMetrics)


    var processMeasurement s.ProcessMeasurement

    c.Do(req, &processMeasurement)


    //fmt.Printf("PROCESS MEASUREMENT SENT: %s\n", u.PrettyPrint(processMeasurement))

}


