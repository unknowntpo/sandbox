CC := gcc
CFLAGS := -g -m32 -Werror -Wall
target := unaligned_set
all: $(target)

%: %.c
	$(CC) $(CFLAGS) -o $@ $<
clean:
	-rm -rf $(target) 
