// generated by jsonenums -type=GridKind; DO NOT EDIT

package sirpent

import (
	"encoding/json"
	"fmt"
)

var (
	_GridKindNameToValue = map[string]GridKind{
		"hex_grid_hexagonal": hex_grid_hexagonal,
	}

	_GridKindValueToName = map[GridKind]string{
		hex_grid_hexagonal: "hex_grid_hexagonal",
	}
)

func init() {
	var v GridKind
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_GridKindNameToValue = map[string]GridKind{
			interface{}(hex_grid_hexagonal).(fmt.Stringer).String(): hex_grid_hexagonal,
		}
	}
}

// MarshalJSON is generated so GridKind satisfies json.Marshaler.
func (r GridKind) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _GridKindValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid GridKind: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so GridKind satisfies json.Unmarshaler.
func (r *GridKind) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("GridKind should be a string, got %s", data)
	}
	v, ok := _GridKindNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid GridKind %q", s)
	}
	*r = v
	return nil
}
