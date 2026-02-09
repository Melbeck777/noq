package prompt_test

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/Melbeck777/noq/internal/presentation/prompt"
)

// Helper function to mock os.Stdin
func setStdin(input string) func() {
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	go func() {
		defer w.Close()
		io.WriteString(w, input)
	}()
	return func() {
		os.Stdin = oldStdin
	}
}

// Helper function to capture os.Stdout
func captureStdout() (string, func()) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	return "", func() {
		w.Close()
		out, _ := io.ReadAll(r)
		_ = out // We might not use the output, but this is how to get it
		os.Stdout = oldStdout
	}
}

func TestScanDetail(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		description string
		quitMark    string
		want        string
	}{
		{
			name:        "normal input",
			input:       "test input\n",
			description: "Enter value:",
			quitMark:    "q",
			want:        "test input",
		},
		{
			name:        "quit input",
			input:       "q\n",
			description: "Enter value:",
			quitMark:    "q",
			want:        "",
		},
		{
			name:        "empty input with no quit mark",
			input:       "\n",
			description: "Enter value:",
			quitMark:    "",
			want:        "",
		},
		{
			name:        "input equals quit mark",
			input:       "exit\n",
			description: "Enter value:",
			quitMark:    "exit",
			want:        "",
		},
		{
			name:        "input not equal to quit mark",
			input:       "continue\n",
			description: "Enter value:",
			quitMark:    "exit",
			want:        "continue",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer setStdin(tt.input)()
			_, restoreStdout := captureStdout()
			defer restoreStdout()

			got := prompt.ScanDetail(tt.description, tt.quitMark)
			if got != tt.want {
				t.Errorf("ScanDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanInput(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		title       string
		currentVal  string
		description string
		wantVal     string
		wantErr     bool
	}{
		{
			name:        "new input",
			input:       "new value\n",
			title:       "Test Input",
			currentVal:  "old value",
			description: "Enter something:",
			wantVal:     "new value",
			wantErr:     false,
		},
		{
			name:        "empty input (keep current)",
			input:       "\n",
			title:       "Test Input",
			currentVal:  "old value",
			description: "Enter something:",
			wantVal:     "old value",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer setStdin(tt.input)()
			_, restoreStdout := captureStdout()
			defer restoreStdout()

			val := tt.currentVal
			input := prompt.NewInput{
				Title:       tt.title,
				Value:       &val,
				Description: tt.description,
			}

			err := prompt.ScanInput(&input)

			if (err != nil) != tt.wantErr {
				t.Errorf("ScanInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if *input.Value != tt.wantVal {
				t.Errorf("ScanInput() got value = %v, want %v", *input.Value, tt.wantVal)
			}
		})
	}
}

func TestSelectInput(t *testing.T) {
	type args struct {
		items     []string
		indexName string
	}
	tests := []struct {
		name         string
		args         args
		input        string
		wantSelected string
		wantOk       bool
		wantErr      bool
	}{
		{
			name: "valid selection",
			args: args{
				items:     []string{"option1", "option2", "option3"},
				indexName: "Choose an option",
			},
			input:        "1\n",
			wantSelected: "option2",
			wantOk:       true,
			wantErr:      false,
		},
		{
			name: "invalid input - not a number",
			args: args{
				items:     []string{"option1", "option2"},
				indexName: "Choose an option",
			},
			input:        "abc\n",
			wantSelected: "",
			wantOk:       false,
			wantErr:      false,
		},
		{
			name: "out of range - negative index",
			args: args{
				items:     []string{"option1", "option2"},
				indexName: "Choose an option",
			},
			input:        "-1\n",
			wantSelected: "",
			wantOk:       false,
			wantErr:      true,
		},
		{
			name: "out of range - too high index",
			args: args{
				items:     []string{"option1", "option2"},
				indexName: "Choose an option",
			},
			input:        "2\n",
			wantSelected: "",
			wantOk:       false,
			wantErr:      true,
		},
		{
			name: "empty input",
			args: args{
				items:     []string{"option1", "option2"},
				indexName: "Choose an option",
			},
			input:        "\n",
			wantSelected: "",
			wantOk:       false,
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer setStdin(tt.input)()
			_, restoreStdout := captureStdout()
			defer restoreStdout()

			gotSelected, gotOk, err := prompt.SelectInput(tt.args.items, tt.args.indexName, func(s string) string { return s })

			if (err != nil) != tt.wantErr {
				t.Errorf("SelectInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSelected != tt.wantSelected {
				t.Errorf("SelectInput() gotSelected = %v, want %v", gotSelected, tt.wantSelected)
			}
			if gotOk != tt.wantOk {
				t.Errorf("SelectInput() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestMapKeySelect(t *testing.T) {
	tests := []struct {
		name     string
		inputMap map[string]string
		input    string
		title    string
		want     string
		wantErr  bool
	}{
		{
			name:     "valid selection",
			inputMap: map[string]string{"key1": "value1", "key2": "value2"},
			input:    "0\n", // Assuming "key1" is at index 0 after sorting
			title:    "Select Key",
			want:     "key1",
			wantErr:  false,
		},
		{
			name:     "cancelled input",
			inputMap: map[string]string{"key1": "value1"},
			input:    "abc\n", // Invalid input, should lead to cancellation
			title:    "Select Key",
			want:     "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer setStdin(tt.input)()
			_, restoreStdout := captureStdout()
			defer restoreStdout()

			got, err := prompt.MapKeySelect(tt.inputMap, tt.title)

			if (err != nil) != tt.wantErr {
				t.Errorf("MapKeySelect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MapKeySelect() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateInputs(t *testing.T) {
	tests := []struct {
		name       string
		inputs     []prompt.Input
		mockInputs []string // for stdin
		wantValues []string
		wantErr    bool
	}{
		{
			name: "update multiple inputs",
			inputs: []prompt.Input{
				{Title: "Input 1", Value: new(string), Description: "Test"},
				{Title: "Input 2", Value: new(string), Description: "Test"},
			},
			mockInputs: []string{"value1\n", "value2\n"},
			wantValues: []string{"value1", "value2"},
			wantErr:    false,
		},
		{
			name: "update with empty input (keep existing value)",
			inputs: []prompt.Input{
				{Title: "Input 1", Value: func() *string { s := "initial1"; return &s }(), Description: "Test"},
				{Title: "Input 2", Value: new(string), Description: "Test"},
			},
			mockInputs: []string{"\n", "value2\n"},
			wantValues: []string{"initial1", "value2"},
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			combinedInput := strings.Join(tt.mockInputs, "")
			defer setStdin(combinedInput)()
			_, restoreStdout := captureStdout()
			defer restoreStdout()

			// Initialize values to avoid nil pointer dereference if not already done
			for i := range tt.inputs {
				if tt.inputs[i].Value == nil {
					tt.inputs[i].Value = new(string)
				}
			}
			fmt.Println("updateInputs call")
			err := prompt.UpdateInputs(tt.inputs)

			fmt.Println("updateInputs finish")
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateInputs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for i := range tt.inputs {
				if *tt.inputs[i].Value != tt.wantValues[i] {
					t.Errorf("UpdateInputs() input %d got value = %v, want %v", i, *tt.inputs[i].Value, tt.wantValues[i])
				}
			}
		})
	}
}

func TestNewInputs(t *testing.T) {
	tests := []struct {
		name       string
		newInputs  []prompt.NewInput
		mockInputs []string
		wantValues []string
		wantErr    bool
	}{
		{
			name: "create multiple new inputs",
			newInputs: []prompt.NewInput{
				{Title: "New Input 1", Value: new(string), Description: "Desc1"},
				{Title: "New Input 2", Value: new(string), Description: "Desc2"},
			},
			mockInputs: []string{"valA\n", "valB\n"},
			wantValues: []string{"valA", "valB"},
			wantErr:    false,
		},
		{
			name: "create new inputs with empty entry",
			newInputs: []prompt.NewInput{
				{Title: "New Input 1", Value: nil, Description: "Desc1"},
				{Title: "New Input 2", Value: nil, Description: "Desc2"},
			},
			mockInputs: []string{"\n", "valB\n"},
			wantValues: []string{"", "valB"}, // Empty input means value remains empty string
			wantErr:    false,
		},
	}

	for j := range tests {
		t.Run(tests[j].name, func(t *testing.T) {
			combinedInput := strings.Join(tests[j].mockInputs, "")
			defer setStdin(combinedInput)()
			_, restoreStdout := captureStdout()
			defer restoreStdout()

			// Initialize values
			for i := range tests[j].newInputs {
				if tests[j].newInputs[i].Value == nil {
					tests[j].newInputs[i].Value = new(string)
				}
			}

			err := prompt.NewInputs(tests[j].newInputs)

			if (err != nil) != tests[j].wantErr {
				t.Errorf("NewInputs() error = %v, wantErr %v", err, tests[j].wantErr)
				return
			}

			for i := range tests[j].newInputs {
				if *tests[j].newInputs[i].Value != tests[j].wantValues[i] {
					t.Errorf("%d NewInputs() input %d got value = %v, want %v", j, i, *tests[j].newInputs[i].Value, tests[j].wantValues[i])
				}
			}
		})
	}
}
