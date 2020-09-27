# golang-ws

```
./client -addr=:9999 -total=1000000 -clients=1
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	67.35s
	Requests per second:	14848.42
	Fastest time for request:	0.04ms
	Average time per request:	0.06ms
	Slowest time for request:	7.27ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.05ms
	5%	time for request:	0.05ms
	10%	time for request:	0.05ms
	25%	time for request:	0.05ms
	50%	time for request:	0.06ms
	75%	time for request:	0.07ms
	90%	time for request:	0.08ms
	95%	time for request:	0.09ms
	99%	time for request:	0.12ms
	99.9%	time for request:	0.30ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	7602388.66 Byte/s (7.60 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	7602388.66 Byte/s (7.60 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=2
Summary:
	Clients:	2
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	39.18s
	Requests per second:	25524.78
	Fastest time for request:	0.03ms
	Average time per request:	0.07ms
	Slowest time for request:	3.41ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.05ms
	5%	time for request:	0.05ms
	10%	time for request:	0.06ms
	25%	time for request:	0.06ms
	50%	time for request:	0.07ms
	75%	time for request:	0.08ms
	90%	time for request:	0.10ms
	95%	time for request:	0.11ms
	99%	time for request:	0.15ms
	99.9%	time for request:	0.30ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	13068685.98 Byte/s (13.07 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	13068685.98 Byte/s (13.07 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=3
Summary:
	Clients:	3
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	25.99s
	Requests per second:	38480.85
	Fastest time for request:	0.03ms
	Average time per request:	0.08ms
	Slowest time for request:	13.48ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.05ms
	10%	time for request:	0.05ms
	25%	time for request:	0.06ms
	50%	time for request:	0.07ms
	75%	time for request:	0.08ms
	90%	time for request:	0.10ms
	95%	time for request:	0.12ms
	99%	time for request:	0.20ms
	99.9%	time for request:	0.44ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	19702195.92 Byte/s (19.70 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	19702195.92 Byte/s (19.70 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
./client -addr=:9999 -total=1000000 -clients=4
Summary:
	Clients:	4
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	21.07s
	Requests per second:	47459.35
	Fastest time for request:	0.03ms
	Average time per request:	0.08ms
	Slowest time for request:	3.70ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.05ms
	10%	time for request:	0.05ms
	25%	time for request:	0.06ms
	50%	time for request:	0.07ms
	75%	time for request:	0.09ms
	90%	time for request:	0.11ms
	95%	time for request:	0.13ms
	99%	time for request:	0.23ms
	99.9%	time for request:	0.52ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	24299185.62 Byte/s (24.30 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	24299185.62 Byte/s (24.30 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)

```
