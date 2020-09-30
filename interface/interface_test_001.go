package main

import (
	"fmt"
)

type User struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"` 
	Role   int     `json:role"`
}

func ( this User ) IsAdmin() bool {
	if this.Role == 2 || this.Role == 1 {
		return true
	} 
	return false
}


func main() {
	terry := User{133,"terry.tsang",2}
	kelly := User{113,"kelly.tsang",12}

	if terry.IsAdmin() {
		fmt.Println("admin")
	} else {
		fmt.Println("user")
	}
	if kelly.IsAdmin() {
		fmt.Println("admin")
	} else {
		fmt.Println("user")
	}
}
