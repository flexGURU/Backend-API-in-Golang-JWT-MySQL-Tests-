package main

import (
	"fmt"

	"github.com/flexGURU/goAPI/types"
)

func main() {

	user := new(types.User)

	user.Email = "mukuna"

	fmt.Println(user)

}