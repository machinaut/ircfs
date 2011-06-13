/*
/ctl
The following commands are understood:
reconnect
connect address nick [password [fromhost]]
disconnect [msg]
quit [msg]
join channel [key]
away [msg]
back
nick nick
umode (to request the user mode)
whois nick
invite nick channel
msg nick|channel msg
time
version
ping (ctcp messages)
ctcp [request]
debug 0|1
Reconnect performs a reconnect with the parameters from the previous
connection. Commands are issued directly to the irc server. Responses
from the irc server can be read from the status directory data file
/0/data. The responses are free-form and cannot be matched to specific
commands (the irc protocol does not contain the necessary information).
*/
package main

import (
    "log"
    "github.com/ajray/go-fuse/fuse"
)

type CtlFile struct {
    fuse.DevNullFile
}

func NewCtlFile() *CtlFile {
    f := new(CtlFile)
    return f
}

func (me *CtlFile) Write(input *fuse.WriteIn, content []byte) (uint32, fuse.Status) {
    log.Println("Write: "+string(content))
    return uint32(len(content)), fuse.OK
}
