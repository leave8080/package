#ifndef QUEUE_H
#define QUEUE_H

#include "nwy_common.h"

typedef int oem_boolean;

#define FEATURE_Q_NO_SELF_QPTR
#define FEATURE_Q_SINGLE_LINK
#define OEM_TRUE  1
#define OEM_FALSE 0

#define INTLOCK()
#define INTFREE()

#define IVOC_LOCK()
#define IVOC_UNLOCK()

typedef struct q_link_struct
{
    struct q_link_struct  *next_ptr;
    /* Ptr to next link in list. If NULL, there is no next link. */
    

#ifndef FEATURE_Q_NO_SELF_QPTR

    void *self_ptr;
    /* Ptr to item which contains this link. */
    
    struct q_struct *q_ptr;
    /* Ptr to the queue on which this item is enqueued. NULL if the
    item is not on a queue. While not strictly necessary, this field
    is helpful for debugging.*/
#endif

#ifndef FEATURE_Q_SINGLE_LINK

    struct q_link_struct  *prev_ptr;

    /* Ptr to prev link in list. If NULL, there is no prev link. */
#endif

} q_link_type;

/* ------------------------------------------------------------------------
** Queue Head Link Structure
**
** When queue items are linked in a singly link list, q_head_link_type is
** used instead of q_link_type. This avoids the overhead of traversing 
** the whole of link list when queueing at the end of the link list.
** This structure should be accessed only by Queue services.
** ------------------------------------------------------------------------ */
typedef struct q_head_link_struct
{
    struct q_link_struct  *next_ptr;
    /* Ptr to head of the link list */
    
#ifndef FEATURE_Q_NO_SELF_QPTR

    void                  *self_ptr;
    /* Ptr to item which contains this link. */

    struct q_struct       *q_ptr;
    /* Ptr to the queue on which this item is enqueued. NULL if the
    item is not on a queue. While not strictly necessary, this field
    is helpful for debugging.*/
#endif
    struct q_link_struct  *prev_ptr;
    /* Ptr to the tail of link list */

} q_head_link_type;


/* ------------------------------------------------------------------------
** Queue Structure
**
** The following structure format is used by the Queue Services to represent
** a queue.
** 
** Do NOT access the fields of a queue directly. Only the Queue Services should
** do this.
** ------------------------------------------------------------------------ */
typedef struct q_struct
{
#ifdef FEATURE_Q_SINGLE_LINK
    q_head_link_type  link;
#else
    q_link_type  link;
#endif
    /* Used for linking items into queue. */

    int cnt;
/* Keeps track of number of items enqueued. Though not necessary, having
this field is tremendously helpful for debugging. */

#ifdef FEATURE_Q_WIN32_LOCKS
    #error code not present
#endif

#ifdef FEATURE_WINCE
    #error code not present
#endif /* FEATURE_WINCE */
} q_type;

/* ------------------------------------------------------------------------
** Queue Generic Item
**   Generic items must have q_link_type as the first element.  This allows
**   the linear search function to traverse the list without knowing
**   anything about the elements
** ------------------------------------------------------------------------ */
typedef struct {
    q_link_type link;
} q_generic_item_type;

/* ------------------------------------------------------------------------
** Queue Compare Function
**    Used by the searching functions to determine if an item is in
**       the queue.
**    Returns non zero if the element should be operated upon, 0 otherwise
**    For linear searching, the operation is to return a pointer to the
**       item and terminate the search
**    For linear deleting, the operation is to remove the item from the
**       queue and continue the traversal
** ------------------------------------------------------------------------ */
typedef int (*q_compare_func_type)( void* item_ptr, void* compare_val );

/* ------------------------------------------------------------------------
** Queue Action Function
**    Used by q_linear_delete to perform an action on an item which is 
**    being deleted from the list AFTER the item is deleted.  To perform
**    an action BEFORE the item is deleted, the user can call the action
**    function directly in the compare function call back.
** ------------------------------------------------------------------------ */
typedef void (*q_action_func_type)( void *item_ptr, void* param );

/*===========================================================================

                             Macro Definitions

===========================================================================*/
#ifdef PC_EMULATOR_H
#ifdef Q_RENAME
    #define PC_EMULATOR_QUEUE
    #include PC_EMULATOR_H
#endif

#ifdef Q_XCEPT
    #define PC_EMULATOR_Q_XCEPT
    #include PC_EMULATOR_H
#endif
#endif

#if !defined(PC_EMULATOR_H) || !defined(Q_XCEPT)
  #define Q_XCEPT_Q_INIT( q_ptr )
  #define Q_XCEPT_Q_LINK( q_ptr, link_ptr )
  #define Q_XCEPT_Q_PUT( q_ptr, link_ptr )
  #define Q_XCEPT_Q_GET( q_ptr )
  #define Q_XCEPT_Q_LAST_GET( q_ptr )
  #define Q_XCEPT_Q_CNT( q_ptr )
  #define Q_XCEPT_Q_CHECK( q_ptr )
  #define Q_XCEPT_Q_LAST_CHECK( q_ptr )
  #define Q_XCEPT_Q_NEXT( q_ptr, link_ptr )
#ifdef FEATURE_Q_NO_SELF_QPTR
    #define Q_XCEPT_Q_INSERT( q_ptr, q_insert_ptr, q_item_ptr )
#else
    #define Q_XCEPT_Q_INSERT( q_insert_ptr, q_item_ptr )
#endif
#ifdef FEATURE_Q_NO_SELF_QPTR
    #define Q_XCEPT_Q_DELETE( q_ptr, q_delete_ptr )
#else
    #define Q_XCEPT_Q_DELETE( q_delete_ptr )
#endif
#endif

#ifdef FEATURE_WINCE
   #error code not present
#endif

/* ========================================================================
MACRO Q_ALREADY_QUEUED
DESCRIPTION
   Evaluates to true if the item passed in is already in a queue and to
   false otherwise.
=========================================================================== */
#define Q_ALREADY_QUEUED( q_link_ptr ) \
   ((q_link_ptr)->next_ptr != NULL)

/*===========================================================================

                            Function Declarations

===========================================================================*/
#ifdef __cplusplus
   extern "C" {
#endif

/*==========================================================================

FUNCTION Q_INIT

DESCRIPTION
  This function initializes a specified queue. It should be called for each
  queue prior to using the queue with the other Queue Services.

DEPENDENCIES
  None.

RETURN VALUE
  A pointer to the initialized queue.

SIDE EFFECTS
  The specified queue is initialized for use with Queue Services.

===========================================================================*/
q_type* q_init ( q_type *q_ptr );


/*===========================================================================

FUNCTION Q_LINK

DESCRIPTION
  This function initializes a specified link. It should be called for each
  link prior to using the link with the other Queue Services.

DEPENDENCIES
  None.

RETURN VALUE
  A pointer to the initialized link.

SIDE EFFECTS
  The specified link is initialized for use with the Queue Services.

===========================================================================*/
q_link_type* q_link ( void *item_ptr, q_link_type *link_ptr );


/*===========================================================================

FUNCTION Q_PUT

DESCRIPTION
  This function enqueues an item onto a specified queue using a specified
  link.

DEPENDENCIES
  The specified queue should have been previously initialized via a call
  to q_init. The specified link field of the item should have been prev-
  iously initialized via a call to q_init_link.

RETURN VALUE
  None.

SIDE EFFECTS
  The specified item is placed at the tail of the specified queue.

===========================================================================*/
void  q_put  ( q_type *q_ptr, q_link_type *link_ptr );


/*===========================================================================

FUNCTION Q_GET

DESCRIPTION
  This function removes an item from the head of a specified queue.

DEPENDENCIES
  The specified queue should have been initialized previously via a call
  to q_init.

RETURN VALUE
  A pointer to the dequeued item. If the specified queue is empty, then
  NULL is returned.

SIDE EFFECTS
  The head item, if any, is removed from the specified queue.

===========================================================================*/
void* q_get ( q_type *q_ptr );

/*===========================================================================

FUNCTION Q_LAST_GET

DESCRIPTION
  This function removes an item from the tail of a specified queue.

DEPENDENCIES
  The specified queue should have been initialized previously via a call
  to q_init.

RETURN VALUE
  A pointer to the dequeued item. If the specified queue is empty, then
  NULL is returned.

SIDE EFFECTS
  The head item, if any, is removed from the specified queue.

===========================================================================*/
void* q_last_get ( q_type *q_ptr );



/*===========================================================================

FUNCTION Q_CNT

DESCRIPTION
  This function returns the number of items currently queued on a specified
  queue.

DEPENDENCIES
  The specified queue should have been initialized previously via a call
  to q_init.

RETURN VALUE
  The number of items currently queued on the specified queue.

SIDE EFFECTS
  None.

===========================================================================*/
int q_cnt  ( q_type *q_ptr );


/*===========================================================================

FUNCTION Q_CHECK

DESCRIPTION
  This function returns a pointer to the data block at the head of the queue.
  The data block is not removed from the queue.

DEPENDENCIES
  The specified queue should have been initialized previously via a call
  to q_init.

RETURN VALUE
  A pointer to the queue item. If the specified queue is empty, then
  NULL is returned.

SIDE EFFECTS
  None

===========================================================================*/
void* q_check (q_type  *q_ptr);


/*===========================================================================

FUNCTION Q_LAST_CHECK

DESCRIPTION
  This function returns a pointer to the data block at the tail of the queue.
  The data block is not removed from the queue.

DEPENDENCIES
  The specified queue should have been initialized previously via a call
  to q_init.

RETURN VALUE
  A pointer to the queue item. If the specified queue is empty, then
  NULL is returned.

SIDE EFFECTS
  The head item, if any, is removed from the specified queue.

===========================================================================*/
void* q_last_check ( q_type *q_ptr );


/*===========================================================================

FUNCTION Q_NEXT

DESCRIPTION
  This function returns a pointer to the next item on the queue.

DEPENDENCIES
  The specified queue should have been initialized previously via a call
  to q_init.

RETURN VALUE
  A pointer to the next item on the queue. If the end of the queue is reached, 
  then NULL is returned.

SIDE EFFECTS
  None.

===========================================================================*/
void* q_next  ( q_type *q_ptr, q_link_type *link_ptr );


/*===========================================================================

FUNCTION Q_INSERT

DESCRIPTION
  This function inserts an item before a specified item on a queue.

DEPENDENCIES
  The specified queue should have been initialized previously via a call
  to q_init.

RETURN VALUE
  None.

SIDE EFFECTS
  Input item is inserted before input item.

===========================================================================*/
#ifdef FEATURE_Q_NO_SELF_QPTR
   void q_insert  ( q_type *q_ptr, q_link_type *q_insert_ptr, q_link_type *q_item_ptr );
#else
   void q_insert  ( q_link_type *q_insert_ptr, q_link_type *q_item_ptr );
#endif


/*===========================================================================

FUNCTION Q_DELETE

DESCRIPTION
  This function removes an item from a specified queue.

DEPENDENCIES
  The specified queue should have been initialized previously via a call
  to q_init.

RETURN VALUE
  None.

SIDE EFFECTS
  Input item is delete from the queue.

===========================================================================*/
#ifdef FEATURE_Q_NO_SELF_QPTR
   void q_delete  ( q_type *q_ptr, q_link_type *q_delete_ptr );
#else
   void q_delete  ( q_link_type *q_delete_ptr );
#endif

/*===========================================================================
FUNCTION Q_DELETE_EXT
DESCRIPTION
  This function removes an item from a specified queue.
DEPENDENCIES
  The specified queue should have been initialized previously via a call
  to q_init.
RETURN VALUE
  FALSE : if the item is not found in the queue.
  TRUE  : if the item is found and removed from the queue.
SIDE EFFECTS
  Input item is deleted from the queue.
===========================================================================*/
#ifdef FEATURE_Q_NO_SELF_QPTR
   oem_boolean q_delete_ext  ( q_type *q_ptr, q_link_type *q_delete_ptr );
#else
   oem_boolean q_delete_ext  ( q_link_type *q_delete_ptr );
#endif

/*===========================================================================

FUNCTION Q_LINEAR_SEARCH

DESCRIPTION
  Given a comparison function, this function traverses the elements in
  a queue, calls the compare function, and returns a pointer to the
  current element being compared if the user passed compare function
  returns non zero.

  The user compare function should return 0 if the current element is
  not the element in which the compare function is interested.

DEPENDENCIES
  The specified queue should have been initialized previously via a call
  to q_init.

  The user's queue elements must have q_link_type as the first element
  of the queued structure.

RETURN VALUE
  A pointer to the found element

SIDE EFFECTS
  None.

===========================================================================*/
void* q_linear_search(
  q_type             *q_ptr,  
  q_compare_func_type compare_func,
  void               *compare_val
);

#if defined FEATURE_Q_NO_SELF_QPTR && defined FEATURE_Q_SINGLE_LINK
/*===========================================================================

FUNCTION Q_LINEAR_DELETE

DESCRIPTION
  Given a comparison function, this function traverses the elements in
  a queue, calls the compare function, and returns a pointer to the
  current element being compared if the user passed compare function
  returns non zero.  In addition, the item will be removed from the queue.

  The user compare function should return 0 if the current element is
  not the element in which the compare function is interested.

DEPENDENCIES
  The specified queue should have been initialized previously via a call
  to q_init.

  The user's queue elements must have q_link_type as the first element
  of the queued structure.

  The user's compare function will be passed NULL for the compare value.

RETURN VALUE
  None

SIDE EFFECTS
  None.

===========================================================================*/
void q_linear_delete(
  q_type             *q_ptr,  
  q_compare_func_type compare_func,
  void               *param,
  q_action_func_type  action_func
);

#endif

#ifdef __cplusplus
   }
#endif

#endif /* QUEUE_H */

