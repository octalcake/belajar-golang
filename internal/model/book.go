package model

import (
	"encoding/json"
)

type BookInput struct {
	Title string `json:"title" binding:"required,"`
	Price int    `json:"price" binding:"required,number"`

	// bisa menggunakan json.[Item] untuk menghindari internal server error
	// akibat error yang dilewati oleh go validator
	PageCount json.Number `json:"pageCount" binding:"required,number"`

	//alias JSON dalam go
	//Subtitle string `json:"sub_title"`
}
