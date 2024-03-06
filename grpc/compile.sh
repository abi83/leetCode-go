protoc --go-grpc_out=. --go_out=. protofile.proto
grpc_tools_node_protoc --js_out=import_style=commonjs,binary:./nodejs-client --grpc_out=grpc_js:./nodejs-client protofile.proto