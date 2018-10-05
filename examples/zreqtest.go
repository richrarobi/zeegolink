// zreqtest.go
// Rich Robinson
// Oct 2018
//
package main

import (
    "github.com/richrarobi/zeegolink"
    "os/signal"
    "syscall"
    "fmt"
    "os"
    "log"
    "time"
)

func main() {
// initialise getout
    running := true
    signalChannel := make(chan os.Signal, 2)
    signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
    go func() {
        sig := <-signalChannel
        switch sig {
        case os.Interrupt:
            log.Println("Stopping on Interrupt")
            running = false
            return
        case syscall.SIGTERM:
            log.Println("Stopping on Terminate")
            running = false
            return
        }
    }()

    srva := "tcp://slither.local:5555"
    sock := zeegolink.New(srva)
// now use the python version on port 5554
    srvb := "tcp://slither.local:5554"
    sockb := zeegolink.New(srvb)
    srvc := "tcp://c.local:5554"
// note port 5554, system c is a raspberry pi 3 using python3 zreply.py
    sockc := zeegolink.New(srvc)
    
    for running {
        fmt.Printf("Received 1 %s\n", zeegolink.Request(sock, "sysType", "" ))
        fmt.Printf("Received 2 %s\n", zeegolink.Request(sockb, "sysType", "" ))
        fmt.Printf("Received 3 %s\n", zeegolink.Request(sock, "cpuTemp", "" ))
// note that zreply.go on the pi is as on the pc
        fmt.Printf("Received 4 %s\n", zeegolink.Request(sockc, "sysType", "" ))
        fmt.Printf("Received 5 %s\n", zeegolink.Request(sockc, "cpuTemp", "" ))
        fmt.Printf("Received 6 %s\n", zeegolink.Request(sock, "adds", "44" ))
        fmt.Printf("Received 7 %s\n", zeegolink.Request(sock, "adds", "44", "33" ))
        fmt.Printf("Received 8 %s\n", zeegolink.Request(sock, "adds", "44", "33", "5.08" ))
        fmt.Printf("Received 9 %s\n", zeegolink.Request(sock, "adds", "44", "33", "fred" ))
        time.Sleep(2 * time.Second)
    }
    
    zeegolink.Stop(sock)
    zeegolink.Stop(sockc)
}
