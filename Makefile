GO=go
GOROOT=/usr/local/go

GO_WINDOWS=~/awork/gosrc/bin/go
GOROOT_WINDOWS=~/awork/gosrc

grpc:
	protoc -I grtest/ --go_out=plugins=grpc:grtest/ grtest/grtest.proto
	python -m grpc_tools.protoc -I grtest/ --python_out=. --grpc_python_out=. grtest/grtest.proto

windows: grpc
	GOROOT=${GOROOT_WINDOWS} CGO_ENABLED=1 GOOS=windows GOARCH=386 CC="x86_64-w64-mingw32-gcc" CXX="x86_64-w64-mingw32-g++" ${GO_WINDOWS} build -o grtest_server server.go

windows_sleep:
	GOROOT=${GOROOT_WINDOWS} CGO_ENABLED=1 GOOS=windows GOARCH=386 CC="x86_64-w64-mingw32-gcc" CXX="x86_64-w64-mingw32-g++" ${GO_WINDOWS} build -o sleep_test_windows sleep.go

linux: grpc
	${GO} build -o grtest_server server.go
	${GO} build -o sleep_test_linux sleep.go

linux_client: grpc
	${GO} build -o grpc_client client.go
