package handlers

import (
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "github.com/mortonar/acs560_course_project/messaging/messages/response"
    "net/url"
    "fmt"
//    "encoding/json"
    "github.com/mortonar/acs560_course_project/database/models"
)

func HandleBookList(searchReq request.BookList) (*response.BookList, error) {
    Name := url.QueryEscape(searchReq.Name)
fmt.Println(Name)
    // TODO super big hack. hard-coded response
    fmt.Println("Book List Response")
    
    bookList := response.BookList{make([]models.Book, 0)}

    bookList.Books = append(bookList.Books, models.Book{ISBN13: "a2salksjd5f5",
        Title: "Dracula", Author: "Bram Stoker"})
    bookList.Books = append(bookList.Books, models.Book{ISBN13: "aasd0f769a7sf",
        Title: "Legacies", Author: "L.E. Modsitt, JR."})
    
   // sr := ListResponse{}

 //   json.Unmarshal(bookList, &sr)

    return &bookList, nil
}

type ListResponse struct {
    Items []ResponseItem2
}

type ResponseItem2 struct {
    VolumeInfo2 VolumeInfo
}

type VolumneInfo2 struct {
    Name string
}


