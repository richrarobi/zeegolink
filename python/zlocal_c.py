#!/usr/bin/python3
# zlocal.py
from time import sleep
import subprocess
import blinkt
import Adafruit_BMP.BMP085 as BMP085

def sysType(a):
    tmp = exeCmd("cat /proc/cpuinfo")
    for line in tmp.splitlines():
        if "model name" in line:
            x = line[13: ]
    return x

def cpuTemp(a):
    if "ARM" in sysType(""):
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

def datim(a):
    import datetime
    now = datetime.datetime.now()
    x = str(now)
    y, z = x.split(" ")
    y, m, d = y.split("-")
    hh, mm, ss = z.split(":")
    ss, dec = ss.split(".")
    stmp = "{}{}{}+{}{}{}".format(y, m, d, hh, mm, ss)
    return stmp

def ledClrAll():
    for x in range(8):
        blinkt.set_pixel(x,0, 0, 0, 0.0)
    blinkt.show()
    return "done"

def ledSet(p, r, g, b, i):
    blinkt.set_pixel(int(p), int(r), int(g), int(b), float(i))
    blinkt.show()
    return "done"

def ledGet(p):
    return  blinkt.get_pixel(int(p))


def getBMP():
    try:
#        import BMP085
        sensor = BMP085.BMP085()
        t = int(sensor.read_temperature())
        p = int(sensor.read_pressure()/100)
#        a = round(sensor.read_altitude(),2)
#        slp = int(sensor.read_sealevel_pressure(260.0)/100)
#    print('Temp = {0:0.2f} *C'.format(t))
#    print('Pressure = {0:0.2f} Pa'.format(p))
#    print('Altitude = {0:0.2f} m'.format(a))
#    print('Sealevel Pressure = {0:0.2f} Pa'.format(slp))
#    reply = {"temp": t, "press": p, "Sea": slp}
        reply = "Temp: {}, Press: {}".format(t, p)
    except:
        reply = "Error in getBMP"
    return reply


def getBaro():
    try:
        sensor = BMP085.BMP085()
        p = int(sensor.read_pressure()/100)
        reply = "Baro : {}".format(p)
    except:
        reply = "Error in getBaro"
    return reply

    
def getRoomTemp():
    try:
        sensor = BMP085.BMP085()
        t = int(sensor.read_temperature())
        reply = "RoomTemp : {}".format(t)
    except:
        reply = "Error in getRoomTemp"
    return reply
    

if __name__ == "__main__":
# test only....
    print(sysType(""))
    print(cpuTemp(""))
    print(datim(""))
#    print(getRoomTemp())
#    print(getBaro())   
    try:
        while True:
            for x in range(8):
                ledSet(str(x), "64", "0", "0", "0.2")
                print(ledGet(str(x)))
                sleep(0.1)
                ledSet(str(x), "0", "0", "0", "0.0")
                print(ledGet(str(x)))
                sleep(0.1)
                
    except KeyboardInterrupt:
        print("Stopping")
        ledClrAll()

