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
        stage('Deploy') {
            steps {
                sh """
mkdir -p ~/.ssh

ssh-keyscan 172.16.0.3 >> ~/.ssh/known_hosts

scp main 172.16.0.3:/home/laborant/
"""
            }
        }
    }
}
