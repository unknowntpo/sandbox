cmd_/home/ubuntu/sandbox/lkm/hello.ko := ld -r -m elf_x86_64 -z max-page-size=0x200000 -T ./scripts/module-common.lds --build-id  -o /home/ubuntu/sandbox/lkm/hello.ko /home/ubuntu/sandbox/lkm/hello.o /home/ubuntu/sandbox/lkm/hello.mod.o ;  true