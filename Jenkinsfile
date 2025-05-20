pipeline {
    agent any

    tools {
       go "1.24.1"
    }

    triggers {
        cron('* * * * *') // Every 1 minute
    }

    stages {
        stage('Build') {
            steps {
                sh "go build main.go"
            }
        }
    }
}
