package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"proto-handler-generator/parser"
)

func writeToFile(
	outputPath string,
	service *parser.Common,
	handlerCode string,
	useCaseCodes []string,
	dtoCode string,
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

	handlerFileName := fmt.Sprintf("%s_handler.go", snakeCase(service.Name))
	dtoFileName := fmt.Sprintf("%s_dto.go", snakeCase(service.Name))

	err = os.WriteFile(filepath.Join(handlerDir, handlerFileName), []byte(handlerCode), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write handler file: %w", err)
	}

	for i, useCaseCode := range useCaseCodes {
		usecaseFileName := fmt.Sprintf("%s_%d_usecase.go", snakeCase(service.Name), i)
		err = os.WriteFile(filepath.Join(usecaseDir, usecaseFileName), []byte(useCaseCode), os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to write usecase file: %w", err)
		}
	}

	err = os.WriteFile(filepath.Join(dtoDir, dtoFileName), []byte(dtoCode), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write dto file: %w", err)
	}

	return nil
}

func snakeCase(input string) string {
	var result []rune
	for i, r := range input {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return string(result)
}
