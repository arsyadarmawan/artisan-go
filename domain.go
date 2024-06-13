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
			fmt.Printf("Domain created \n")
			return nil
		},
	}
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
	parentfolder := fmt.Sprintf(template.DomainName, dirName)
	if errCreateDelivery := createDelivery(parentfolder, dirName); errCreateDelivery != nil {
		return errCreateDelivery
	}
	if errCreateUsecase := createUsecase(parentfolder, dirName); errCreateUsecase != nil {
		return errCreateUsecase
	}

	if errCreateRepository := createRepositoryFolder(parentfolder, dirName); errCreateRepository != nil {
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
	path := fmt.Sprintf("./internal/app/%s/repository/", folderName)
	fullPath := path + fileName
	result := pkg.ReplacePlaceholders(template.RepositoryTemplate, values)
	err := WriteToFile(fullPath, result)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return nil
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
	fullPath := path + fileName
	result := pkg.ReplacePlaceholders(template.RegistryTemplate, values)
	err := WriteToFile(fullPath, result)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return nil
	}
	fmt.Printf("delivery is created\n")
	return nil
}

func createUsecase(parentFolderName, folderName string) error {
	usecaseMockPath := parentFolderName + "/usecase/usecasemock"
	usecaseImplPath := parentFolderName + "/usecase/usecaseimpl"
	if errUsecaseImpl := os.MkdirAll(usecaseImplPath, 0755); errUsecaseImpl != nil {
		return errUsecaseImpl
	}
	if errusecaseMock := os.MkdirAll(usecaseMockPath, 0755); errusecaseMock != nil {
		return errusecaseMock
	}

	if errCreateFile := createFileUsecaseName(folderName); errCreateFile != nil {
		return errCreateFile
	}
	fmt.Printf("usecase is created\n")
	return nil
}

func createFileUsecaseName(folderName string) error {
	fileName := folderName + ".go"
	values := map[string]string{
		"{{structName}}": pkg.ToPascalCase(folderName),
		"{{sourceFile}}": folderName,
	}
	path := fmt.Sprintf("./internal/app/%s/usecase/", folderName)
	fullPath := path + fileName
	result := pkg.ReplacePlaceholders(template.UsecaseTemplate, values)
	err := WriteToFile(fullPath, result)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return nil
	}
	return nil
}
