package front

import (
	"fmt"
	"github.com/MakeNowJust/heredoc/v2"
)

type (
	Console struct{}
)

func (this *Console) GetInt(desc string) int {
	fmt.Print("\n")
	fmt.Println("---------------")
	fmt.Println(heredoc.Doc(desc))
	fmt.Println("---------------")
	fmt.Print(":")

	var intVal int
	fmt.Scan(&intVal)

	return intVal
}

func (this *Console) Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

func (this *Console) Printf(format string, a ...any) (n int, err error) {
	return fmt.Printf(format, a...)
}
