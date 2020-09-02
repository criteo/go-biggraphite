package main

import (
	"fmt"

	"github.com/criteo/go-biggraphite/version"
)

func main() {
	fmt.Println("Project: " + version.ProjectName)
	fmt.Println("Description: " + version.ProjectDescription)
	fmt.Println("Copyright: " + version.ProjectCopyright)
	fmt.Println("Version: " + version.Version)
	fmt.Println("Revision: " + version.Revision)
	fmt.Println("Branch: " + version.Branch)
}
