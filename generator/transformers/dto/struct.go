package dto

type DTOField struct {
	Name string
	Type string
}

type DTO struct {
	Name   string
	Fields []DTOField
}

type DTOTemplateData struct {
	DTOs []DTO
}

type DTOResultData struct {
	Name              string
	GeneratedTemplate string
}
