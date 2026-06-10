pipeline {
    agent any

    environment {
        IMAGE_NAME = 'johnwickk99/go-server'
    }

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
                        sh 'go mod download'
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

                // stage('GolangCI Lint') {
                // steps {
                //         sh '''
                //             curl -sSfL https://golangci-lint.run/install.sh \
                //               | sh -s -- -b $(go env GOPATH)/bin v2.12.2

                //             $(go env GOPATH)/bin/golangci-lint --version
                //             $(go env GOPATH)/bin/golangci-lint run ./...
                //         '''
                //     }
                // }

                stage('Run Tests') {
                    steps {
                        sh 'go test -v ./...'
                    }
                }
            }
        }
        stage('Build Docker Image') {
                            steps {
                                script {
                                    env.SHORT_SHA = sh(
                                        script: 'git rev-parse --short=8 HEAD',
                                        returnStdout: true
                                    ).trim()

                                    env.IMAGE_TAG = "dev-${env.SHORT_SHA}"

                                    echo "Building image: ${IMAGE_NAME}:${IMAGE_TAG}"

                                    docker.build(
                                        "${IMAGE_NAME}:${IMAGE_TAG}",
                                        "."
                                    )
                                }
                            }
                        }

                        stage('Push Docker Image') {
                            steps {
                                script {
                                    docker.withRegistry(
                                        'https://index.docker.io/v1/',
                                        'dockerhub-cred'
                                    ) {

                                        def image = docker.image(
                                            "${IMAGE_NAME}:${IMAGE_TAG}"
                                        )

                                        image.push()

                                        // Optional moving tag
                                        image.push('dev')
                                    }
                                }
                            }
                        }
    }

    post {
        success {
            build (
                job: 'go-server-dev-deploy',
                parameters: [
                    string(
                        name: 'IMAGE_TAG',
                        value: env.IMAGE_TAG
                    )
                ]
            )
        }
    }
}
