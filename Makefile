VERSION=master
BASH_DIR=$(shell pwd)/bash

deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/progrium/go-basher/...

build:
	rm bindata.go || true
	rm -rf $(BASH_DIR) || true
	mkdir $(BASH_DIR)
	wget -O $(BASH_DIR)/$(VERSION).tar.gz https://github.com/wffls/waffles/archive/$(VERSION).tar.gz
	tar xzf $(BASH_DIR)/$(VERSION).tar.gz --strip-components 1 --wildcards -C $(BASH_DIR) */init.sh */resources */functions
	rm $(BASH_DIR)/$(VERSION).tar.gz
	go-bindata bash/...
	go install
