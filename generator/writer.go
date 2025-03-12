package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"unicode"

	"proto-handler-generator/generator/transformers/dto"
	"proto-handler-generator/generator/transformers/handler"
	"proto-handler-generator/generator/transformers/usecase"
	"proto-handler-generator/parser"
)

func writeToFile(
	outputPath string,
	service *parser.Common,
	handler handler.HandlerResultData,
	useCase usecase.UseCaseResultData,
	dto dto.DTOResultData,
) error {
	handlerDir := filepath.Join(outputPath, "handler")
	usecaseDir := filepath.Join(outputPath, "usecase")
	dtoDir := filepath.Join(outputPath, "dto")

	err := os.MkdirAll(handlerDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create handler directory: %w", err)
	}
	err = os.MkdirAll(usecaseDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create usecase directory: %w", err)
	}
	err = os.MkdirAll(dtoDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create dto directory: %w", err)
	}

	err = os.WriteFile(filepath.Join(handlerDir, handler.Name), []byte(handler.GeneratedTemplate), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write handler file: %w", err)
	}

	for _, useCaseCode := range useCase.GeneratedTemplates {
		err = os.WriteFile(filepath.Join(usecaseDir, useCaseCode.Name), []byte(useCaseCode.Template), os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to write usecase file: %w", err)
		}
	}

	err = os.WriteFile(filepath.Join(dtoDir, dto.Name), []byte(dto.GeneratedTemplate), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write dto file: %w", err)
	}

	return nil
}

func camelToSnakeCase(input string) string {
	var result []rune
	for i, r := range input {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}
