# Go 设计模式图解

用 Go 深入浅出讲解 GoF 23 个经典设计模式：**图先行、代码可运行、不长篇大论**。

- 每章一个模式：章首类比大图 + 4 个短小节（问题与思路 → Go 实现 → 专题 → 小结），每页一两屏读完
- 代码只用标准库，每个模式都附测试，可直接 `go test`

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

## 检查与发布

向 `master` 提交或发起拉取请求时，GitHub Actions 会检查 Go 示例、严格构建站点并校验 Playground 链接。

要发布在线站点，维护者先在仓库 **Settings → Pages → Source** 选择 **GitHub Actions**，再到 **Actions → 检查与发布 → Run workflow** 选择 `master` 手动运行。自动检查和拉取请求不会发布站点。

## 说明

- 本项目的正文、代码与插图均由 AI 生成，供人工审查。
- 吉祥物参考了 [Go Gopher 原作](https://go.dev/blog/gopher)，加入黄色安全帽和工具腰带。原作由 Renée French 设计，采用 [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/) 许可；这不是对整书的授权声明。
- 致敬并参考了 [hello-algo](https://github.com/krahets/hello-algo) 的组织方式。
