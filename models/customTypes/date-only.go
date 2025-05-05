package customTypes

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type DateOnly time.Time

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		return nil
	}
	parsed, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = DateOnly(parsed)
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	return []byte(`"` + t.Format("2006-01-02") + `"`), nil
}

// database/sql/driver.Valuer
func (d DateOnly) Value() (driver.Value, error) {
	t := time.Time(d)
	return t.Format("2006-01-02"), nil
}

// sql.Scanner
func (d *DateOnly) Scan(value interface{}) error {
	if value == nil {
		*d = DateOnly(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = DateOnly(v)
	case []byte:
		t, err := time.Parse("2006-01-02", string(v))
		if err != nil {
			return err
		}
		*d = DateOnly(t)
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		*d = DateOnly(t)
	default:
		return fmt.Errorf("cannot convert %T to DateOnly", value)
	}
	return nil
}
