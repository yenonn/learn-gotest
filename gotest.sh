test:
go test -v ./... && git commit -am "WIP" || git reset --hard
