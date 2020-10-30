CC := gcc
target := bt
objs := bt.o
CFLAGS := -g -Wall -Werror
LDFLAGS :=

all: $(target)

$(target): $(objs)
	$(CC) -o $@ $(LDFLAGS) $^

%.o: %.c
	$(CC) -o $@ $(CFLAGS) -c $<

.PHONY: clean
clean:
	-rm $(target) *.o
