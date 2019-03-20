
package handlers


import (
    "encoding/json"
    "net/http"
    "fmt"
    "os"

    s "github.com/zilard/metrix/metrix/structs"
    "github.com/gorilla/mux"
)



var nodeMetricsMap = make(s.NodeMetricsMap)
var processMetricsArray = []s.ProcessMetricsByName{}



func CreateProcessMetrics(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    nodeName := params["nodename"]
    processName := params["processname"]

    var processMeasurement s.ProcessMeasurement
    result := json.NewDecoder(r.Body).Decode(&processMeasurement)

    fmt.Printf("result: %v\n", result)
    if result != nil {
        fmt.Fprintf(os.Stderr, "result=%v\n", result)
        return
    }

    fmt.Printf("GOT => processMeasurement %v for NODE %v and PROCESS %v\n", processMeasurement, nodeName, processName)


    processMetricsData := s.ProcessMetricsByName{}
    processMetricsData.ProcessName = processName
    processMetricsData.MetricsData = processMeasurement

    processMetricsArray = append(processMetricsArray, processMetricsData)


    if _, ok := nodeMetricsMap[nodeName]; ok {

        nodeData := nodeMetricsMap[nodeName]

        processMetricsMap := nodeData.ProcessMeasurementMap

        if _, ok := processMetricsMap[processName]; ok {

            processMeasurementArray := processMetricsMap[processName]
            processMeasurementArray = append(processMeasurementArray, processMeasurement)

            processMetricsMap[processName] = processMeasurementArray

        } else {

           processMeasurementArray := []s.ProcessMeasurement{}
           processMeasurementArray = append(processMeasurementArray, processMeasurement)

           processMetricsMap[processName] = processMeasurementArray

        }

        nodeData.ProcessMeasurementMap = processMetricsMap

        nodeMetricsMap[nodeName] = nodeData

    } else {

        nodeData := s.NodeData{}
        processMetricsMap := make(s.ProcessMetricsMap)

        processMeasurementArray := []s.ProcessMeasurement{}
        processMeasurementArray = append(processMeasurementArray, processMeasurement)

        processMetricsMap[processName] = processMeasurementArray

        nodeData.ProcessMeasurementMap = processMetricsMap

        nodeMetricsMap[nodeName] = nodeData

    }


    fmt.Printf("SET nodeMetricsMap %v\n", nodeMetricsMap)

    fmt.Printf("\n")
    json.NewEncoder(w).Encode(processMeasurement)



}





func CreateNodeMetrics(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    nodeName := params["nodename"]

    var nodeMeasurement s.NodeMeasurement
    result := json.NewDecoder(r.Body).Decode(&nodeMeasurement)

    fmt.Printf("result: %v\n", result)
    if result != nil {
        fmt.Fprintf(os.Stderr, "result=%v\n", result)
        return
    }


    fmt.Printf("GOT => nodeMeasurement %v for NODE %v\n", nodeMeasurement, nodeName)


    if _, ok := nodeMetricsMap[nodeName]; ok {

        nodeData := nodeMetricsMap[nodeName]
        nodeData.NodeMeasurementArray = append(nodeData.NodeMeasurementArray, nodeMeasurement)
        nodeMetricsMap[nodeName] = nodeData

    } else {

        nodeData := s.NodeData{}
        nodeData.NodeMeasurementArray = append(nodeData.NodeMeasurementArray, nodeMeasurement)
        processMetricsMap := make(s.ProcessMetricsMap)
        nodeData.ProcessMeasurementMap = processMetricsMap
        nodeMetricsMap[nodeName] = nodeData

    }


    fmt.Printf("SET nodeMetricsMap %v\n", nodeMetricsMap)

    fmt.Printf("\n")
    json.NewEncoder(w).Encode(nodeMeasurement)

}





