target := Merge_Sorted_Array Merge_Two_Sorted_Lists # just add target file at here
CC := gcc
CFLAGS := -g -Werror -Wall
all: $(target)

%: %.c
	$(CC) $(CFLAGS) -o $@ $<

.PHONY: clean
clean:
	-rm -f $(target)
