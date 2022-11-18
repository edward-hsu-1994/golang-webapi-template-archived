swag init --parseDependency -g ./cmd/main.go
rm -f -r ./assets/swagger/docs
mv ./docs ./assets/swagger
rm -f ./assets/swagger/docs/docs.go
