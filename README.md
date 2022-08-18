# MongoDB-Filter-Struct
![Test workflow](https://github.com/Kamran151199/mongo-filter-struct/.github/workflows/test.yml/badge.svg)

## Introduction
`mongo-filter-struct` is a library for creating filtration bson for MongoDB using golang's native struct type.
It is inspired by [django-rest-framework-filters](https://github.com/philipn/django-rest-framework-filters)

## Installation

```bash
go get github.com/Kamran151199/mongo-filter-struct
```

## Usage

```go
package main

import (
	"fmt"
	filterbuilder "github.com/Kamran151199/mongo-filter-struct"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type SampleFilter struct {
	Name  string `lookup:"name" operator:"$regex"`
	Age   int    `lookup:"age" operator:"$gt"`
	Adult bool   `lookup:"adult" operator:"$eq"`
}

func main() {
	filter := SampleFilter{
		Name:  "John",
		Age:   30,
		Adult: true,
	}
	builder := filterbuilder.NewBuilder()
	query, err := builder.BuildQuery(filter)
	if err != nil {
		fmt.Println(err)
	}
	expected := bson.M{
		"name":  bson.M{"$regex": "John"},
		"age":   bson.M{"$gt": 30},
		"adult": bson.M{"$eq": true},
	}
	fmt.Printf("Expected is equal to res: %v\n", reflect.DeepEqual(expected, query))
}
```

## License
-------

GNU License Version 3, see [LICENSE](LICENSE)

Copyright (c) 2022 Kamran
