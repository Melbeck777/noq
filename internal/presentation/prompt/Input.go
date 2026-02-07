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

type MapInput struct {
	KeyDescritpiton   string
	ValueDescritpiton string
	QuitMark          string
	Map               map[string]string
}

func MapInputs(input MapInput) {
	flag := true
	for flag {
		key := ScanDetail(input.KeyDescritpiton, input.QuitMark)
		if key == "" {
			return
		}
		if v, ok := input.Map[key]; ok {
			// 同じkeyやvalueが入った時にはじくようにする
			fmt.Printf("This %s is already exisit, value is %s\n- Input 'c' for continue\n- Input 'i' for ignore this input\n- Input 'q' for quiting input database\n", key, v)
		}

		value := ScanDetail(input.ValueDescritpiton, input.QuitMark)
		if value == "" {
			return
		}
		input.Map[key] = value
	}
}

func ScanDetail(description, quitMark string) string {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Printf(description)
	sc.Scan()
	if sc.Err() != nil {
		return ""
	}
	text := sc.Text()
	if text == "q" {
		return ""
	}
	return text
}
