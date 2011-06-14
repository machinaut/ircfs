// Copyright 2011 Alex Ray. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "log"
    "github.com/ajray/go-fuse/fuse"
    "time"
    "os/signal"
)

type StoopidFs struct{
    fuse.DefaultFileSystem
}

func main() {
    go func() {
        for {
            sig := <-signal.Incoming
            time.Sleep(1) // FIXME TODO XXX hack to make the goroutine scheduler switch
            log.Print("Reading signal: " + sig.String() )
        }
    }()
    _, _, _ = fuse.MountFileSystem("foo", &StoopidFs{}, nil)
    for {
        time.Sleep(1e9)
        log.Print("ho")
    }
}
