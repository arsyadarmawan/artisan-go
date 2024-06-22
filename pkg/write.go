package pkg

import (
	"fmt"
	"os"
)

func CreateFile(path, content string) error {
	err := WriteToFile(path, content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return nil
	}
	return nil
}

func WriteToFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
