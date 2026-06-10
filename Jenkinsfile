pipeline {
    agent any

    options {
        timeout(time: 15, unit: 'MINUTES')
    }

    stages {
        stage('CI Env') {
            agent {
                docker {
                    image 'golang:1.25.3'
                }
            }

            stages {
                stage('Download Dependencies') {
                    steps {
                        sh 'go mod tidy'
                    }
                }

                stage('Check Formatting') {
                    steps {
                        sh 'test -z "$(gofmt -s -l .)"'
                    }
                }

                stage('Go Vet') {
                    steps {
                        sh 'go vet ./...'
                    }
                }

                stage('GolangCI Lint') {
                    steps {
                        sh '''
                            curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
                              | sh -s -- -b $(go env GOPATH)/bin v1.64.5

                            $(go env GOPATH)/bin/golangci-lint run ./...
                        '''
                    }
                }

                stage('Run Tests') {
                    steps {
                        sh 'go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...'
                    }
                }
            }
        }
    }

    post {
        always {
            archiveArtifacts artifacts: 'coverage.txt', allowEmptyArchive: true
        }
    }
}
