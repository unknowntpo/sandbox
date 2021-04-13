
# specify what we wanna build
target := factorial
objs := main.o factorial.o
deps := $(objs:%.o=%.o.d)
CC := gcc
CFLAGS := -Wall -Werror
DEPFLAGS := -MMD -MF

# specify the rules to build target

all: $(target)

target: $(objs)
	@echo "check point 1"
	$(CC) $(CFLAGS) -o $@ $^

%.o: %.c
	@echo "dep file: $(deps)"
	$(CC) $(CFLAGS) -o -c $(DEPFLAGS) $@.d $<

.PHONY: generate
generate:
	touch main.c factorial.c factorial.h
	@echo "int main(){}" > main.c

.PHONY: clean
clean:
	-rm -f $(target) $(objs)

-include $(deps)
