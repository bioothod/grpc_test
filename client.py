import argparse

import grpc
import grtest_pb2
import grtest_pb2_grpc

import math

import datetime as dt

parser = argparse.ArgumentParser(description='gRPC test.')
parser.add_argument('--remote', default='localhost:12345')
parser.add_argument('--ping', default='ping data')
parser.add_argument('--aux', default='ping aux')

args = parser.parse_args()

channel = grpc.insecure_channel(args.remote)
stub = grtest_pb2_grpc.TestServiceStub(channel)

req = grtest_pb2.Ping(ping=args.ping, aux=args.aux)
reply = stub.PingRequest(req)
print "ping reply: %s:%s" % (reply.pong, reply.aux)

prev_time = dt.datetime.now()
for r in stub.Stream(req):
    t = dt.datetime.now()
    if t > prev_time + dt.timedelta(seconds=10):
        print "stream reply: %s:%s" % (r.pong, r.aux)
        prev_time = t
