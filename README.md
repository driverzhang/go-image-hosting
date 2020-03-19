# go-image-hosting
Image Hosting for golang (用golang实现的图床工具,github做免费永久存储,配合微信截图保存到粘贴板)

# Why ?

之前的用的图床工具 Ipic 免费默认的是 微博图床，不过后来微博禁止了外链于是乎个人博客的所有图片都挂调了，没错就是提出来访问都无法显示那种。

鉴于 Ipic 的工作原理和功能需求实现，本人打算亲自打造一款类似的可以永久免费的存储的图床工具。所以该开源项目就此诞生了。

# What?

- 永久免费存储 基于 Github.
- 利用微信截图，自动保存到粘贴板功能（当然你也可以找其他这样的工具）
- 直接运行命令行 命令如：bd ，粘贴板内就已经存在 markdown 版的图片链接地址。

# How?


要完美使用该工具，你本地配置必须具备如下条件：

1. 有一个已经授权的好的 github repo项目，用于专门存放你的图片。

2. 有一个截图工具，要求是截图后自动会保存到你的粘贴板中（之后会扩展其他方式获取图片来源）

---

## 注意事项流程说明：

### 如何让自己图片在国内能正常访问？

主要是把你的repo项目 进行setting 将其添加进入 github pages 中

参考链接：

[GitHub Pages](https://lab.github.com/githubtraining/github-pages)

[github pages微博图床图片无法正确显示](https://counter2015.com/2019/04/29/bug2/)

[Github图床](https://www.jianshu.com/p/c794bad425e5)

