package usecase

type UseCaseField struct {
	Name string
	Type string
}

type UseCaseMethod struct {
	MethodName     string
	UseCaseName    string
	UseCaseDTO     string
	ResponseDTO    string
	RequestFields  []UseCaseField
	ResponseFields []UseCaseField
}

type UseCaseTemplateData struct {
	Methods []UseCaseMethod
}
