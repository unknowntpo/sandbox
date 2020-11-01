CC = gcc
CFLAGS = -std=c11 -g -static
TARGET = 9cc
OBJ = 9cc.o

all: $(TARGET)

# Verbose mode
ifeq ("$(VERBOSE)", "1")
    Q :=
    VECHO = @true
else
    Q := @
    VECHO = @printf
endif

$(TARGET): $(OBJ)
	$(VECHO) "  LD\t$@\n"
	$(Q)$(CC) -o $@ $(LDFLAGS) $^

%.o: %.c
	$(VECHO) "  CC\t$@\n"
	$(Q)$(CC) -o $@ $(CFLAGS) -c $<

test: 9cc
	./scripts/test.sh

clean:
	rm -f 9cc *.o *~ tmp*

.PHONY: test clean
