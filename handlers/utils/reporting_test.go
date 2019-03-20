
package utils

import (
    "strconv"
    "bufio"
    "bytes"
    "flag"
    "io/ioutil"
    "path/filepath"
    "encoding/json"
    "testing"

    s "github.com/zilard/metrix/structs"
)


var testNodeMetricsMap = make(s.NodeMetricsMap)

var update = flag.Bool("update", false, "update .golden files")

const TIMESLICE = 360

func CreateDummyNodeMetricsForTest(nodeMetricsMap s.NodeMetricsMap) {
    for i := 1; i <= 2; i++ {
        nodeData := s.NodeData{}
        for j := 1; j <= 10; j++ {
            nodeData.NodeMeasurementArray = append(nodeData.NodeMeasurementArray,
                         s.NodeMeasurement{
                             TimeSlice: float64(j * 10),
                             Cpu: float64(j * 5),
                             Mem: float64(j * 6),
                         })
        }
        nodeMetricsMap["n" + strconv.Itoa(i)] = nodeData
     }
}


func TestCreateNodeAverageReport(t *testing.T) {

    timeSlice := float64(TIMESLICE)

    testtable := []struct {
        tname string
    }{
        {
            tname: "ok",
        },
    }

    CreateDummyNodeMetricsForTest(testNodeMetricsMap)

    for _, tc := range testtable {

        t.Run(tc.tname, func(t *testing.T) {

            var buffer bytes.Buffer
            writer := bufio.NewWriter(&buffer)

            err := json.NewEncoder(writer).Encode(CreateNodeAverageReport(testNodeMetricsMap, timeSlice))
            if err != nil {
                t.Fatalf("failed writing json: %s", err)
            }
            writer.Flush()

            goldenPath := filepath.Join("testdata", filepath.FromSlash(t.Name()) + ".golden")


            if *update {

	        t.Log("update golden file")
	        if err := ioutil.WriteFile(goldenPath, buffer.Bytes(), 0644); err != nil {
                    t.Fatalf("failed to update golden file %s: %s", goldenPath, err)
                }

             }


             goldenData, err := ioutil.ReadFile(goldenPath)

             if err != nil {
                 t.Fatalf("failed reading .golden file %s: %s", goldenPath, err)
             }

             t.Log(string(buffer.Bytes()))

             if !bytes.Equal(buffer.Bytes(), goldenData) {
                 t.Errorf("bytes do not match .golden file %s", goldenPath)
             }

         })
    }

}



