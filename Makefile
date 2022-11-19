dev-local:
	go mod edit -replace github.com/grimoh/go-test-server/article/delivery/http=./article/delivery/http
	go mod edit -replace github.com/grimoh/go-test-server/article/delivery/http/middleware=./article/delivery/http/middleware
	go mod edit -replace github.com/grimoh/go-test-server/article/repository/mysql=./article/repository/mysql
	go mod edit -replace github.com/grimoh/go-test-server/article/usecase=./article/usecase
	go mod edit -replace github.com/grimoh/go-test-server/author/repository/mysql=./author/repository/mysql

gen-grpc:
	protoc --go_out=pb ./proto/test.proto
