package shared

import "encoding/json"

// Struct generico para manejar valores nulos o string
type Nullable[T any] struct {
	Value T
	Valid bool
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Value)
}
