package main

import (
	"fmt"
	"user_cli/menu"
	"user_cli/utils"
)

func main() {
	fmt.Println("Start user CLI")

	users := utils.Database()
	menu.Menu(&users)
}
