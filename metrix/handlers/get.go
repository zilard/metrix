
package handlers

import (
    "encoding/json"
    "net/http"
    "fmt"
    "strconv"

    u "github.com/zilard/metrix/metrix/handlers/utils"

    "github.com/gorilla/mux"
)



// GetAllNodeAverageMetrics - handler for path /v1/analytics/nodes/average
// parses the URL Query parameter timeslice and invokes CreateNodeAverageReport function
// to get the Node Average Analytics
// The analytics data is obtained through a NodeAverageReport struct
// Returns back the JSON encoded responde through http ResponseWriter
func GetAllNodeAverageMetrics(w http.ResponseWriter, r *http.Request) {

    paramArray, ok := r.URL.Query()["timeslice"]

    var timeSlice float64

    if !ok || len(paramArray[0]) < 1 {
        timeSlice = 60
    } else {
        timeSlice, _ = strconv.ParseFloat(paramArray[0], 64)
    }

    totalNodeAverageReport := u.CreateNodeAverageReport(nodeMetricsMap, timeSlice)

    fmt.Printf("TOTAL NODE AVERAGE ANALYTICS %v\n", totalNodeAverageReport)
    json.NewEncoder(w).Encode(totalNodeAverageReport)

}



// GetProcessAverageMetricsAllNodes - handler for path /v1/analytics/processes/{processname}
// parses the URL Query parameters processname and timeslice and invokes CreateProcessAverageReport function
// to get the Process Average Analytics for a specific process calculated based on the info collected on all 
// nodes where this process is/was running. The analytics data is obtained through a ProcessAverageReport struct
// Returns back the JSON encoded responde through http ResponseWriter
func GetProcessAverageMetricsAllNodes(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    processName := params["processname"]

    paramArray, ok := r.URL.Query()["timeslice"]

    fmt.Printf("GOT => PROCESSNAME: %v\n", processName)

    var timeSlice float64

    if !ok || len(paramArray[0]) < 1 {
        timeSlice = 60
    } else {
        timeSlice, _ = strconv.ParseFloat(paramArray[0], 64)
    }

    fmt.Printf("GOT => TIMESLICE: %v\n", timeSlice)

    totalProcessAverageReport := u.CreateProcessAverageReport(nodeMetricsMap, processName, timeSlice)

    fmt.Printf("PROCESS AVERAGE ANALYTICS ALL NODES %v\n", totalProcessAverageReport)
    json.NewEncoder(w).Encode(totalProcessAverageReport)

}




// GetMostRecentProcesses - handler for path /v1/analytics/processes
// parses the URL Query parameter timeslice and invokes CreateProcessHistoryReport function
// to get the most recent history of process metrics in a ProcessHistoryReport struct
// Returns back the JSON encoded response through http ResponseWriter
func GetMostRecentProcesses(w http.ResponseWriter, r *http.Request) {

    paramArray, ok := r.URL.Query()["timeslice"]

    var timeSlice float64

    if !ok || len(paramArray[0]) < 1 {
        timeSlice = 60
    } else {
        timeSlice, _ = strconv.ParseFloat(paramArray[0], 64)
    }

    fmt.Printf("GOT => TIMESLICE: %v\n", timeSlice)

    mostRecentProcessHistoryReport := u.CreateProcessHistoryReport(processMetricsArray, timeSlice)

    fmt.Printf("MOST RECENT PROCESS HISTORY REPORT %v\n", mostRecentProcessHistoryReport)
    json.NewEncoder(w).Encode(mostRecentProcessHistoryReport)

}




