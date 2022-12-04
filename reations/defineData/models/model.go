package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type CUser struct {
	ID    uint
	CInfo CInfo
}

type CInfo struct {
	Name string
	Age  int
}

func (c CInfo) Value() (driver.Value, error) {

	str, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return string(str), err

}
func (c *CInfo) Scan(value interface{}) error {
	str, ok := value.([]byte)
	if !ok {
		return errors.New("不匹配的数据类型")
	}
	err := json.Unmarshal(str, c)
	return err
}
