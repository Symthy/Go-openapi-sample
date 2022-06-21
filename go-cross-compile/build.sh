go build cmd/main.go

pushd sub >> /dev/null
go build cmd/main.go
popd >> /dev/null