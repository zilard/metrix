
package handlers


import (
    "encoding/json"
    "net/http"
    "fmt"
    "os"

    s "github.com/zilard/metrix/metrix/structs"
    "github.com/gorilla/mux"
)


// nodeMetricsMap is where all data is stored
var nodeMetricsMap = make(s.NodeMetricsMap)

// processMetricsArray required to collect the process metrics history
var processMetricsArray = []s.ProcessMetricsByName{}




// CreateProcessMetrics - handler for path /v1/metrics/nodes/{nodename}/process/{processname}/
// Populates the NodeMetricsMap struct and it's internals with the received data, process metrics like
// timeselice, cpu usage and mem usage regarding the specific Node and Process
func CreateProcessMetrics(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    nodeName := params["nodename"]
    processName := params["processname"]

    var processMeasurement s.ProcessMeasurement
    result := json.NewDecoder(r.Body).Decode(&processMeasurement)

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




// CreateNodeMetrics - handler for path /v1/metrics/node/{nodename}/
// Populates the NodeMetricsMap struct and it's internals with the received data, node metrics like
// timeselice, cpu percentage and mem percentage regarding the specific Node
func CreateNodeMetrics(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    nodeName := params["nodename"]

    var nodeMeasurement s.NodeMeasurement
    result := json.NewDecoder(r.Body).Decode(&nodeMeasurement)

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





