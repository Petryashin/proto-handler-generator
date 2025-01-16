package parser

type ProtoMethod struct {
	Name       string
	InputType  string
	OutputType string
}

type ProtoService struct {
	Name    string
	Methods []ProtoMethod
}

type Attribute struct {
	Name string
	Type string
}

type Message struct {
	Name       string
	Attributes []Attribute
}

type Messages map[string]Message

type Enum struct {
	Name       string
	Attributes []string
}

type Enums map[string]Enum

type Common struct {
	*ProtoService
	Messages
	Enums
}
