
package handlers

import (
    "encoding/json"
    "net/http"
    "fmt"
    "strconv"

    s "github.com/zilard/metrix/structs"
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




type NodeAverageAnalytics map[string]NodeAverageReport

type NodeAverageReport struct {
        TimeSlice float64     `json:"timeslice,omitempty"`
        CpuUsed   float64     `json:"cpu_used,omitempty"`
        MemUsed   float64     `json:"mem_used,omitempty"`
}



func GetAllNodeAverageMetrics(w http.ResponseWriter, r *http.Request) {

    paramArray, ok := r.URL.Query()["timeslice"]

    var timeSlice float64

    if !ok || len(paramArray[0]) < 1 {
        timeSlice = 60
    } else {
        timeSlice, _ = strconv.ParseFloat(paramArray[0], 64)
    }

    fmt.Printf("TIMESLICE %v\n", timeSlice)


    CreateDummyNodeMetrics()


    fmt.Printf("NODE METRICS %v\n", nodeMetricsMap)


    nodeAverageAnalytics := make(NodeAverageAnalytics)

    for nodeName, nodeData := range nodeMetricsMap {

        fmt.Printf("NODE NAME %v\n", nodeName)


        nodeAverageReport := NodeAverageReport{
                                 TimeSlice: 0,
                                 CpuUsed: 0,
                                 MemUsed: 0,
                             }

        timeS := timeSlice

        for i := range nodeData.NodeMeasurementArray {
            nodeMeasurement := nodeData.NodeMeasurementArray[len(nodeData.NodeMeasurementArray)-1-i]
            fmt.Printf("__NODE MEASUREMENT %v\n", nodeMeasurement)

            if nodeMeasurement.TimeSlice >= timeS {
                nodeAverageReport.TimeSlice += timeS
                nodeAverageReport.CpuUsed += nodeMeasurement.Cpu * timeS
                nodeAverageReport.MemUsed += nodeMeasurement.Mem * timeS
                break
            } else {
                timeS -= nodeMeasurement.TimeSlice
                nodeAverageReport.TimeSlice += nodeMeasurement.TimeSlice
                nodeAverageReport.CpuUsed += nodeMeasurement.Cpu * nodeMeasurement.TimeSlice
                nodeAverageReport.MemUsed += nodeMeasurement.Mem * nodeMeasurement.TimeSlice
            }
        }

        nodeAverageReport.CpuUsed = nodeAverageReport.CpuUsed/nodeAverageReport.TimeSlice
        nodeAverageReport.MemUsed = nodeAverageReport.MemUsed/nodeAverageReport.TimeSlice


        nodeAverageAnalytics[nodeName] = nodeAverageReport

    }

    fmt.Printf("PER NODE AVERAGE ANALYTICS %v\n", nodeAverageAnalytics)



    totalNodeAverageReport := NodeAverageReport{
                                  TimeSlice: 0,
                                  CpuUsed: 0,
                                  MemUsed: 0,
                              }


    for _, nodeAverageReport := range nodeAverageAnalytics {

        if totalNodeAverageReport.TimeSlice == 0 {
            totalNodeAverageReport.TimeSlice = nodeAverageReport.TimeSlice
        } else {
            if totalNodeAverageReport.TimeSlice > nodeAverageReport.TimeSlice {
                totalNodeAverageReport.TimeSlice = nodeAverageReport.TimeSlice
            }
        }

        totalNodeAverageReport.CpuUsed += nodeAverageReport.CpuUsed
        totalNodeAverageReport.MemUsed += nodeAverageReport.MemUsed
    }

    totalNodeAverageReport.CpuUsed = totalNodeAverageReport.CpuUsed/float64(len(nodeAverageAnalytics))
    totalNodeAverageReport.MemUsed = totalNodeAverageReport.MemUsed/float64(len(nodeAverageAnalytics))



    fmt.Printf("TOTAL NODE AVERAGE ANALYTICS %v\n", totalNodeAverageReport)


    json.NewEncoder(w).Encode(totalNodeAverageReport)


}








