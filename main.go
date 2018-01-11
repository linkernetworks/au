package main

import (
	"bitbucket.org/linkernetworks/aurora/src/aurora"
	"bitbucket.org/linkernetworks/aurora/src/cmd/au/cmd"
)

func main() {
	cmd.SetBuildRevision(aurora.BuildRevision)
	cmd.SetBuildNumber(aurora.BuildNumber)
	cmd.SetBuildDate(aurora.BuildDate)
	cmd.Execute()
}
