bar:
	GOOS=linux GOARCH=amd64 go build -o ./bar ./cmd/bar

tree:
	GOOS=linux GOARCH=amd64 go build -o ./tree ./cmd/tree

app:
	GOOS=linux GOARCH=amd64 go build -o ./app ./cmd/app

all:
	make app && make tree && make bar


.PHONY: bar tree app all