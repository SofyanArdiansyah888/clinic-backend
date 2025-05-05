package customTypes

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type DateTime time.Time

func (d *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		return nil
	}
	parsed, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	*d = DateTime(parsed)
	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	return []byte(`"` + t.Format("2006-01-02 15:04:05") + `"`), nil
}

// database/sql/driver.Valuer
func (d DateTime) Value() (driver.Value, error) {
	t := time.Time(d)
	return t.Format("2006-01-02 15:04:05"), nil
}

// sql.Scanner
func (d *DateTime) Scan(value interface{}) error {
	if value == nil {
		*d = DateTime(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = DateTime(v)
	case []byte:
		t, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		*d = DateTime(t)
	case string:
		t, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		*d = DateTime(t)
	default:
		return fmt.Errorf("cannot convert %T to DateTime", value)
	}
	return nil
}
