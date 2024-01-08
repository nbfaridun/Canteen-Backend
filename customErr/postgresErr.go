package customErr

import (
	"strings"
)

// IsDuplicateKeyError checks if the error indicates a duplicate key violation.
// It also returns the name of the column causing the violation (if available).
func IsDuplicateKeyError(err error) (bool, string) {
	if strings.Contains(err.Error(), "duplicate key") {
		// Example: duplicate key value violates unique constraint "users_username_key"
		parts := strings.Split(err.Error(), `"`)
		if len(parts) == 3 {
			return true, strings.Split(parts[1], "_")[1]
		}
	}

	return false, ""
}
