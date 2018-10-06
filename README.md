# zeegolink
go and python intercommunication using zeromq4 ("github.com/pebbe/zmq4")

go get github.com/richrarobi/zeegolink

reply servers written in go and python3 allow requests in go or python3 to run functions
in either language on any (linux based?) locally connected systems.

Python is the "major" language used on Raspberry Pi systems, and so a python interface
is useful. zreply.py uses port 5554. zreply.go uses port 5555.
zeegolink can send requests to either port, so a Pi running BOTH zreply servers can
be called from go or python, as can other systems running zreply. Multiple systems may
be interconnected

the go version of reply uses internal functions, whereas the python reply uses a small
library in zlocal.py - see the example in zreqtest.go, and the zbeat.py.

NOTE that all json parameters are string, and variables such as integers need to
be converted for use. Also note that there are dummy parameters in some of the python
functions.
In the python folder, a file zlocal_c.py is the zlocal.py on my pi system c. It drives a Pimoroni blinkt. I have already published a go version of this, that I could have used.

Currently having an issue with zmq when a reply server is stopped, the request will not then reconnect after restart. This doesn't seem to happen with nanomsg- mangos.... 

Output from example zreqtest.go is in output.txt

To run the python3 programs, you will need to :-

sudo apt-get install python-zmq

sudo pip3 install pyzmq
