proto:
	rm -f pb/*.go
	protoc --proto_path=calculator/proto --go_out=calculator/pb --go_opt=paths=source_relative \
		--go-grpc_out=calculator/pb --go-grpc_opt=paths=source_relative \
		calculator/proto/*.proto 
