# zeegolink
go and python intercommunication using zeromq4

reply servers written in go and python3 allow requests in go or pyhon to run functions in either language on locally connected systems.

Python is the "major" language used on Raspberry Pi systems, and so a python interface is useful. zreply.py uses port 5554. zreply.go uses port 5555.
zeegolink can send requests to either port, so a Pi running BOTH zreply servers can be called from go or python in a single program, as can other systems running zreply.

the go version of reply uses internal functions, whereas the python reply uses a small library in zlocal.py - see the example in zreqtest.go, and the zbeat.py.

NOTE that all json parameters are string, and variables such as integers need to be converted for use. Also note that there are dummy parameters in some of the python functions.


Output from example zreqtest.go follows:

./zreqtest
Received 1 Intel(R) Core(TM) i5-4670K CPU @ 3.40GHz
Received 2 Intel(R) Core(TM) i5-4670K CPU @ 3.40GHz
Received 3 cpuTemp: SysType Not ARM
Received 4 ARMv7 Processor rev 4 (v7l)
Received 5 temp=40.8'C
Received 6 44
Received 7 77
Received 8 77
Received 9 77
Received 1 Intel(R) Core(TM) i5-4670K CPU @ 3.40GHz
Received 2 Intel(R) Core(TM) i5-4670K CPU @ 3.40GHz
Received 3 cpuTemp: SysType Not ARM
Received 4 ARMv7 Processor rev 4 (v7l)
Received 5 temp=41.3'C
Received 6 44
Received 7 77
Received 8 77
Received 9 77
^C2018/10/05 14:15:34 Stopping on Interrupt
