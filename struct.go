package main

import (
	"encoding/json"
	"io"
)

type DirectoryEntry struct {
	Filename string `json:"f"`
	INodeID  uint64 `json:"i"`
}

type Directory []DirectoryEntry

type INode struct {
	Size  uint64 `json:"s"`
	Atime uint64 `json:"at"`
	Mtime uint64 `json:"mt"`
	Ctime uint64 `json:"ct"`
	Mode  uint32 `json:"m"`
	Uid   uint32 `json:"u"`
	Gid   uint32 `json:"g"`

	Blocks []uint64 `json:"b"`
}

func (d Directory) Marshal() []byte {
	res, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return res
}

func UnmarshalDirectory(data []byte) (Directory, error) {
	d := make(Directory, 0)
	err := json.Unmarshal(data, d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func UnmarshalDirectoryFromReader(r io.Reader) (Directory, error) {
	d := make(Directory, 0)
	dec := json.NewDecoder(r)
	err := dec.Decode(d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (ino INode) Marshal() []byte {
	res, err := json.Marshal(ino)
	if err != nil {
		panic(err)
	}
	return res
}

func UnmarshalINode(data []byte) (INode, error) {
	var ino INode
	err := json.Unmarshal(data, &ino)
	if err != nil {
		return INode{}, err
	}
	return ino, nil
}
