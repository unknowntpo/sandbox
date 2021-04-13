#!/bin/bash

for i in `seq 0 1 200`; do \
		echo 3 > /proc/sys/vm/drop_caches ;\
		printf "%d," $$i;\
		./alignment; \
	done > clock_gettime.txt	
