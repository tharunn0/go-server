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

            environment {
                GOCACHE = "${WORKSPACE}/.gocache"
                HOME = "${WORKSPACE}"
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
                            curl -sSfL https://golangci-lint.run/install.sh \
                              | sh -s -- -b $(go env GOPATH)/bin v2.12.2

                            $(go env GOPATH)/bin/golangci-lint --version
                            $(go env GOPATH)/bin/golangci-lint run ./...
                        '''
                    }
                }

                stage('Run Tests') {
                    steps {
                        sh 'go test -v  -coverprofile=coverage.txt -covermode=atomic ./...'
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
