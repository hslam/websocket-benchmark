# gorilla-ws

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

./client -addr=:9999 -total=1000000 -clients=2
Summary:
	Clients:	2
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	36.08s
	Requests per second:	27718.85
	Fastest time for request:	0.03ms
	Average time per request:	0.07ms
	Slowest time for request:	3.77ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.05ms
	10%	time for request:	0.05ms
	25%	time for request:	0.06ms
	50%	time for request:	0.07ms
	75%	time for request:	0.08ms
	90%	time for request:	0.09ms
	95%	time for request:	0.10ms
	99%	time for request:	0.14ms
	99.9%	time for request:	0.28ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	14192050.71 Byte/s (14.19 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	14192050.71 Byte/s (14.19 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=3
Summary:
	Clients:	3
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	23.67s
	Requests per second:	42249.12
	Fastest time for request:	0.03ms
	Average time per request:	0.07ms
	Slowest time for request:	4.63ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.05ms
	25%	time for request:	0.05ms
	50%	time for request:	0.06ms
	75%	time for request:	0.08ms
	90%	time for request:	0.09ms
	95%	time for request:	0.11ms
	99%	time for request:	0.17ms
	99.9%	time for request:	0.39ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	21631551.31 Byte/s (21.63 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	21631551.31 Byte/s (21.63 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=4
Summary:
	Clients:	4
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	19.27s
	Requests per second:	51892.06
	Fastest time for request:	0.03ms
	Average time per request:	0.08ms
	Slowest time for request:	26.38ms

Time:
	0.1%	time for request:	0.03ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.05ms
	25%	time for request:	0.06ms
	50%	time for request:	0.07ms
	75%	time for request:	0.08ms
	90%	time for request:	0.10ms
	95%	time for request:	0.12ms
	99%	time for request:	0.21ms
	99.9%	time for request:	0.49ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	26568734.59 Byte/s (26.57 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	26568734.59 Byte/s (26.57 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)

```

