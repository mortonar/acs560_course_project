package main

import (
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/mortonar/acs560_course_project/sever"
)

func main() {
    s := &sever.Server{}
    s.Start()
}