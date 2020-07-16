# factorial - factorial.o
# 	    	|_ factorial.c factorial.h
# 	    - main.o
# 	    	|_ main.c factorial.h

target := factorial
objs := main.o factorial.o
deps := $(objs:%.o=.%.o.d)
gen_file := main.c factorial.c factorial.h


CC := gcc
CFLAGS := -g -Werror -Wall

all: $(target)

$(target): $(objs)
	$(VECHO) "LD\t$@\n"
	$(Q)$(CC) $(CFLAGS) -o $@ $^

%.o: %.c
	$(VECHO) "CC\t$@\n"
	$(Q)$(CC) $(CFLAGS) -o $@ -c -MMD -MF .$@.d $<


ifeq ("$(verbose)", "1")
    Q :=
    VECHO = @true
else
    Q := @
    VECHO = @printf
endif




.PHONY: generate
generate:
	touch $(gen_file)
	echo '#include "factorial.h" \n int main(){} ' > main.c
	echo '#include "factorial.h"' > factorial.c

.PHONY: clean
clean:
	-rm -f $(target) $(objs) $(deps)

-include $(deps)
