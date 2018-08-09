MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1")
VERSION:=${MAIN_VERSION}\-$(shell git log -n 1 --pretty=format:"%h")

LDFLAGS_USER:=-ldflags "-X github.com/lukasjarosch/educonn-platform/user/internal/platform/config.Version=${VERSION}"
LDFLAGS_USER_API:=-ldflags "-X github.com/lukasjarosch/educonn-platform/api/user/internal/platform/config.Version=${VERSION}"
LDFLAGS_MAIL:=-ldflags "-X github.com/lukasjarosch/educonn-platform/mail/internal/platform/config.Version=${VERSION}"
LDFLAGS_VIDEO:=-ldflags "-X github.com/lukasjarosch/educonn-platform/video/internal/platform/config.Version=${VERSION}"
LDFLAGS_VIDEO_API:=-ldflags "-X github.com/lukasjarosch/educonn-platform/api/video/internal/platform/config.Version=${VERSION}"
LDFLAGS_TRANSCODE:=-ldflags "-X github.com/lukasjarosch/educonn-platform/transcode/internal/platform/config.Version=${VERSION}"

default: run

test:
	@go test -v ./..

clean:
	rm -rf ./coverage.out ./coverage-all.out ./video/cmd/videod/videod ./course/cmd/coursed/coursed \
		./mail/cmd/maild/maild ./user/cmd/userd/userd ./transcode/cmd/transcoded/transcoded \
		./api/user/cmd/user-apid/user-apid ./api/video/cmd/video-apid/video-apid


dev: user mail video transcode user-api video-api
docker: user-docker mail-docker video-docker transcode-docker user-api-docker video-api-docker
docker-push-dev: user-docker-push-dev mail-docker-push-dev video-docker-push-dev transcode-docker-push-dev
proto: user-proto mail-proto video-proto transcode-proto user-api-proto video-api-proto

# --------- LESSON ---------
lesson-proto: lesson-video-proto
	@echo protoc LESSON
	@protoc -I. --go_out=plugins=micro:. --micro_out=. lesson/proto/lesson.proto

# --------- VIDEO LESSON ---------
lesson-video-run:
	@echo Starting the VIDEO LESSON service
	@cd lesson-video/cmd/lesson-videod && go run main.go

lesson-video-proto:
	@echo protoc VIDEO LESSON
	@protoc -I. --go_out=plugins=micro:. --micro_out=. lesson-video/proto/lesson-video.proto
	@cd lesson-video/proto && protoc -I. --go_out=plugins=micro:${GOPATH}/src  --micro_out=:${GOPATH}/src lesson-video.proto

# --------- VIDEO API ---------
video-api: clean
	@echo Buildung VIDEO API service ...
	@cd api/video/cmd/video-apid && CGO_ENABLED=0 go build ${LDFLAGS_VIDEO_API} -a -installsuffix cgo -o video-apid main.go

video-api-proto:
	@echo protoc VIDEO API
	@protoc -I. --go_out=plugins=micro:. --micro_out=. api/video/proto/video_api.proto

video-api-docker:
	@echo Building VIDEO API docker image ...
	@cd api/video && docker build -t derwaldemar/educonn-video-api:${VERSION} -t derwaldemar/educonn-video-api:dev .

video-api-run:
	@echo Starting the VIDEO API service
	@cd api/video/cmd/video-apid && go run main.go

# --------- USER API ---------
user-api: clean
	@echo Buildung USER API service ...
	@cd api/user/cmd/user-apid && CGO_ENABLED=0 go build ${LDFLAGS_USER_API} -a -installsuffix cgo -o user-apid main.go

user-api-proto:
	@echo protoc USER API
	@protoc -I. --go_out=plugins=micro:. --micro_out=. api/user/proto/user_api.proto

user-api-docker:
	@echo Building USER API docker image ...
	@cd api/user && docker build -t derwaldemar/educonn-user-api:${VERSION} -t derwaldemar/educonn-user-api:dev .

user-api-run:
	@echo Starting the USER API service
	@cd api/user/cmd/user-apid && go run main.go

# --------- USER ---------
user: clean
	@echo Building USER service...
	@cd user/cmd/userd && CGO_ENABLED=0 go build ${LDFLAGS_USER} -a -installsuffix cgo -o userd main.go

user-proto:
	@echo protoc USER
	@cd user/proto && protoc -I. --go_out=plugins=micro:${GOPATH}/src  --micro_out=:${GOPATH}/src user.proto

user-run:
	@echo Starting the USER service
	@cd user/cmd/userd && go run main.go

user-docker:
	@echo Building USER docker image ...
	@cd user && docker build -t derwaldemar/educonn-user:${VERSION} -t derwaldemar/educonn-user:dev .

user-docker-push-dev:
	@echo Pushing educonn-user:dev image ...
	docker push derwaldemar/educonn-user:dev

# --------- MAIL ---------
mail: clean
	@echo Building MAIL service...
	@cd mail/cmd/maild && CGO_ENABLED=0 go build ${LDFLAGS_MAIL} -a -installsuffix cgo -o maild main.go

mail-proto:
	@echo protoc MAIL
	@cd mail/proto && protoc -I. --go_out=plugins=micro:${GOPATH}/src  --micro_out=:${GOPATH}/src mail.proto

mail-run:
	@echo Starting the MAIL service
	@cd mail/cmd/maild && go run main.go

mail-docker:
	@echo Building MAIL docker image ...
	@cd mail && docker build -t derwaldemar/educonn-mail:${VERSION} -t derwaldemar/educonn-mail:dev .

mail-docker-push-dev:
	@echo Pushing educonn-mail:dev image ...
	docker push derwaldemar/educonn-mail:dev

# --------- VIDEO ---------
video: clean
	@echo Building VIDEO service...
	@cd video/cmd/videod && CGO_ENABLED=0 go build ${LDFLAGS_VIDEO} -a -installsuffix cgo -o videod main.go

video-proto:
	@echo protoc VIDEO
	@cd video/proto && protoc -I. --go_out=plugins=micro:${GOPATH}/src  --micro_out=:${GOPATH}/src video.proto

video-run:
	@echo Starting the VIDEO service
	@cd video/cmd/videod && go run main.go

video-docker:
	@echo Building VIDEO docker image ...
	@cd video && docker build -t derwaldemar/educonn-video:${VERSION} -t derwaldemar/educonn-video:dev .

video-docker-push-dev:
	@echo Pushing educonn-mail:dev image ...
	docker push derwaldemar/educonn-mail:dev

# --------- TRANSCODE ---------
transcode: clean
	@echo Building TRANSCODE service...
	@cd transcode/cmd/transcoded && CGO_ENABLED=0 go build ${LDFLAGS_TRANSCODE} -a -installsuffix cgo -o transcoded main.go

transcode-proto:
	@echo protoc TRANSCODE
	@cd transcode/proto && protoc -I. --go_out=plugins=micro:${GOPATH}/src  --micro_out=:${GOPATH}/src transcode.proto

transcode-run:
	@echo Starting the TRANSCODE service
	@cd transcode/cmd/transcoded && go run main.go

transcode-docker:
	@echo Building TRANSCODE docker image ...
	@cd transcode && docker build -t derwaldemar/educonn-transcode:${VERSION} -t derwaldemar/educonn-transcode:dev .

transcode-docker-push-dev:
	@echo Pushing educonn-transcode:dev image ...
	docker push derwaldemar/educonn-transcode:dev

