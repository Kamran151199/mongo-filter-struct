package filterbuilder

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestParam_IsValidOperator(t *testing.T) {
	type args struct {
		fieldType reflect.Kind
		operator  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string can use `equals` operator",
			args: args{
				fieldType: reflect.String,
				operator:  "$eq",
			},
			want: true,
		},
		{
			name: "string can use `in` operator",
			args: args{
				fieldType: reflect.String,
				operator:  "$in",
			},
			want: true,
		},
		{
			name: "int can use `equals` operator",
			args: args{
				fieldType: reflect.Int,
				operator:  "$eq",
			},
			want: true,
		},
		{
			name: "int can use `greater than` operator",
			args: args{
				fieldType: reflect.Int,
				operator:  "$gt",
			},
			want: true,
		},
		{
			name: "int can use `less than` operator",
			args: args{
				fieldType: reflect.Int,
				operator:  "$lt",
			},
			want: true,
		},
		{
			name: "int can use `greater than or equal to` operator",
			args: args{
				fieldType: reflect.Int,
				operator:  "$gte",
			},
			want: true,
		},
		{
			name: "int can use `less than or equal to` operator",
			args: args{
				fieldType: reflect.Int,
				operator:  "$lte",
			},
			want: true,
		},
		{
			name: "int can use `in` operator",
			args: args{
				fieldType: reflect.Int,
				operator:  "$in",
			},
			want: true,
		},
		{
			name: "bool can use `equals` operator",
			args: args{
				fieldType: reflect.Bool,
				operator:  "$eq",
			},
			want: true,
		},
		{
			name: "bool can use `in` operator",
			args: args{
				fieldType: reflect.Bool,
				operator:  "$in",
			},
			want: true,
		},
		{
			name: "string can use `regex` operator",
			args: args{
				fieldType: reflect.String,
				operator:  "$regex",
			},
			want: true,
		},
		{
			name: "int can not use `regex` operator",
			args: args{
				fieldType: reflect.Int,
				operator:  "$regex",
			},
			want: false,
		},
		{
			name: "bool can not use `regex` operator",
			args: args{
				fieldType: reflect.Bool,
				operator:  "$regex",
			},
			want: false,
		},
		{
			name: "bool can not use `greater` operator",
			args: args{
				fieldType: reflect.Bool,
				operator:  "$gt",
			},
			want: false,
		},
		{
			name: "bool can not use `less` operator",
			args: args{
				fieldType: reflect.Bool,
				operator:  "$lt",
			},
			want: false,
		},
		{
			name: "bool can not use `greater or equal to` operator",
			args: args{
				fieldType: reflect.Bool,
				operator:  "$gte",
			},
			want: false,
		},
		{
			name: "bool can not use `less or equal to` operator",
			args: args{
				fieldType: reflect.Bool,
				operator:  "$lte",
			},
			want: false,
		},
		{
			name: "string can not use greater operator",
			args: args{
				fieldType: reflect.String,
				operator:  "$gt",
			},
			want: false,
		},
		{
			name: "string can not use less operator",
			args: args{
				fieldType: reflect.String,
				operator:  "$lt",
			},
			want: false,
		},
		{

			name: "string can not use greater or equal to operator",
			args: args{
				fieldType: reflect.String,
				operator:  "$gte",
			},
			want: false,
		},
		{
			name: "string can not use less or equal to operator",
			args: args{
				fieldType: reflect.String,
				operator:  "$lte",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			param := Param{
				FieldType: reflect.String,
				FieldName: "field",
				Operator:  tt.args.operator,
				Value:     "value",
			}
			isValid := param.IsValidOperator(tt.args.fieldType, tt.args.operator)
			require.Equal(t, tt.want, isValid)
		})
	}
}

func TestParam_IsValidValue(t *testing.T) {
	type args struct {
		fieldType reflect.Kind
		operator  string
		value     interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string can use `equals` operator",
			args: args{
				fieldType: reflect.String,
				operator:  "$eq",
				value:     "value",
			},
			want: true,
		},
		{
			name: "string can use `in` operator",
			args: args{
				fieldType: reflect.String,
				operator:  "$in",
				value:     []string{"value"},
			},
			want: true,
		},
		{
			name: "int can use `equals` operator",
			args: args{
				fieldType: reflect.Int,
				operator:  "$eq",
				value:     1,
			},
			want: true,
		},
		{
			name: "int can use `greater than` operator",
			args: args{
				fieldType: reflect.Int,
				operator:  "$gt",
				value:     1,
			},
			want: true,
		},
		{
			name: "int can use `less than` operator",
			args: args{
				fieldType: reflect.Int,
				operator:  "$lt",
				value:     1,
			},
			want: true,
		},
		{
			name: "int can use `greater than or equal to` operator",
			args: args{
				fieldType: reflect.Int,
				operator:  "$gte",
				value:     1,
			},
			want: true,
		},
		{
			name: "int can use `less than or equal to` operator",
			args: args{
				fieldType: reflect.Int,
				operator:  "$lte",
				value:     1,
			},
			want: true,
		},
		{
			name: "bool can use `equals` operator",
			args: args{
				fieldType: reflect.Bool,
				operator:  "$eq",
				value:     true,
			},
			want: true,
		},
		{

			name: "bool can use `in` operator",
			args: args{

				fieldType: reflect.Bool,
				operator:  "$in",
				value:     []bool{true},
			},
			want: true,
		},
		{
			name: "bool can not use `greater` operator",
			args: args{
				fieldType: reflect.Bool,
				operator:  "$gt",
				value:     true,
			},
			want: false,
		},
		{
			name: "bool can not use `less` operator",
			args: args{
				fieldType: reflect.Bool,
				operator:  "$lt",
				value:     true,
			},
			want: false,
		},
		{
			name: "bool can not use `greater or equal to` operator",
			args: args{
				fieldType: reflect.Bool,
				operator:  "$gte",
				value:     true,
			},
			want: false,
		},
		{
			name: "bool can not use `less or equal to` operator",
			args: args{
				fieldType: reflect.Bool,
				operator:  "$lte",
				value:     true,
			},
			want: false,
		},
		{
			name: "string can not use `greater` operator",
			args: args{
				fieldType: reflect.String,
				operator:  "$gt",
				value:     "value",
			},
			want: false,
		},
		{
			name: "string can not use `less` operator",
			args: args{
				fieldType: reflect.String,
				operator:  "$lt",
				value:     "value",
			},
			want: false,
		},
		{
			name: "string can not use `greater or equal to` operator",
			args: args{
				fieldType: reflect.String,
				operator:  "$gte",
				value:     "value",
			},
			want: false,
		},
		{
			name: "string can not use `less or equal to` operator",
			args: args{
				fieldType: reflect.String,
				operator:  "$lte",
				value:     "value",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			param := &Param{
				FieldName: "field",
				Operator:  tt.args.operator,
				FieldType: tt.args.fieldType,
				Value:     tt.args.value,
			}
			err := param.Validate()
			if !tt.want {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
