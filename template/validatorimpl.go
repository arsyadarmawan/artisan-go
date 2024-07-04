package template

const ValidatorImplTemplate = "package validatorimpl \n \n" +
	"type {{name}}ImplOptions struct{}\n" +
	"type {{name}}Impl struct{\n" +
	"\toptions {{name}}ImplOptions" +
	"\n}\n" +
	"func New{{name}}Registry(opt {{name}}ImplOptions) *{{name}}Impl" +
	"{\n" +
	"\t return &{{name}}Impl{opt}" +
	"\n}"
