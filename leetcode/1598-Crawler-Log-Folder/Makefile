CC = gcc
TARGET = main
CFLAGS = -g -Werror -Wall -std=c99

all: $(TARGET)

$(TARGET): main.c
	$(CC) $(CFLAGS) -o $@ $<

.PHONY: clean
clean:
	rm -f $(TARGET)
