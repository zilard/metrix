
package main

import (
    "log"
    "net/http"
    "fmt"
    "strconv"
    "os"

    h "github.com/zilard/metrix/metrix/handlers"
    "github.com/gorilla/mux"
    "github.com/spf13/cobra"

)

const PORT = 8080

var Port int


var RootCmd = &cobra.Command{
    Use: "metrix",
    Run: func(cmd *cobra.Command, args []string) {

        if Port < 1024 || Port > 65535 {
            fmt.Printf("Port number out of range!\nPlease use a port number between 1024 and 65535\n")
            return
        }

        Run()

    },
}


// Adding flag for optional port number that can be specified
func init() {

    RootCmd.PersistentFlags().IntVarP(&Port, "port", "p", PORT, "port")

}



func main() {

    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

}


// Run calling HTTP Handlers functions for each specific API path
func Run() {

    router := mux.NewRouter()

    router.HandleFunc("/v1/metrics/node/{nodename}/", h.CreateNodeMetrics).Methods("POST")
    router.HandleFunc("/v1/metrics/nodes/{nodename}/process/{processname}/", h.CreateProcessMetrics).Methods("POST")

    router.HandleFunc("/v1/analytics/nodes/average", h.GetAllNodeAverageMetrics).Methods("GET")

    router.HandleFunc("/v1/analytics/processes/{processname}", h.GetProcessAverageMetricsAllNodes).Methods("GET")
    router.HandleFunc("/v1/analytics/processes/{processname}/", h.GetProcessAverageMetricsAllNodes).Methods("GET")
    router.HandleFunc("/v1/analytics/processes", h.GetMostRecentProcesses).Methods("GET")
    router.HandleFunc("/v1/analytics/processes/", h.GetMostRecentProcesses).Methods("GET")

    fmt.Printf("SERVER LISTENING ON :%d\n", Port)
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(Port), router))
}



