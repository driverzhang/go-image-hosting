package entities

import (
	"fmt"
	"github.com/driverzhang/go-image-hosting/clipboard"
	"github.com/urfave/cli"
)

const url = "https://raw.githubusercontent.com/"

type Cli struct {
	// repo 名称 如： driverzhang/image-storage/master/1
	// master代表你设置github pages 的分支名 1代表内部建立文件夹名
	Repo string `json:"repo"` // repo 名称
	AP   string `json:"ap"`   // 图床存放对应本地的 repo 绝对路径
}

var Clip *Cli

func RunClipBoard(c *cli.Context) {
	fmt.Println(1111)
	j, err := clipboard.ReadAll()
	if err != nil {
		return
	}

	fmt.Println(222)
	fmt.Printf("粘贴板输入：%+v\n", j)
	if Clip == nil {
		Clip = &Cli{}
		Clip.Repo = "driverzhang/image-storage/master/1/"
	}

	r := fmt.Sprintf("%s%s%s", url, Clip.Repo, j)
	fmt.Printf("粘贴板输出 %+v\n", r)
	err = clipboard.WriteAll(r)
	fmt.Println("image hosting from clipboard success")
	return
}
