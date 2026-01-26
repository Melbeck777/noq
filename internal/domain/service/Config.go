package service

import "fmt"

func EmptyCheck(target, title string, errorType error) error {
	if target == "" {
		return fmt.Errorf("%w: %s is empty\n", errorType, title)
	}
	return nil
}
