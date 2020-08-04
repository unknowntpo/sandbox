target := ll
objs := ll.o
CC := gcc
deps := $(objs:%.o=.%.o.d)
CFLAGS := -g -Wall -Werror
DEPFLAGS := -MMD -MF
all: $(target)

target: $(objs)
	$(CC) -o $@ $^

%.o: %.c
	$(CC) $(CFLAGS) -o $@ -c $(DEPFLAGS) .$@.d $<

.PHONY: clean
clean:
	-rm -f $(objs) $(deps) $(target)

-include $(deps)
