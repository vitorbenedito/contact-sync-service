rm -rf bin
mkdir bin
go mod tidy
go build -ldflags '-s -w' -o bin/api ./
cp api/.env bin
./bin/api -v