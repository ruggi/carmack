SHELL=/bin/bash -o pipefail

CURRENT_VERSION_MAJOR = 1
CURRENT_VERSION_MINOR = 0
CURRENT_VERSION_BUG = 0

.PHONY: clean build run install publish publish-major publish-minor publish-bug update-master test

clean:
	rm -rf bin/*

build: clean
	go build -ldflags="-w -s -X github.com/ruggi/carmack/util.Version=`git describe`" -o ./bin/carmack

run: build
	./bin/carmack

install:
	go install

update-master:
	git checkout master
	git pull

publish: 
	@if [ "$(VERSION)" = "" ] ; then echo You should define the version like so: make publish VERSION=x.y.z ; exit 1 ; fi
	@git diff --exit-code --cached || { git status ; echo You have changes that are staged but not committed ; false ; };
	@git diff --exit-code || { git status ; echo You have changes that are not committed ; false ; };
	@git diff --exit-code Makefile || { echo You have made changes to the Makefile that were not committed, please stash or commit them ; false ; };
	$(eval dots = $(subst ., ,$(VERSION)))
	$(eval new_major = $(word 1, $(dots)))
	$(eval new_minor = $(word 2, $(dots)))
	$(eval new_bug = $(word 3, $(dots)))
	sed -i.bak -e 's/^\(var Version = \).*/\1"$(VERSION)"/g' util/version.go
	sed -i.bak -e 's/^\(CURRENT_VERSION_MAJOR = \).*/\1$(new_major)/g' Makefile
	sed -i.bak -e 's/^\(CURRENT_VERSION_MINOR = \).*/\1$(new_minor)/g' Makefile
	sed -i.bak -e 's/^\(CURRENT_VERSION_BUG = \).*/\1$(new_bug)/g' Makefile
	rm Makefile.bak util/version.go.bak

	git commit -am 'Bump version to v$(VERSION)'
	git tag v$(VERSION)
	git push --follow-tags

publish-major: update-master
	@make publish VERSION=$$(($(CURRENT_VERSION_MAJOR) + 1)).0.0
publish-minor: update-master
	@make publish VERSION=$(CURRENT_VERSION_MAJOR).$$(($(CURRENT_VERSION_MINOR) + 1)).0
publish-bug: update-master
	@make publish VERSION=$(CURRENT_VERSION_MAJOR).$(CURRENT_VERSION_MINOR).$$(($(CURRENT_VERSION_BUG) + 1))

test:
	go test -cover ./...
