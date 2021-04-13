# play with variable substitution

foo := a.o b.o l.a c.o

bar := $(foo:.o=.c)

.PHONY: print
print:
	@echo "foo's value: $(foo)"
	@echo "bar's value: $(bar)"
    

.PHONY: touch
touch: 
	touch a.o b.o l.a c.o
