
package main

import (
    "log"
    "net/http"
    "fmt"
    "strconv"

    h "github.com/zilard/metrix/handlers"
    "github.com/gorilla/mux"
)

const PORT = 8080


func main() {
    router := mux.NewRouter()

    router.HandleFunc("/v1/metrics/node/{nodename}", h.CreateNodeMetrics).Methods("POST")
    router.HandleFunc("/v1/metrics/node/{nodename}/process/{processname}", h.CreateProcessMetrics).Methods("POST")

    router.HandleFunc("/v1/analytics/nodes/average", h.GetAllNodeAverageMetrics).Methods("GET")
    router.HandleFunc("/v1/analytics/processes/{processname}", h.GetProcessAverageMetricsAllNodes).Methods("GET")
    router.HandleFunc("/v1/analytics/processes/", h.GetMostRecentProcesses).Methods("GET")

    fmt.Printf("SERVER LISTENING ON :%d\n", PORT)
    log.Fatal(http.ListenAndServe(":" + strconv.Itoa(PORT), router))
}



