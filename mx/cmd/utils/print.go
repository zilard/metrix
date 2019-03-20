
package utils

import (
    "fmt"
    "bytes"
    "encoding/json"
)


// PrettyPrint - pretty print a given struct in Json style
func PrettyPrint(data interface{}) string {

    var out bytes.Buffer

    b, _ := json.Marshal(data)
    json.Indent(&out, b, "", "  ")

    return fmt.Sprintf("\n\n%s\n", out.String())

}



