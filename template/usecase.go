package template

const UsecaseTemplate = "package {{packageName}} \n \n" +
	"//go:generate mockgen --source={{sourceFile}}.go --destination=usecasemock/{{mockFile}}_mock.go --package=usecasemock\n" +
	"type {{interfaceName}} interface{}\n"
