@protoc --proto_path=./test --go_out=./test data.proto

@protoc --proto_path=./test --gotagger_out=xxx="bson+\"-\"",output_path=./test:./test data.proto
