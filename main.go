package main

import "./gitcmd"

func main() {

	gitlub := gitcmd.Commands{}

	// _, err := gitlub.Clone("git@github.com:rchicoli/gitlub.git", "/tmp/test2")
	// if err != nil {
	// 	panic(err)
	// }

	// err := gitlub.Fetch("/tmp/test2", "origin")
	// if err != nil {
	// 	panic(err)
	// }

	err := gitlub.Pull("/tmp/test2", "master")
	if err != nil {
		panic(err)
	}

}
