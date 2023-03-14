package main

import "github.com/benmizrahi/godist/master"

func main() {

	sc := master.NewMaster(true, "localhost", 9999, 2).Context()

	sc.
		Extract("").
		Transform("").
		Load("")
}
