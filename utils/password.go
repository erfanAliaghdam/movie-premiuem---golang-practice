package utils

import "fmt"

func HashPassword(password string) string {
	return fmt.Sprintf("hashed_%s", password)
}
