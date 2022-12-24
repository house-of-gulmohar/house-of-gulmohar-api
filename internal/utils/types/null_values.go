package types

import (
	"database/sql/driver"
	"errors"
)

type NullString string

func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	strVal, ok := value.(string)
	if !ok {
		return errors.New("column is not a string")
	}
	*s = NullString(strVal)
	return nil
}
func (s NullString) Value() (driver.Value, error) {
	if len(s) == 0 { // if nil or empty string
		return nil, nil
	}
	return string(s), nil
}

type NullEnum string

func (s *NullEnum) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	strVal, _ := value.([]byte)
	str := string(strVal)
	*s = NullEnum(str)
	return nil
}
func (s NullEnum) Value() (driver.Value, error) {
	if len(s) == 0 { // if nil or empty string
		return nil, nil
	}
	return string(s), nil
}

type NullInt int64

func (i *NullInt) Scan(value interface{}) error {
	if value == nil {
		*i = 0
		return nil
	}
	intVal, ok := value.(int64)
	if !ok {
		return errors.New("column is not a int")
	}
	*i = NullInt(intVal)
	return nil
}
func (i NullInt) Value() (driver.Value, error) {
	if i == 0 { // if no value
		return nil, nil
	}
	return int64(i), nil
}

type NullFloat float64

func (i *NullFloat) Scan(value interface{}) error {
	if value == nil {
		*i = 0
		return nil
	}
	floatVal, ok := value.(float64)
	if !ok {
		return errors.New("column is not a float")
	}
	*i = NullFloat(floatVal)
	return nil
}
func (i NullFloat) Value() (driver.Value, error) {
	if i == 0 { // if no value
		return nil, nil
	}
	return float64(i), nil
}
