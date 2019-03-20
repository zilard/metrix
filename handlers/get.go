
package handlers

import (
    "encoding/json"
    "net/http"
    "fmt"
    "strconv"

    u "github.com/zilard/metrix/handlers/utils"

    "github.com/gorilla/mux"
)



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




