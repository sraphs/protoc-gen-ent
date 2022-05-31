# protoc-gen-go_orm

[![CI](https://github.com/sraphs/protoc-gen-go_orm/actions/workflows/ci.yml/badge.svg)](https://github.com/sraphs/protoc-gen-go_orm/actions/workflows/ci.yml)

>  Simply generate ORM model from proto

## Install

```bash
go install github.com/sraphs/protoc-gen-go_orm@latest
```

## Usage

Define your proto file in `.proto` format.

```proto
// internal/testdata/fields.proto
syntax = "proto3";

package testdata;

import "orm/opts.proto";

option go_package = "ent/testdata";

message Pet {
  option (orm.opts).gen = true;
  string name = 1 [(orm.field) = { tag: [ "db", "id", "pk", "autoincrement" ] }];
}
```

Run `protoc` to generate `.pb.go` and `_orm.pb.go` files.

```bash
protoc \
    --proto_path ./internal/testdata \
    --proto_path . \
    --go_out=paths=source_relative:./out \
    --go_orm_out=paths=source_relative:./out \
    ./internal/testdata/fields.proto
```

This will generate `.pb.go` and `_orm.pb.go` files in `./out` directory.

```go
// out/fields_orm.pb.go

package testdata

import (
	_ " github.com/sraphs/protoc-gen-go_orm/orm"
)

type PetORM struct {
	Name string `db:"id,pk,autoincrement"`
}

```

> fist tag is used as tag key, and other tags are used as tag value.

## Contributing

We alway welcome your contributions :clap:

1.  Fork the repository
2.  Create Feat_xxx branch
3.  Commit your code
4.  Create Pull Request


## CHANGELOG
See [Releases](https://github.com/sraphs/protoc-gen-go_orm/releases)

## License
[MIT © sraph.com](./LICENSE)
