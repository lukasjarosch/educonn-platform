#!/usr/bin/env groovy

pipeline {

	agent any

	stages {
		stage('Build') {
			steps {
                echo 'Building educonn-user'
				sh("make educonn-user-proto")
				sh("make educonn-user")
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
