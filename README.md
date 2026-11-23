# Go 设计模式图解

用 Go 深入浅出讲解 GoF 23 个经典设计模式：**图先行、代码可运行、不长篇大论**。

- 每章一个模式：章首类比大图 + 4 个短小节（问题与思路 → Go 实现 → 专题 → 小结），每页一两屏读完
- 代码只用标准库，每个模式都附测试，可直接 `go test`

## Docker 部署

公开镜像发布后，无需登录即可拉取并启动（支持 Linux `amd64`、`arm64`）：

```bash
docker pull ghcr.io/caticat/design-pattern:latest
docker run -d --name design-pattern --restart unless-stopped -p 8080:8080 ghcr.io/caticat/design-pattern:latest
```

打开 <http://localhost:8080/>。如需用当前源码构建，先执行 `docker build -t design-pattern .`，再把运行命令里的镜像名改为 `design-pattern`。

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

## 更新在线书站

- **自动发布**：编辑 `docs/index.md`（首页）、`docs/patterns.md`（目录）或相应章节，将改动推送或合并到 `master`。GitHub Actions 检查并构建通过后，自动更新 [在线书站](https://caticat.github.io/design-pattern/)；容器检查通过后也会独立发布 Docker 镜像。拉取请求只检查，不发布。
- **手动重发**：在 [Actions → 检查与发布](https://github.com/caticat/design-pattern/actions/workflows/site.yml) 点击 **Run workflow**，选择 `master` 并运行。不需要空提交；手动重发只更新 Pages，不发布 Docker 镜像。

在运行记录中确认「发布到 GitHub Pages」成功，再刷新书站。日常更新无需重新设置 Pages 来源。新增页面时，在 `mkdocs.yml` 的 `nav` 里加入口；只改正文不需要刷新 Playground 链接。

如果修改了带 Playground 链接的 Go 示例包，先把代码提交推到临时分支，再按 `scripts/playground.py` 的说明运行 `python -I scripts/playground.py --update`。提交更新后的链接，再合入 `master`；保留原代码提交，不要压缩或变基，否则链接绑定的版本会变。

**个人主页的项目入口**统一维护在 [`caticat/caticat.github.io` 的链接配置](https://github.com/caticat/caticat.github.io/blob/master/src/config/site.config.js)。添加其他项目的链接时，只改这份配置；更新本书内容时，不用改个人主页。

## 说明

- 本项目的正文、代码与插图均由 AI 生成，供人工审查。
- 吉祥物参考了 [Go Gopher 原作](https://go.dev/blog/gopher)，加入黄色安全帽和工具腰带。原作由 Renée French 设计，采用 [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/) 许可；这不是对整书的授权声明。
- 致敬并参考了 [hello-algo](https://github.com/krahets/hello-algo) 的组织方式。
