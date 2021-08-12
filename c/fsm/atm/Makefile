CC = gcc
CFLAGS = -g -Wall
LDFLAGS =
TARGET = atm
OBJS = atm.o \
       test.o

ifeq ("$(V)", "1")
    Q :=
    VECHO := @true
else
    Q := @
    VECHO := @printf
endif


all: $(TARGET)

$(TARGET): $(OBJS)
	$(VECHO) "  LD\t$@\n"
	$(Q)$(CC) $(LDFLAGS) -o $@ $^

%.o: %.c
	$(VECHO) "  CC\t$@\n"
	$(Q)$(CC) -c $(CFLAGS) -o $@ $<

.PHONY: clean
clean:
	-rm -rf $(TARGET) $(OBJS) *.dSYM
