
package cmd

import (
    "fmt"

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

    },

}


