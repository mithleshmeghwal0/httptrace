httptrace

The main idea is the ability to track a set of events that occur within a lifecycle of a http request. For example, dns resolution, tcp connection creation, data written to the tcp connection, data received from the tcp connection and so on.

client request lifecycle
% chmod +x runner.sh                              
% ./runner.sh

developer@Developers-MacBook-Air httptrace % ./runner.sh
starting to create conn  www.bing.com:443
starting to look up dns {"Host":"www.bing.com"}
done looking up dns {"Addrs":[{"IP":"204.79.197.200","Zone":""},{"IP":"13.107.21.200","Zone":""},{"IP":"2620:1ec:c11::200","Zone":""}],"Err":null,"Coalesced":false}
starting tcp connection {"addr":"204.79.197.200:443","network":"tcp"}
tcp connection created {"addr":"204.79.197.200:443","err":null,"network":"tcp"}
connection established {"Conn":{},"Reused":false,"WasIdle":false,"IdleTime":0}

body closed
starting to create conn  www.bing.com:443
connection established {"Conn":{},"Reused":true,"WasIdle":true,"IdleTime":1868251459}

body closed
starting to create conn  www.bing.com:443
connection established {"Conn":{},"Reused":true,"WasIdle":true,"IdleTime":1923763459}

body closed
starting to create conn  www.google.com:443
starting to look up dns {"Host":"www.google.com"}
done looking up dns {"Addrs":[{"IP":"142.250.199.36","Zone":""},{"IP":"2404:6800:4001:808::2004","Zone":""}],"Err":null,"Coalesced":false}
starting tcp connection {"addr":"142.250.199.36:443","network":"tcp"}
tcp connection created {"addr":"142.250.199.36:443","err":null,"network":"tcp"}
connection established {"Conn":{},"Reused":false,"WasIdle":false,"IdleTime":0}

body closed
starting to create conn  www.google.com:443
connection established {"Conn":{},"Reused":true,"WasIdle":true,"IdleTime":1680313625}

body closed
starting to create conn  www.google.com:443
connection established {"Conn":{},"Reused":true,"WasIdle":true,"IdleTime":1692113542}

body closed
starting to create conn  www.bing.com:443
connection established {"Conn":{},"Reused":true,"WasIdle":true,"IdleTime":8235616584}

body closed
starting to create conn  www.bing.com:443
connection established {"Conn":{},"Reused":true,"WasIdle":true,"IdleTime":1839693708}

body closed
developer@Developers-MacBook-Air httptrace % 
