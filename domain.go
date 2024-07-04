package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"mycli/pkg"
	"mycli/template"
	"os"
)

func domainCommand() *cli.Command {
	return &cli.Command{
		Name:      template.DomainCommand,
		Aliases:   []string{"d"},
		Usage:     "create domain model",
		ArgsUsage: "[model]",
		Action: func(c *cli.Context) error {
			name := c.Args().Get(0)
			if name == "" {
				panic("domain name is required")
			}
			if err := createParentDirectory(name); err != nil {
				return err
			}

			if err := createDomainChildDirectory(name); err != nil {
				return err
			}

			// create wire file
			if err := createWire(name); err != nil {
				return err
			}
			fmt.Printf("Domain created \n")
			return nil
		},
	}
}

func createWire(dirName string) error {
	appFolder := fmt.Sprintf(template.DomainName, dirName)
	fileName := appFolder + "/wire.go"

	values := map[string]string{
		"{{packageName}}": dirName,
	}
	content := pkg.ReplacePlaceholders(template.WireTemplate, values)
	if err := pkg.CreateFile(fileName, content); err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}

func createParentDirectory(dirName string) error {
	folderName := fmt.Sprintf(template.DomainName, dirName)
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		return err
	}
	return nil
}

func createDomainChildDirectory(dirName string) error {
	appFolder := fmt.Sprintf(template.DomainName, dirName)
	if errCreateDelivery := createDelivery(appFolder, dirName); errCreateDelivery != nil {
		return errCreateDelivery
	}
	if errCreateUsecase := createUsecase(appFolder, dirName); errCreateUsecase != nil {
		return errCreateUsecase
	}

	if errCreateRepository := createRepositoryFolder(appFolder, dirName); errCreateRepository != nil {
		return errCreateRepository
	}
	return nil
}

func createRepositoryFolder(parentFolder, structName string) error {
	repositoryPath := parentFolder + "/repository"
	repositoryMockPath := parentFolder + "/repository/repositorymock"

	if errRepository := os.MkdirAll(repositoryPath, 0755); errRepository != nil {
		return errRepository
	}
	if errRepositoryMock := os.MkdirAll(repositoryMockPath, 0755); errRepositoryMock != nil {
		return errRepositoryMock
	}
	if errCreateFile := createRepositoryFile(structName); errCreateFile != nil {
		return errCreateFile
	}
	fmt.Printf("repository is created \n")
	return nil
}

func createRepositoryFile(folderName string) error {
	fileName := folderName + ".go"
	values := map[string]string{
		"{{structName}}": pkg.ToPascalCase(folderName),
		"{{sourceFile}}": folderName,
	}
	repositoryPath := fmt.Sprintf("./internal/app/%s/repository/", folderName)
	fullPath := repositoryPath + fileName
	content := pkg.ReplacePlaceholders(template.RepositoryTemplate, values)
	if err := pkg.CreateFile(fullPath, content); err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}

func createDelivery(parentFolder, structName string) error {
	deliveryPath := parentFolder + "/delivery/web"
	if errDelivery := os.MkdirAll(deliveryPath, 0755); errDelivery != nil {
		return errDelivery
	}

	if errFile := createFileRegistryName(structName); errFile != nil {
		return errFile
	}
	return nil
}

func createFileRegistryName(structName string) error {
	fileName := "registry.go"
	values := map[string]string{
		"{{domainName}}": pkg.ToPascalCase(structName),
	}
	path := fmt.Sprintf("./internal/app/%s/delivery/web/", structName)
	registryPath := path + fileName
	registryContent := pkg.ReplacePlaceholders(template.RegistryTemplate, values)
	if err := pkg.CreateFile(registryPath, registryContent); err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	fmt.Printf("delivery is created\n")
	return nil
}

func createUsecase(parentFolderName, domainName string) error {
	usecasePaths := []string{
		"/usecase/usecasemock",
		"/usecase/usecaseimpl",
		"/usecase/usecasedecorator",
	}

	for _, v := range usecasePaths {
		path := parentFolderName + v
		if errCreateFolder := os.MkdirAll(path, 0755); errCreateFolder != nil {
			return errCreateFolder
		}
	}

	if errCreateFile := createFileUsecaseName(domainName); errCreateFile != nil {
		return errCreateFile
	}

	if errUsecaseImpl := createFileUsecaseImplName(domainName); errUsecaseImpl != nil {
		return errUsecaseImpl
	}
	fmt.Printf("usecase is created\n")
	return nil
}

func createFileUsecaseName(file string) error {
	fileName := file + ".go"
	values := map[string]string{
		"{{structName}}": pkg.ToPascalCase(file),
		"{{sourceFile}}": file,
	}
	path := fmt.Sprintf("./internal/app/%s/usecase/", file)
	usecasePath := path + fileName
	usecaseContent := pkg.ReplacePlaceholders(template.UsecaseTemplate, values)
	if err := pkg.CreateFile(usecasePath, usecaseContent); err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}

func createFileUsecaseImplName(file string) error {
	fileName := file + ".go"
	values := map[string]string{
		"{{name}}": pkg.ToPascalCase(file),
	}
	usecaImplPath := fmt.Sprintf("./internal/app/%s/usecase/usecaseimpl/", file)
	fullPath := usecaImplPath + fileName
	usecaseImplContent := pkg.ReplacePlaceholders(template.UsecaseImplTemplate, values)
	if err := pkg.CreateFile(fullPath, usecaseImplContent); err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}
	return nil
}
