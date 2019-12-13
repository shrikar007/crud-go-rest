node{
    stage("SCM"){
       deleteDir()
       git 'https://github.com/shrikar007/crud-go-rest.git'
    }
    stage("SonarQube"){
        def scannerHome = tool 'SonarScanner';
         withSonarQubeEnv('SonarQube') { 
          sh "${scannerHome}/bin/sonar-scanner"
         }
         status()
    }
   
}
def status(){
     sh "curl -X GET http://192.168.1.69:9000/api/qualitygates/project_status\\?projectKey\\=crud-pipe-app-go-dev-shrikar > crud.json"
        def json = readJSON file:'crud.json'
        if ("${json.projectStatus.status}" == "ERROR") {
             currentBuild.result = 'FAILURE'
             error('Sonar Failed')
         }
}
