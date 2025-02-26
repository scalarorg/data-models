package util

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"

	"gorm.io/gorm/schema"
)

// HexSerializer implements GORM's Serializer interface for encoding/decoding `[]byte`
type HexSerializer struct{}
type JSONSerializer struct{}

// Scan implements serializer interface
func (HexSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	fieldValue := reflect.New(field.FieldType)

	if dbValue != nil {
		if v, ok := dbValue.(string); ok {
			bytes, err := hex.DecodeString(v)
			if err != nil {
				return err
			}
			fieldValue.SetBytes(bytes)
		}
		return fmt.Errorf("invalid data type")
	}

	field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
	return
}

// Scan implements serializer interface
func (JSONSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	fieldValue := reflect.New(field.FieldType)

	if dbValue != nil {
		var bytes []byte
		switch v := dbValue.(type) {
		case []byte:
			bytes = v
		case string:
			bytes = []byte(v)
		default:
			return fmt.Errorf("failed to unmarshal JSONB value: %#v", dbValue)
		}

		err = json.Unmarshal(bytes, fieldValue.Interface())
	}

	field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
	return
}

// Value implements serializer interface
func (HexSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	if v, ok := fieldValue.([]byte); ok {
		return hex.EncodeToString(v), nil
	}
	return nil, fmt.Errorf("invalid data type")
}

// Value implements serializer interface
func (JSONSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return json.Marshal(fieldValue)
}
