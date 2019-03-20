
package utils

import (
    "fmt"
    "strconv"

    s "github.com/zilard/metrix/structs"
)




func CreateDummyNodeMetrics(nodeMetricsMap s.NodeMetricsMap) {


    for i := 1; i <= 2; i++ {

        nodeData := s.NodeData{}

        for j := 1; j <= 10; j++ {
            nodeData.NodeMeasurementArray = append(nodeData.NodeMeasurementArray,
                         s.NodeMeasurement{
                             TimeSlice: float64(j*10),
                             Cpu: float64(j*5),
                             Mem: float64(j*6),
                         })
        }

        nodeMetricsMap["n"+strconv.Itoa(i)] = nodeData

     }

}





func CreateNodeAverageReport(nodeMetricsMap s.NodeMetricsMap, timeSlice float64) s.NodeAverageReport {


    fmt.Printf("TIMESLICE %v\n", timeSlice)


    CreateDummyNodeMetrics(nodeMetricsMap)


    fmt.Printf("NODE METRICS %v\n", nodeMetricsMap)


    nodeAverageAnalytics := make(s.NodeAverageAnalytics)

    for nodeName, nodeData := range nodeMetricsMap {

        fmt.Printf("NODE NAME %v\n", nodeName)


        nodeAverageReport := s.NodeAverageReport{
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



    totalNodeAverageReport := s.NodeAverageReport{
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


    return totalNodeAverageReport

}



