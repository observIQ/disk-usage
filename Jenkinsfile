pipeline {
	agent any
    stages {
        stage('Run parallel scripts'){
            parallel{
                stage ('Linux'){
                    agent {
                        label "linux && docker"
                    }
                    stages{
                        stage('Checkout SCM') {
                            steps {
                                checkout scm
                            }
                        }
                        stage('Run linux build script') {
                            steps {
                                sh '''make'''
                            }
                        }
                    }
                }
                stage ('Windows'){
                    agent {
                        label "windows && golang"
                    }
                    stages {
                        stage('Checkout SCM') {
                            steps {
                                checkout scm
                            }
                        }
                        stage('Run windows unit tests') {
                            steps {
                                sh '''c:/Go/bin/go test ./...'''
                            }
                        }
                        stage('Run windows build') {
                            steps {
                                sh '''c:/Go/bin/go build'''
                            }
                        }
                    }
                }
            }
        }
    }
}
