all:
compile:
	as -o cpuid.o cpuid.s
	ld -o cpuid cpuid.o
debug:
	as -gstabs -o cpuid.o cpuid.s
	ld -o cpuid cpuid.o
list:
	as -adhlns=cpuid.lst -o cpuid.o cpuid.s
clean:
	rm -f *.lst *.o cpuid
