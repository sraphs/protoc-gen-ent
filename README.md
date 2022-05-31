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
syntax = "proto3";

package testdata;

import "orm/opts.proto";

option go_package = "ent/testdata";

message Pet {
  option (orm.opts).gen = true;
  string name = 1 [(orm.field) = { tag: [ "db", "pk" ] }];
}
```

Run `protoc` to generate `.pb.go` and `_orm.pb.go` files.

```bash
protoc \
    --proto_path ./internal/testdata \
    --proto_path . \
    --go_out=paths=source_relative:. \
    --go_orm_out=paths=source_relative:. \
    ./internal/testdata/fields.proto
```

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
