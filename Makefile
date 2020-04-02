# # Go parameters
PROJECT_NAME = "github.com/alallema/picture_dictionnary.git"
PKG = "$(PROJECT_NAME)"
PKGS = collector core redis vision-client
LIST := $(foreach pkg,$(PKGS),$(PROJECT_NAME)/$(pkg)/...)
PKG_LIST := $(shell env GO111MODULE=on go list $(LIST) )

test: ## Run tests
	go test -race -covermode=atomic -coverprofile=c.temp -coverpkg=$(PROJECT_NAME)/... ${PKG_LIST}
	cat c.temp | grep -v "/cmd/" > c.out;
	@go tool cover -func=c.out | cat;
	rm -f c.out c.temp