pipeline {
    agent any

    stages {
        stage('Packaging/Pushing imagae') {

            steps {
                withDockerRegistry(credentialsId: 'dockerhub', url: 'https://index.docker.io/v1/') {
                    sh 'docker build -t duybroooo/lightweight-echo .'
                    sh 'docker push duybroooo/lightweight-echo'
                }
            }
        }

        stage('Pull and Run') {
            steps {
                echo 'Deploying and cleaning'
                sh 'docker pull duybroooo/lightweight-echo'
                sh 'docker run -d --rm --name duybroooo-lightweight-echo -e PORT=3000 -p 3000:3000 duybroooo/lightweight-echo'
            }
        }
    }
}