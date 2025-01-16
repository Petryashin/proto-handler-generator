package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func ParseProtoFiles(protoPath string) (*Common, error) {
	absProtoPath, err := filepath.Abs(protoPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	var protoFiles []string
	err = filepath.Walk(absProtoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking the path %s: %w", path, err)
		}
		if filepath.Ext(info.Name()) == ".proto" {
			protoFiles = append(protoFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk directory %s: %w", absProtoPath, err)
	}

	m, e, err := parseProtoMessages(protoFiles[0])

	ps, err := parseProtoService(protoFiles[1])

	return &Common{Messages: m, Enums: e, ProtoService: ps}, err
}

func parseProtoService(protoPath string) (*ProtoService, error) {
	data, err := os.ReadFile(protoPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", protoPath, err)
	}

	var service *ProtoService

	serviceNameRegex := regexp.MustCompile(`service\s+(\w+)\s*\{`)
	serviceRegex := regexp.MustCompile(`rpc\s+\w+\s*\(.*\)\s*returns\s*\(.*\)\s*\{\}`)
	methodRegex := regexp.MustCompile(`rpc\s+(\w+)\((\w+).*\((\w+)\)\s+{`)

	matchName := serviceNameRegex.FindAllSubmatch(data, -1)
	matches := serviceRegex.FindAllSubmatch(data, -1)
	serviceName := string(matchName[0][1])
	service = &ProtoService{Name: serviceName}

	var methods []ProtoMethod
	for _, match := range matches {
		serviceBody := string(match[0])

		methodMatches := methodRegex.FindAllStringSubmatch(serviceBody, -1)
		for _, methodMatch := range methodMatches {
			methods = append(methods, ProtoMethod{
				Name:       methodMatch[1],
				InputType:  methodMatch[2],
				OutputType: methodMatch[3],
			})
		}
	}

	service.Methods = methods

	return service, nil
}

func parseProtoMessages(protoPath string) (Messages, Enums, error) {
	data, err := os.ReadFile(protoPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read file %s: %w", protoPath, err)
	}

	structRegex := regexp.MustCompile(`(enum|message)\s+(\w+)\s+{([\s\S]*?)\n\}\n`)
	messAttrRegex := regexp.MustCompile(`(\w+)\s+(\w+)\s+\=`)
	enumAttrRegex := regexp.MustCompile(`(\w+)\s+\=`)

	matches := structRegex.FindAllSubmatch(data, -1)

	enums := make(Enums)
	messages := make(Messages)

	for _, match := range matches {
		structType := string(match[1])
		structName := string(match[2])
		structBody := match[3]

		if structType == "enum" {
			enumAttributes := enumAttrRegex.FindAllSubmatch(structBody, -1)

			en := Enum{Name: structName}
			for _, attr := range enumAttributes {
				en.Attributes = append(en.Attributes, string(attr[1]))
			}

			enums[structName] = en
		}

		if structType == "message" {
			messageAttributes := messAttrRegex.FindAllSubmatch(structBody, -1)

			m := Message{Name: structName}
			for _, attr := range messageAttributes {
				m.Attributes = append(m.Attributes, Attribute{
					Type: string(attr[1]),
					Name: string(attr[2]),
				})
			}

			messages[structName] = m
		}
	}

	return messages, enums, nil
}
