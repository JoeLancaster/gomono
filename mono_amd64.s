#include "go_asm.h"
#include "textflag.h"

// func gomono.getMono() (ret int64)
TEXT ·getMono(SB),NOSPLIT,$16-8
	MOVQ	SP, R12
	SUBQ	$16, SP //space for one timespec
	ANDQ	$~15, SP //align

	// call vdso version of clock_gettime, go runtime already has the location
	MOVL	$1, DI // CLOCK_MONOTONIC
	LEAQ	8(SP), SI
	MOVQ	runtime·vdsoClockgettimeSym(SB), AX
	CALL	AX
	
	//fetch result
	MOVQ	16(SP), AX //timespec.tv_nsec
	MOVQ	8(SP), CX //timespec.tv_sec
	IMULQ	$1000000000, CX //secs to nanosecs
	ADDQ	AX, CX
	
	MOVQ	R12, SP
	MOVQ	CX, ret+0(FP) //put result in ret
	RET

//func checkVdsoavailable() (available bool)
TEXT ·checkVdsoAvailable(SB),NOSPLIT,$0-1
	MOVQ	runtime·vdsoClockgettimeSym(SB), AX
	CMPQ	AX, $0 //if symbol is not set a.k.a zero value then we can't use it
	MOVB	AX, available+0(FP)
	RET
