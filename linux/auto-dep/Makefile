# ver1. 0:33

all: factorial

factorial: main.o factorial.o
	gcc -Wall -Werror -o factorial main.o factorial.o

main.o: main.c factorial.h
	gcc -Wall -Werror -c main.c

factorial.o: factorial.c factorial.h
	gcc -Wall -Werror -c factorial.c

.PHONY: generate

generate:
	touch main.c factorial.c factorial.h
	@echo "int main(){}" > main.c


clean:
	rm -f factorial main.o factorial.o

# ver2. 2:03

CC := gcc
CFLAGS := -Wall -Werror

all: factorial

factorial: main.o factorial.o
	$(CC) $(CFLAGS) -o factorial main.o factorial.o

main.o: main.c factorial.h
	$(CC) $(CFLAGS) -c main.c

factorial.o: factorial.c factorial.h
	$(CC) $(CFLAGS) -c factorial.c

.PHONY: generate

generate:
	touch main.c factorial.c factorial.h
	@echo "int main(){}" > main.c


clean:
	rm -f factorial main.o factorial.o

# ver3. 2:33

CC := gcc
CFLAGS := -Wall -Werror

all: factorial

factorial: main.o factorial.o
	$(CC) $(CFLAGS) -o $@ $^ 

main.o: main.c factorial.h
	$(CC) $(CFLAGS) -c $<

factorial.o: factorial.c factorial.h
	$(CC) $(CFLAGS) -c $<

.PHONY: generate

generate:
	touch main.c factorial.c factorial.h
	@echo "int main(){}" > main.c


clean:
	rm -f factorial main.o factorial.o
# ver4 3:13

CC := gcc
CFLAGS := -Wall -Werror

all: factorial

factorial: main.o factorial.o
	$(CC) $(CFLAGS) -o $@ $^ 

%.o: %.c factorial.h
	$(CC) $(CFLAGS) -c $<

.PHONY: generate

generate:
	touch main.c factorial.c factorial.h
	@echo "int main(){}" > main.c


clean:
	rm -f factorial main.o factorial.o

# ver5 3:34

target := factorial
objs = main.o factorial.o

CC := gcc
CFLAGS := -Wall -Werror

all: $(target)

factorial: $(objs) 
	$(CC) $(CFLAGS) -o $@ $^ 

%.o: %.c factorial.h
	$(CC) $(CFLAGS) -c $<

.PHONY: generate

generate:
	touch main.c factorial.c factorial.h
	@echo "int main(){}" > main.c


clean:
	rm -f $(target) $(objs) 


