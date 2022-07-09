.PHONY: build
build:
	@scripts/build.sh

install: build
	@scripts/install.sh

alterXdg: install
	@scripts/alterXdg.sh
