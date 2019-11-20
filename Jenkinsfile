pipeline {
	agent any
    environment {

    }
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
                        label "windows && docker"
                    }
                    stages {
                        stage('Checkout SCM') {
                            steps {
                                checkout scm
                            }
                        }
                        stage('Run windows unit tests') {
                            steps {
                                sh '''go test ./...'''
                            }
                        }
                        stage('Run windows build') {
                            steps {
                                sh '''go build'''
                            }
                        }
                    }
                }
            }
        }
    }
}
