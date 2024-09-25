package main

import (
	"fmt"

	"github.com/eldonmacdonald/GO-DB-System/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
