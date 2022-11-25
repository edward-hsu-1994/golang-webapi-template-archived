rm -f -r ./assets/swagger/docs
swag init --parseDependency -g ./cmd/main.go -o ./assets/swagger/docs --ot json,yaml
