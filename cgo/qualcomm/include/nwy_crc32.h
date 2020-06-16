/*====*====*====*====*====*====*====*====*====*====*====*====
   Copyright (c) 2018 Neoway Technologies, Inc.
   All rights reserved.
   Confidential and Proprietary - Neoway Technologies, Inc.
   Author: wangjie@neoway.com
   Date: 2018.06
*====*====*====*====*====*====*====*====*====*====*====*====*/

#ifndef  _NWY_CRC32_H  
#define  _NWY_CRC32_H

typedef unsigned int uint32_t;
typedef unsigned char uint8_t;

uint32_t nwy_calc_crc32(uint32_t crc, const void *buf, size_t size);

#endif /*_NWY_CRC32_H*/

