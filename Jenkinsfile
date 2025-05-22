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
                withCredentials([sshUserPrivateKey(credentialsId: 'target-ssh-key',
                                                   keyFileVariable: 'ssh_key',
                                                   usernameVariable: 'ssh_user')]) {
                    sh """

cat ${ssh_key}

scp -o StrictHostKeyChecking=no main ${ssh_user}@target: -i ${ssh_key}
"""

}
            }
        }
    }
}
