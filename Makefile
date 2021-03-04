prod: main.go
	go build -o dist/apricot_darwin_amd64 .
	GOOS=linux go build -o dist/apricot_linux_amd64 .
	GOOS=windows go build -o dist/apricot_windows_amd64.exe .
