TARGET_MODULE := hello
obj-m := hello.o
KDIR := /lib/modules/$(shell uname -r)/build/

all:
	make -C $(KDIR) M=$(PWD) modules
load:
	sudo insmod $(TARGET_MODULE).ko
unload:
	sudo rmmod $(TARGET_MODULE) || true >/dev/null
clean:
	make -C $(KDIR) M=$(PWD) clean

