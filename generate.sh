protoc -I src/ --go_out=src/ src/simple/simple.proto
protoc -I src/ --go_out=src/ src/enum_example/enum_example.proto
protoc -I src/ --go_out=src/ src/complex/complex.proto