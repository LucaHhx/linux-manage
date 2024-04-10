# sudo npm install -g protoc-gen-js
# 安装 https://github.com/grpc/grpc-web/releases
# sudo mv ~/Downloads/protoc-gen-grpc-web-1.5.0-darwin-x86_64 \
#     /usr/local/bin/protoc-gen-grpc-web
# chmod +x /usr/local/bin/protoc-gen-grpc-web

# npm install egf-protobuf -g

cd `dirname $0`

DIR=../server

CommonFile=../server/protobuf/proto/*.proto

OUT_DIR=./src

# protoc -I=$DIR $PBFILE --js_out=import_style=commonjs,binary:$OUT_DIR --grpc-web_out=import_style=typescript,mode=grpcweb:$OUT_DIR
# npx protoc --ts_out=$OUT_DIR --proto_path $DIR $CommonFile $LudoFile
protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto --ts_proto_opt=esModuleInterop=true --ts_proto_opt=forceLong=long --ts_proto_out=$OUT_DIR --proto_path $DIR $CommonFile 