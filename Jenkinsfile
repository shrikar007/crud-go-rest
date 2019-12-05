node('master') {
    docker.withServer('unix:///var/run/docker.sock') {
        stage('Code Checkout') {
            docker
                .image('golang:latest')
                .inside('--volumes-from jenkins-ci') {
                    checkout scm
                }
        }
       
        stage('Build') {
            docker
                .image('golang:latest')
                .inside('--volumes-from jenkins-ci') {
                    echo "Deployinh"
                }
        }
    }
}
