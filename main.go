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
var nick = "ajray"
var ctl = NewCtlFile()

type IrcFs struct {
    fuse.DefaultFileSystem
}

func (me *IrcFs) GetAttr(name string) (*os.FileInfo, fuse.Status) {
    log.Print("GetAttr " + name)
    switch name {
    case "file.txt":
        return &os.FileInfo{Mode: fuse.S_IFREG | 0444, Size: int64(len(name))}, fuse.OK
    case "ctl":
        return &os.FileInfo{Mode: fuse.S_IFREG | 0222, Size: int64(len(name))}, fuse.OK
    case "event":
        return &os.FileInfo{Mode: fuse.S_IFREG | 0444, Size: int64(len(name))}, fuse.OK
    case "nick":
        return &os.FileInfo{Mode: fuse.S_IFREG | 0444, Size: int64(len(nick))}, fuse.OK
    case "raw":
        return &os.FileInfo{Mode: fuse.S_IFREG | 0666, Size: int64(len(name))}, fuse.OK
    case "pong":
        return &os.FileInfo{Mode: fuse.S_IFREG | 0444, Size: int64(len(name))}, fuse.OK
    case "":
        return &os.FileInfo{Mode: fuse.S_IFDIR | 0755}, fuse.OK
    }
    return nil, fuse.ENOENT
}

func (me *IrcFs) OpenDir(name string) (stream chan fuse.DirEntry, code fuse.Status) {
    log.Print("OpenDir " + name)
    if name == "" {
        stream = make(chan fuse.DirEntry, 6) // , n + 5) // MUAHAHA NO BUFFER
        stream <- fuse.DirEntry{Name: "file.txt", Mode: fuse.S_IFREG}
        stream <- fuse.DirEntry{Name: "ctl", Mode: fuse.S_IFREG}
        stream <- fuse.DirEntry{Name: "event", Mode: fuse.S_IFREG}
        stream <- fuse.DirEntry{Name: "nick", Mode: fuse.S_IFREG}
        stream <- fuse.DirEntry{Name: "raw", Mode: fuse.S_IFREG}
        stream <- fuse.DirEntry{Name: "pong", Mode: fuse.S_IFREG}
        close(stream)
        return stream, fuse.OK
    }
    return nil, fuse.ENOENT
}

func (me *IrcFs) Open(name string, flags uint32) (file fuse.File, code fuse.Status) {
    log.Print("Open " + name)
    switch name {
    case "nick":
        return fuse.NewReadOnlyFile([]byte(nick)), fuse.OK
    case "file.txt":
        return fuse.NewReadOnlyFile([]byte(name)), fuse.OK
    case "ctl":
        return ctl, fuse.OK
    }
    return nil, fuse.ENOENT
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
