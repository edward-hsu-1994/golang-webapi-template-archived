rm -f -r ./dist
mkdir -p ./dist
sh ./scripts/update-swagger.sh
go mod download

if [ "$(expr substr $(uname -s) 1 10)" == "MINGW32_NT" ] || [ "$(expr substr $(uname -s) 1 10)" == "MINGW64_NT" ]; then
    # Windows
    go build -o ./dist/$(basename "$PWD").exe ./cmd 
else
    go build -o ./dist/$(basename "$PWD") ./cmd 
fi


go build -o ./dist/$(basename "$PWD").exe ./cmd 
cp -r ./assets ./dist
cp -r ./configs ./dist
