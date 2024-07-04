package template

const UsecaseTemplate = "package usecase \n \n" +
	"//go:generate mockgen --source={{sourceFile}}.go --destination=usecasemock/{{sourceFile}}_mock.go --package=usecasemock\n" +
	"//go:generate amargo-wrapper new --source={{sourceFile}}.go --decorator=logger --output=./usecasedecorator/{{sourceFile}}_decorator_logger_gen.go --package=usecasedecorator\n" +
	"//go:generate amargo-wrapper new --source={{sourceFile}}.go --decorator=nrtracer --output=./usecasedecorator/{{sourceFile}}_decorator_nrtracer_gen.go --package=usecasedecorator\n" +
	"type {{structName}} interface{\n" +
	"\t Get()error\n" +
	"}\n"
