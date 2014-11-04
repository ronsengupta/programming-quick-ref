all:
compile:
	as -o cpuid2.o cpuid2.s
	ld -dynamic-linker /lib64/ld-linux-x86-64.so.2 -o cpuid2 -lc cpuid2.o
debug:
	as -gstabs -o cpuid2.o cpuid2.s
	ld -dynamic-linker /lib64/ld-linux-x86-64.so.2 -o cpuid2 -lc cpuid2.o
list:
	as -adhlns=cpuid2.lst -o cpuid2.o cpuid2.s
clean:
	rm -f *.lst *.o cpuid2
