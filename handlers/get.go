
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



func CreateDummyNodeMetrics() {


    for i := 1; i <= 2; i++ {

        nd := s.NodeData{}

        for j := 1; j <= 10; j++ {
            nd.NodeMeasurementArray = append(nd.NodeMeasurementArray,
                         s.NodeMeasurement{
                             TimeSlice: float64(j*10),
                             Cpu: float64(j*5),
                             Mem: float64(j*6),
                         })
        }

        nodeMetricsMap["n"+strconv.Itoa(i)] = nd

     }

}



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



