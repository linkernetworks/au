package main

import "bitbucket.org/linkernetworks/aurora/src/cmd/au/cmd"

var revision string

func main() {
	cmd.SetRevision(revision)
	cmd.Execute()
}
