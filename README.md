
ircfs - irc file server
=======================
SYNOPSIS
--------

    mount { ircfs [ -dt ] [ -a addr ] [ -f fromhost ] [ -n nick ] [ -p password ] [ -l logpath ] netname } mtpt
DESCRIPTION
-----------

Ircfs maintains a connection to an irc server and exports it using styx. Communication channels (irc channels or queries to users) are represented by directories with a unique sequence number as name. The file data in these directories allow text to be read and written. Other files, such as ctl allow execution of irc commands such as setting the topic and kicking people. Directory ``0'' is always present, providing status messages such as irc connection errors and certain irc messages.
###Options
-a addr
Use addr as irc server. Note that ircfs does not automatically dial the irc server at startup. Write the `connect` or `reconnect` command to ctl after startup to connect. 
-f fromhost
Make the connection from ip address or host name fromhost. 
-n nick
Use nick as nick when connecting. 
-l logpath
Log contents of the data files of all irc directories to files named logpath/name.log, where name is the channel or user name. 
-t
Use long timestamps, in both logging and the data file. Also removes text type (see section about the data file) from messages before logging. 
-d
Enable debugging.
###Hierarchy
The following files are exported:

####/ctl
The following commands are understood: reconnect, connect address nick [password [fromhost]], disconnect [msg], quit [msg], join channel [key], away [msg], back, nick nick, umode (to request the user mode), whois nick, invite nick channel, msg nick|channel msg, time, version, ping (ctcp messages), and ctcp [request], and finally debug 0|1. Reconnect performs a reconnect with the parameters from the previous connection. Commands are issued directly to the irc server. Responses from the irc server can be read from the status directory data file /0/data. The responses are free-form and cannot be matched to specific commands (the irc protocol does not contain the necessary information). 
####/event
Reads from this read-only file will block until an event occurs, one event per line. The first word indicates the type of event. ``new n name'' indicates that a new directory n with channel or user name name is available. For example after a join or an incoming user query. ``del n'' indicates directory n has been closed. For example after a part. Irc connection status changes are: connecting, connected nick, disconnected', and nick nick , with nick being our current or new nick name. 
####/nick
This read-only file returns our current nick or the nick that will be used when connecting. 
####/raw
Reading blocks until an irc message is read from or written to the irc server. Such messages are returned with <<< or >>> prepended. Writes to this file are passed on to the irc connection verbatim, such writes will also be read by readers of this file. 
####/pong
Read-only file to check whether the connection to ircfs and the irc server is responsive. Ircfs sends PING messages to the irc server. When the server replies with a PONG, a read on this file returns a line containing ``pong n'', with n the number of seconds between the ping and pong messages. If the irc server is not responding in timely fashion, reads will return lines of the form ``nopong n'', with n the number of elapsed seconds since the ping message. These ``nopong''-messages are repeated every 5 seconds. This allows clients like wm/irc to determine whether the connection to the ircfs is functioning. 
####/n
Connection directory. Each channel or user query is represented by a directory. The special directory ``/0'' is reserved for connection status messages. 
####/n/ctl
Commands accepted by /ctl are also understood by /n/ctl, modulo some commands that do not work on the status directory. Additional commands: names, n (shorthand for names), me msg, notice msg, mode [mode], part [msg], remove [msg], topic [topic]. Remove is like part but also sends a hint on the event file that the target is to be removed. The following commands only work on channels and require one or more user names as parameters: kick, op, deop, ban, unban, voice, devoice . 
####/n/data
Data written to this file is sent to the channel or user query. A write can contain multiple lines. Reads on this file block until activity occurs on the channel or query. After opening, reads will first return a backlog of text. Each line will start with two characters, ``+ '' for normal text, ``- '' for normal text written by the client itself, ``# '' for meta text, and ``! '' for non-irc protocol meta text. Meta text is information about the channel, such as users joining or leaving a channel, new topics, mode changes, etc. These characters are nearly always followed by a time stamp of the form ``hh:mm '', followed by the text itself. Date changes are written with the ``! '' meta text line. 
####/n/name
Reads return the name of the channel or the user. For the status directory ``0'' the string ``(netname)'' is returned, with netname from the command-line. Read-only. 
####/n/users
Only useful if the directory represents an irc channel. Reads return lines of the form ``+user'' and ``-user'' for joined, and parted or quit users. This allows clients to keep track of who is currently in a channel. Read-only.

EXAMPLE
-------
To run ircfs and export its files:
    # note: make sure the paths exist
    mount {ircfs -l $home/irclog/freenode -a net!irc.freenode.net!6667 
        -n ircfsuser freenode} /mnt/irc/freenode
        styxlisten -A net!*!6872 {export /mnt/irc}

SEE ALSO
--------
wm-irc(1).

SOURCE
------
/appl/cmd/ircfs.b

BUGS
----
Ircfs does not provide information to readers whether data has been read.

**IRCFS(4 )**
