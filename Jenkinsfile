pipeline {
    agent {
        dockerContainer {
            image 'jenkins/agent'
        }
    }
    
    environment {
        DOCKER_HOST = 'tcp://192.168.122.42:4243'
    }

    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'main', url: 'https://github.com/snwzt/url-shortener'
            }
        }

        stage('Initialize Docker') {
            steps {
                script {
                    def dockerHome = tool 'docker'
                    env.PATH = "${dockerHome}/bin:${env.PATH}"
                }
            }
        }

        stage('Build') {
            steps {
                sh "docker build . -t ursh:${env.BUILD_ID}"
            }
        }
        
        stage('Push') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'e301d761-0cdf-44de-bda6-a634432bc4b7', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                    sh 'echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin'
                    sh "docker tag ursh:${env.BUILD_ID} notmde/urshsmth:${env.BUILD_ID}"
                    sh "docker push notmde/urshsmth:${env.BUILD_ID}"
                }
            }
        }
    }
}
