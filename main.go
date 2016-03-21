package main

import (
	"flag"
	"log"

	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

//Exfs is a implementation of pathfs.FileSystem
type ExFs struct {
	pathfs.FileSystem
}

// A exfs will be mount at the path in the first arg
func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		log.Fatal("Usage:\n  exfs MOUNTPOINT NEEDED")
	}

	nfs := pathfs.NewPathNodeFs(&ExFs{FileSystem: pathfs.NewDefaultFileSystem()}, nil)
	server, _, err := nodefs.MountRoot(flag.Arg(0), nfs.Root(), nil)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}
	server.Serve()
}
