# hslam-ws

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

./client -addr=:9999 -total=1000000 -clients=2
Summary:
	Clients:	2
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	33.68s
	Requests per second:	29688.85
	Fastest time for request:	0.03ms
	Average time per request:	0.06ms
	Slowest time for request:	4.61ms

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
	99.9%	time for request:	0.23ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	15200690.68 Byte/s (15.20 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	15200690.68 Byte/s (15.20 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=3
Summary:
	Clients:	3
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	22.29s
	Requests per second:	44865.37
	Fastest time for request:	0.03ms
	Average time per request:	0.06ms
	Slowest time for request:	9.00ms

Time:
	0.1%	time for request:	0.03ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.04ms
	25%	time for request:	0.05ms
	50%	time for request:	0.06ms
	75%	time for request:	0.07ms
	90%	time for request:	0.09ms
	95%	time for request:	0.10ms
	99%	time for request:	0.16ms
	99.9%	time for request:	0.38ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	22971067.04 Byte/s (22.97 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	22971067.04 Byte/s (22.97 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=4
Summary:
	Clients:	4
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	17.79s
	Requests per second:	56210.09
	Fastest time for request:	0.03ms
	Average time per request:	0.07ms
	Slowest time for request:	27.68ms

Time:
	0.1%	time for request:	0.03ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.04ms
	25%	time for request:	0.05ms
	50%	time for request:	0.06ms
	75%	time for request:	0.08ms
	90%	time for request:	0.10ms
	95%	time for request:	0.11ms
	99%	time for request:	0.18ms
	99.9%	time for request:	0.45ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	28779564.89 Byte/s (28.78 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	28779564.89 Byte/s (28.78 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```


netpoll
```
./client -addr=:9999 -total=1000000 -clients=1
Summary:
	Clients:	1
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	50.33s
	Requests per second:	19870.18
	Fastest time for request:	0.03ms
	Average time per request:	0.05ms
	Slowest time for request:	2.43ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.04ms
	25%	time for request:	0.04ms
	50%	time for request:	0.04ms
	75%	time for request:	0.05ms
	90%	time for request:	0.06ms
	95%	time for request:	0.06ms
	99%	time for request:	0.09ms
	99.9%	time for request:	0.23ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	10173529.82 Byte/s (10.17 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	10173529.82 Byte/s (10.17 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=2
Summary:
	Clients:	2
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	33.08s
	Requests per second:	30228.32
	Fastest time for request:	0.03ms
	Average time per request:	0.06ms
	Slowest time for request:	4.85ms

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
	99.9%	time for request:	0.23ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	15476897.31 Byte/s (15.48 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	15476897.31 Byte/s (15.48 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=3
Summary:
	Clients:	3
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	21.75s
	Requests per second:	45974.82
	Fastest time for request:	0.03ms
	Average time per request:	0.06ms
	Slowest time for request:	14.64ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.05ms
	25%	time for request:	0.05ms
	50%	time for request:	0.06ms
	75%	time for request:	0.07ms
	90%	time for request:	0.08ms
	95%	time for request:	0.09ms
	99%	time for request:	0.16ms
	99.9%	time for request:	0.39ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	23539107.58 Byte/s (23.54 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	23539107.58 Byte/s (23.54 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)

./client -addr=:9999 -total=1000000 -clients=4
Summary:
	Clients:	4
	Parallel calls per client:	1
	Total calls:	1000000
	Total time:	17.44s
	Requests per second:	57354.01
	Fastest time for request:	0.03ms
	Average time per request:	0.07ms
	Slowest time for request:	5.92ms

Time:
	0.1%	time for request:	0.04ms
	1%	time for request:	0.04ms
	5%	time for request:	0.04ms
	10%	time for request:	0.05ms
	25%	time for request:	0.05ms
	50%	time for request:	0.06ms
	75%	time for request:	0.07ms
	90%	time for request:	0.09ms
	95%	time for request:	0.10ms
	99%	time for request:	0.18ms
	99.9%	time for request:	0.41ms

Request:
	Total request body sizes:	512000000
	Average body size per request:	512.00 Byte
	Request rate per second:	29365252.29 Byte/s (29.37 MByte/s)

Response:
	Total response body sizes:	512000000
	Average body size per response:	512.00 Byte
	Response rate per second:	29365252.29 Byte/s (29.37 MByte/s)

Result:
	Response ok:	1000000 (100.00%)
	Errors:	0 (0.00%)
```
