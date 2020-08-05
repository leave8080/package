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
#include "nwy_client.h"
static int fd;



int Serive_open(){
	fd = -1;
	char *dev_string = NULL;
	dev_string = "/dev/ttyHSL0";
	unsigned int baudrate = 0;
	baudrate = 115200;
	fd = nwy_uart_open(dev_string, baudrate, FC_NONE);
	printf("Open com success!!!!!!!!!!!open fd=%d\n",fd);

	if (fd < 0)
	{
		printf("Error device open failed,code:%d\n", fd);
	}
	return fd;
}
void Serive_close(){
	if (fd <= 0)
	{
 	printf("Serive_closer fd <0 \n");
		return;
	}

	nwy_uart_close(fd);
}
void* ReadData(void* arg){
 char     buff[1024] = {0};
    uint32_t count      = 0;

	int      fd         = (*(int *)arg);
    int      ret;

    //fd_set fds;
    //FD_ZERO(&fds);
    //FD_SET(fd, &fds);
    while (1)
    {
        //ret = select(fd + 1, &fds, NULL, NULL, NULL);
        //if (ret <= 0)
        //{
        //    continue;
        //}
		ret = nwy_uart_read(fd,buff,1024);
       // ret = read(fd, buff, 1024);
        if (ret > 0)
        {
            count++;
            printf("[%4u]:%s\n", count, buff);
        }
        else
        {
            printf("read error\n");
			Serive_close();
        }
    }

    return arg;

}

*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	var fd int
	for {
		fd = (int)(C.Serive_open())
		if fd > 0 {
			break
		} else {
			fmt.Println("open err")
			time.Sleep(time.Second * 3)
			continue
		}

	}

	cfd := unsafe.Pointer(&fd)
	C.ReadData(cfd)

}
