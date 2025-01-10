pipeline {
  agent { label 'ecs-golang-1_22_1' }
  stages {
    stage('Test') {
      steps {
        sh """
          echo "Testing..."
          cd services/golang
          go run github.com/onsi/ginkgo/v2/ginkgo -r
        """
      }
    }
    stage('Deploy') {
      steps {
        sh 'echo "Deploying..."'
      }
    }
  }
}