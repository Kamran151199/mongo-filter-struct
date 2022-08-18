# MongoDB-Filter-Struct

## Introduction
`mongodb-filter-struct` is a library for creating filtration bson for MongoDB using golang's native struct type.
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
    "github.com/Kamran151199/mongo-filter-struct"
)

type SampleFilter struct {
    Name string `lookup:"name" operator:"$regex"`
    Age int `lookup:"age" operator:"$gt"`
    Adult bool `lookup:"adult" operator:"$eq"`
}

func main() {
    filter := SampleFilter{
        Name: "John",
        Age: 30,
        Adult: true,
    }
    bson, err := filter.BuildQuery(filter)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(bson)
}

// Output:
// {
//     "age": bson.M{
//         "$gt": 30
//     },
//     "adult": bson.M{
//          "$eq": true
//    },
//    "name": bson.M{
//          "$regex": "John"
//    },
// }
```

## License
-------

GNU License Version 3, see [LICENSE](LICENSE)

Copyright (c) 2020 Kamran