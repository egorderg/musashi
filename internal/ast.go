package internal

type AstItem interface{}
type AstDatumId int

const (
	datumInt AstDatumId = iota
	datumFloat
	datumString
	datumBool
	datumNil
)

type AstProgram struct {
	forms []AstForm
}

type AstForm struct {
	items []AstItem
}

type AstSymbol struct {
	name string
}

type AstDatum struct {
	id    AstDatumId
	value string
}
