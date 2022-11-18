rm -f -r ./dist
mkdir -p ./dist
sh ./scripts/update-swagger.sh
go build -o ./dist/$(basename "$PWD").exe ./cmd 
cp -r ./assets ./dist
cp -r ./configs ./dist
