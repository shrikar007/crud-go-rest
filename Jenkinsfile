node('master') {
    docker.withServer('unix:///var/run/docker.sock') {
        stage('Code Checkout') {
            docker
                .image('golang:latest')
                .inside('--volumes-from jenkins-run_jenkins_1') {
                    checkout scm
                }
        }
       
        stage('Build') {
            docker
                .image('golang:latest')
                .inside('--volumes-from jenkins-run_jenkins_1') {
                    echo "Deployinh"
                }
        }
    }
}
