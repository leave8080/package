package main

/*
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <errno.h>
#include <time.h>
#include <unistd.h>
#include <sys/epoll.h>
#include <sys/wait.h>
#include <termios.h>
#include "nwy_uart.h"
#include "nwy_common.h"
static int fd;



int Serive_open(){
	fd = -1;
	char *dev_string = NULL;
	dev_string = "/dev/ttyHSL0";
	unsigned long baudrate = 0;
	baudrate = 115200;
	fd = nwy_uart_open(dev_string, baudrate, FC_NONE);
	printf("open fd=%d\n",fd);

	if (fd < 0)
	{
		printf("Error device open failed,code:%d\n", fd);
	}
	return fd;
}

void* ReadData(void* arg){
 char     buff[1024] = {0};
    uint32_t count      = 0;
    int      fd         = *((int*) arg);
    int      ret;

    fd_set fds;
    FD_ZERO(&fds);
    FD_SET(fd, &fds);
    while (1)
    {
        ret = select(fd + 1, &fds, NULL, NULL, NULL);
        if (ret <= 0)
        {
            continue;
        }
        ret = read(fd, buff, 1024);
        if (ret > 0)
        {
            count++;
            printf("[%4u]:%s\n", count, buff);
        }
        else
        {
            printf("read error\n");
        }
    }

    return arg;

}
void Serive_close(){
	if (fd <= 0)
	{
		return;
	}
	close(fd);
}
*/
import "C"
import "fmt"

func main() {
	fd := (int)(C.Serive_open())
	if fd == -1{
		fmt.Println("open err")
		return
	}
	C.ReadData(C.int(fd))


}
