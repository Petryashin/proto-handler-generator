package generator

import (
	"fmt"
	"os"

	"proto-handler-generator/generator/transformers/dto"
	"proto-handler-generator/generator/transformers/handler"
	"proto-handler-generator/generator/transformers/usecase"
	"proto-handler-generator/parser"

	"github.com/joho/godotenv"
)

func GenerateCode(protoPath, outputPath string) error {
	handlerTemplate := "templates/grpc_handler.tmpl"
	usecaseTemplate := "templates/usecase.tmpl"
	dtoTemplate := "templates/dto.tmpl"

	err := godotenv.Load(".env")

	common, err := parser.ParseProtoFiles(protoPath)
	if err != nil {
		return fmt.Errorf("error parsing proto files: %w", err)
	}

	transformer1 := handler.TransformToHandlerTemplate(common,
		os.Getenv("GO_PACKAGE_NAME"),
		os.Getenv("PROTO_PACKAGE_NAME"),
	)

	handlerCode, err := generateCode(handlerTemplate, transformer1)
	if err != nil {
		return fmt.Errorf("error generating handler code: %w", err)
	}
	handlerFileName := fmt.Sprintf("%s_handler.go", camelToSnakeCase(common.Name))

	handlerRes := handler.HandlerResultData{
		Name:              handlerFileName,
		GeneratedTemplate: handlerCode,
	}

	var useCaseRes usecase.UseCaseResultData
	transformer2 := usecase.TransformToUseCaseTemplate(common)

	for _, method := range transformer2.Methods {
		useCaseCode, err := generateCode(usecaseTemplate, method)
		if err != nil {
			return fmt.Errorf("error generating usecase code: %w", err)
		}
		useCaseFileName := fmt.Sprintf("%s_usecase.go", camelToSnakeCase(method.MethodName))

		useCaseRes.GeneratedTemplates = append(
			useCaseRes.GeneratedTemplates, usecase.GeneratedTemplate{
				Name:     useCaseFileName,
				Template: useCaseCode,
			})
	}

	transformer3 := dto.TransformToDTOTemplate(common)

	dtoCode, err := generateCode(dtoTemplate, transformer3)
	if err != nil {
		return fmt.Errorf("error generating DTO code: %w", err)
	}
	dtoFileName := fmt.Sprintf("%s_dto.go", camelToSnakeCase(common.Name))

	dtoRes := dto.DTOResultData{
		Name:              dtoFileName,
		GeneratedTemplate: dtoCode,
	}

	err = writeToFile(outputPath, common, handlerRes, useCaseRes, dtoRes)
	if err != nil {
		return fmt.Errorf("error writing generated files: %w", err)
	}

	return nil
}
