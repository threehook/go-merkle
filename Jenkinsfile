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
        stage('Initializing docker') {
            steps {
                script {
                    def dockerHome = tool 'docker'
                    env.PATH = "${dockerHome}/bin:${env.PATH}"
                    sh 'docker run -d --name my-jenkins -v /var/jenkins_home:~/.jenkins -v /var/run/docker.sock:/var/run/docker.sock -p 8080:8080 jenkins'
                }
            }
        }
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
