package e

import "fmt"

func Wrap(msgErr string, err error) error {
	return fmt.Errorf("%w: %w", msgErr, err)
}

// Для случая, когда ошибка не нулевая
func WrapIfErr(msgErr string, err error) error {
	if err != nil {
		return Wrap(msgErr, err)
	}

	return nil
}
