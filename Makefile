SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += -j$(NPROCS)

CURRENT_VERSION_MAJOR = 1
CURRENT_VERSION_MINOR = 19
CURRENT_VERSION_BUG = 0

GIT_DESCRIBE := $(shell git describe)
GITHUB_HEAD_REF ?= $(shell git branch --show-current)

update-main:
	git checkout main
	git pull

publish: $(DEPS)
	@if [ "$(VERSION)" = "" ]; then echo You should define the version like so: make publish VERSION=x.y.z; exit 1; fi
	@git diff --exit-code --cached || { git status; echo You have changes that are staged but not committed ; false ; };
	@git diff --exit-code || { git status; echo You have changes that are not committed ; false ; };
	@git diff --exit-code Makefile || { echo You have made changes to the Makefile that were not committed, please stash or commit them; false ; };
	$(eval dots = $(subst ., ,$(VERSION)))
	$(eval new_major = $(word 1, $(dots)))
	$(eval new_minor = $(word 2, $(dots)))
	$(eval new_bug = $(word 3, $(dots)))
	sed -i.bak -e 's/^\(CURRENT_VERSION_MAJOR = \).*/\1$(new_major)/g' Makefile
	sed -i.bak -e 's/^\(CURRENT_VERSION_MINOR = \).*/\1$(new_minor)/g' Makefile
	sed -i.bak -e 's/^\(CURRENT_VERSION_BUG = \).*/\1$(new_bug)/g' Makefile
	rm Makefile.bak

	git checkout -b 'release/v$(VERSION)'
	git commit -am 'Bump version to v$(VERSION)'
	git push --set-upstream origin 'release/v$(VERSION)'

publish-major: update-main
	make publish VERSION=$$(($(CURRENT_VERSION_MAJOR) + 1)).0.0
publish-minor: update-main
	make publish VERSION=$(CURRENT_VERSION_MAJOR).$$(($(CURRENT_VERSION_MINOR) + 1)).0
publish-patch publish-bug: update-main
	make publish VERSION=$(CURRENT_VERSION_MAJOR).$(CURRENT_VERSION_MINOR).$$(($(CURRENT_VERSION_BUG) + 1))

publish-staging:
	@if [ "$(SUFFIX)" = "" ]; then echo You should define the version like so: make publish SUFFIX=test-ratelimit; exit 1; fi
	@git diff --exit-code --cached || { git status; echo You have changes that are staged but not committed ; false ; };
	@git diff --exit-code || { git status; echo You have changes that are not committed ; false ; };
	@git diff --exit-code Makefile || { echo You have made changes to the Makefile that were not committed, please stash or commit them; false ; };
	$(eval VERSION := $(CURRENT_VERSION_MAJOR).$(CURRENT_VERSION_MINOR).$(CURRENT_VERSION_BUG)-$(SUFFIX))
	git tag -a -m "RELEASED BY: $$(git config user.name)" v$(VERSION)
	git push --follow-tags

release-changelog:
	$(eval VERSION := $(CURRENT_VERSION_MAJOR).$(CURRENT_VERSION_MINOR).$(CURRENT_VERSION_BUG))
	git log v$(VERSION)..HEAD  --oneline --shortstat

.PHONY: build-docker-image
build-docker-image:
	@docker buildx build \
			-f Dockerfile \
			--platform linux/amd64 --load \
			-t ghcr.io/getstream/protobuf-generate:latest \
			.

.PHONY: push-docker-image
push-docker-image:
	@docker buildx build \
			-f Dockerfile \
			--platform linux/arm64,linux/amd64 \
			-t ghcr.io/getstream/protobuf-generate:latest \
			. --push
