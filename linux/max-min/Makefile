CC = gcc
CFLAGS = -g -Wall
TARGET = max0 max2 max3 typeof

all: $(TARGET)

# How to support multiple target ?
max0: max0.c
	$(CC) $(CFLAGS) -o $@ $<

max2: max2.c
	$(CC) $(CFLAGS) -o $@ $<

max3: max3.c
	$(CC) $(CFLAGS) -o $@ $<

typeof: typeof.c
	$(CC) $(CFLAGS) -o $@ $<


.PHONY: clean
clean:
	-rm -rf $(TARGET) *.dSYM
