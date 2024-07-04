package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"mycli/pkg"
	"mycli/template"
)

func validatorCommand() *cli.Command {
	return &cli.Command{
		Name:      template.ValidatorTemplate,
		Aliases:   []string{"c"},
		Usage:     "create entity model",
		ArgsUsage: "[model]",
		Action: func(c *cli.Context) error {
			name := c.Args().Get(0)
			if name == "" {
				return errors.New("validator name is required")
			}
			filename := name + ".go"
			modelPath := fmt.Sprintf(template.ModelDirectory, filename)
			values := map[string]string{
				"{{modelName}}":   pkg.ToPascalCase(name),
				"{{packageName}}": "ent",
			}
			content := pkg.ReplacePlaceholders(template.ModelTemplate, values)
			if err := pkg.WriteToFile(modelPath, content); err != nil {
				fmt.Println("Error writing to file:", err)
				return err
			}
			fmt.Printf("File written successfully:", filename)

			return nil
		},
	}
}
