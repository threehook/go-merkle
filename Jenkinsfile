pipeline {
    agent any
    tools {
        withEnv(['LANG=en_US.UTF-8', 'LANGUAGE=en_US.UTF-8', 'LC_ALL=en_US.UTF-8', LC_CTYPE=en_US.UTF-8']) {
            go 'go-1.21'
        }
    }
    environment {
      JAVA_OPTS = '-Dfile.encoding=UTF-8 -Dsun.jnu.encoding=UTF-8'
      LANG = 'en_US.UTF-8'
      LANGIAGE = 'en_US.UTF-8'
      LC_ALL = 'en_US.UTF-8'
      LC_CTYPE = 'en_US.UTF-8'
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
