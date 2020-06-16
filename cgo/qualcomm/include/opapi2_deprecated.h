#ifndef __OPAPI2_DEPRECATED_H
#define __OPAPI2_DEPRECATED_H

#if defined(__WIN32) || defined(__WIN32__) || defined(WIN32)
    #ifndef MAGUS_WINDOWS
        #define MAGUS_WINDOWS
    #endif
#else
    #ifndef MAGUS_POSIX
        #define MAGUS_POSIX
    #endif
#endif

#ifndef GMacroAPI
    #ifdef MAGUS_WINDOWS
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


/**
* @brief 报警数据结构 - 已过时
*/
typedef struct
{
    int id;				// 点ID号
    OPTime t_firstAlarm;	// 首次报警时间
    OPTime t_levelChanged;	// 变到当前级别的时间

    OPTime time;           // 实时值时间
    short status;       // 实时状态
    double value;       // 实时值
}AlarmVal;
/** @} */


/* ============================================================================================================================ */
/**
 * @defgroup apiv2_op_buffer opWriter 属性设置接口
 * @brief   对OPBuffer属性值进行设置
 * @details
 *      key、type、意义、[默认值]
 *  @li transfer_protocol                int     使用tcp传输         无默认值
 *  @li server_address                   str     服务器地址          无默认值
 *  @li server_port                      int     服务器端口号        无默认值
 *  @li client_name                      str     客户端用户名        无默认值
 *  @li client_password                  str     客户端密码          无默认值
 *  @li storage_location                 str     缓存数据存储路径    无默认值
 *  @li storage_capacity                 int     缓存数据最大容量    无默认值
 *  @li option                           int     见op_writer_open    无默认值
 *  @li isolator_enabled                 int     过隔离器            无默认值
 *  @li net_probe_interval               int     网络探测间隔        5 second
 *  @li upload_history_interval          int     上传历史间隔        100 ms
 *  @li realtime_filter_enabled          int     是否过滤实时值      0 bool
 *  @li filter_bool_upload_interval      int     开关量上传间隔      30 second
 *  @li filter_bool_storage_interval     int     开关量存储间隔      900 second
 *  @li size_for_each_upload             int     每次上传文件大小    10 MB
 *  @li upload_history_per_time          int     每次上传的记录数    50000 条
 *  @li single_file_capacity             int     每个缓存文件大小    1024 MB
 *  @li storage_file_quantity            int     缓冲的最大文件个数  无默认值
 *  @li storage_time_limit               int     缓冲的最大天数      无默认值
 *  @li history_task_interval            int     后台线程sleep时间   1000 ms
 *  @li record_filename                  str     记录文件名          "op.writer.json"
 *
 * @{
 */

/**
* @brief 本地缓冲，用户透明
*/
typedef void* OPLocalBuffer;

/**
* @brief 获取OPBuffer的内部状态
*
* @param op 数据库句柄
* @param key 关键字(key):  network_connection, try_connection_count, cache_status, upload_status
* @param value 对应属性的值
*/
GMacroAPI(int) op2_buffer_get_status(OpenPlant op, const char* key);

/**
* @brief 对OPBuffer属性值进行设置，其属性值类型为int
*
* @param op 数据库句柄
* @param name 属性关键字(key),如:transfer_protocol (使用tcp传输)
* @param value 对应属性的值
*/
GMacroAPI(void) op2_buffer_set_int   (OpenPlant op, const char* name, int value);

/**
* @brief 对OPBuffer属性值进行设置，其属性值类型为string
*
* @param
* @param name 属性关键字(key)
* @param value 对应属性的值

*/
GMacroAPI(void) op2_buffer_set_string(OpenPlant op, const char* name, const char* value);

/**
* @brief 对点进行设置
* @param op 数据库句柄
* @param id  点的ID
* @param type 点的类型
* @param deadband 死区值
*/
GMacroAPI(void) op2_buffer_set_point (OpenPlant op, int id, int type, double deadband);

/** 
 * @brief 打开本地缓冲，按照给定时间段从中检索数据
 * 
 * @param directroy     缓冲文件所在目录
 * @param beginTime     要检索数据的开始时间
 * @param endTime       要检索数据的结束时间
 * @param error         错误字符串
 *
 * @return 成功:OPLocalBuffer; 失败:NULL
 */
GMacroAPI(OPLocalBuffer) op2_buffer_open(const char* directroy, OPTime beginTime, OPTime endTime, const char** error);

/** 
 * @brief 获取缓冲中的下一块
 * 
 * @param opBuffer op2_buffer_open返回的句柄
 *
 * @return 有数据:返回一个数据块(结果集); 无数据:NULL
 */
GMacroAPI(OPResult) op2_buffer_next(OPLocalBuffer opBuffer);

/** 
 * @brief 
 * 
 * @param result    op2_buffer_next 返回的数据块
 * @param id        数据块中下一个记录中的id字段
 * @param time      数据块中下一个记录中的time字段
 * @param status    数据块中下一个记录中的status字段
 * @param value     数据块中下一个记录中的value字段
 */
GMacroAPI(int) op2_buffer_get_froresult(OPResult result, int *id, int *time, short *status, double *value);

/** 
 * @brief 释放op2_buffer_next 返回的结果集
 * 
 * @param 将被释放的结果集
 */
GMacroAPI(void) op2_buffer_free_result(OPResult result);

/** 
 * @brief 关闭op2_buffer_open打开的本地缓冲
 * 
 * @param opBuffer 将被关闭的本地缓冲
 */
GMacroAPI(void) op2_buffer_close(OPLocalBuffer opBuffer);

/** @} */


/**
* @brief 创建查询对象
*
* @param op   数据库句柄
* @param id   对象ID
* @param fields 对象字段列表（逗号分隔）， 如：TM,AS,AV
* @return
*   @li 成功: 对象描述字
*   @li 失败: NULL
*/
GMacroAPI(OPObject) op2_new_query_byid(OpenPlant op, int id);


/**
* @brief 创建查询对象
*
* @param op   数据库句柄
* @param id   对象全局名称W3
* @param fields 对象字段列表（逗号分隔）， 如：TM,AS,AV
* @return
*   @li 成功: 对象描述字
*   @li 失败: NULL
*/
GMacroAPI(OPObject) op2_new_query_byname(OpenPlant op, const char *global_name);


/**
* @brief 更新对象字段数据
*
* @param op 数据库服务器
* @param fields 对象字段列表（逗号分隔）， 如：ID,TM,AS,AV 或 GN(全局名称),TM,AS,AV
* @param gh 名字组
* @param objects 对象数组
* @param errors 错误码
* @return
*   @li 成功: 0
*   @li 失败: -1
*/
GMacroAPI(int) op2_update(OpenPlant op, const char *fieldlist, int num, OPObject *objects, int *errors);


/**
* @brief 取得结果集尺寸
* @param result 结果集
* @return 结果集尺寸
*/
GMacroAPI(int) op2_alarnurows(OPResult result);

/**
* @brief 取得结果集中下一个记录的值
*
* @param result 结果集
* @param aval 用于返回的报警记录值
*
* @return
*   @li 0: 已经到结果集末尾
*   @li 1: 还有下一条记录
*   @li 其它: 出错
*/
GMacroAPI(int) op2_alarfetch_value(OPResult result, AlarmVal *aval);

/**
* @brief 释放一个报警结果集
*
* @param result 结果集
* @return 本函数没有返回值
*/
GMacroAPI(void) op2_alarfree_result(OPResult result);

/**
* @brief 通过ID获取对象实时报警信息
*
* @param op 数据库服务器句柄
* @param num 对象个数
* @param id 对象ID数组
* @param result 报警结果集
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_get_real_alarbyid(OpenPlant op, int num, int *id, OPResult *result);

/**
* @brief 通过点名获取对象实时报警信息
*
* @param op 数据库服务器句柄
* @param gh 点名组
* @param result 报警结果集
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_get_real_alarbyname(OpenPlant op, OPGroup gh, OPResult *result);

/**
* @brief 通过ID获取子对象实时报警信息
*
* @param op 数据库服务器句柄
* @param id 对象ID号
* @param result 报警结果集
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_get_real_child_alarbyid(OpenPlant op, int id, OPResult *result);

/**
* @brief 通过点名获取子对象实时报警信息
*
* @param op 数据库服务器句柄
* @param obj_name 对象名称
* @param result 报警结果集
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_get_real_child_alarbyname(OpenPlant op, const char *obj_name, OPResult *result);

/**
* @brief 通过ID获取对象历史报警信息
*
* @param op 数据库服务器句柄
* @param num 对象个数
* @param id 对象ID数组
* @param result 报警结果集
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_get_hist_alarbyid(OpenPlant op, int num, int *id, OPTime begin_tm, OPTime end_tm, OPResult *result);

/**
* @brief 通过点名获取对象历史报警信息
*
* @param op 数据库服务器句柄
* @param gh 点名组
* @param result 报警结果集
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_get_hist_alarbyname(OpenPlant op, OPGroup gh, OPTime begin_tm, OPTime end_tm, OPResult *result);

/**
* @brief 通过ID获取子对象历史报警信息
*
* @param op 数据库服务器句柄
* @param id 对象ID号
* @param result 报警结果集
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_get_hist_child_alarbyid(OpenPlant op, int id, OPTime begin_tm, OPTime end_tm, OPResult *result);

/**
* @brief 通过点名获取子对象历史报警信息
*
* @param op 数据库服务器句柄
* @param obj_name 对象名称
* @param result 报警结果集
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_get_hist_child_alarbyname(OpenPlant op, const char *obj_name, OPTime begin_tm, OPTime end_tm, OPResult *result);

/**
* @brief 同时更新多个点的报警
*
* @param op      数据库句柄
* @param pt      value打包类型: float, char, short, int, double
* @param num     点的个数
* @param id      指向点ID的数组
* @param time    指向点的报警时间标签的数组
* @param status  指向对应与时间标签点状态的数组
* @param value   指向对于与时间标签点值的数组
* @param errors  指向回返错误码的数组，若数组中有为非零的值，则表明该点写实时值错误
*
* @return 正确 返回 0
*         错误 返回 其他
*/
GMacroAPI(int) op2_update_alarm(OpenPlant op, int pt, int num, const int *id, const OPTime *time, const short *status, const double *value, int *errors);

/**
* @brief 通过ID确认对象报警
*
* @param op 数据库服务器句柄
* @param num 对象个数
* @param id 对象ID数组
* @param errors 错误码
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_alarack_byid(OpenPlant op, int num, int *id, int *errors);

/**
* @brief 通过点名确认对象报警
*
* @param op 数据库服务器句柄
* @param gh 点名组
* @param errors 错误码
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_alarack_byname(OpenPlant op, OPGroup gh, int *errors);


/**
* @brief 通过ID抑制对象报警
*
* @param op 数据库服务器句柄
* @param num 对象个数
* @param id 对象ID数组
* @param errors 错误码
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_alarinhibit_byid(OpenPlant op, int num, int *id, int inhibit, int *errors);

/**
* @brief 通过点名抑制报警
*
* @param op 数据库服务器句柄
* @param gh 点名组
* @param errors 错误码
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_alarinhibit_byname(OpenPlant op, OPGroup gh, int inhibit, int *errors);

/**
* @brief 通过ID控制定制状态位
*
* @param op		数据库句柄
* @param num	测点个数
* @param id		测点ID数组
* @param on		设定值{0,1}
* @param errors 错误码
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_custocontrol(OpenPlant op, int num, int *id, int on, int *errors);

/**
* @brief 通过点名控制定制状态位
*
* @param op		数据库句柄
* @param num	测点个数
* @param tags	测点名称数组
* @param on		设定值{0,1}
* @param errors 错误码
* @return
*   @li 成功: 0
*   @li 失败: 错误码
*/
GMacroAPI(int) op2_custocontrol_tags(OpenPlant op, int num, const char *tags[], int on, int *errors);


/**
* @brief 手工强制测点数值
*
* @param op      数据库句柄
* @param num     点个数
* @param id      指向点ID数组
* @param value   指向点的值的数组
* @param errors  指向回返错误码的数组，数组中值为0或者1都表示写入成功，其中1表示第一次写入成功；0表示对当前实时值修改成功。
*
* @return 正确 返回 0
*         错误 返回 其他
*/
GMacroAPI(int) op2_force(OpenPlant op, int num, const int *id, const double *value, int *errors);

GMacroAPI(int) op2_force_tags(OpenPlant op, int num, const char *tags[], const double *value, int *errors);


/**
* @brief 清除强制测点
*
* @param op      数据库句柄
* @param num     点个数
* @param id      指向点ID数组
* @param errors  指向回返错误码的数组，数组中值为0或者1都表示写入成功，其中1表示第一次写入成功；0表示对当前实时值修改成功。
*
* @return 正确 返回 0
*         错误 返回 其他
*/
GMacroAPI(int) op2_clear_force(OpenPlant op, int num, const int *id, int *errors);

GMacroAPI(int) op2_clear_force_tags(OpenPlant op, int num, const char *tags[], int *errors);


/**
* @brief 订阅开关量事件---【向后兼容，不建议使用】
*
* @param op 		OpenPlant 连接
* @param type 		订阅类型：AX, DX, I2, I4, R8
* @param snapshot 	启动时是否发送全部测点的快照：0-不启用，1-启用
* @param cb 		回调函数
* @param userdata 	回调函数使用的用户参数
* @param error 		错误码
*/
typedef void event_callback(void *userdata, int id, const char *name, tsv_t *value);
typedef void* OPAsyncHandler;
GMacroAPI(OPAsyncHandler) op2_realtime_open_async(void *op, const char *type, int withSnapshot, event_callback *cb, void *userdata, int *error);


#ifdef __cplusplus
}
#endif

#endif //__OPAPI2_DEPRECATED_H


