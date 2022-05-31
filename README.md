# protoc-gen-go_orm

[![CI](https://github.com/sraphs/protoc-gen-go_orm/actions/workflows/ci.yml/badge.svg)](https://github.com/sraphs/protoc-gen-go_orm/actions/workflows/ci.yml)

>  Simply generate ORM model from proto

## Install

```bash
go install github.com/sraphs/protoc-gen-go_orm@latest
```

## Usage

protoc \
    --proto_path ./internal/testdata \
    --proto_path . \
    --go_out=paths=source_relative:. \
    --go_orm_out=paths=source_relative:. \
    ./internal/testdata/basic.proto

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
