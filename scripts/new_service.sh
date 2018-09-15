#!/usr/bin/env bash

SERVICE=$1
TYPE=$2



echo "Bootstrapping service $SERVICE"

# Docker and Jenkins
echo " + $SERVICE/build"
mkdir -p ./$SERVICE/build/ci
tee $SERVICE/build/Dockerfile << EOF > /dev/null
FROM golang:1.10 as builder

MAINTAINER lukas.jarosch@mail.com

# dep
RUN wget https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -O /usr/bin/dep && chmod +x /usr/bin/dep

# gopath
RUN mkdir -p /go/src/github.com/lukasjarosch/educonn-platform/${SERVICE}/
WORKDIR /go/src/github.com/lukasjarosch/educonn-platform/${SERVICE}/

# vendoring
COPY ${SERVICE}/Gopkg.lock ${SERVICE}/Gopkg.toml ./
RUN dep ensure -vendor-only -v

# build
COPY ${SERVICE}/ ./
RUN ${SERVICE}/build.sh

# ---------

FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/lukasjarosch/educonn-platform/${SERVICE}/cmd/${SERVICE}/${SERVICE} .
ENTRYPOINT ["./${SERVICE}"]
CMD ["./${SERVICE}"]
EOF

tee $SERVICE/build/ci/Jenkinsfile << EOF > /dev/null
node {
	ws("\${JENKINS_HOME}/jobs/\${JOB_NAME}/builds/}\${BUILD_ID}") {
		withEnv(["GOPATH=\${JENKINS_HOME}/jobs/\${JOB_NAME}/builds/\${BUILD_ID}"]) {
            env.PATH="\${GOPATH}/bin:\$PATH"
                stage('Checkout'){
                echo 'Checking out SCM'
                checkout scm
                }
            stage('Build') {
                echo 'Building'
                dir('src/github.com/lukasjarosch/educonn-platform/') {
                    sh('make ${SERVICE}-docker')
                }
            }
            stage('Publish') {
                withDockerRegistry([ credentialsId: "1f664f6b-8f57-435a-bda6-5de629641b79", url: "" ]) {
                    dir('src/github.com/lukasjarosch/educonn-platform/') {
                        sh('make ${SERVICE}-publish')
                    }
                }
            }
            stage('Deploy') {
                dir('src/github.com/lukasjarosch/educonn-platform/') {
                    tag = sh(script: "git describe --abbrev=0 --tags || echo '0.1'", returnStdout: true).trim()
                    commit = sh(script: 'git log -n 1 --pretty=format:"%h"', returnStdout: true).trim()
                    sh "ssh -o StrictHostKeyChecking=no manager@159.89.3.151 'bash edu-update-svc.sh educonn_${SERVICE} derwaldemar/educonn-${SERVICE}:\${tag}-\${commit}'"
                }
            }
		}
	}
}
EOF

# create main.go
SERVICE_CC=$(echo "$SERVICE" | sed -r 's/(^|-)(\w)/\U\2/g')

echo " + $SERVICE/cmd"
mkdir -p ./$SERVICE/cmd/$SERVICE/
tee $SERVICE/cmd/$SERVICE/main.go << EOF > /dev/null
package main

import (
	"os"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/micro/go-micro"
	"github.com/lukasjarosch/educonn-platform/${SERVICE}/internal/platform/config"
	pb${SERVICE_CC} "github.com/lukasjarosch/educonn-platform/${SERVICE}/proto"
	"time"
	"github.com/lukasjarosch/educonn-platform/${SERVICE}/internal/service"
)

func main() {
	if os.Getenv("DEV_ENV") != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// setup micro service
	svc := micro.NewService(
		micro.Name(config.ServiceName),
		micro.Version(config.Version),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	svc.Init()

	// service handler
	pb${SERVICE_CC}.Register${SERVICE_CC}Handler(
		svc.Server(),
		service.New${SERVICE_CC}Service(),
	)

	// fire
	if err := svc.Run(); err != nil {
		panic(err)
	}
}
EOF

echo " + $SERVICE/internal/platform/config"
mkdir -p ./$SERVICE/internal/platform/config

# create version.go
SERVICE_NAMESPACED="educonn.$TYPE.$SERVICE"

echo " + $SERVICE/internal/platform/config/version.go"
tee $SERVICE/internal/platform/config/version.go <<EOF > /dev/null
package config

// Version is used to define the application version and should be overridden during the
// compilation phase.
var Version = "(undefined)"

// ServiceName is the discovery identity for this service.
const ServiceName = "$SERVICE_NAMESPACED"
EOF
echo " ! Discoverable service name: $SERVICE_NAMESPACED"


# create service
SERVICE_NAME="$(tr '[:lower:]' '[:upper:]' <<< ${SERVICE:0:1})${SERVICE:1}"

echo " + $SERVICE/internal/service"
mkdir -p $SERVICE/internal/service
tee $SERVICE/internal/service/service.go <<EOF > /dev/null
package service

import (
	_ "github.com/rs/zerolog"
	_ "github.com/rs/zerolog/log"
)

type $(sed 's/./\l&/1' <<< "${SERVICE_CC}")Service struct {
}

func New${SERVICE_CC}Service() *$(sed 's/./\l&/1' <<< "${SERVICE_CC}")Service {
	return &$(sed 's/./\l&/1' <<< "${SERVICE_CC}")Service{}
}
EOF
echo " + $SERVICE/internal/service/service.go"

# create scripts
echo " + $SERVICE/scripts"
mkdir -p $SERVICE/scripts
echo " + $SERVICE/scripts/build.sh"
tee $SERVICE/scripts/build.sh <<EOF > /dev/null
#!/usr/bin/env bash

# Set the service's name
SERVICE_NAME=$SERVICE

VERSION=\$1

LD_FLAGS="-X github.com/lukasjarosch/educonn-platform/\${SERVICE_NAME}/internal/platform/config.Version=\${VERSION}"

# Get the parent directory of where this script is.
SOURCE="\${BASH_SOURCE[0]}"
while [ -h "\$SOURCE" ] ; do SOURCE="\$(readlink "\$SOURCE")"; done
DIR="\$( cd -P "\$( dirname "\$SOURCE" )/.." && pwd )"

# Change into that directory
cd "\$DIR"

# Delete old binary
echo "==> Removing old \$SERVICE_NAME binary..."
rm  -f cmd/\$SERVICE_NAME/\$SERVICE_NAME

# We want a statically linked binary
export CGO_ENABLED=0

# Build!
echo "==> Building \$SERVICE_NAME..."
cd cmd/\$SERVICE_NAME
go build -ldflags="\${LD_FLAGS}" -a -installsuffix cgo -o \$SERVICE_NAME main.go
EOF
chmod +x $SERVICE/scripts/build.sh

echo " + $SERVICE/scripts/proto.sh"
tee $SERVICE/scripts/proto.sh <<EOF > /dev/null
#!/usr/bin/env bash

# Get the parent directory of where this script is.
SOURCE="\${BASH_SOURCE[0]}"
while [ -h "\$SOURCE" ] ; do SOURCE="\$(readlink "\$SOURCE")"; done
DIR="\$( cd -P "\$( dirname "\$SOURCE" )/.." && pwd )"

cd "\$DIR/proto"

echo "==> Building protocol buffers for \$(basename \$(dirname \$(dirname "\$SOURCE")))"
protoc  \\
    -I. \\
    --go_out=plugins=micro:\$GOPATH/src \\
    --micro_out=:\$GOPATH/src \\
    *.proto
EOF
chmod +x $SERVICE/scripts/proto.sh

# create proto
echo " + $SERVICE/proto"
mkdir -p $SERVICE/proto


tee $SERVICE/proto/$SERVICE.proto <<EOF > /dev/null
syntax = "proto3";

package educonn.${SERVICE/-/.};

option go_package = "github.com/lukasjarosch/educonn-platform/$SERVICE/proto";

service ${SERVICE_CC}{
}

// ----------------------------
// ${SERVICE^^} RESOURCE
// ----------------------------


// ----------------------------
// EVENTS
// ----------------------------


// ----------------------------
// MISC
// ----------------------------
message Error {
    int32 code = 1;
    string description = 2;
}

EOF


# update Makefile
echo " + Updating Makefile"
tee -a  Makefile << EOF > /dev/null

# --------- ${SERVICE^^} ---------
${SERVICE}:
	@sh -c "'\$(CURDIR)'/${SERVICE}/scripts/build.sh" \${VERSION}

${SERVICE}-proto:
	@sh -c "'\$(CURDIR)'/${SERVICE}/scripts/proto.sh"

${SERVICE}-run:
	@echo Starting the ${SERVICE^^} service
	cd ${SERVICE}/cmd/${SERVICE} && go run main.go

${SERVICE}-docker:
	@echo "==> Building ${SERVICE} docker image..."
	docker build -t derwaldemar/educonn-${SERVICE}:\${VERSION} -t derwaldemar/educonn-${SERVICE}:latest-staging -f ${SERVICE}/build/Dockerfile .

${SERVICE}-publish:
	@echo "==> Publishing latest image version"
	docker push derwaldemar/educonn-${SERVICE}:latest-staging
	docker push derwaldemar/educonn-${SERVICE}:\${VERSION}

${SERVICE}-deploy:
	@echo "==> Deploying image version: \${VERSION}"
	@echo "==> Deploying image version: latest-staging"
EOF


make ${SERVICE}-proto

echo "!! DONE: Don't forget to execute: dep init -v"
