MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1")
VERSION:=${MAIN_VERSION}\-$(shell git log -n 1 --pretty=format:"%h")

LDFLAGS_USER:=-ldflags "-X github.com/lukasjarosch/educonn-platform/user/internal/platform/config.Version=${VERSION}"
LDFLAGS_USER_API:=-ldflags "-X github.com/lukasjarosch/educonn-platform/api/user/internal/platform/config.Version=${VERSION}"
LDFLAGS_MAIL:=-ldflags "-X github.com/lukasjarosch/educonn-platform/mail/internal/platform/config.Version=${VERSION}"
LDFLAGS_VIDEO:=-ldflags "-X github.com/lukasjarosch/educonn-platform/video/internal/platform/config.Version=${VERSION}"
LDFLAGS_VIDEO_API:=-ldflags "-X github.com/lukasjarosch/educonn-platform/api/video/internal/platform/config.Version=${VERSION}"
LDFLAGS_TRANSCODE:=-ldflags "-X github.com/lukasjarosch/educonn-platform/transcode/internal/platform/config.Version=${VERSION}"
LDFLAGS_LESSION:=-ldflags "-X github.com/lukasjarosch/educonn-platform/transcode/lesson/platform/config.Version=${VERSION}"

default: run

test:
	@go test -v ./..

clean:
	rm -rf ./coverage.out ./coverage-all.out ./video/cmd/videod/videod ./course/cmd/coursed/coursed \
		./mail/cmd/maild/maild ./user/cmd/userd/userd ./transcode/cmd/transcoded/transcoded \
		./api/user/cmd/user-apid/user-apid ./api/video/cmd/video-apid/video-apid \
		./lesson/cmd/lessond/lessond


dev: user mail video transcode user-api video-api
docker: user-docker mail-docker video-docker transcode-docker user-api-docker video-api-docker
docker-push-dev: user-docker-push-dev mail-docker-push-dev video-docker-push-dev transcode-docker-push-dev
proto: user-proto mail-proto video-proto transcode-proto user-api-proto video-api-proto

# --------- LESSON ---------
lesson:
	@sh -c "'$(CURDIR)'/lesson/scripts/build.sh" ${VERSION}

lesson-proto:
	@sh -c "'$(CURDIR)'/lesson/scripts/proto.sh"

lesson-run:
	@echo Starting the USER service
	docker-compose -f docker-compose.local.yml up lesson

lesson-docker:
	@echo "==> Building USER docker image..."
	docker build -t derwaldemar/educonn-lesson:${VERSION} -t derwaldemar/educonn-lesson:latest-staging -f lesson/build/Dockerfile .

lesson-publish:
	@echo "==> Publishing latest image version"
	docker push derwaldemar/educonn-lesson:latest-staging
	docker push derwaldemar/educonn-lesson:${VERSION}

lesson-deploy:
	@echo "==> Deploying image version: ${VERSION}"
	@echo "==> Deploying image version: latest-staging"

# --------- VIDEO API ---------
video-api:
	@sh -c "'$(CURDIR)'/video-api/scripts/build.sh" ${VERSION}

video-api-proto:
	@sh -c "'$(CURDIR)'/video-api/scripts/proto.sh"

video-api-run:
	@echo Starting the VIDEO-API service
	@cd user-api/cmd/video-api && go run main.go

video-api-docker:
	@echo "==> Building VIDEO-API docker image..."
	docker build -t derwaldemar/educonn-video-api:${VERSION} -t derwaldemar/educonn-video-api:latest-staging -f video-api/build/Dockerfile .

video-api-publish:
	@echo "==> Publishing latest image version"
	docker push derwaldemar/educonn-video-api:latest-staging
	docker push derwaldemar/educonn-video-api:${VERSION}

video-api-deploy:
	@echo "==> Deploying image version: ${VERSION}"
	@echo "==> Deploying image version: latest-staging"

# --------- USER API ---------
user-api:
	@sh -c "'$(CURDIR)'/user-api/scripts/build.sh" ${VERSION}

user-api-proto:
	@sh -c "'$(CURDIR)'/user-api/scripts/proto.sh"

user-api-run:
	@echo Starting the USER-API service
	@cd user-api/cmd/user-api && go run main.go plugins.go

user-api-docker:
	@echo "==> Building USER-API docker image..."
	docker build -t derwaldemar/educonn-user-api:${VERSION} -t derwaldemar/educonn-user-api:latest-staging -f user-api/build/Dockerfile .

user-api-publish:
	@echo "==> Publishing latest image version"
	docker push derwaldemar/educonn-user-api:latest-staging
	docker push derwaldemar/educonn-user-api:${VERSION}

user-api-deploy:
	@echo "==> Deploying image version: ${VERSION}"
	@echo "==> Deploying image version: latest-staging"

# --------- USER ---------
user:
	@sh -c "'$(CURDIR)'/user/scripts/build.sh" ${VERSION}

user-proto:
	@sh -c "'$(CURDIR)'/user/scripts/proto.sh"

user-run:
	@echo Starting the USER service
	#@cd user/cmd/user && go run main.go
	docker-compose -f docker-compose.local.yml up user

user-docker:
	@echo "==> Building USER docker image..."
	docker build -t derwaldemar/educonn-user:${VERSION} -t derwaldemar/educonn-user:latest-staging -f user/build/Dockerfile .

user-publish:
	@echo "==> Publishing latest image version"
	docker push derwaldemar/educonn-user:latest-staging
	docker push derwaldemar/educonn-user:${VERSION}

user-deploy:
	@echo "==> Deploying image version: ${VERSION}"
	@echo "==> Deploying image version: latest-staging"

# --------- MAIL ---------
mail:
	@sh -c "'$(CURDIR)'/mail/scripts/build.sh" ${VERSION}

mail-proto:
	@sh -c "'$(CURDIR)'/mail/scripts/proto.sh"

mail-run:
	@echo Starting the MAIL service
	@cd mail/cmd/maild && go run main.go

mail-docker:
	@echo "==> Building MAIL docker image..."
	docker build -t derwaldemar/educonn-mail:${VERSION} -t derwaldemar/educonn-mail:latest-staging -f mail/build/Dockerfile .

mail-publish:
	@echo "==> Publishing latest image version"
	docker push derwaldemar/educonn-mail:latest-staging
	docker push derwaldemar/educonn-mail:${VERSION}

mail-deploy:
	@echo "==> Deploying image version: ${VERSION}"
	@echo "==> Deploying image version: latest-staging"

# --------- VIDEO ---------
video:
	@sh -c "'$(CURDIR)'/video/scripts/build.sh" ${VERSION}

video-proto:
	@sh -c "'$(CURDIR)'/video/scripts/proto.sh"

video-run:
	@echo Starting the VIDEO service
	@cd video/cmd/video && go run main.go

video-docker:
	@echo "==> Building VIDEO docker image..."
	docker build -t derwaldemar/educonn-video:${VERSION} -t derwaldemar/educonn-video:latest-staging -f video/build/Dockerfile .

video-publish:
	@echo "==> Publishing latest image version"
	docker push derwaldemar/educonn-video:latest-staging
	docker push derwaldemar/educonn-video:${VERSION}

video-deploy:
	@echo "==> Deploying image version: ${VERSION}"
	@echo "==> Deploying image version: latest-staging"

# --------- TRANSCODE ---------
transcode:
	@sh -c "'$(CURDIR)'/transcode/scripts/build.sh" ${VERSION}

transcode-proto:
	@sh -c "'$(CURDIR)'/transcode/scripts/proto.sh"

transcode-run:
	@echo Starting the TRANSCODE service
	cd transcode/cmd/transcode && go run main.go

transcode-docker:
	@echo "==> Building TRANSCODE docker image..."
	docker build -t derwaldemar/educonn-transcode:${VERSION} -t derwaldemar/educonn-transcode:latest-staging -f transcode/build/Dockerfile .

transcode-publish:
	@echo "==> Publishing latest image version"
	docker push derwaldemar/educonn-transcode:latest-staging
	docker push derwaldemar/educonn-transcode:${VERSION}

transcode-deploy:
	@echo "==> Deploying image version: ${VERSION}"
	@echo "==> Deploying image version: latest-staging"


# --------- LESSON-API ---------
lesson-api:
	@sh -c "'$(CURDIR)'/lesson-api/scripts/build.sh" ${VERSION}

lesson-api-proto:
	@sh -c "'$(CURDIR)'/lesson-api/scripts/proto.sh"

lesson-api-run:
	@echo Starting the LESSON-API service
	cd lesson-api/cmd/lesson-api && go run main.go

lesson-api-docker:
	@echo "==> Building lesson-api docker image..."
	docker build -t derwaldemar/educonn-lesson-api:${VERSION} -t derwaldemar/educonn-lesson-api:latest-staging -f lesson-api/build/Dockerfile .

lesson-api-publish:
	@echo "==> Publishing latest image version"
	docker push derwaldemar/educonn-lesson-api:latest-staging
	docker push derwaldemar/educonn-lesson-api:${VERSION}

lesson-api-deploy:
	@echo "==> Deploying image version: ${VERSION}"
	@echo "==> Deploying image version: latest-staging"

