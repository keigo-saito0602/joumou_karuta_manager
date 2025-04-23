package main

import (
	"os"

	"github.com/keigo-saito0602/joumou_karuta_manager/cmd"
)

// @title Joumou Karuta Manager API
// @version 1.0
// @description 上毛カルタAPIのSwaggerドキュメント
// @BasePath /
func main() {
	os.Exit(cmd.Run())
}
