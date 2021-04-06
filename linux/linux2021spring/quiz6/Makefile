CC = gcc
CFLAGS = -g
BIN = twoSum-hash twoSum-bt

all: $(BIN)

%: %.c
	$(CC) $(CFLAGS) -o $@ $<

.PHONY: clean

clean:
	-rm -f $(BIN) *.o
