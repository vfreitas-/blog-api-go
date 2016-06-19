package models

import (
    "github.com/jinzhu/gorm"
)

type Category struct {
    gorm.Model
    Name string `json:"name"`
    ParentID int64 `json:"parent_id"`
}
