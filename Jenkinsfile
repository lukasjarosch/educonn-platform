node {
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

						dir('src/github.com/lukasjarosch/educonn-platform/') {
							sh 'dep ensure -v'
						}

					}
		}
	}
}
