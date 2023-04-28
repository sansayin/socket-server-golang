# socket-server-go
This is just reinvent the wheel, keep in mind that golang already has epoll support natively.
## Use buffered channel to limit number of go routines
## Apache Bench Stress Test(nor for long-connection)
```
ab -c 1000 -n 2000 -m GET http://localhost:9988
```

## Tesing Tool
```
sudo apt install hping3
```

### TCP Stress Test
```

```

### UDP Stress Test
```
sudo hping3 --udp [UDP SERVER IP] -p 9988 --rand-source --flood 
```
### Pprof

```
#Run program with option "-pprof=true"
go tool pprof -http=:8001 http://127.0.0.1:6060/debug/pprof/profilee 
```
### CURL Flood
```
#!/bin/bash
target=${1:-http://localhost:9988}
while true # loop forever, until ctrl+c pressed.
do
	for i in $(seq 100) # perfrom the inner command 100 times.
	do
		curl $target > /dev/null & # send out a curl request, the & indicates not to wait for the response.
	done

	wait # after 100 requests are sent out, wait for their processes to finish before the next iteration.
done
```
### With Metrics
```
docker-compose -f ./docker-compose.yml up

```
