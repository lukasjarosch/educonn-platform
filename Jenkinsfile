node {
	ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/src/github.com/lukasjarosch/educonn-platform/") {
		withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
					env.PATH="${GOPATH}/bin:$PATH"
						stage('Checkout'){
						echo 'Checking out SCM'
						checkout scm
						}
					stage('Dependencies') {
						echo 'Pulling Dependencies'

						sh 'go version'
						sh 'go get -u github.com/golang/dep/cmd/dep'
						sh 'dep ensure -v'
					}
					stage('Build') {
						echo 'Building'
						sh('make user-proto')
						sh('make user')
					}
		}
	}
}
