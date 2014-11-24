gcc hello.c -o hello $(pkg-config --cflags --libs gtk+-3.0)
