clean: main.go
	rm -f dist/apricot_linux_amd64
	rm -f dist/apricot_linux_arm64
	rm -f dist/apricot_darwin_amd64
	rm -f dist/apricot_cgo_linux_arm64
	rm -f dist/apricot_cgo_linux_amd64
	rm -f dist/apricot_cgo_darwin_amd64
	rm -f dist/apricot_windows_amd64.exe

build: main.go
	hack/scripts/build.sh

cgo-build: main.go
	@echo "\nwarn: cgo build only support darwin(current) and linux\n"
	CGO_ENABLED=1 go build -v -o dist/apricot_cgo_darwin_amd64 .
	docker run --rm -v $$(pwd):/app golang /app/hack/scripts/cgo_build.sh

prod: main.go
	make clean build cgo-build

test: main.go
	go test ./...

release:
	goreleaser release --rm-dist
	[[ -d dist ]] && rm -r dist
