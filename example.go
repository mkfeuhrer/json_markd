package json_markd

import (
	"fmt"
)

// Example usage for json_markd
func main() {
	SetTabSpaceValue(2)
	SetupLogger()
	result, err := ParseMarkdown("data/sample_api.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
	return
}
