#!/usr/bin/python3
# Filename: zreply.py
#
from time import sleep
import threading
import zmq
import signal
import sys
import zlocal
import importlib

def worker_routine(worker_url, context=None):
    context = context or zmq.Context.instance()
# Socket to talk to dispatcher
    socket = context.socket(zmq.REP)
    socket.connect(worker_url)

    while True:
        data = socket.recv_json()
        func = str(data['C'])
        args = data['A']
# reload the zlocal library if changed
        importlib.reload(zlocal)
        module = importlib.import_module('zlocal')
        try:
# get the function if callable
            fn = getattr(module, func)
            if callable(fn):
#                print("Calling fn: {}, args: {} ".format(fn, *args))
                x = fn(*args)
                reply = {'R': x }
            else:
                reply = {'R': "not available" }
        except:
            reply = {'R' : "Function Error"}
        socket.send_json(reply)

if __name__ == "__main__":
    try:
        url_worker = "inproc://workers"
        url_client = "tcp://*:5554"
        context = zmq.Context.instance()

# Socket to talk to clients
        clients = context.socket(zmq.ROUTER)
        clients.bind(url_client)
# Socket to talk to workers
        workers = context.socket(zmq.DEALER)
        workers.bind(url_worker)

# Launch pool of worker threads
        for i in range(3):
            thread = threading.Thread(target=worker_routine, args=(url_worker,))
            thread.start()

        zmq.proxy(clients, workers)
    
    except KeyboardInterrupt:
        print("Stopping!")
# We should never get here but clean up anyhow
        clients.close()
        workers.close()
        context.term()
        sleep(2)
