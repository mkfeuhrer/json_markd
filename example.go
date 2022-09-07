package json_markd

import (
	"fmt"
	"os"
)

// Example usage for json_markd
func main() {
	SetTabSpaceValue(2)
	SetupLogger()
	f, err := os.Open("data/sample_api.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := ParseMarkdown(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	return
}
