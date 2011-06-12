// Copyright 2011 Alex Ray. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "log"
    "github.com/ajray/go-fuse/fuse"
    "os"
)

var n = 3

type IrcFs struct {
    fuse.DefaultFileSystem
}

func (me *IrcFs) GetAttr(name string) (*os.FileInfo, fuse.Status) {
    switch name {
    case "file.txt":
        return &os.FileInfo{Mode: fuse.S_IFREG | 0644, Size: int64(len(name))}, fuse.OK
    case "":
        return &os.FileInfo{Mode: fuse.S_IFDIR | 0755}, fuse.OK
    }
    return nil, fuse.ENOENT
}

func (me *IrcFs) OpenDir(name string) (stream chan fuse.DirEntry, code fuse.Status) {
    if name == "" {
        stream = make(chan fuse.DirEntry) // , n + 5) // MUAHAHA NO BUFFER
        stream <- fuse.DirEntry{Name: "file.txt", Mode: fuse.S_IFREG}
        close(stream)
        return stream, fuse.OK
    }
    return nil, fuse.ENOENT
}

func (me *IrcFs) Open(name string, flags uint32) (file fuse.File, code fuse.Status) {
    if name != "file.txt" {
        return nil, fuse.ENOENT
    }
    if flags&fuse.O_ANYWRITE != 0 {
        return nil, fuse.EPERM
    }
    return fuse.NewReadOnlyFile([]byte(name)), fuse.OK
}

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Usage:  ircfs MOUNTPOINT")
    }
    state, _, err := fuse.MountFileSystem(os.Args[1], &IrcFs{}, nil)
    if err != nil {
        log.Fatal("Mount fail:", err)
    }
    state.Loop(true)
}
