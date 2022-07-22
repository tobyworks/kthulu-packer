#!/bin/sh
JAR="$1"
BIN="$2"

echo "$JAR"
if ! command -v go &> /dev/null
then
    brew install go
fi
go run grab.go
echo $JAR
cp $JAR resources/app.jar

go build -ldflags "-X main.binary=$BIN" pack.go
rm -rf out
mkdir out
mv pack out/$BIN
rm -rf resources