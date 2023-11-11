pipeline {
    environment {
      GO111MODULE = 'on'
      registry = "threehook/go-merkle"
      registryCredential = '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'
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
