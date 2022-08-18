package filterbuilder

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

// Builder is a helper struct for building query strings.
type Builder struct {
	err    error
	output bson.M
}

// GetOutput returns the output of the builder if there are no errors
func (b *Builder) GetOutput() (bson.M, error) {
	if b.err != nil {
		return nil, b.err
	}
	return b.output, nil
}

// AddParam adds a param to the resulting query - Output.
func (b *Builder) AddParam(query *Param) *Builder {
	if err := query.Validate(); err != nil {
		b.err = err
		return b
	}
	b.setParam(query.FieldName, query.Operator, query.Value)
	return b
}

// setParam sets a param to the resulting query.
func (b *Builder) setParam(fieldName string, action string, value interface{}) *Builder {
	if b.output[fieldName] != nil {
		if _, ok := b.output[fieldName].(bson.M); ok {
			b.output[fieldName].(bson.M)[action] = value
		} else {
			b.err = fmt.Errorf("field %s has incorrect type value %T", fieldName, b.output[fieldName])
		}
	} else {
		b.output[fieldName] = bson.M{
			action: value,
		}
	}
	return b
}

// BuildQuery composes a query string from provided args for the specified storage type.
func (b *Builder) BuildQuery(filterStruct interface{}) (bson.M, error) {
	reflectionVal := reflect.ValueOf(filterStruct)
	reflectionType := reflectionVal.Type()
	if reflectionVal.Kind() != reflect.Struct {
		return nil, fmt.Errorf("filter must be a struct")
	}

	for i := 0; i < reflectionVal.NumField(); i++ {
		fieldType := reflectionType.Field(i)
		fieldVal := reflectionVal.Field(i)

		fieldNameTag := fieldType.Tag.Get(lookup)
		operatorTag := fieldType.Tag.Get(operator)

		if operatorTag != Equals && operatorTag != Regex && operatorTag != GreaterThan && operatorTag != LessThan &&
			operatorTag != GreaterThanOrEqualTo && operatorTag != LessThanOrEqualTo && operatorTag != In {
			return nil, fmt.Errorf("invalid/unsupported action %s for field %s", operatorTag, fieldNameTag)
		}

		if fieldVal.IsZero() {
			continue
		}

		b.AddParam(&Param{
			FieldType: fieldVal.Kind(),
			FieldName: fieldNameTag,
			Operator:  operatorTag,
			Value:     fieldVal.Interface(),
		})
	}

	res, err := b.GetOutput()

	if err != nil {
		return nil, err
	}
	return res, nil
}

// NewBuilder returns a new Builder with empty Output.
func NewBuilder() *Builder {
	return &Builder{
		output: bson.M{},
	}
}
