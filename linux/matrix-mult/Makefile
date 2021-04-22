CC = gcc
CFLAGS = -g -Wall
BIN = mm

all: $(BIN)

$(BIN): mm.c
	$(CC) $(CFLAGS) -o $@ $< 
format:
	@clang-format -i mm.c

.PHONY: clean
clean:
	-rm -f $(BIN)
