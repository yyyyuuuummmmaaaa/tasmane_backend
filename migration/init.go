package main

import (
  "gorm.io/gorm"
  postgres "packages/database"
)

type userService struct {
  gorm.Model
  Name string
  Description string
  Role string
  Email string
  Phone int
  Team string
}

func NewUser() *userService {
	return &userService{}
}

func main() {

  db := postgres.Connect()
	db.AutoMigrate(NewUser())
}