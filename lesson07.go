package main

import (
	"fmt"
)

/* map（object） */

func main() {
	user := map[string]string{
		"name": "pienman",
		"age":  "7", // string
	}
	fmt.Println(user["name"])

	type userType struct {
		name  string
		age   int
		login bool
	}

	user1 := userType{
		name:  "pienman",
		age:   7,
		login: true,
	}

	fmt.Println(user1.name, "is login.")

}
