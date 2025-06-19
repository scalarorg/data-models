package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// StringArray is a custom type to handle PostgreSQL text[] arrays
type StringArray []string

// Value implements the driver.Valuer interface
func (sa StringArray) Value() (driver.Value, error) {
	if sa == nil {
		return nil, nil
	}
	if len(sa) == 0 {
		return "{}", nil
	}

	// Escape and quote each string, then join with commas
	quoted := make([]string, len(sa))
	for i, s := range sa {
		// Escape quotes and backslashes
		escaped := strings.ReplaceAll(s, "\\", "\\\\")
		escaped = strings.ReplaceAll(escaped, "\"", "\\\"")
		quoted[i] = fmt.Sprintf("\"%s\"", escaped)
	}

	return fmt.Sprintf("{%s}", strings.Join(quoted, ",")), nil
}

// Scan implements the sql.Scanner interface
func (sa *StringArray) Scan(value interface{}) error {
	if value == nil {
		*sa = nil
		return nil
	}

	var str string
	switch v := value.(type) {
	case string:
		str = v
	case []byte:
		str = string(v)
	default:
		return fmt.Errorf("cannot scan %T into StringArray", value)
	}

	// Handle empty array
	if str == "{}" {
		*sa = StringArray{}
		return nil
	}

	// Remove braces and split by comma
	str = strings.Trim(str, "{}")
	if str == "" {
		*sa = StringArray{}
		return nil
	}

	// Parse the array elements
	var result []string
	var current strings.Builder
	inQuotes := false
	escapeNext := false

	for i := 0; i < len(str); i++ {
		char := str[i]

		if escapeNext {
			current.WriteByte(char)
			escapeNext = false
			continue
		}

		if char == '\\' {
			escapeNext = true
			continue
		}

		if char == '"' {
			inQuotes = !inQuotes
			continue
		}

		if char == ',' && !inQuotes {
			result = append(result, strings.TrimSpace(current.String()))
			current.Reset()
			continue
		}

		current.WriteByte(char)
	}

	// Add the last element
	if current.Len() > 0 {
		result = append(result, strings.TrimSpace(current.String()))
	}

	*sa = StringArray(result)
	return nil
}

// MarshalJSON implements json.Marshaler
func (sa StringArray) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string(sa))
}

// UnmarshalJSON implements json.Unmarshaler
func (sa *StringArray) UnmarshalJSON(data []byte) error {
	var arr []string
	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}
	*sa = StringArray(arr)
	return nil
}
