#!/usr/bin/env groovy

pipeline {
	agent any

		stages {
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
}
