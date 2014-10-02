package loader

import (
	"encoding/json"
	"io"
)

type Basic struct{}

func NewBasic() *Basic {
	return &Basic{}
}

func (b *Basic) Load(r io.Reader) (*Spec, error) {
	s := &Spec{}

	dec := json.NewDecoder(r)
	err := dec.Decode(s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
