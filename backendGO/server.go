package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/////// Why does this doesn't work?
// type Books struct {
// 	Kind       string `json:"kind"`
// 	TotalItems int    `json:"totalItems"`
// 	Items      []struct {
// 		Info VolumeInfo `json:volumeInfo, omitempty`
// 	} `json:"items"`
// }

// type VolumeInfo struct {
// 	Title     string   `json:"title"`
// 	Authors   []string `json:"authors"`
// 	PageCount string   `json:"pageCount"`
// 	Language  string   `json:"language"`
// }

type Books struct {
	Items []struct {
		VolumeInfo struct {
			Title     string   `json:"title"`
			PageCount int      `json:"pageCount"`
			Authors   []string `json:"authors"`
			Language  string   `json:"language"`
		} `json:"volumeInfo,omitempty"`
	} `json:"items"`
}

type GetBooksBody struct {
	BookName string
}

func getBooks(w http.ResponseWriter, req *http.Request) {
	// var bookName GetBooksBody

	// err := json.NewDecoder(req.Body).Decode(&bookName)

	// if err != nil {
	// 	panic(err)
	// }
	bookName := req.URL.Query().Get("bookName")
	url := "https://www.googleapis.com/books/v1/volumes?q=" + bookName
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

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
