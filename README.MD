# CS425-MP0 Event Logger

**To start the logger,** 

*if on a Linux platform,* run: 

./mp0-go/logger [port]


*if on a Windows platform,* run: 

go run ./mp0-go/logger.go [port]


The port argument is where you specify the port to listen to, default is 9999. You can also specify a log file name after the port argument; otherwise a log file named with the current timestamp will be generated (for analyzing bandwidth and delays).

**To start the event generator and node,** 

*if on a Linux platform,* run: 

python -u ./mp0-py/generator.py [rate] | ./mp0-go/node [node_name] [address] [port]


*if on a Windows platform,* run: 

python -u ./mp0-py/generator.py [rate] | go run ./mp0-go/node.go [node_name] [address] [port]


The rate argument works as explained in the MP prompt. The node_name specifies the name of the node which will be included in the event messages. The address and port arguments should be provided together. Default address and port is 127.0.0.1:9999 which is for local testing only.
