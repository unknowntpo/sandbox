CC = gcc
CFLAGS = -g -Wall
TARGET = test_getopt

all: $(TARGET)

test_getopt: test_getopt.c
	$(CC) $(CFLAGS) -o $@ $<

.PHONY: check
check:
	./test_getopt -a -emy_args -c dddd  123 456 7890 -b -f -d

.PHONY: clean
clean:
	-rm -rf $(TARGET) *.dSYM
