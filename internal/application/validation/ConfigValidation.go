package validation

import "fmt"

func EmptyCheck(target, title string, errorType error) error {
	if target == "" {
		return fmt.Errorf("%w: %s is empty", errorType, title)
	}
	return nil
}
