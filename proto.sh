set -e
set -x

protoc \
 --proto_path=$GOPATH/pkg/protobuf/third_party \
 --proto_path=$GOPATH/src \
 --proto_path=$GOPATH/src/github.com/pengzhong2010/go-server-base \
 --plugin=$GOBIN/protoc-gen-gofast \
 --gofast_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,plugins=grpc,paths=source_relative:. \
 app/app1/api/v1/common.proto

protoc \
 --proto_path=$GOPATH/pkg/protobuf/third_party \
 --proto_path=$GOPATH/src \
 --proto_path=$GOPATH/src/github.com/pengzhong2010/go-server-base \
 --plugin=$GOBIN/protoc-gen-gofast \
 --gofast_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,plugins=grpc,paths=source_relative:. \
 app/app1/api/v1/robot.proto

#protoc \
# --proto_path=$GOPATH/pkg/protobuf/third_party \
# --proto_path=$GOPATH/src \
# --proto_path=$GOPATH/src/github.com/pengzhong2010/go-server-base \
# --plugin=$GOBIN/protoc-gen-swagger \
# --swagger_out=explicit_http=true:. \
# app/app1/api/v1/robot.proto

protoc \
 --proto_path=$GOPATH/pkg/protobuf/third_party \
 --proto_path=$GOPATH/src \
 --proto_path=$GOPATH/src/github.com/pengzhong2010/go-server-base \
 --openapiv2_out . \
 --openapiv2_opt logtostderr=true \
 app/app1/api/v1/robot.proto