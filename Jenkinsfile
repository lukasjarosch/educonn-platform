#!/usr/bin/env groovy

pipeline {

	agent {
		docker {
            image 'node'
            args '-u root'
		}
	}

	stages {
		stage('Build') {
			steps {
                echo 'Building...'
			}
		}
		stage('Test') {
			steps {
                echo 'Testing...'
			}
		}
	}
}
