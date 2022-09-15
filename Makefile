OS = darwin linux windows
ARCH = amd64 arm64


default:
	make asset
	@echo "Build current platform executable..."
	go build -o DailyGakki .

static:
	make asset
	CGO_ENABLED=0 go build -a -ldflags '-s -w -extldflags "-static"' -o DailyGakki .


asset:
	@echo "Installing dependencies and building static files......"
	@go get -u github.com/go-bindata/go-bindata/...
	@go install github.com/go-bindata/go-bindata/...
	@~/go/bin/go-bindata  -o assets.go images/...


dev:
	@~/go/bin/go-bindata  -o assets.go images/default.gif


clean:
	@rm -rf builds
	@rm -f assets.go
	@rm -f Gakki
	@rm -f DailyGakki

all:
	make clean
	make asset
	@echo "Build all platform executables..."
	@for o in $(OS) ; do            \
        		for a in $(ARCH) ; do     \
        		  	echo "Building $$o-$$a..."; \
        		  	if [ "$$o" = "windows" ]; then \
                    	CGO_ENABLED=0 GOOS=$$o GOARCH=$$a go build -ldflags="-s -w" -o builds/DailyGakki-$$o-$$a.exe .;    \
                    else \
        				CGO_ENABLED=0 GOOS=$$o GOARCH=$$a go build -ldflags="-s -w" -o builds/DailyGakki-$$o-$$a .;    \
        			fi; \
        		done   \
        	done

	@make universal
	@make checksum


checksum: builds/*
	@echo "Generating checksums..."
	if [ "$(shell uname)" = "Darwin" ]; then \
		shasum -a 256 $^ >>  builds/checksum-sha256sum.txt ;\
	else \
		sha256sum  $^ >> builds/checksum-sha256sum.txt; \
	fi


universal:
	@echo "Building macOS universal binary..."
	docker run --rm -v $(shell pwd)/builds:/app/ bennythink/lipo-linux -create -output \
		DailyGakki-darwin-universal \
		DailyGakki-darwin-amd64    DailyGakki-darwin-arm64

	file builds/DailyGakki-darwin-universal

release:
	git tag $(git rev-parse --short HEAD)
	git push --tags