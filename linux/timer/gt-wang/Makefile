CC = gcc
CFLAGS = -g -Wall
TARGET = clock_monotonic

clock_monotonic: clock_monotonic.c
	$(CC) $(CFLAGS) -o $@ $<

.PHONY: clean
clean:
	-rm -rf $(TARGET) *.dSYM
