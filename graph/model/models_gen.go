// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Memo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type NewMemo struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
