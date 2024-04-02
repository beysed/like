package parsers

var parsers = map[string]DataPraser{
	"json": JsonParser{},
	"env":  EnvParser{},
	"yaml": YamlParser{}}

func GetParser(fmt string) DataPraser {
	return parsers[fmt]
}
