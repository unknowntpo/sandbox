
objects = program.o foo.o utils.o

program : $(objects)
	cc -o program $(objects)

new :
	touch program.c foo.c utils.c
%.o: %.c
	cc -c -o $@ $<
clean :
	-rm $(objects)
