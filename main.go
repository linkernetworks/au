package main

import "bitbucket.org/linkernetworks/aurora/src/cmd/au/cmd"

var version string

func main() {
	cmd.SetBuildNumber(version)
	cmd.Execute()
}
