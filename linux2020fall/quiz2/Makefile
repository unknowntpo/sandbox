CC := gcc
CFLAGS := -g -Wall

targets := isascii hexchar2val fastdiv
all: $(targets)

isascii: isascii.o

hexchar2val: hexchar2val.o

fastdiv: fastdiv.o

plotisascii: isascii scripts/plot.py 
	./isascii > plot/out.txt
	python3 scripts/plot.py
.PHONY: clean
clean:
	-rm -f *.o $(targets) 
