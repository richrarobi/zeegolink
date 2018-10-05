// zreply.go
// Rich Robinson
// Oct 2018
//
package main

import (
        zmq "github.com/pebbe/zmq4"
        "encoding/json"
        "fmt"
        "log"
        "time"
        "strings"
        "os/exec"
        "context"
        "bufio"
        "strconv"
)

type Reqst struct {
    Cmnd  string  `json:cmnd`
    Args  []string `json:args`
    Rslt  string  `json:"rslt,omitempty"`
}

func worker_routine() {
        receiver, _ := zmq.NewSocket(zmq.REP)
        defer receiver.Close()
        
        receiver.Connect("inproc://workers")
        reqst := Reqst{}
        for {
                var msg []byte
                var e error
                msg, e = receiver.RecvBytes(0)
                if e != nil {
                        fmt.Println("Error in Receive")
                        break
                }
                err := json.Unmarshal(msg, &reqst)
                if err != nil {
                        fmt.Println("error:", err)
                }
//                fmt.Printf("Recd %s\n",  reqst)
//                fmt.Println("Received Command :", reqst.Cmnd)
//                fmt.Println("Received Arguments :", reqst.Args)
                
// call the function and loadup the result
                if reqst.Cmnd != "" {
                    switch reqst.Cmnd {
                    case "cpuTemp":
                        reqst.Rslt = cpuTemp()
                    case "sysType":
                        reqst.Rslt = sysType()
                    case "adds":
                        args := reqst.Args
                        reqst.Rslt = adds(args ...)
                    default:
                        reqst.Rslt = none()
                    }
                }

//  Send reply back to client
                b, err := json.Marshal(&reqst)
                if err != nil {
                    b = []byte("error in json Marshal")
                    }
                receiver.Send(string(b), 0)
        }
}

func main() {
//  Socket to talk to clients
        clients, _ := zmq.NewSocket(zmq.ROUTER)
        defer clients.Close()
        clients.Bind("tcp://*:5555")

//  Socket to talk to workers
        workers, _ := zmq.NewSocket(zmq.DEALER)
        defer workers.Close()
        workers.Bind("inproc://workers")

//  Launch pool of worker goroutines
        for thread_nbr := 0; thread_nbr < 2; thread_nbr++ {
                go worker_routine()
        }
//  Connect work threads to client threads via a queue proxy
        err := zmq.Proxy(clients, workers, nil)
        log.Println("Proxy interrupted:", err)
}

func adds ( args... string ) ( res string) {
    var z int
    for i := 0 ; i < len(args); i++ {
            x, err := strconv.Atoi(args[i])
            if err == nil {
                z = x + z
            }
        }
    res = strconv.Itoa(z)
    return res
}

func cpuTemp()(res string) {
    var x string
    if strings.Contains( sysType(), "ARM" ) {
        x = exeCmd( "/opt/vc/bin/vcgencmd", "measure_temp")
    } else {
        x = "cpuTemp: SysType Not ARM"
    }
    return x
}

func sysType()(res string) {
    var x string
    scanner := bufio.NewScanner(strings.NewReader(exeCmd( "cat", "/proc/cpuinfo")))
    for scanner.Scan() {
        if strings.Contains( scanner.Text(), "model name") {
            x = scanner.Text() [ 13:len( scanner.Text() ) ]
            return x
//            fmt.Println( x )
            break
        }
    }
     return "SysType Not Found"
}

func exeCmd(command string, args... string) (res string) {
// e.g.   log.Println( exeCmd("ls", "..", "-l",) )
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    cmd := exec.CommandContext(ctx, command, args... )
    out, err := cmd.Output()

    if ctx.Err() == context.DeadlineExceeded {
        fmt.Println("exeCmd: Command timed out")
        return
    }
    if err != nil {
        fmt.Println("Non-zero exit code:", err)
        return "exeCmd: Error in external command"
        }
    return string(out)
}

func none () (res string) {
    return ("Function Not Found")
}
