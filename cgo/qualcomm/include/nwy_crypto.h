/*====*====*====*====*====*====*====*====*====*====*====*====*====*====*====
	Copyright (c) 2018 Neoway Technologies, Inc.
	All rights reserved.
	Confidential and Proprietary - Neoway Technologies, Inc.
	Author: lijingbei@neoway.com
	Date: 2018.06
*====*====*====*====*====*====*====*====*====*====*====*====*====*====*====*/

#ifndef __NWY_CRYPTO_H
#define __NWY_CRYPTO_H
#include <errno.h>
#include <stdint.h>
#include <sys/ioctl.h> 

/*---------------------------Macro Definition---------------------------*/
#define SVC_TEST_SIZE_1K			(1024)
#define CRYPTO_ASYM_KEY_SIZE_MAX	(512+16)
#define CRYPTO_ASYM_IV_LENGTH		(16)
#define CRYPTO_ASYM_HMAC_LENGTH		(32)

///< Keyblob size in bytes
#define KEYBLOB_SIZE				(32+16+8)

///< seal_buf_len = input_buf_len + 24
#define SEAL_DATA_PLU				(24)

///< import aes_key size, 32BYtes, 256bits
#define SVC_AES256_KEY_SIZE			(32)

///<Minimum signature size in bytes
#define SVC_RSA_MIN_SIG_LEN			(128)

/*---------------------------Type Definition ---------------------------*/
typedef unsigned int uint32_t;
typedef unsigned char uint8;
typedef uint8 boolean;

/**
 * Type of test mode for RSA operations.
 */
typedef enum {
	GEN_KEY_TEST = 0, //gen key test
	IMP_KEY_TEST, //import key test
	CRYPTO_MAX_VAL = 0x7FFFFFFF,
} nwy_usr_test_mode_t;

/**
 * Type of padding used for RSA operations.
 */
typedef enum {
	TZ_RSA_PKCS115_SHA2_256=0x00000001,
	TZ_RSA_PSS_SHA2_256=0x00000002,
	TZ_RSA_MAX_VAL=0x7FFFFFFF, //not in use.
} nwy_rsa_digest_pad_algo_t;

/**
 * Parameters needed to generate an RSA key.
 */
typedef struct {
	uint32_t modulus_size;
	uint64_t public_exponent;
	nwy_rsa_digest_pad_algo_t digest_pad_type;
} nwy_asym_rsa_keygen_params_t;

typedef struct {
	uint32_t magic_num; /* Validate the integrity of keyblob content */
	uint32_t version; /* API version number */
	nwy_rsa_digest_pad_algo_t digest_padding;
	uint8 modulus[CRYPTO_ASYM_KEY_SIZE_MAX]; /* Modulus (N) array [MSB...LSB] */
	uint32_t modulus_size; /* Modulus array length */
	uint8 public_exponent[CRYPTO_ASYM_KEY_SIZE_MAX]; /* public exponent (e) array [MSB...LSB] */
	uint32_t public_exponent_size; /* public exponent array length */
	uint8 iv[CRYPTO_ASYM_IV_LENGTH]; /* Initial vector */
	uint8 encrypted_private_exponent[CRYPTO_ASYM_KEY_SIZE_MAX]; /* Encrypted Private Exponent (d) array [MSB...LSB] */
	uint32_t encrypted_private_exponent_size; /* Encrypted Private Exponent array length */
	uint8 hmac[CRYPTO_ASYM_HMAC_LENGTH]; /* HMAC array */
} nwy_rsa_key_t;

typedef struct _rsa_svc_ {
	uint8 *modulus;
	uint32_t modulus_size;
	uint8 *pub_exp;
	uint32_t pub_exp_size;
	uint8 *priv_exp;
	uint32_t priv_exp_size;
	nwy_rsa_digest_pad_algo_t padding;
} nwy_rsa_svc_t;

/*---------------------------Function Definition--------------------------*/
int nwy_crypto_stor_drv_test
(
	nwy_usr_test_mode_t mode
);

int nwy_crypto_stor_gen_sym_key
(
	void *key,
	uint32_t key_len
);

int nwy_crypto_stor_seal_data
(
	char *key,
	uint32_t key_len,
	char *in_buf,
	uint32_t in_len,
	char *out_buf,
	uint32_t out_len
);

int nwy_crypto_stor_unseal_data
(
	char *key,
	uint32_t key_len,
	char *in_buf,
	uint32_t in_len,
	char *out_buf,
	uint32_t out_len
);

int nwy_crypto_stor_imp_sym_key
(
	char *imp_key,
	uint32_t imp_key_len,
	char *key,
	uint32_t key_len
);

int nwy_crypto_rsa_drv_test
(
	nwy_usr_test_mode_t mode
);

int nwy_crypto_rsa_gen_asym_keypair
(
	nwy_rsa_key_t *usr_rsa_key,
	uint32_t usr_rsa_key_len,
	nwy_rsa_digest_pad_algo_t padding,
	uint32_t modulus_size,
	uint64_t pub_exp
);

int nwy_crypto_rsa_sign_data
(
	nwy_rsa_key_t *usr_rsa_key,
	uint32_t usr_rsa_key_len,
	uint8_t *data,
	uint32_t data_len,
	uint8_t *signature,
	uint32_t sig_len
);

int nwy_crypto_rsa_verify_data
(
	nwy_rsa_key_t* usr_rsa_key,
	uint32_t usr_rsa_key_len,
	uint8_t* data,
	uint32_t data_len,
	uint8_t* signature,
	uint32_t signed_len,
	uint32_t sig_len
);

int nwy_crypto_rsa_imp_asym_keypair
(
	nwy_rsa_key_t* usr_rsa_key,
	uint32_t usr_rsa_key_len,
	nwy_rsa_digest_pad_algo_t padding,
	uint8_t *modulus,
	uint32_t modulus_size,
	uint8_t *pub_exp,
	uint32_t pub_exp_size,
	uint8_t *priv_exp,
	uint32_t priv_exp_size
);

int nwy_crypto_rsa_exp_pubkey_and_modulu
(
	nwy_rsa_key_t *usr_rsa_key,
	uint32_t usr_rsa_key_len,
	uint8_t *pub_exp,
	uint32_t *pub_exp_size,
	uint8_t *modulus,
	uint32_t *modulus_size
);


#endif /* __NWY_CRYPTO_H */


