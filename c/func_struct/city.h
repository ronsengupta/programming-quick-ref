int nagoya(int);
int tokyo(int);

struct my_favorite_city
{
	char *name;
	int (*city_func)(int);
};


struct my_favorite_city nagoya_city={
	.name		= "nagoya",
	.city_func	= nagoya,
};

struct my_favorite_city tokyo_city={
	.name		= "tokyo",
	.city_func	= tokyo,
};
