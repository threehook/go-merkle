pipeline {
    environment {
      GO111MODULE = 'on'
      registry = "gustavoapolinario/docker-test"
      registryCredential = 'dockerhub'
      dockerImage = ''
    }
    agent any
    tools {
        go 'go-1.21'
    }
    stages {
        stage('Building app') {
            steps {
                sh 'go build'
            }
        }
        // don't forget to include a stage for unit testing right here
        stage('Building image') {
          steps{
            script {
              dockerImage = docker.build registry + ":$BUILD_NUMBER"
            }
          }
        }
    }
}
