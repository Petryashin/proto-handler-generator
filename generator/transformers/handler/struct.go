package handler

type HandlerTemplateData struct {
	GoPackage    string
	ProtoPackage string
	ServiceName  string
	Methods      []HandlerMethod
}

type HandlerMethod struct {
	MethodName     string
	RequestType    string
	ResponseType   string
	UseCaseName    string
	UseCaseDTO     string
	RequestFields  []Field
	ResponseFields []Field
}

type Field struct {
	Name string
	Type string
}
