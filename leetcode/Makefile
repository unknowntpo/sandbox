target := Merge_Sorted_Array  # just add target file at here 
CC := gcc
CFLAGS := -g -Werror -Wall
all: $(target)

%: %.c
	$(CC) -o $@ $<

.PHONY: clean
clean:
	-rm -f $(target)
