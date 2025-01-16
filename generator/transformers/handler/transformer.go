package handler

import (
	"proto-handler-generator/generator/transformers"
	"proto-handler-generator/parser"
)

func TransformToHandlerTemplate(common *parser.Common, goPackage, protoPackage string) HandlerTemplateData {
	var methods []HandlerMethod

	for _, method := range common.ProtoService.Methods {
		requestMessage := common.Messages[method.InputType]
		responseMessage := common.Messages[method.OutputType]

		// Преобразуем поля запросов
		var requestFields []Field
		for _, attr := range requestMessage.Attributes {
			requestFields = append(requestFields, Field{
				Name: transformers.ToCamelCase(attr.Name),
				Type: attr.Type,
			})
		}

		// Преобразуем поля ответов
		var responseFields []Field
		for _, attr := range responseMessage.Attributes {
			responseFields = append(responseFields, Field{
				Name: attr.Name,
				Type: attr.Type,
			})
		}

		// Генерация имени UseCase и DTO
		useCaseName := method.Name + "UseCase"
		useCaseDTO := method.Name + "RequestDTO"

		methods = append(methods, HandlerMethod{
			MethodName:     method.Name,
			RequestType:    method.InputType,
			ResponseType:   method.OutputType,
			UseCaseName:    useCaseName,
			UseCaseDTO:     useCaseDTO,
			RequestFields:  requestFields,
			ResponseFields: responseFields,
		})
	}

	return HandlerTemplateData{
		GoPackage:    goPackage,
		ProtoPackage: protoPackage,
		ServiceName:  common.ProtoService.Name,
		Methods:      methods,
	}
}
