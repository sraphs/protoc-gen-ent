package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Run(gen *protogen.Plugin) error {
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}

	}

	return nil
}

func run(gen *protogen.Plugin, f *protogen.File) {
	panic("unimplemented")
}

func shouldGenerate(msg *protogen.Message) bool {
	opts, ok := msg.Desc.Options().(*descriptorpb.MessageOptions)
	if !ok {
		return false
	}
	extension := proto.GetExtension(opts, orm.E_Gen)
	mop, ok := extension.(*orm.Gen)
	return mop, ok
}
