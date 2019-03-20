
package handlers

import (
    "encoding/json"
    "net/http"
    "fmt"
    "strconv"

    s "github.com/zilard/metrix/structs"
    u "github.com/zilard/metrix/handlers/utils"

    "github.com/gorilla/mux"
)



func GetNodeMetrics(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    nodeName := params["nodename"]

    fmt.Printf("Metrics for node %v: %v\n", nodeName, nodeMetricsMap[nodeName])

    json.NewEncoder(w).Encode(nodeMetricsMap[nodeName])

}


func GetAllNodeMetrics(w http.ResponseWriter, r *http.Request) {

    fmt.Printf("Metrics for all nodes: %v\n", nodeMetricsMap)

    json.NewEncoder(w).Encode(nodeMetricsMap)

}



func GetAllNodeAverageMetrics(w http.ResponseWriter, r *http.Request) {

    paramArray, ok := r.URL.Query()["timeslice"]

    var timeSlice float64

    if !ok || len(paramArray[0]) < 1 {
        timeSlice = 60
    } else {
        timeSlice, _ = strconv.ParseFloat(paramArray[0], 64)
    }

    //ONLY FOR TESTING
    //u.CreateDummyNodeMetrics(nodeMetricsMap)

    totalNodeAverageReport := u.CreateNodeAverageReport(nodeMetricsMap, timeSlice)

    fmt.Printf("TOTAL NODE AVERAGE ANALYTICS %v\n", totalNodeAverageReport)
    json.NewEncoder(w).Encode(totalNodeAverageReport)

}



func CreateDummyProcessMetrics() {

    for i := 1; i <= 2; i++ {
        nodeData := s.NodeData{}

        processMetricsMap := make(s.ProcessMetricsMap)

        for j := 1; j <= 1; j++ {
            processMeasurementArray := []s.ProcessMeasurement{}
            for k := 1; k <= 2; k++ {
                processMeasurementArray = append(processMeasurementArray,
                             s.ProcessMeasurement{
                                 TimeSlice: float64(k * 10),
                                 CpuUsed: float64(k * 5),
                                 MemUsed: float64(k * 6),
                              })
            }
            processMetricsMap["proc" + strconv.Itoa(j)] = processMeasurementArray
        }
        nodeData.ProcessMeasurementMap = processMetricsMap
        nodeMetricsMap["node" + strconv.Itoa(i)] = nodeData
     }
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


    //ONLY FOR TESTING
    CreateDummyProcessMetrics()

    //fmt.Printf("NODE METRICS MAP %v\n\n", nodeMetricsMap)


    totalProcessAverageReport := u.CreateProcessAverageReport(nodeMetricsMap, processName, timeSlice)

    fmt.Printf("PROCESS AVERAGE ANALYTICS ALL NODES %v\n", totalProcessAverageReport)
    json.NewEncoder(w).Encode(totalProcessAverageReport)


}




