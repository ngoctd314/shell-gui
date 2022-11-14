bar:
	GOOS=linux GOARCH=amd64 go build -o ./shgui_bar ./cmd/bar

tree:
	GOOS=linux GOARCH=amd64 go build -o ./shgui_tree ./cmd/tree

app:
	GOOS=linux GOARCH=amd64 go build -o ./shgui ./cmd/app

all:
	make app && make tree && make bar


.PHONY: bar tree app all