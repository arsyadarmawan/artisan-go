package template

const RegistryTemplate = "package web \n \n" +
	"type {{domainName}}HandlerOptions struct{}\n" +
	"type {{domainName}}HandlerRegistry struct{\n" +
	"\toptions {{domainName}}HandlerOptions" +
	"\n}\n" +
	"func New{{domainName}}Registry(opt {{domainName}}HandlerOptions) *{{domainName}}HandlerRegistry" +
	"{\n" +
	"\t return &{{domainName}}HandlerRegistry{opt}" +
	"\n}"
