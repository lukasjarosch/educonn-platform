#!/usr/bin/env groovy

pipeline {

	agent {
		docker {
            args '-u root'
		}
	}

	stages {
		stage('Build') {
			steps {
                echo 'Building educonn-user'
				make educonn-user-proto
				make educonn-user
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
				make user-docker
			}
		}
	}
}
