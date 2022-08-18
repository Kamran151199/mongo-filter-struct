package filterbuilder

import (
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type TestFilterWithoutNamingCollision struct {
	Field1 int    `json:"field-1" default:"0" validate:"min=0" lookup:"field_1" operator:"$gte"`
	Field2 int    `json:"field-2" default:"10" validate:"min=1" lookup:"field_2" operator:"$lte"`
	Field3 string `json:"field-3" lookup:"field_3" operator:"$regex"`
	Field4 []int  `json:"field-4" lookup:"field_4" operator:"$in"`
	Field5 int    `json:"field-5" lookup:"field_5" operator:"$eq"`
	Field6 int    `json:"field-6" lookup:"field_6" operator:"$gt"`
	Field7 int    `json:"field-7" lookup:"field_7" operator:"$lt"`
}

type TestFilterWithNamingCollision struct {
	Field1 int    `json:"field-1" default:"0" validate:"min=0" lookup:"field_1" operator:"$gte"`
	Field2 int    `json:"field-2" default:"10" validate:"min=1" lookup:"field_1" operator:"$lte"`
	Field3 string `json:"field-3" lookup:"field_3" operator:"$regex"`
	Field4 []int  `json:"field-4" lookup:"field_4" operator:"$in"`
	Field5 int    `json:"field-5" lookup:"field_5" operator:"$eq"`
	Field6 int    `json:"field-6" lookup:"field_6" operator:"$gt"`
	Field7 int    `json:"field-7" lookup:"field_7" operator:"$lt"`
}

func TestBuilder_BuildQuery(t *testing.T) {
	tests := []struct {
		name string
		f    interface{}
		want bson.M
	}{
		{
			name: "all fields without lookup collision and nulls",
			f: TestFilterWithoutNamingCollision{
				Field1: 1,
				Field2: 2,
				Field3: "3",
				Field4: []int{4, 5},
				Field5: 6,
				Field6: 7,
				Field7: 8,
			},
			want: bson.M{
				"field_1": bson.M{
					"$gte": 1,
				},
				"field_2": bson.M{
					"$lte": 2,
				},
				"field_3": bson.M{
					"$regex": "3",
				},
				"field_4": bson.M{
					"$in": []int{4, 5},
				},
				"field_5": bson.M{
					"$eq": 6,
				},
				"field_6": bson.M{
					"$gt": 7,
				},
				"field_7": bson.M{
					"$lt": 8,
				},
			},
		},
		{
			name: "all fields without lookup collision and with nulls",
			f: TestFilterWithoutNamingCollision{
				Field1: 1,
				Field2: 2,
				Field3: "3",
				Field4: []int{4, 5},
				Field6: 7,
				Field7: 8,
			},
			want: bson.M{
				"field_1": bson.M{
					"$gte": 1,
				},
				"field_2": bson.M{
					"$lte": 2,
				},
				"field_3": bson.M{
					"$regex": "3",
				},
				"field_4": bson.M{
					"$in": []int{4, 5},
				},
				"field_6": bson.M{
					"$gt": 7,
				},
				"field_7": bson.M{
					"$lt": 8,
				},
			},
		},
		{
			name: "all fields with lookup collision and without nulls",
			f: TestFilterWithNamingCollision{
				Field1: 1,
				Field2: 2,
				Field3: "3",
				Field4: []int{4, 5},
				Field5: 6,
				Field6: 7,
				Field7: 8,
			},
			want: bson.M{
				"field_1": bson.M{
					"$gte": 1,
					"$lte": 2,
				},
				"field_3": bson.M{
					"$regex": "3",
				},
				"field_4": bson.M{
					"$in": []int{4, 5},
				},
				"field_5": bson.M{
					"$eq": 6,
				},
				"field_6": bson.M{
					"$gt": 7,
				},
				"field_7": bson.M{
					"$lt": 8,
				},
			},
		},
		{
			name: "all fields with lookup collision and with nulls",
			f: TestFilterWithNamingCollision{
				Field1: 1,
				Field2: 2,
				Field3: "3",
				Field4: []int{4, 5},
				Field6: 7,
				Field7: 8,
			},
			want: bson.M{
				"field_1": bson.M{
					"$gte": 1,
					"$lte": 2,
				},
				"field_3": bson.M{
					"$regex": "3",
				},
				"field_4": bson.M{
					"$in": []int{4, 5},
				},
				"field_6": bson.M{
					"$gt": 7,
				},
				"field_7": bson.M{
					"$lt": 8,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder()
			got, err := b.BuildQuery(tt.f)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
