package json_markd

import (
	"fmt"
	"json_markd/logger"
)

// Example usage for json_markd
func main() {
	SetTabSpaceValue(2)
	logger.SetupLogger()
	result, err := ParseMarkdown("data/sample_api.md")
	if err != nil {
		logger.Log.Panic(err)
		return
	}
	fmt.Println(result)
	return
}
