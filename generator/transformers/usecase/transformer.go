package usecase

import (
	"proto-handler-generator/generator/transformers"
	"proto-handler-generator/parser"
)

func TransformToUseCaseTemplate(common *parser.Common) UseCaseTemplateData {
	var methods []UseCaseMethod

	for _, method := range common.ProtoService.Methods {
		requestMessage := common.Messages[method.InputType]
		responseMessage := common.Messages[method.OutputType]

		var requestFields []UseCaseField
		for _, attr := range requestMessage.Attributes {
			requestFields = append(requestFields, UseCaseField{
				Name: transformers.ToCamelCase(attr.Name),
				Type: attr.Type,
			})
		}

		var responseFields []UseCaseField
		for _, attr := range responseMessage.Attributes {
			responseFields = append(responseFields, UseCaseField{
				Name: transformers.ToCamelCase(attr.Name),
				Type: attr.Type,
			})
		}

		useCaseName := method.Name + "UseCase"
		useCaseDTO := method.Name + "RequestDTO"
		responseDTO := method.Name + "ResponseDTO"

		methods = append(methods, UseCaseMethod{
			MethodName:     method.Name,
			UseCaseName:    useCaseName,
			UseCaseDTO:     useCaseDTO,
			ResponseDTO:    responseDTO,
			RequestFields:  requestFields,
			ResponseFields: responseFields,
		})
	}

	return UseCaseTemplateData{
		Methods: methods,
	}
}
