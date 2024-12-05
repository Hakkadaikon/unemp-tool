package front

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
)

type (
	Console struct{}
)

func (this *Console) GetInt(desc string) int {
	validate := func(input string) error {
		_, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    desc,
		Validate: validate,
	}

	intVal := 0
	result, err := prompt.Run()
	if err == nil {
		intVal, _ = strconv.Atoi(result)
	}

	return intVal
}

func (this *Console) Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

func (this *Console) Printf(format string, a ...any) (n int, err error) {
	return fmt.Printf(format, a...)
}
