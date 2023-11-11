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
            environment {
                LC_ALL = 'en_US.UTF-8'
                LANG    = 'en_US.UTF-8'
                LANGUAGE = 'en_US.UTF-8'
            }
            steps {
                sh 'go build'
            }
        }
        // don't forget to include a stage for unit testing right here
        stage('Building image') {
            steps {
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
