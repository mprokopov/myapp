pipeline {
    agent any

    tools {
       go "1.24.1"
    }

    triggers {
        pollSCM('*/1 * * * *') // Poll Git repository every 1 minute
    }

    stages {
        stage('Build') {
            steps {
                sh "go build main.go"
            }
        }
    }
}
