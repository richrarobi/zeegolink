// zeegolink.go
// Rich Robinson
// Oct 2018
//
package zeegolink

import (
    zmq "github.com/pebbe/zmq4"
    "encoding/json"
    "time"
)

type Reqst struct {
    Cmnd  string  `json:cmnd`
    Args  []string `json:args`
    Rslt  string  `json:"rslt,omitempty"`
}

func New( srv string )( sock *zmq.Socket ){
    sock, _ = zmq.NewSocket(zmq.REQ)
    sock.SetLinger(0)
    sock.SetRcvtimeo(time.Second * 3)
    sock.Connect(srv)
//    fmt.Printf("Socket: %s\n",  sock)
    return sock
}

func Stop( sock *zmq.Socket ) (res string) {
    res = "closed"
    sock.Close()
    return res
}

func Request( sock *zmq.Socket, command string, args... string) (res string) {
    var msg []byte
    var e error

    reqst := Reqst{}
    reqst.Cmnd = command
    reqst.Args = args
//    fmt.Printf("Sending: %s\n",  reqst)
    b, err := json.Marshal(reqst)
    if err != nil {
        b = []byte("error in json Marshal")
        }
//    fmt.Println(string(b))
    sock.Send(string(b), zmq.DONTWAIT)
    var reply string
    msg, e = sock.RecvBytes(0)
    if e == nil {
// received data
        err = json.Unmarshal(msg, &reqst)
        if err != nil {
            reply = "Unmarshal Error"
        } else {
            reply = reqst.Rslt
        }
    } else { 
        reply = "Lost Connection"
        }
    return reply
}

