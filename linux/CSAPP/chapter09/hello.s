	.file	"hello.c"
	.intel_syntax noprefix
	.text
	.section	.rodata
.LC0:
	.string	"hello, world\n"
	.text
	.globl	main
	.type	main, @function
main:
.LFB0:
	.cfi_startproc
	push	rbp
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	mov	rbp, rsp
	.cfi_def_cfa_register 6
	mov	edx, 13
	lea	rsi, .LC0[rip]
	mov	edi, 1
	mov	eax, 0
	call	write@PLT
	mov	edi, 0
	call	_exit@PLT
	.cfi_endproc
.LFE0:
	.size	main, .-main
	.ident	"GCC: (Ubuntu 7.5.0-3ubuntu1~18.04) 7.5.0"
	.section	.note.GNU-stack,"",@progbits
