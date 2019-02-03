package gegegengogogo

import "time"

// Gengo is an entity of 元号
type Gengo struct {
	C    GengoCode `json:"code"`
	Name string    `json:"name"`
	Kana string    `json:"kana"`
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}
