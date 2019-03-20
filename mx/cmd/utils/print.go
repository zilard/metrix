
package utils

import (
    "fmt"
    "bytes"
    "encoding/json"
)

func PrettyPrint(data interface{}) string {

    var out bytes.Buffer

    b, _ := json.Marshal(data)
    json.Indent(&out, b, "", "  ")

    return fmt.Sprintf("\n\n%s\n", out.String())

}



