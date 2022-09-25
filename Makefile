proto:
	rm -f pb/*.go
	protoc --proto_path=greet/proto --go_out=greet/pb --go_opt=paths=source_relative \
		--go-grpc_out=greet/pb --go-grpc_opt=paths=source_relative \
		greet/proto/*.proto 
calproto:
	rm -f pb/*.go
	protoc --proto_path=calculator/proto --go_out=calculator/pb --go_opt=paths=source_relative \
		--go-grpc_out=calculator/pb --go-grpc_opt=paths=source_relative \
		calculator/proto/*.proto 


.PHONY: calproto proto