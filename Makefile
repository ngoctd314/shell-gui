build_linux:
	GOOS=linux GOARCH=amd64 go build -o ./run_linux .
build_macos:
	GOOS=darwin GOARCH=amd64 go build -o ./run_mac .
build_win:
	GOOS=windows GOARCH=amd64 go build -o ./run_win .