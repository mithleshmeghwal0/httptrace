httptrace

The main idea is the ability to track a set of events that occur within a lifecycle of a http request. For example, dns resolution, tcp connection creation, data written to the tcp connection, data received from the tcp connection and so on.

client request lifecycle
% chmod +x runner.sh                              
% ./runner.sh 
starting to create conn  www.example.com:443
starting to look up dns {"Host":"www.example.com"}
done looking up dns {"Addrs":[{"IP":"93.184.216.34","Zone":""},{"IP":"2606:2800:220:1:248:1893:25c8:1946","Zone":""}],"Err":null,"Coalesced":false}
starting tcp connection tcp 93.184.216.34:443
tcp connection created tcp 93.184.216.34:443 <nil>
connection established {"Conn":{},"Reused":false,"WasIdle":false,"IdleTime":0}
