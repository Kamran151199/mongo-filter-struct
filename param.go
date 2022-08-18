package filterbuilder

import (
	"fmt"
	"reflect"
)

// Param is a single field's parameter for the query.
type Param struct {
	FieldType reflect.Kind
	FieldName string
	Operator  string
	Value     interface{}
}

// Validate runs all the validation checks on the query.
func (q *Param) Validate() error {
	if !q.IsValidOperator(q.FieldType, q.Operator) {
		return fmt.Errorf("invalid/unsupported operator %s for field %s", q.Operator, q.FieldName)
	}
	return nil
}

// IsValidOperator returns true if the operator is valid for the field type.
func (q *Param) IsValidOperator(fieldType reflect.Kind, operator string) bool {
	switch fieldType {
	case reflect.String:
		return operator == Regex || operator == Equals || operator == In
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return operator == Equals || operator == GreaterThan || operator == LessThan ||
			operator == GreaterThanOrEqualTo || operator == LessThanOrEqualTo || operator == In
	case reflect.Bool:
		return operator == Equals || operator == In
	case reflect.Slice:
		return operator == In
	default:
		return false
	}
}
