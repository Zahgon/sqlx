package types

import (
	"database/sql/driver"
	"encoding/json"
)

type GzippedText []byte

func (g GzippedText) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

func (g *GzippedText) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }

//lint:ignore ST1005 changing this could break consumers of this package

type JSONText json.RawMessage

var emptyJSON = JSONText("{}")

func (j JSONText) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (j *JSONText) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (j JSONText) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

func (j *JSONText) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }

//lint:ignore ST1005 changing this could break consumers of this package

func (j *JSONText) Unmarshal(v interface{}) error { _ = "STUB: not implemented"; return nil }

func (j JSONText) String() string { _ = "STUB: not implemented"; return "" }

type NullJSONText struct {
	JSONText
	Valid bool
}

func (n *NullJSONText) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

func (n NullJSONText) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

type BitBool bool

func (b BitBool) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

func (b *BitBool) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }
