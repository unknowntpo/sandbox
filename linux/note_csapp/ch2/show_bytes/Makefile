OBJ := show_bytes.o
CFLAGS := -g -Wall -Werror
CC := gcc
LDFLAGS :=

ifeq ("$(SANITIZER)", "1")
    CFLAGS  += -fsanitize=alignment
    LDFLAGS += -fsanitize=alignment
endif

all: show_bytes
	
show_bytes: $(OBJ) 
	$(CC) $(LDFLAGS) -o $@ $^

%.o: %.c
	$(CC) -o $@ $(CFLAGS) -c $<	

.PHONY: clean
clean:
	-rm -f show_bytes *.o
