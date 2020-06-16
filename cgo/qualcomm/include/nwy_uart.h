/*====*====*====*====*====*====*====*====*====*====*====*====*====*====*====
    Copyright (c) 2018 Neoway Technologies, Inc.
    All rights reserved.
    Confidential and Proprietary - Neoway Technologies, Inc.
    Author: Shi.Shaogang
    Date: 2018.06
*====*====*====*====*====*====*====*====*====*====*====*====*====*====*====*/

#ifndef __NWY_UART_H__
#define __NWY_UART_H__

/*----------------------------Type Definition-----------------------------*/
typedef enum {
    FC_NONE = 0,  // None Flow Control
    FC_RTSCTS,    // Hardware Flow Control (rtscts)
    FC_XONXOFF    // Software Flow Control (xon/xoff)
}nwy_flowctrl_t;

typedef enum {
    PB_NONE = 0,
    PB_ODD,
    PB_EVEN,
    PB_SPACE,
    PB_MARK
}nwy_paritybits_t;

typedef struct {
    unsigned int baudrate; 
    unsigned int databits;
    unsigned int stopbits;
    nwy_paritybits_t parity;
    nwy_flowctrl_t flowctrl;
}nwy_uartdcb_t;

typedef enum {
    BR_300 	    = 300,
    BR_600 	    = 600,
    BR_1200 	= 1200,
    BR_2400 	= 2400,
    BR_4800     = 4800,
    BR_9600 	= 9600,
    BR_19200 	= 19200,
    BR_38400 	= 38400,
    BR_57600 	= 57600,
    BR_115200   = 115200,
    BR_230400   = 230400,
    BR_460800   = 460800
}nwy_baudrate_t;

/*---------------------------Function Definition--------------------------*/
int nwy_uart_open(const char* port, unsigned int baudrate, nwy_flowctrl_t flowctrl);
int nwy_uart_read (int fd,unsigned char* buf, unsigned int buf_len);
int nwy_uart_write(int fd, const unsigned char* buf, unsigned int buf_len);
int nwy_uart_setdcb(int fd, struct termios *term);
int nwy_uart_getdcbconfig(int fd, struct termios *term);
int nwy_uart_ioctl(int fd, unsigned int cmd, void* pvalue);
int nwy_uart_close(int fd);

#endif 
