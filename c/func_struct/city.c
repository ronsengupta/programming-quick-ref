#include <stdio.h>
#include "city.h"

int main(int argc, char *argv[])
{

	struct my_favorite_city *my_favorite_city = &tokyo_city;
	printf("City name is = %s\n", my_favorite_city->name);
	my_favorite_city->city_func(1);

	printf ("========================\n");
	my_favorite_city = &nagoya_city;

	printf("City name is = %s\n", my_favorite_city->name);
	my_favorite_city->city_func(2);
//
	return 0;
}
