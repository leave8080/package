/*
*  oplicense.h
*  Copyright (C) 2004-2014 Magus Technology Co.,Ltd.

   Change history:
   1.0  2014  First version

*/
#ifndef __OP_LICENSE_H
#define __OP_LICENSE_H


#ifndef GMacroAPI
	#if defined(__WIN32) || defined(__WIN32__) || defined(WIN32)
        #ifdef MAGUS_IMPLETMENT_SHARED
            #define GMacroAPI(type) __declspec(dllexport) type _stdcall
        #elif defined MAGUS_IMPLETMENT_STATIC || defined MAGUS_USE_STATIC
            #define GMacroAPI(type) type
        #else
            #define GMacroAPI(type) __declspec(dllimport) type _stdcall
        #endif
    #else
        #define GMacroAPI(type) type
    #endif
#endif


#ifdef __cplusplus
extern "C" {
#endif


// 授权类型
#define LNONE     0
#define LFILE     1
#define LUKEY   	2
#define LANY     -1


/**
* @brief 授权管理服务
*/
#ifndef OPLM
typedef void* OPLM;
#endif


/**
* @brief 产品授权
*/
#ifndef OPLicense
typedef void* OPLicense;
#endif


/**
 * @brief 获取授权服务
 * @param type 	授权设备码的来源：LFILE-系统自动生成，LUKEY-外部USB硬件狗，LALL-按顺序自动查找
 * @param path 	授权服务的路径
 * @param cert 	授权证书的路径，如果为空，缺省为授权文件名加后缀.cer
 * @param error 返回的错误码
 *
 * @return 授权服务，用完需要调用lclose()释放。
*/
GMacroAPI(OPLM) op2_lopen(int type, const char *path, const char *cert, int *error);


/**
 * @brief 关闭授权服务
*/
GMacroAPI(void) op2_lclose(OPLM lm);


/**
 * @brief 授权服务的类型
*/
GMacroAPI(int) op2_lget_type(OPLM lm);


/**
 * @brief 授权发布时间
*/
GMacroAPI(void) op2_lget_issue_date(OPLM lm, char *buf, int len);


/**
 * @brief 获取设备码
 * @param type 	设备码的来源：LFILE-系统自动生成，LUKEY-外部USB硬件狗
 * @param uid 	字符数组，用来存放设备码的值
 * @param len 	字符数组可用的最大长度
 *
 * @return 成功:0; 失败:错误码
*/
GMacroAPI(int) op2_lget_uid(int type, char *uid, int len);


/**
 * @brief 获取给定项目、产品和设备的授权
 * @param lm 		授权服务
 * @param project 	授权项目，如果为空，缺省为第一个项目
 * @param product 	授权产品
 * @param host 		授权设备ID
 *
 * @return 产品授权, 用完需要调用lfree_license()释放
*/
GMacroAPI(OPLicense) op2_lget_license(OPLM lm, const char *project, const char *product, const char *uid);


/**
 * @brief 释放授权
*/
GMacroAPI(void) op2_lfree_license(OPLicense lic);


/**
 * @brief 获取授权的属性值
 * @param lic 	产品授权
 * @param key 	授权属性
 * @param uid 	字符数组，用来存放授权属性的值
 * @param len 	字符数组可用的最大长度
 *
 * @return 返回码：0-成功，-1-错误，未找到
*/
GMacroAPI(int) op2_lget_license_value(OPLicense lic, const char *key, char *buf, int len);


/**
 * @brief 返回授权记录数
*/
GMacroAPI(int) op2_lget_count(OPLM lm);


/**
 * @brief 返回迭代器对应的授权
 * @param lm 		授权服务
 * @param index 	授权索引
 *
 * @return 产品授权, 用完需要调用lfree_license()释放
*/
GMacroAPI(OPLicense) op2_lget_at(OPLM lm, int index);


#ifdef __cplusplus
}
#endif

#endif // __OP_LICENSE_H

