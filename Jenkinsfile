#!/usr/bin/env groovy

node {
	// Install the desired Go version
	def root = tool name: 'Go 1.10.1', type: 'go'

	// Export environment variables pointing to the directory where Go was installed
	withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
		sh 'go version'
	}

	stage('Dep') {
		steps {
			echo 'Ensure dependencies'
			sh('dep ensure -v')
		}
	}
	stage('Build') {
		steps {
			echo 'Building educonn-user'
			sh("make user-proto")
			sh("make user")
		}
	}
	stage('Test') {
		steps {
			echo 'Testing...'
		}
	}
	stage('Docker') {
		steps {
			echo 'Building educonn-user'
			sh("make user-docker")
		}
	}
}
