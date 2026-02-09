package prompt

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Input struct {
	Title       string
	Value       *string
	Description string
}

// scanAndApplyInput scans one line from the provided scanner and applies it to the input value.
func scanAndApplyInput(input *NewInput, sc *bufio.Scanner) error {
	currentValue := ""
	if input.Value != nil {
		currentValue = *input.Value
	}
	fmt.Printf("%s (current: %s) %s> ", input.Title, currentValue, input.Description)

	sc.Scan()
	if err := sc.Err(); err != nil {
		return err
	}
	if text := sc.Text(); text != "" {
		if input.Value == nil {
			input.Value = new(string)
		}
		*input.Value = text
	}
	return nil
}

func UpdateInputs(inputs []Input) error {
	fmt.Println("Update start")
	sc := bufio.NewScanner(os.Stdin) // Create scanner once
	for i := range inputs {
		if inputs[i].Value == nil {
			inputs[i].Value = new(string)
		}
		newInput := NewInput{
			Title:       inputs[i].Title,
			Value:       inputs[i].Value,
			Description: inputs[i].Description,
		}
		if err := scanAndApplyInput(&newInput, sc); err != nil { // Pass scanner to helper
			return err
		}
	}
	fmt.Println("Update sucess")
	return nil
}

type NewInput struct {
	Title       string
	Value       *string
	Description string
}

func NewInputs(inputs []NewInput) error {
	fmt.Println("New input start")
	sc := bufio.NewScanner(os.Stdin) // Create scanner once
	for i := range inputs {
		if inputs[i].Value == nil {
			inputs[i].Value = new(string)
		}
		if err := scanAndApplyInput(&inputs[i], sc); err != nil { // Pass scanner to helper
			return err
		}
	}
	fmt.Println("New input sucess")
	return nil
}

func MapKeySelect(m map[string]string, title string) (string, error) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	selectedValue, ok, err := SelectInput[string](keys, title, func(s string) string { return s })
	if err != nil {
		return "", err
	}
	if !ok {
		return "", fmt.Errorf("Cancelled input")
	}
	return selectedValue, nil
}

func ScanInput(input *NewInput) error {
	sc := bufio.NewScanner(os.Stdin)
	return scanAndApplyInput(input, sc)
}

type MapInput struct {
	KeyDescription   string
	ValueDescription string
	QuitMark         string
	Map              map[string]string
}

func MapInputs(input MapInput) {
	sc := bufio.NewScanner(os.Stdin)
	for {
		key := ScanDetail(input.KeyDescription, input.QuitMark)
		if key == "" {
			return
		}

		if v, ok := input.Map[key]; ok {
			fmt.Printf("This key '%s' already exists with value '%s'.\n", key, v)
			fmt.Print("- Input 'c' to continue and overwrite.\n- Input 'i' to ignore this key and enter a new one.\n- Input 'q' to quit.\n> ")
			sc.Scan()
			choice := strings.TrimSpace(sc.Text())
			switch choice {
			case "c":
				// Proceed to get value
			case "i":
				continue // Restart loop for a new key
			case "q":
				return
			default:
				fmt.Println("Invalid choice. Ignoring and entering a new key.")
				continue
			}
		}

		value := ScanDetail(input.ValueDescription, input.QuitMark)
		if value == "" && input.QuitMark == "" {
			// allow empty value if no quit mark
		} else if value == "" {
			return
		}
		input.Map[key] = value
	}
}

func ScanDetail(description, quitMark string) string {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print(description)
	sc.Scan()
	if sc.Err() != nil {
		return ""
	}
	text := sc.Text()
	if quitMark != "" && text == quitMark {
		return ""
	}
	return text
}

func SelectInput[T any](items []T, indexName string, label func(T) string) (selected T, ok bool, err error) {
	fmt.Printf("Index: %s\n", indexName)
	for i, item := range items {
		fmt.Printf("%d %s\n", i, label(item))
	}

	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Input the index you want to set: ")
	sc.Scan()
	if err := sc.Err(); err != nil {
		return selected, false, err
	}

	raw := strings.TrimSpace(sc.Text())
	if raw == "" {
		return selected, false, nil // Treat empty input as cancellation
	}

	index, err := strconv.Atoi(raw)
	if err != nil {
		// Non-integer input is also treated as cancellation
		return selected, false, nil
	}

	if index < 0 || index >= len(items) {
		return selected, false, fmt.Errorf("out of range: index %d is not within [0, %d)", index, len(items))
	}

	return items[index], true, nil
}
