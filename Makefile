NAME=wafflescript
WAFFLESCRIPT_VERSION=master
WAFFLES_VERSION=master
BASH_DIR=$(shell pwd)/bash
ARCH=$(shell uname -m)

deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/progrium/go-basher/...

build:
	rm bindata.go || true
	rm -rf $(BASH_DIR) || true
	mkdir $(BASH_DIR)
	wget -O $(BASH_DIR)/$(WAFFLES_VERSION).tar.gz https://github.com/wffls/waffles/archive/$(WAFFLES_VERSION).tar.gz
	tar xzf $(BASH_DIR)/$(WAFFLES_VERSION).tar.gz --strip-components 1 --wildcards -C $(BASH_DIR) */init.sh */resources */functions
	rm $(BASH_DIR)/$(WAFFLES_VERSION).tar.gz
	go-bindata bash/...
	go install

build_dev:
	rm bindata.go || true
	go-bindata -ignore=\\.swp bash/...
	go install

release:
	rm -rf release && mkdir release
	mkdir -p build/Linux && GOOS=linux go build -o build/Linux/$(NAME)
	tar -zcf release/$(NAME)_$(WAFFLESCRIPT_VERSION)-$(WAFFLES_VERSION)_linux_$(ARCH).tgz -C build/Linux $(NAME)
	gh-release create wffls/$(NAME) $(WAFFLESCRIPT_VERSION)-$(WAFFLES_VERSION) $(shell git rev-parse --abbrev-ref HEAD) $(WAFFLESCRIPT_VERSION)-$(WAFFLES_VERSION)

.PHONY: deps build release
