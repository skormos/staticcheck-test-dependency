package utils

import "fmt"

// Awesome makes everything so.
func Awesome(s string) string {
	return fmt.Sprintf("your %s is awesome", s)
}

// SoSad is meant to be sad.
// Deprecated: Not sad anymore, use Awesome instead.
func SoSad(s string) string {
	return fmt.Sprintf("so sad your %s is not awesome", s)
}
