package main

import (
	"errors"
	"fmt"
	xerrors "github.com/pkg/errors"
)

var ErrNoRows = errors.New("no rows ")

// 即使error是具有模糊界限的错误，也应该 warp 后返回给上层，由调用者决定是否是个错误
func dao() error {
	return xerrors.Wrap(ErrNoRows, "调用dao接口错误")
}

// 抛给上层
func service() error {
	return dao()
}

// 接口顶层入口，统一记录日志以及处理
func controller() {
	if err := service(); err != nil {
		fmt.Printf("controller error, type(%T), msg(%v)\r\n trace(%+v)", err, err, err)
		// do something
		return
	}
	return
}

func main() {
	controller()
}
