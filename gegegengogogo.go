package gegegengogogo

import "time"

// GengoCode is a type alias for using Codes of 元号
type GengoCode int8

const (
	code GengoCode = iota
)

// Gengo is an entity of 元号
type Gengo struct {
	Name string
	From time.Time
	To   time.Time
}
