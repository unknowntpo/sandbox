	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 11, 0	sdk_version 11, 1
	.globl	_utilfunc                       ## -- Begin function utilfunc
	.p2align	4, 0x90
_utilfunc:                              ## @utilfunc
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	movq	%rdi, -8(%rbp)
	movq	%rsi, -16(%rbp)
	movq	%rdx, -24(%rbp)
	movq	-8(%rbp), %rax
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function
	.globl	_myfunc                         ## -- Begin function myfunc
	.p2align	4, 0x90
_myfunc:                                ## @myfunc
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	subq	$96, %rsp
	movq	24(%rbp), %rax
	movq	16(%rbp), %r10
	movq	%rdi, -8(%rbp)
	movq	%rsi, -16(%rbp)
	movq	%rdx, -24(%rbp)
	movq	%rcx, -32(%rbp)
	movq	%r8, -40(%rbp)
	movq	%r9, -48(%rbp)
	movq	-8(%rbp), %rcx
	imulq	-16(%rbp), %rcx
	imulq	-24(%rbp), %rcx
	imulq	-32(%rbp), %rcx
	imulq	-40(%rbp), %rcx
	imulq	-48(%rbp), %rcx
	imulq	16(%rbp), %rcx
	imulq	24(%rbp), %rcx
	movq	%rcx, -56(%rbp)
	movq	-8(%rbp), %rcx
	addq	-16(%rbp), %rcx
	addq	-24(%rbp), %rcx
	addq	-32(%rbp), %rcx
	addq	-40(%rbp), %rcx
	addq	-48(%rbp), %rcx
	addq	16(%rbp), %rcx
	addq	24(%rbp), %rcx
	movq	%rcx, -64(%rbp)
	movq	-56(%rbp), %rdi
	movq	-64(%rbp), %rsi
	movq	-56(%rbp), %rcx
	movq	%rax, -80(%rbp)                 ## 8-byte Spill
	movq	%rcx, %rax
	cqto
	idivq	-64(%rbp)
	movq	%r10, -88(%rbp)                 ## 8-byte Spill
	callq	_utilfunc
	movq	%rax, -72(%rbp)
	movq	-72(%rbp), %rax
	addq	$20, %rax
	addq	$96, %rsp
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function
.subsections_via_symbols
