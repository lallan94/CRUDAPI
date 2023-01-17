package main

import "gorm.io/gorm"

type Studentdata struct {
	gorm.Model
	StudentName string `json:"studentname"`
	StudentID   string `json:"studentid"`
	BranchName  string `json:"branchname"`
	PassingYear int    `json:"passingyear"`
}
