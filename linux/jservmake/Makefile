target := fac
objs := fac.o main.o
deps := $(objs:%.o=.%.o.d)
gen_file = main.c fac.c fac.h

CC = gcc
CFLAGS = -g -Werror -Wall

ifeq ("$(verbose)", "1")
    Q :=
    VECHO = @true
else
    Q := @
    VECHO = @printf
endif

ifeq ("$(sanity)", "1")
    CFLAGS += -fsanitize=address -fno-omit-frame-pointer -fno-common
    LDFLAGS += -fsanitize=address
endif


all: $(target)

$(target): $(objs)
	$(VECHO) "LD\t$@\n"
	$(Q)$(CC) $(LDFLAGS) -o $@ $^

%.o: %.c
	$(VECHO) "CC\t$@\n"
	$(Q)$(CC) -o $@ -c $(CFLAGS) -MMD -MF .$@.d $<

.PHONY: debug
debug:
	echo "CFLAGS: $(CFLAGS)"



.PHONY: generate
generate:
	@echo "generate files ..."
	touch $(gen_file)
	echo '#include "fac.h" \n int main(){}' > main.c
	echo '#include "fac.h"' > fac.c
	
.PHONY: clean
clean:
	rm -f $(targets) $(objs) $(deps)

-include $(deps)
