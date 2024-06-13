package template

const RepositoryTemplate = "package repository \n \n" +
	"//go:generate mockgen --source={{sourceFile}}.go --destination=repositorymock/{{sourceFile}}_mock.go --package=repositorymock\n" +
	"type {{structName}} interface{}\n"
