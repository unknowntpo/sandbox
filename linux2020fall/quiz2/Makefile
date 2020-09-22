CC := gcc
CFLAGS := -g -Werror -Wall

targets := isascii hexchar2val
all: $(targets)

isascii: isascii.o

hexchar2val: hexchar2val.o

.PHONY: clean
clean:
	-rm -f *.o $(targets) 
