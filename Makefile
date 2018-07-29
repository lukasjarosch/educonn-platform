MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1")
VERSION:=${MAIN_VERSION}\-$(shell git log -n 1 --pretty=format:"%h")

LDFLAGS_USER:=-ldflags "-X github.com/lukasjarosch/educonn-master-thesis/user/internal/platform/config.Version=${VERSION}"
LDFLAGS_MAIL:=-ldflags "-X github.com/lukasjarosch/educonn-master-thesis/mail/internal/platform/config.Version=${VERSION}"
LDFLAGS_VIDEO:=-ldflags "-X github.com/lukasjarosch/educonn-master-thesis/video/internal/platform/config.Version=${VERSION}"
LDFLAGS_TRANSCODE:=-ldflags "-X github.com/lukasjarosch/educonn-master-thesis/transcode/internal/platform/config.Version=${VERSION}"

default: run

test:
	@go test -v ./..

clean:
	@rm -rf ./coverage.out ./coverage-all.out ./video/cmd/videod/videod ./course/cmd/coursed/coursed ./mail/cmd/maild/maild ./user/cmd/userd/userd ./transcode/cmd/transcoded/transcoded

user: clean
	@echo Building USER service...
	@cd user/cmd/userd && CGO_ENABLED=0 go build ${LDFLAGS_USER} -a -installsuffix cgo -o userd main.go

user-docker:
	@echo Building USER docker image ...
	@cd user && docker build -t derwaldemar/educonn-user:${VERSION} -t derwaldemar/educonn-user:dev .

mail: clean
	@echo Building MAIL service...
	@cd mail/cmd/maild && CGO_ENABLED=0 go build ${LDFLAGS_MAIL} -a -installsuffix cgo -o maild main.go

mail-docker:
	@echo Building MAIL docker image ...
	@cd mail && docker build -t derwaldemar/educonn-mail:${VERSION} -t derwaldemar/educonn-mail:dev .

video: clean
	@echo Building VIDEO service...
	@cd video/cmd/videod && CGO_ENABLED=0 go build ${LDFLAGS_VIDEO} -a -installsuffix cgo -o videod main.go

video-docker:
	@echo Building VIDEO docker image ...
	@cd video && docker build -t derwaldemar/educonn-video:${VERSION} -t derwaldemar/educonn-video:dev .

transcode: clean
	@echo Building TRANSCODE service...
	@cd transcode/cmd/transcoded && CGO_ENABLED=0 go build ${LDFLAGS_TRANSCODE} -a -installsuffix cgo -o transcoded main.go

transcode-docker:
	@echo Building TRANSCODE docker image ...
	@cd transcode && docker build -t derwaldemar/educonn-transcode:${VERSION} -t derwaldemar/educonn-transcode:dev .


all: user mail video transcode
docker: user-docker mail-docker video-docker transcode-docker
	@echo "Building all docker images"
proto: user-proto mail-proto video-proto transcode-proto
	@echo "All protobufs regenerated"

user-proto:
	@echo protoc USER
	@cd user/proto && protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lukasjarosch/educonn-master-thesis/user/proto  --micro_out=. user.proto

mail-proto:
	@echo protoc MAIL
	@cd mail/proto && protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lukasjarosch/educonn-master-thesis/mail/proto  --micro_out=. mail.proto

video-proto:
	@echo protoc VIDEO
	@cd video/proto && protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lukasjarosch/educonn-master-thesis/video/proto --micro_out=. video.proto

transcode-proto:
	@echo protoc TRANSCODE
	@cd transcode/proto && protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/lukasjarosch/educonn-master-thesis/transcode/proto --micro_out=. transcode.proto
