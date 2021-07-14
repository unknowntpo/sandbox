CC = gcc
CFLAGS = -g -Wall
TARGET = max0

all: $(TARGET)

# How to support multiple target ?
$(TARGET): max0.c
	$(CC) $(CFLAGS) -o $@ $<

.PHONY: clean
clean:
	-rm -f $(TARGET)
