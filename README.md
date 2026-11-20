# Go 设计模式图解

用 Go 深入浅出讲解 GoF 23 个经典设计模式：**图先行、代码可运行、不长篇大论**。

- 每章：一张类比插图 → 反例 → 结构图 → Go 实现 → 适用场景 → 小结
- 进阶内容（Go 惯用写法、标准库实例、相似模式对比）放在折叠块中
- 代码只用标准库，每个模式都附测试，可直接 `go test`

## 本地预览

```bash
python -m venv .venv
.venv/Scripts/pip install -r requirements.txt   # macOS/Linux: .venv/bin/pip
.venv/Scripts/mkdocs serve                      # 打开 http://127.0.0.1:8000
```

## 运行代码

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

## 说明

- 本项目的正文、代码与插图均由 AI 生成，经人工审查。
- 吉祥物形象受 Go Gopher 启发。Go Gopher 由 [Renée French](https://reneefrench.blogspot.com/) 设计，以 [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/) 授权。
- 致敬并参考了 [hello-algo](https://github.com/krahets/hello-algo) 的组织方式。
