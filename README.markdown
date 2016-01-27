## About

Photo website, currently unusable.

> Calhoun was right about the camera: it destroys everything it shoots.

  \- The Tale of the Curious Camera

### Run tests

For now, there is one package per layer (data/application/UI), and their test suites can be run
independently using absolute paths: `go test ./app ./db`.  To get test output, you can pass `-v`, or
run them one package at a time `go test` from the package directory.

