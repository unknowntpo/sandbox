target := ll
objs := ll.o
CC := gcc
deps := $(objs:%.o=.%.o.d)
CFLAGS := -g -Wall -Werror
DEPFLAGS := -MMD -MF
all: $(target)

# Add verbose mode
ifeq ("$(v)","1")
    Q :=     # quiet
    VECHO := @true
else
    Q := @
    VECHO := @printf
endif

$(target): $(objs)
	$(VECHO) "LD\t $@\n"
	$(Q)$(CC) -o $@ $^

%.o: %.c
	$(VECHO) "CC\t $@\n"
	$(Q)$(CC) $(CFLAGS) -o $@ -c $(DEPFLAGS) .$@.d $<

.PHONY: clean
clean:
	-rm -f $(objs) $(deps) $(target)

-include $(deps)
