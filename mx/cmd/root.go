
package cmd

import (
    "fmt"

     g "github.com/zilard/metrix/mx/cmd/get"
     s "github.com/zilard/metrix/mx/cmd/send"

    "github.com/spf13/cobra"
)



// Main command parser
var RootCmd = &cobra.Command{
    Use: "mx",
    Run: func(cmd *cobra.Command, args []string) {

        fmt.Printf("Use:" +
                   "\n\t1.) \"mx send <params>\" to upload metrics" +
                   "\n\t2.) \"mx get <params>\" to get analytics\n")

    },
}


// Adding subcommands: get, send
func init() {

    RootCmd.AddCommand(g.GetCmd)
    RootCmd.AddCommand(s.SendCmd)

}


