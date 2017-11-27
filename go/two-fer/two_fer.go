// Describing share with others
package twofer

import (
	"fmt"
)

// Returns a string describes the share with others.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
