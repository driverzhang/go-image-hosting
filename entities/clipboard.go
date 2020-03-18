package entities

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/atotto/clipboard"
	
)

const url  =  "https://raw.githubusercontent.com/"

type Cli struct {
	// repo 名称 如： driverzhang/image-storage/master/1
	// master代表你设置github pages 的分支名 1代表内部建立文件夹名
	Repo string `json:"repo"` // repo 名称
	AP string `json:"ap"` // 图床存放对应本地的 repo 绝对路径
}

var Clip *Cli

func RunClipBoard(c *cli.Context)  {
	j, err := clipboard.ReadAll()
	if err != nil {
		return
	}
	
	fmt.Printf("%+v", j)
	if Clip.Repo == "" {
		Clip.Repo = "driverzhang/image-storage/master/1/"
	}
	
	r := fmt.Sprintf("%s%s%s", url, Clip.Repo,j)
	fmt.Printf("%+v", r)
	err = clipboard.WriteAll(r)
	fmt.Println("image hosting from clipboard success")
	return
}