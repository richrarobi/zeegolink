#!/usr/bin/python3
# zlocal.py
#
from time import sleep
import subprocess

def sysType():
    tmp = exeCmd("cat /proc/cpuinfo")
    for line in tmp.splitlines():
        if "model name" in line:
            x = line[13: ]
    return x

def cpuTemp():
    if "ARM" in sysType():
        import subprocess
        tmp = exeCmd("/opt/vc/bin/vcgencmd measure_temp")
        return tmp
    else:
        return "notARM"

def exeCmd(cmnd):
# internal use only
    from subprocess import Popen, STDOUT, PIPE
    import shlex
    try:
        args = shlex.split(cmnd)
        p = Popen(args, stdout=PIPE)
        out, err = p.communicate(timeout=30)
        out = out.decode("utf-8").rstrip("\n")
    except:
        p.kill()
    return out

def datim():
    import datetime
    now = datetime.datetime.now()
    x = str(now)
    y, z = x.split(" ")
    y, m, d = y.split("-")
    hh, mm, ss = z.split(":")
    ss, dec = ss.split(".")
    stmp = "{}-{}-{} {}:{}:{}".format(y, m, d, hh, mm, ss)
    return stmp

if __name__ == "__main__":
    print(sysType())
    print(cpuTemp())
    print(datim())
