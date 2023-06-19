package model

import "gorm.io/gorm"

type Blog struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Message string `json:"message"`
	gorm.Model
}
