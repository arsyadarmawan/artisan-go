package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"mycli/pkg"
	"mycli/template"
)

func usecaseCommand() *cli.Command {
	return &cli.Command{
		Name:    template.UsecaseCommand,
		Aliases: []string{"t"},
		Usage:   "options for task templates",
		Action: func(c *cli.Context) error {
			domainName := c.String("domain")
			domainPath := fmt.Sprintf(template.DomainName, domainName)
			if errDomain := checkFolderIsExist(domainName, domainPath); errDomain != nil {
				return errDomain
			}

			usecaseFolderPath := fmt.Sprintf(template.UsecaseFolder, domainName)
			usecaseName := c.String("usecase")
			if errUsecase := checkFolderIsExist(domainName, usecaseFolderPath); errUsecase != nil {
				return errUsecase
			}
			if errCreateUsecaseImpl := createUsecaseImplFile(domainName, usecaseName); errCreateUsecaseImpl != nil {
				return errCreateUsecaseImpl
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "domain",
				Aliases: []string{"dmn"},
				Usage:   "check domain is exist",
			},
			&cli.StringFlag{
				Name:    "usecase",
				Aliases: []string{"usc"},
				Usage:   "create usecase file",
			},
		},
	}
}

func checkFolderIsExist(domainName, path string) error {
	if domainName == "" {
		return errors.New("domain is required\n")
	}
	if err := pkg.CheckFolderIsExist(path); err != nil {
		return err
	}
	return nil
}

func createUsecaseFile(domainName, usecaseName string) error {
	fileName := usecaseName + ".go"
	values := map[string]string{
		"{{structName}}": pkg.ToPascalCase(usecaseName),
		"{{sourceFile}}": usecaseName,
	}
	usecasePath := fmt.Sprintf(template.UsecaseDetailName, domainName)
	fullPath := usecasePath + fileName
	contentPath := pkg.ReplacePlaceholders(template.UsecaseTemplate, values)
	if errCreateUsecase := pkg.CreateFile(fullPath, contentPath); errCreateUsecase != nil {
		fmt.Println("Error writing to file:")
		return errCreateUsecase
	}
	return nil
}

func createUsecaseImplFile(domainName, usecaseName string) error {
	fileName := usecaseName + ".go"
	if errCreateFile := createUsecaseFile(domainName, usecaseName); errCreateFile != nil {
		return errCreateFile
	}

	values := map[string]string{
		"{{name}}": pkg.ToPascalCase(usecaseName),
	}

	usecaseImplPath := fmt.Sprintf(template.UsecaseImplFile, domainName)
	usecaseImplFullPath := usecaseImplPath + fileName
	contentUsecaseImpl := pkg.ReplacePlaceholders(template.UsecaseImplTemplate, values)
	if errCreateUsecaseimpl := pkg.CreateFile(usecaseImplFullPath, contentUsecaseImpl); errCreateUsecaseimpl != nil {
		return errCreateUsecaseimpl
	}
	return nil
}
