package models

import (
    "github.com/jinzhu/gorm"
)

type Category struct {
    gorm.Model
    Name string `json:"name"`
    Parent Category
}
