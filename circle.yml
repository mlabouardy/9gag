version: "2"
jobs:
  build:
    docker:
      - image: circleci/golang:1.8

    working_directory: /go/src/github.com/mlabouardy/9gag

    steps:
      - checkout
      - run:
          name: Get dependencies
          command: go get -v

      - run:
          name: Run vet commands
          command: go vet ./...

      - run:
          name: Run unit tests
          command: go test -v -race ./...
