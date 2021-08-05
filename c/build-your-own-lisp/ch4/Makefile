CC = clang
CFLAGS = -g -Wall
TARGET = lispy

all: $(TARGET)

lispy: lispy.c
	$(CC) $(CFLAGS) -o $@ $<

.PHONY: clean
clean:
	-rm -rf $(TARGET) *.dSYM
