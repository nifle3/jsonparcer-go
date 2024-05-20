package jsonparcer_go

type JsonTokenType int

const (
	TOKEN_BRACKET_RIGHT JsonTokenType = iota
	TOKEN_BRACKET_LEFT
	TOKENN_ARRAY_BRACKET_LEFT
	TOKEN_ARRAY_BRACKET_RIGHT
	TOKEN_COLON
	TOKEN_COMMA
	TOKEN_STRING
	TOKEN_NUMBER
	TOKEN_TRUE
	TOKEN_FALSE
	TOKEN_NULL
)

type AST struct {
	thisValue string
	right     *AST
	left      *AST
}

type JsonError struct {
	position int
	message  string
}

type JsonToken struct {
	position int
	token    string
	value    interface{}
}

type Json struct {
	errors []JsonError
}

func (j Json) HasError() bool {
	return len(j.errors) > 0
}

func (j Json) GetErrors() []JsonError {
	copying := make([]JsonError, 0, len(j.errors))
	copy(copying, j.errors)
	return copying
}
