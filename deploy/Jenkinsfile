pipeline {
    agent any
    
    stages {
        stage('Build') {
            steps {
                // Pull the Docker image
                script {
                    docker.image('duybroooo/lightweight-echo:latest').pull()
                }
            }
        }
        
        stage('Run') {
            steps {
                // Run the Docker container
                script {
                    docker.image('duybroooo/lightweight-echo:latest').run('--rm -p 8080:8080')
                }
            }
        }
    }
}
