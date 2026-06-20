CC      = cc
CFLAGS  = -Wall -Wextra -std=c11 -g
TARGET  = httpd
SRCS    = src/c/main.c
OBJS    = $(SRCS:.c=.o)

$(TARGET): $(OBJS)
	$(CC) $(CFLAGS) -o $@ $^

%.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $<

clean:
	rm -f $(OBJS) $(TARGET)

.PHONY: clean
