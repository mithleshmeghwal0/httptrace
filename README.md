httptrace

The main idea is the ability to track a set of events that occur within a lifecycle of a http request. For example, dns resolution, tcp connection creation, data written to the tcp connection, data received from the tcp connection and so on.

client request lifecycle
% chmod +x runner.sh                              
% ./runner.sh
starting to create conn  www.bing.com:443
starting to look up dns {"Host":"www.bing.com"}
done looking up dns {"Addrs":[{"IP":"204.79.197.200","Zone":""},{"IP":"13.107.21.200","Zone":""},{"IP":"2620:1ec:c11::200","Zone":""}],"Err":null,"Coalesced":false}
starting tcp connection {"addr":"204.79.197.200:443","network":"tcp"}
tcp connection created {"addr":"204.79.197.200:443","err":null,"network":"tcp"}
connection established {"Conn":{},"Reused":false,"WasIdle":false,"IdleTime":0}

starting to create conn  www.bing.com:443
connection established {"Conn":{},"Reused":true,"WasIdle":false,"IdleTime":0}

starting to create conn  www.bing.com:443
connection established {"Conn":{},"Reused":true,"WasIdle":false,"IdleTime":0}

starting to create conn  www.google.com:443
starting to look up dns {"Host":"www.google.com"}
done looking up dns {"Addrs":[{"IP":"172.217.25.196","Zone":""},{"IP":"2404:6800:4001:810::2004","Zone":""}],"Err":null,"Coalesced":false}
starting tcp connection {"addr":"172.217.25.196:443","network":"tcp"}
tcp connection created {"addr":"172.217.25.196:443","err":null,"network":"tcp"}
connection established {"Conn":{},"Reused":false,"WasIdle":false,"IdleTime":0}

starting to create conn  www.google.com:443
connection established {"Conn":{},"Reused":true,"WasIdle":false,"IdleTime":0}

starting to create conn  www.google.com:443
connection established {"Conn":{},"Reused":true,"WasIdle":false,"IdleTime":0}

starting to create conn  www.bing.com:443
connection established {"Conn":{},"Reused":true,"WasIdle":true,"IdleTime":542637868}

starting to create conn  www.bing.com:443
connection established {"Conn":{},"Reused":true,"WasIdle":false,"IdleTime":0}
