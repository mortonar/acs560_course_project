package main

import (
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/mortonar/acs560_course_project/server"
)

func main() {
    s := &server.Server{}
    s.Start()
}