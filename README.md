# Go 设计模式图解

用 Go 深入浅出讲解 GoF 23 个经典设计模式：**图先行、代码可运行、不长篇大论**。

**[在线阅读 →](https://caticat.github.io/design-pattern/)**

- 每章一个模式：章首类比大图 + 4 个短小节（问题与思路 → Go 实现 → 专题 → 小结），每页一两屏读完
- 每章示例可在 Go Playground 一键运行；小结附「动手试试」练习
- 不知道该用哪个模式：查[模式目录](https://caticat.github.io/design-pattern/patterns/)的「我想……」场景表和[易混模式对比](https://caticat.github.io/design-pattern/compare/)
- 代码只用标准库，每个模式都附测试，可直接 `go test`

## Docker 部署

公开镜像发布后，无需登录即可拉取并启动（支持 Linux `amd64`、`arm64`）：

```bash
docker pull ghcr.io/caticat/design-pattern:latest
docker run -d --name design-pattern --restart unless-stopped -p 8080:8080 ghcr.io/caticat/design-pattern:latest
```

打开 <http://localhost:8080/>。图表库已随站点打包，内网或离线环境也能正常显示。如需用当前源码构建，先执行 `docker build -t design-pattern .`，再把运行命令里的镜像名改为 `design-pattern`。

## 本地预览

```bash
python -m venv .venv
.venv/Scripts/pip install -r requirements.txt   # macOS/Linux: .venv/bin/pip
.venv/Scripts/mkdocs serve                      # 打开 http://127.0.0.1:8000
```

## 运行代码

每章示例下方有「在 Go Playground 运行」按钮，可在浏览器中直接运行和修改。本地运行：

```bash
go test ./...                                    # 需要 Go 1.23+
go test -v -run Example ./codes/behavioral/strategy
```

## 目录

```text
docs/    正文（MkDocs Material）
codes/   Go 代码，按 creational / structural / behavioral 分类
spec/    写作、代码、插图规范
```

## 参与维护

发布流程、Playground 链接刷新等维护说明见 [CONTRIBUTING.md](CONTRIBUTING.md)。

## 许可与鸣谢

本项目的正文、示例代码和原创插图采用 [CC BY-NC-SA 4.0](LICENSE)，另有说明的第三方素材按原许可使用。非商业转载须署名、链接原站及许可，并注明是否修改；分享改编版须沿用相同或兼容的许可。详见[许可与鸣谢](https://caticat.github.io/design-pattern/credits/)。

本书内容由 AI 生成，供人工审查。感谢 [《Hello 算法》](https://github.com/krahets/hello-algo) 提供图解与章节组织的参考，以及 Renée French 设计的 [Go Gopher 原作](https://go.dev/blog/gopher)（[CC BY 4.0](https://creativecommons.org/licenses/by/4.0/)）。
