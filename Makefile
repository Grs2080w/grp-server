# generate the docs
doc:
	swag init --output docs --generalInfo cmd/main.go

# run the server
run:
	go run cmd/main.go