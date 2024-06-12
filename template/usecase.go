package template

const UsecaseTemplate = "package usecase \n \n" +
	"//go:generate mockgen --source={{sourceFile}}.go --destination=usecasemock/{{sourceFile}}_mock.go --package=usecasemock\n" +
	"type {{structName}} interface{}\n"
