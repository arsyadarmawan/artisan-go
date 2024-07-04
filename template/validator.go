package template

//go:generate mockgen -source=creatorvalidator.go -destination=validatormock/creatorvalidator_mock.go -package=validatormock
//go:generate amargo-wrapper new --source=creatorvalidator.go --decorator=logger --output=./validatordecorator/creatorvalidator_decorator_logger_gen.go --package=validatordecorator
//go:generate amargo-wrapper new --source=creatorvalidator.go --decorator=nrtracer --output=./validatordecorator/creatorvalidator_decorator_nrtracer_gen.go --package=validatordecorator

const ValidatorTemplate = "package usecase \n \n" +
	"//go:generate mockgen --source={{structName}}.go --destination=usecasemock/{{sourceFile}}_mock.go --package=validatormock\n" +
	"//go:generate amargo-wrapper new --source={{structName}}.go --decorator=logger --output=./usecasedecorator/{{structName}}_decorator_logger_gen.go --package=validatordecorator\n" +
	"//go:generate amargo-wrapper new --source={{structName}}.go --decorator=nrtracer --output=./usecasedecorator/{{structName}}_decorator_nrtracer_gen.go --package=validatordecorator\n" +
	"type {{structName}} interface{\n" +
	"\t Get()error\n" +
	"}\n"
