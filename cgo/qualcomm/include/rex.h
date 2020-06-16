//#include "oem_api.h"
#include <pthread.h>

typedef unsigned long rex_sigs_type;

typedef struct rex_tcb_type
{
    pthread_mutex_t mutex;
    pthread_cond_t cond;
    rex_sigs_type sigs;
}rex_tcb_type;

void rex_init(void * tcb);
rex_sigs_type rex_set_sigs( rex_tcb_type *p_tcb,rex_sigs_type p_sigs);
rex_sigs_type rex_clr_sigs(rex_tcb_type *p_tcb,rex_sigs_type p_sigs);
rex_sigs_type rex_wait(rex_tcb_type *p_tcb,rex_sigs_type p_sigs );
