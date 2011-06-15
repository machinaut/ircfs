include $(GOROOT)/src/Make.inc

TARG=ircfs
GOFILES=main.go ctl.go nick.go

include $(GOROOT)/src/Make.cmd
