#compile object file to executable

objects = main.o kbd.o command.o display.o \
          insert.o search.o files.o utils.o

edit : $(objects)
        cc -o edit $(objects)
main.o : main.c defs.h
        cc -c main.c



test : $(objects) 
	cc -pthread -o test $(objects)
cpfile.o : cpfile.
