
package structs


type NodeAverageAnalytics map[string]NodeAverageReport

type NodeAverageReport struct {
        TimeSlice float64     `json:"timeslice,omitempty"`
        CpuUsed   float64     `json:"cpu_used,omitempty"`
        MemUsed   float64     `json:"mem_used,omitempty"`
}


type ProcessAverageReport struct {
        TimeSlice    float64     `json:"timeslice,omitempty"`
        CpuUsed      float64     `json:"cpu_used,omitempty"`
        MemUsed      float64     `json:"mem_used,omitempty"`
        NumInstances float64     `json:"num_instances,omitempty"`
}




