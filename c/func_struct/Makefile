SRC = city.c nagoya.c tokyo.c
OBJ = $(SRC:%.c=%.o)

all: $(OBJ)
	gcc -o city $(OBJ)

city.o: city.c
	gcc -c city.c

nagoya.o: nagoya.c
	gcc -c nagoya.c

tokyo.o: tokyo.c
	gcc -c tokyo.c

clean:
	rm -f city $(OBJ)

.PHONY: clean
