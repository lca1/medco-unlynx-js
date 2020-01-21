PWD           = $(shell pwd)
CURRENT_USER  = $(shell id -u)
CURRENT_GROUP = $(shell id -g)
JSBUILD		  = docker-compose run -e USER_ID=$(CURRENT_USER) -e GROUP_ID=$(CURRENT_GROUP) build

help:
	echo "available targets: clean, js, publish"

.PHONY: help clean js npm publish

clean:
	cd $(PWD)/npm; rm -f wasm_exec.js medco-unlynx-js.wasm

js: clean
	cd $(PWD)/build; $(JSBUILD)

publish: npm
	cd $(PWD)/npm; npm publish --access public

%: help
