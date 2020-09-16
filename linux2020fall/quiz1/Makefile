targets := ll
CC := gcc
CFLAGS := -g -Werror -Wall


all: $(targets)

%: %.c
	$(CC) $(CFLAGS) -o $@ $<
.PHONY: clean

clean:
	-rm -f $(targets)

