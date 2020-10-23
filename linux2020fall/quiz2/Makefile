CC := gcc
CFLAGS := -g -Wall

targets := isascii hexchar2val fastdiv
all: $(targets)

isascii: isascii.o

hexchar2val: hexchar2val.o

fastdiv: fastdiv.o

plotisascii_gnuplot: isascii scripts/perf_ascii.gp
	./isascii > plot/out.txt
	gnuplot scripts/perf_ascii.gp
plotisascii_pyplot: isascii scripts/plot.py 
	./isascii > plot/out.txt
	python3 scripts/plot.py
.PHONY: clean
clean:
	-rm -f *.o $(targets) 
