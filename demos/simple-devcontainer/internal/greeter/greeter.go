package greeter

import "fmt"

func Greetings(to string) string {
	return fmt.Sprintf("Hello dear %s. I wish you a wonderful day!", to)
}
