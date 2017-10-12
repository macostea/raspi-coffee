DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)

.PHONY: all clean get-deps

all: get-deps out/bin/raspi-coffee

out/bin/raspi-coffee:
						go build \
						-tags release \
						-ldflags '-X raspi-coffee/cmd.Version=$(VERSION) -X raspi-coffee/cmd.BuildDate=$(DATE)' \
						-o out/bin/raspi-coffee

clean:
						rm -rf out/
						echo Clean complete

get-deps:
						go get

install:
						cp out/bin/raspi-coffee /usr/local/bin/raspi-coffee
						cp scripts/raspi-coffee-service.sh /etc/init.d/raspi-coffee
						chmod +x /etc/init.d/raspi-coffee
						update-rc.d raspi-coffee defaults
