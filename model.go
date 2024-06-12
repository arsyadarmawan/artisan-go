package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"mycli/pkg"
	"mycli/template"
	"os"
)

func modelCommand() *cli.Command {
	return &cli.Command{
		Name:      template.ModelCommand,
		Aliases:   []string{"c"},
		Usage:     "create entity model",
		ArgsUsage: "[model]",
		Action: func(c *cli.Context) error {
			name := c.Args().Get(0)
			if name == "" {
				fmt.Errorf("model name is required")
				return nil
			}
			filename := name + ".go"
			path := "./internal/app/ent/" + filename
			values := map[string]string{
				"{{modelName}}":   pkg.ToPascalCase(name),
				"{{packageName}}": "ent",
			}
			result := pkg.ReplacePlaceholders(template.ModelTemplate, values)
			err := WriteToFile(path, result)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return nil
			}
			fmt.Printf("File written successfully:", filename)

			return nil
		},
	}
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
