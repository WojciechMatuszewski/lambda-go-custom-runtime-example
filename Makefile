build-Hello:
	GOOS=linux GOARCH=amd64 go build -o bootstrap bootstrap.go
	cp ./bootstrap $(ARTIFACTS_DIR)
