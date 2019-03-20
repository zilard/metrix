
package cmd

import (
    "fmt"
    "reflect"

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



        fmt.Printf("SEND NODE METRICS: received=> nodeName %v  type: %v\n", nodeName, reflect.TypeOf(nodeName))

        fmt.Printf("SEND NODE METRICS: received=> timeSlice %v  type: %v\n", timeSlice, reflect.TypeOf(timeSlice))

        fmt.Printf("SEND NODE METRICS: received=> cpu %v  type: %v\n", cpu, reflect.TypeOf(cpu))

        fmt.Printf("SEND NODE METRICS: received=> mem %v  type: %v\n", mem, reflect.TypeOf(mem))

    },

}


