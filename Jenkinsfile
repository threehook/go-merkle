pipeline {
    agent any
    tools {
        go 'go-1.21'
    }
    environment {
      GO111MODULE = 'on'
      ghcr-credential = credentials('ghcr-pat')
      registry = "threehook/go-merkle"
      registryCredential = 'ghcr-credential'
      dockerImage = ''
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
            script {
                docker.withRegistry( '', registryCredential ) {
                dockerImage.push()
            }
         }
    }
}
