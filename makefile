run:
	go run cmd/api/main.go

wire:
	cd pkg/di && wire

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest 
	go get -u github.com/swaggo/swag/cmd/swag 
	go get -u github.com/swaggo/gin-swagger 
	go get -u github.com/swaggo/files

swag:
	swag init -g pkg/api/server.go -o ./docs
