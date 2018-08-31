ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
	withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
                env.PATH="${GOPATH}/bin:$PATH"
					stage('Checkout'){
                    echo 'Checking out SCM'
                    checkout scm
					}
				stage('Pre Test'){
                    echo 'Pulling Dependencies'

                    sh 'go version'
                    sh 'go get -u github.com/golang/dep/cmd/dep'
                    sh 'go get -u github.com/golang/lint/golint'
                    sh 'go get github.com/tebeka/go2xunit'

                    //or -update
                    sh 'cd ${GOPATH}/src/cmd/project/ && dep ensure' 

				}
	}
}
