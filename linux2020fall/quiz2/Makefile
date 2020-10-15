CC := gcc
CFLAGS := -g -Wall

targets := isascii hexchar2val fastdiv
all: $(targets)

isascii: isascii.o

hexchar2val: hexchar2val.o

fastdiv: fastdiv.o

.PHONY: clean
clean:
	-rm -f *.o $(targets) 
