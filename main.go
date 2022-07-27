package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"hello/db"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
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

func CreateDBConnection() (*sql.DB, error) {
	connStr := "postgres://shahroz:mysecretpassword@localhost:5452/test?sslmode=disable"

	c, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return c, err
}

func addBooks(w http.ResponseWriter, req *http.Request) {
	bookName := req.URL.Query().Get("bookName")
	escapedBook := url.PathEscape(bookName)

	url := "https://www.googleapis.com/books/v1/volumes?q=" + escapedBook
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	conn, err := CreateDBConnection()
	queries := db.New(conn)

	if err != nil {
		fmt.Println("err = ", err)
	}

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

	for _, item := range books.Items {
		name := item.VolumeInfo.Title
		id := uuid.New()
		params := db.AddBooksParams{ID: id, Name: sql.NullString{String: name, Valid: true}}
		queries.AddBooks(ctx, params)
	}

	json.NewEncoder(w).Encode(books)

}

func getBooksFromDB(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	conn, err := CreateDBConnection()
	if err != nil {
		fmt.Println("err = ", err)
	}

	queries := db.New(conn)

	books, err := queries.ListBooks(ctx)

	if err != nil {
		fmt.Println("err = ", err)
	}

	json.NewEncoder(w).Encode(books)
}

func main() {
	http.HandleFunc("/addBooks", addBooks)
	http.HandleFunc("/getBooks", getBooksFromDB)
	http.ListenAndServe(":8090", nil)
}
