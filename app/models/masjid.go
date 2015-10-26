package models

import  "github.com/jinzhu/gorm"

type Masjid struct {
    gorm.Model
    Name string
}