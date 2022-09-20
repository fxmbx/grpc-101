🚀🚀
Http2bb
Types of grpc API

- uniary
  service GreetService{
  rpc Greet(GreetingRequest) returns (GreetResonse) {};

- streaming server
  rpc GreetManyTimes(GreetManayTimesRequest) returns (stream GreetManyTimeRespnse){};

- streaming client
  rpc LongGreet(stream LongGreetRequest) returns (LongGreetResponse) {};

- bidirectional streaming
  rpc GreatEveryone(stream GreetEveryOneRequest) returns (stream GreetEveryOneResponse) {};

}

# scalability

grpc is asynchronous by default . each server can serve millions of request in parallel
client side load balancing

# security

ssl

grpc-GO
https://github.com/grpc/grpc-go
