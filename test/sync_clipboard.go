/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : fankaljead
#   Email         : fankaljead@gmail.com
#   File Name     : sync_clipboard.go
#   Last Modified : 2018-12-12 13:27
#   Describe      :
#
# ====================================================*/

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os/exec"
)

var (
	content string
)

func main() {
	for {
		temp := content
		content = ReadClipboard()
		if temp != content {
			log.Println("正在上传")
			log.Println(content)
			UploadClip(ReadClipboard())
			log.Println("上传成功")
		}
	}
}

// ReadClipboard read the clip board from computor
func ReadClipboard() string {
	// 执行系统命令
	// 第一个参数是命令名称
	// 后面参数可以有多个，命令参数
	cmd := exec.Command("xclip", "-o")
	// 获取输出对象，可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)

	}
	return string(opBytes)
}

// UploadClip upload the clip content to server
func UploadClip(content string) {
	db, err := sql.Open("mysql",
		"root:212kawhi@tcp(45.40.201.117:3306)/fankaljead")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("insert into clipboard(content) values(?)")
	res, err := stmt.Exec(content)

	if err != nil {
		log.Fatal(err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID=%d\n", lastID)
}
