#!/usr/bin/python3
# zbeat.py
#
from time import sleep
from zreq import ZReqst

if __name__ == "__main__":
    zp = ZReqst("c.local:5554")
    led = 3

    try:
        while True:
            print("sysType : {}".format(zp.req("sysType", "")))
            print("getTemp : {}".format(zp.req("cpuTemp", "")))
            
            print("system c : invalid : {}".format(zp.req("invalid")))
            zp.req("ledSet",str(led), "64", "0", "0", "0.2")
            print(" system c : led {}: {}".format(led, zp.req("ledGet",str(led))))
            sleep(1)
            
            zp.req("ledSet",str(led), "0", "64", "0", "0.2")
            print(" system c : led {}: {}".format(led, zp.req("ledGet",str(led))))
            sleep(1)
            
            zp.req("ledSet",str(led), "0", "0", "64", "0.2")
            print(" system c : led {}: {}".format(led, zp.req("ledGet",str(led))))
            sleep(1)
 
    except KeyboardInterrupt:
        print("Stopping!")
        sleep(2)
        zp.req("ledSet",str(led), "0", "0", "0", "0.0" )
        sleep(1)
