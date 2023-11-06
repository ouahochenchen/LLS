proto:
	set -ue
	cd protocol/grpc/ && mkdir -p go && protoc --proto_path=${GOPATH}/src --proto_path=${GOPATH}/src/github.com/google/protobuf/src  --proto_path=. \
                                                        --go_out=plugins=grpc:./go/ *.proto --govalidators_out=./go/