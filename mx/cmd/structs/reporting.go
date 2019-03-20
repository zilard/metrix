
package structs


type NodeAnalytics struct {
    TimeSlice float64     `json:"timeslice,omitempty"`
    Cpu       float64     `json:"cpu_used,omitempty"`
    Mem       float64     `json:"mem_used,omitempty"`
}


type Process struct {
    Name   string   `json:"name,omitempty"`
    Url    string   `json:"url,omitempty"`
}


type ProcessAnalytics struct {
    TimeSlice   float64    `json:"timeslice,omitempty"`
    Processes   []Process  `json:"processes,omitempty"`
}









