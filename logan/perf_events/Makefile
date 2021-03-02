CC = gcc
CFLAGS = -Wall -O2 -g -fno-omit-frame-pointer
BS = 8388608 # 8 * 1024 * 1024
DATA = input.dat
TARGET = matrix-v1 matrix-v2

all: $(TARGET)

%: %.c
	gcc $< -o $@ $(CFLAGS) 

.PHONY: gendata clean

gendata: matrix-v1
	@echo "Generating input data..."
	dd if=/dev/urandom of=$(DATA) bs=$(BS) count=1

clean:
	-rm $(DATA) $(TARGET)
