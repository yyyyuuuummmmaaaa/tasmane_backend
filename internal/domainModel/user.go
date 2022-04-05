package model

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Name string
  Description string
  Role string
  Email string
  Phone int
  Team: string
}