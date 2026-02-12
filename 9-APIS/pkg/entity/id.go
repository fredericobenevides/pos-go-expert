package entity

import (
	"database/sql/driver"
	"errors"

	"github.com/google/uuid"
)

type ID uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}

func (id ID) String() string {
	return uuid.UUID(id).String()
}

func (id ID) Value() (driver.Value, error) {
	return id.String(), nil
}

func (id *ID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return errors.New("invalid type for ID")
	}

	parsed, err := uuid.Parse(str)
	if err != nil {
		return err
	}
	*id = ID(parsed)
	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + id.String() + `"`), nil
}

func (id *ID) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" {
		return nil
	}
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return errors.New("invalid format for ID")
	}

	parsed, err := ParseID(s[1 : len(s)-1])
	if err != nil {
		return err
	}
	*id = parsed
	return nil
}
