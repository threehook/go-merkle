pipeline {
    agent any
    tools {
        go 'go-1.21'
    }
    environment {
      GO111MODULE = 'on'
      ghcrCredential = credentials('ghcr-pat')
      registry = 'threehook/go-merkle'
      //registryCredential = 'ghcrCredential'
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
            steps {
                sh 'sudo groupadd docker'
                sh 'sudo usermod -aG docker $USER'
                script {
                    dockerImage = docker.build registry + ":$BUILD_NUMBER"
                }
            }
        }
        stage('Deploying image') {
            steps {
                script {
                    docker.withRegistry( 'https://ghcr.io', ghcrCredential ) {
                        dockerImage.push()
                    }
                }
            }
        }
    }
}
