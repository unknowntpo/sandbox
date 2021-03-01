CC = gcc
CFLAGS = -Wall -O2 -g -fno-omit-frame-pointer
BS = 8388608 # 8 * 1024 * 1024
DATA = input.dat

all: matrix-v1
	
matrix-v1: matrix-v1.c
	gcc $< -o $@ $(CFLAGS) 

.PHONY: gendata clean

gendata: matrix-v1
	@echo "Generating input data..."
	dd if=/dev/urandom of=$(DATA) bs=$(BS) count=1

clean:
	-rm $(DATA) matrix-v1
