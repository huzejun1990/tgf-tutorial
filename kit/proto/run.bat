
::protoc --go_out=./../../common/pb --plugin=./proto/protoc-gen-go --proto_path=../../hall/internal/proto ../../hall/internal/proto/hall.proto



:: protoc --go_out=./../common/pb --plugin=./proto/protoc-gen-go.bat --proto_path=../hall/internal/proto ../hall/internal/proto/hall.proto

::   protoc --go_out=./../../common/pb --plugin=./../proto/protoc-gen-go.bat --proto_path=../../hall/internal/proto ../../hall/internal/proto/hall.proto

::protoc --go_out=./../common/pb --plugin=./proto/protoc-gen-go.bat --proto_path=../hall/internal/proto ../hall/internal/proto/hall.proto


protoc --go_out=../../common/pb --plugin=./proto/protoc-gen-go.bat --proto_path=../../hall/internal/proto ../../hall/internal/proto/hall.proto

