target := read
objs := read.o csapp.o
deps := $(objs:%.o=.%.o.d)

CC := gcc
CFLAGS := -g -pthread

all: $(target)

$(target): $(objs)
	$(CC) $(CFLAGS) -o $@ $^
%.o: %.c
	$(CC) $(CFLAGS) -o $@ -c -MMD -MF .$@.d $<
.PHONY: generate
generate:
	@echo "generate foo.txt..."
	-rm -f foo.txt
	@wget https://randomword.com/paragraph -O foo.txt

.PHONY: clean
clean:
	-rm -f $(objs) $(target) $(deps)

-include $(deps)
