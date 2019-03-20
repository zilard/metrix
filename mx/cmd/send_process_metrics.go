
package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)



func init() {

    SendProcessMetricsCmd.Flags().StringVarP(&nodeName, "nodename", "n", "", "Node Name")
    SendProcessMetricsCmd.Flags().StringVarP(&processName, "processname", "p", "", "Process Name")

    SendProcessMetricsCmd.Flags().Float64VarP(&cpu, "cpu", "c", 0, "cpu percentage")
    SendProcessMetricsCmd.Flags().Float64VarP(&mem, "mem", "m", 0, "memory percentage")

}



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


    },

}


