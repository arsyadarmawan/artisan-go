package pkg

import (
	"fmt"
	"os"
)

func WriteFile(path string, values map[string]string, template string) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()
	content := ReplacePlaceholders(template, values)
	fmt.Printf(content)
	_, err = file.WriteString(content)
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}
}
