OS = darwin linux windows
ARCH = amd64
default:
	git pull
	@echo "Installing dependencies..."
	@go get -u github.com/go-bindata/go-bindata/...
	@echo "Build static files..."
	make asset
	@echo "Build current platform executable..."
	go build -o DailyGakki .

static:
	@echo "Installing dependencies..."
	@go get -u github.com/go-bindata/go-bindata/...
	@echo "Build static files..."
	make asset
	CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o Gakki .

all:
	git pull
	make asset
	@echo "Build all platform executables..."
	@for o in $(OS) ; do            \
    		for a in $(ARCH) ; do     \
    			CGO_ENABLED=0 GOOS=$$o GOARCH=$$a go build -ldflags="-s -w" -o builds/DailyGakki-$$o-$$a .;    \
    		done                              \
    	done


asset:
	@~/go/bin/go-bindata  -o assets.go images/...


dev:
	@~/go/bin/go-bindata  -o assets.go images/default.gif


clean:
	@rm -rf builds
	@rm -f assets.go
	@rm -f Gakki
	@rm -f DailyGakki
