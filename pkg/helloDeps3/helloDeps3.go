package helloDeps3

import "fmt"

func PrintPhrase(phrase string) string {
	fmt.Println(phrase, "called helloDeps3.PrintPhrase")
	return phrase
}
