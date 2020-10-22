package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func main() {
	password := []byte("thisIsPassWord")
	nowG := time.Now()
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	fmt.Println("加密后", string(hashedPassword), "耗时", time.Now().Sub(nowG))
	nowC := time.Now()
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println("验证耗费时间", time.Now().Sub(nowC))
	fmt.Println(err)
}

// 结果
// 加密后 $2a$10$ESkb/bwSyISLgq1bOH0C2utXdb.hcH9oBQD1hUnfDOzm4bMKK6EX2 耗时 67.9985ms
// 验证耗费时间 66.0008ms
// <nil>
