package messaging

import (
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "encoding/json"
    "fmt"
)

func ParseMessage(baseMessage request.Base, intendedType interface{}) error {
    bytes, error := json.Marshal(baseMessage.Payload)
    fmt.Println(bytes)
    if error != nil {
        fmt.Println("Parse Error: ", error)
        decoded := string(bytes)
        fmt.Printf("DECODED: %+v (%T)\n", decoded)
        return error
    }
    error = json.Unmarshal(bytes, intendedType)
    if error != nil {
        fmt.Println("UnMarshal Error: ", error)
        return error
    }
    return nil
}
