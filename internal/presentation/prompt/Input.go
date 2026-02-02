package prompt

import (
	"bufio"
	"fmt"
	"os"
)

type Input struct {
	Title string
	Value *string
}

func UpdateInputs(inputs []Input) error {
	for _, input := range inputs {
		newInput := NewInput{
			Title:       input.Title,
			Value:       input.Value,
			Description: "Input full path ",
		}
		if err := ScanInput(newInput); err != nil {
			return err
		}
	}
	return nil
}

type NewInput struct {
	Title       string
	Value       *string
	Description string
}

func NewInputs(inputs []NewInput) error {
	for _, input := range inputs {
		if err := ScanInput(input); err != nil {
			return err
		}
	}
	return nil
}

func ScanInput(input NewInput) error {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s (current %s) %s> ", input.Title, *input.Value, input.Description)

	sc.Scan()
	if err := sc.Err(); err != nil {
		return err
	}
	if text := sc.Text(); text != "" {
		*input.Value = text
	}
	return nil
}

type KeyValueInput struct {
	Title             string
	Key               *string
	Value             *string
	KeyDescritpiton   string
	ValueDescritpiton string
	QuitMark          string
}

func KeyValueInputs(input KeyValueInput) error {
	flag := true
	sc := bufio.NewScanner(os.Stdin)
	for flag {
		sc.Scan()
		if err := sc.Err(); err != nil {
			return err
		}
	}
	return nil
}
