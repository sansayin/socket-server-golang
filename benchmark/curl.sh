#!/bin/bash
target=${1:-http://localhost:8899/post}
while true # loop forever, until ctrl+c pressed.
do
	for i in $(seq 10) # perfrom the inner command 100 times.
	do
		curl -X POST  $target -d "sad" -v > /dev/null & # send out a curl request, the & indicates not to wait for the response.
	done

	wait # after 100 requests are sent out, wait for their processes to finish before the next iteration.
done

