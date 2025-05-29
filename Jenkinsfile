pipeline {
    agent any

    tools {
       go "1.24.1"
    }

    triggers {
        pollSCM('*/1 * * * *') // Poll Git repository every 1 minute
    }

    stages {
        stage('Unit Test') {
            steps {
                sh "go test -v ./..."
            }
        }
        stage('Build') {
            steps {
                sh "go build main.go"
            }
        }
        stage('Docker Image') {
            steps {
                sh "docker build . --tag ttl.sh/myapp:1h"
            }
        }
        stage('Docker Push') {
            steps {
                sh "docker push ttl.sh/myapp:1h"
            }
        }
        stage('Deploy to Kubernetes') {
            steps {
                sh "kubectl apply -f pod.yaml"
            }
        }
    }
}
