CC = gcc
CFLAGS = -g -Wall
TARGET = max0 max2 typeof

all: $(TARGET)

# How to support multiple target ?
max0: max0.c
	$(CC) $(CFLAGS) -o $@ $<

max2: max2.c
	$(CC) $(CFLAGS) -o $@ $<


typeof: typeof.c
	$(CC) $(CFLAGS) -o $@ $<


.PHONY: clean
clean:
	-rm -f $(TARGET)
