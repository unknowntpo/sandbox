CC = gcc
CFLAGS = -pthread -Wall
PROGS = loop

all: $(PROGS)

loop: loop.c
	$(CC) $(CFLAGS) -o $@ $<

.PHONY: clean
clean:
	-rm -f $(PROGS)
