# websocket-benchmark

## hslam-ws
```
./client -addr=:9999 -total=1000000 -clients=1
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	61.96s
	Requests per second:	16139.43
	Fastest time for request:	0.03ms
	Average time per request:	0.06ms
	Slowest time for request:	8.34ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.05ms
	10%	time for request:	0.05ms
	25%	time for request:	0.05ms
	50%	time for request:	0.06ms
	75%	time for request:	0.06ms
	90%	time for request:	0.07ms
	95%	time for request:	0.08ms
	99%	time for request:	0.11ms
	99.9%	time for request:	0.24ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	8263390.40 Byte/s (8.26 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	8263390.40 Byte/s (8.26 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=64
Summary:
	Clients:	64
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	10.05s
	Requests per second:	99479.84
	Fastest time for request:	0.03ms
	Average time per request:	0.64ms
	Slowest time for request:	64.52ms

Time:
	0.1%	time for request:	0.05ms
	1%	time for request:	0.09ms
	5%	time for request:	0.31ms
	10%	time for request:	0.41ms
	25%	time for request:	0.52ms
	50%	time for request:	0.59ms
	75%	time for request:	0.67ms
	90%	time for request:	0.80ms
	95%	time for request:	0.97ms
	99%	time for request:	2.01ms
	99.9%	time for request:	8.06ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	50933677.98 Byte/s (50.93 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	50933677.98 Byte/s (50.93 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

## golang-ws
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

./client -addr=:9999 -total=1000000 -clients=64
Summary:
	Clients:	64
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	12.38s
	Requests per second:	80776.60
	Fastest time for request:	0.03ms
	Average time per request:	0.79ms
	Slowest time for request:	47.47ms

Time:
	0.1%	time for request:	0.05ms
	1%	time for request:	0.06ms
	5%	time for request:	0.12ms
	10%	time for request:	0.23ms
	25%	time for request:	0.44ms
	50%	time for request:	0.63ms
	75%	time for request:	0.83ms
	90%	time for request:	1.19ms
	95%	time for request:	1.88ms
	99%	time for request:	4.81ms
	99.9%	time for request:	12.36ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	41357618.76 Byte/s (41.36 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	41357618.76 Byte/s (41.36 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

## gorilla-ws
```
./client -addr=:9999 -total=1000000 -clients=1
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	64.68s
	Requests per second:	15460.79
	Fastest time for request:	0.04ms
	Average time per request:	0.06ms
	Slowest time for request:	3.92ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.05ms
	10%	time for request:	0.05ms
	25%	time for request:	0.05ms
	50%	time for request:	0.06ms
	75%	time for request:	0.07ms
	90%	time for request:	0.08ms
	95%	time for request:	0.09ms
	99%	time for request:	0.12ms
	99.9%	time for request:	0.29ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	7915925.70 Byte/s (7.92 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	7915925.70 Byte/s (7.92 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=64
Summary:
	Clients:	64
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	10.52s
	Requests per second:	95078.60
	Fastest time for request:	0.03ms
	Average time per request:	0.67ms
	Slowest time for request:	67.22ms

Time:
	0.1%	time for request:	0.05ms
	1%	time for request:	0.11ms
	5%	time for request:	0.32ms
	10%	time for request:	0.41ms
	25%	time for request:	0.52ms
	50%	time for request:	0.60ms
	75%	time for request:	0.69ms
	90%	time for request:	0.87ms
	95%	time for request:	1.11ms
	99%	time for request:	2.42ms
	99.9%	time for request:	8.67ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	48680242.50 Byte/s (48.68 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	48680242.50 Byte/s (48.68 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```

