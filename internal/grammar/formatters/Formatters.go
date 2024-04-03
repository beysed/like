package formatters

var formatters = map[string]DataFormatter{
	"json": JsonFormatter{},
	"env":  EnvFormatter{},
	"yaml": YamlFormatter{}}

func GetFormatter(fmt string) DataFormatter {
	return formatters[fmt]
}
