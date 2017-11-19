package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/mortonar/acs560_course_project/database/models"
	"github.com/mortonar/acs560_course_project/messaging/messages/request"
	"github.com/mortonar/acs560_course_project/messaging/messages/response"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

func HandleBookSearch(searchReq request.BookSearch) (*response.BookSearch, error) {
	query := url.QueryEscape(searchReq.Title + " " + searchReq.Author)
	// TODO super big hack. hard-coded api key, results limit, return fields, etc.
	searchString := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s&"+
		"key=AIzaSyCD-MeEYy5AxnvsH--Tm--g3VOPedkAW8s"+
		"&fields=kind,items(volumeInfo(title,authors,description,industryIdentifiers,imageLinks))"+
		"&maxResults=10", query)
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

	fmt.Println("Returning results: ", bookSearch)
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
	thumbnail, image := extractImageURLs(info.ImageLinks)
	return models.Book{ISBN13: isbn13,
		Title:        info.Title,
		Author:       author,
		Description:  info.Description,
		ThumbnailURL: thumbnail,
		ImageURL:     image,
	}
}

var thumbnail_types = []string{"Small", "Thumbnail", "SmallThumbnail"}
var image_types = []string{"ExtraLarge", "Large", "Medium"}

func extractImageURLs(imageLinks ImageLinks) (thumbnail, image string) {
	thumbnail = extractUrl(imageLinks, thumbnail_types)
	image = extractUrl(imageLinks, image_types)
	return
}

// search through 'imageLinks' via reflection for non-empty fields in order of preference given by 'urls'
func extractUrl(imageLinks ImageLinks, urls []string) (chosenURL string) {
	chosenURL = ""
	val := reflect.ValueOf(imageLinks)
	for _, url := range urls {
		urlField := val.FieldByName(url)
		if len(urlField.String()) > 0 {
			chosenURL = urlField.String()
			break
		}
	}
	return
}

type SearchResponse struct {
	Items []ResponseItem
}

type ResponseItem struct {
	VolumeInfo VolumeInfo
}

type VolumeInfo struct {
	Title               string
	Authors             []string
	Description         string
	IndustryIdentifiers []Isbn
	ImageLinks          ImageLinks
}

type Isbn struct {
	Type       string
	Identifier string
}

type ImageLinks struct {
	SmallThumbnail string
	Thumbnail      string
	Small          string
	Medium         string
	Large          string
	ExtraLarge     string
}
