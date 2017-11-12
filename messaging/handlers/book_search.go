package handlers

import (
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "github.com/mortonar/acs560_course_project/messaging/messages/response"
    "net/url"
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/mortonar/acs560_course_project/database/models"
    "strings"
)

func HandleBookSearch(searchReq request.BookSearch) (*response.BookSearch, error) {
    title := url.QueryEscape(searchReq.Title)
    author := url.QueryEscape(searchReq.Author)
    // TODO super big hack. hard-coded api key, results limit, return fields, etc.
    searchString := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=intitle:%s+inauthor:%s&" +
        "key=AIzaSyCD-MeEYy5AxnvsH--Tm--g3VOPedkAW8s" +
            "&fields=kind,items(volumeInfo(title,authors,description,industryIdentifiers))" +
                "&maxResults=10", title, author)
    resp, error := http.Get(searchString)
    fmt.Println("Response: ", resp)
    defer resp.Body.Close()
    if error != nil {
        return nil, error
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    sr := SearchResponse{}
    json.Unmarshal(body, &sr)
    fmt.Println("SearchResponse:", sr)
    return makeResponse(sr), nil
}

func makeResponse(searchResponse SearchResponse) *response.BookSearch {
    bookSearch := response.BookSearch{make([]models.Book, 0)}
    for _, item := range searchResponse.Items {
        book := makeBook(item.VolumeInfo)
        if book.ISBN13 != "" && book.Title != "" && book.Author != "" {
            fmt.Println("Book being added: ", book)
            bookSearch.Books = append(bookSearch.Books, book)
        }
    }

    fmt .Println("Returning results: ", bookSearch)
    return &bookSearch
}

func makeBook(info VolumeInfo) models.Book {
    isbn13 := ""
    for _, isbn := range info.IndustryIdentifiers {
       if isbn.Type == "ISBN_13" {
           isbn13 = isbn.Identifier
           break
       }
    }
    author := strings.Join(info.Authors, ",")
    return models.Book{ISBN13: isbn13, Title: info.Title, Author: author}
}

type SearchResponse struct {
    Items []ResponseItem
}

type ResponseItem struct {
    VolumeInfo VolumeInfo
}

type VolumeInfo struct {
    Title string
    Authors []string
    IndustryIdentifiers []Isbn
}

type Isbn struct {
    Type string
    Identifier string
}
