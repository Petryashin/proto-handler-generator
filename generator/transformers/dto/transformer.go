package dto

import (
	"strings"

	"proto-handler-generator/generator/transformers"
	"proto-handler-generator/parser"
)

func TransformToDTOTemplate(common *parser.Common) DTOTemplateData {
	var dtos []DTO

	for _, message := range common.Messages {
		var fields []DTOField

		for _, attr := range message.Attributes {
			fields = append(fields, DTOField{
				Name: transformers.ToCamelCase(attr.Name),
				Type: attr.Type,
			})
		}

		dtoName := strings.Title(message.Name)

		dtos = append(dtos, DTO{
			Name:   dtoName,
			Fields: fields,
		})
	}

	return DTOTemplateData{
		DTOs: dtos,
	}
}
