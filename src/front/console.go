package front 

import (
	"fmt"
	"github.com/MakeNowJust/heredoc/v2"
)

type (
	Console struct {}
)

var console Console

func (this Console) GetInt(desc string) int {
	fmt.Print("\n")
	fmt.Println("---------------")
	fmt.Println(heredoc.Doc(desc))
	fmt.Println("---------------")
	fmt.Print(":")

	var intVal int
	fmt.Scan(&intVal)

	return intVal
}

func (this Console) Println(a...interface{}) (n int, err error) {
    return fmt.Println(a...)
}