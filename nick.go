package main

import (
//    "log"
    "github.com/ajray/go-fuse/fuse"
//    "strings"
)

type NickFile struct {
    Nick string
    fuse.DefaultFile
}

func NewNickFile(nick string) *NickFile {
    f := new(NickFile)
    f.Nick = nick
    return f
}

func (me *NickFile) Read(input *fuse.ReadIn, bp fuse.BufferPool) ([]byte, fuse.Status) {
    end := int(input.Offset) + int(input.Size)
    if end > len(me.Nick) {
        end = len(me.Nick)
    }

    return []byte(me.Nick[input.Offset:end]), fuse.OK
}
