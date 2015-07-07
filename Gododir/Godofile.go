package main

import (
	. "gopkg.in/godo.v1"
)

func tasks(p *Project) {
	Env = `GOPATH=.vendor::$GOPATH`

	p.Task("build", D{"lint"}, func() error {
		return Bash(`gom exec gox ./src/...`)
	})

	p.Task("lint", func() error {
		//Bash(`gom exec golint ./src/...`)
		Run("golint ./src/...")
		Run("gofmt -w -s ./src/")
		return Run("go vet ./src/...")
	}).Watch("**/*.go")

	p.Task("test", D{"lint"}, func() error {
		return Run(`gom exec go test -v ./src/...`)
	}).Watch("**/*.go")
}

func main() {
	Godo(tasks)
}