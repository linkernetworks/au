package main

import "bitbucket.org/linkernetworks/aurora/src/cmd/au/cmd"

var buildNumber string

func main() {
	cmd.SetBuildNumber(buildNumber)
	cmd.Execute()
}
