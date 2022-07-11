package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Books struct {
	Items []Items `json:"items"`
}

type Items struct {
	VolumeInfo VolumeInfo `json:"volumeInfo,omitempty"`
}

type VolumeInfo struct {
	Title     string   `json:"title"`
	PageCount int      `json:"pageCount"`
	Authors   []string `json:"authors"`
	Language  string   `json:"language"`
}

type GetBooksBody struct {
	BookName string
}

func getBooks(w http.ResponseWriter, req *http.Request) {
	bookName := req.URL.Query().Get("bookName")
	escapedBook := url.PathEscape(bookName)

	url := "https://www.googleapis.com/books/v1/volumes?q=" + escapedBook
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	books := Books{}
	err = json.Unmarshal(body, &books)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	http.HandleFunc("/getBooks", getBooks)
	http.ListenAndServe(":8090", nil)
}
