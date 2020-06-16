#--------------------------------------------------------------
#               Neoway OpenLinux
#   
#--------------------------------------------------------------
APP_NAME = boxsupport-q

#CROSS_COMPILER=arm-oe-linux-gnueabi-

#---------------------
# Include definition
#---------------------
INC_DIR=  -I /home/leave/N720_OPEN_LINUX_SDK_25/include

USR_INC_DIR = -I/home/leave/N720_OPEN_LINUX_SDK_25/include \
			  -I ./include/

#--------------------------------------------------------------
# LIB
#--------------------------------------------------------------
STD_LIB_FILES =	-L. -ldl -lpthread
USR_LIB_FILES =	-L/home/leave/N720_OPEN_LINUX_SDK_25/libs -lnwy_device -lnwy_service -lnwy_common \
				-L/home/leave/N720_OPEN_LINUX_SDK_25/libs/share  \
				-L./libs -lopapi4

#--------------------------------------------------------------
# Source code files
#--------------------------------------------------------------
DIRS = $(shell find . -maxdepth 3 -type d)
LOCAL_SRC_FILES = $(foreach dir, $(DIRS), $(wildcard $(dir)/*.c)) 
#LOCAL_SRC_FILES = nwy_dm_test.c
LOCAL_OBJECTS = $(notdir $(patsubst %.c, %.o,$(LOCAL_SRC_FILES)))

#--------------------------------------------------------------
# Compilation options
#--------------------------------------------------------------
LOCAL_CFLAGS = -march=armv7-a -mfloat-abi=soft -mfpu=neon -g


.PHONY: $(APP_NAME)

$(APP_NAME) : $(LOCAL_SRC_FILES)
	$(CC) -g  $(LOCAL_CFLAGS) $(LOCAL_INC_DIR) $(USR_INC_DIR) $(LOCAL_SRC_FILES) $(STD_LIB_FILES) $(USR_LIB_FILES)  -o $@
	rm -rf  *.o *.a
	@echo ---- finished ----

.PHONY: clean
clean:
	rm -rf $(APP_NAME)
	rm -rf  *.o *.a
