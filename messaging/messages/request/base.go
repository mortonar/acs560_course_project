package request

type Base struct {
    Token string
    Action string
    Payload interface{}
}