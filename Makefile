MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1")
VERSION:=${MAIN_VERSION}\#$(shell git log -n 1 --pretty=format:"%h")

LDFLAGS_USER:=-ldflags "-X github.com/lukasjarosch/educonn-master-thesis/user/internal/platform/config.Version=${VERSION}"
LDFLAGS_MAIL:=-ldflags "-X github.com/lukasjarosch/educonn-master-thesis/mail/internal/platform/config.Version=${VERSION}"

default: run

test:
	@go test -v ./..

clean:
	@rm -rf ./coverage.out ./coverage-all.out ./video/cmd/videod/videod ./course/cmd/coursed/coursed ./mail/cmd/maild/maild ./user/cmd/userd/userd

user: clean
	@echo Building USER service...
	@cd user/cmd/userd && CGO_ENABLED=0 go build ${LDFLAGS_USER} -a -installsuffix cgo -o userd main.go

mail: clean
	@echo Building MAIL service...
	@cd mail/cmd/maild && CGO_ENABLED=0 go build ${LDFLAGS_MAIL} -a -installsuffix cgo -o maild main.go

all: user mail

proto: user-proto mail-proto
	@echo "All protobufs regenerated"

user-proto:
	@cd user/proto && protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lukasjarosch/educonn-master-thesis/user/proto  --micro_out=. user.proto

mail-proto:
	@cd mail/proto && protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lukasjarosch/educonn-master-thesis/mail/proto  --micro_out=. mail.proto
