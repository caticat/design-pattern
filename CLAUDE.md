# 项目规范：Go 设计模式图解

参照 [hello-algo](https://github.com/krahets/hello-algo)，用 Go 深入浅出讲解设计模式。

## 已定决策

| 项 | 决定 |
|---|---|
| 读者 | 初学者为主；进阶内容放 `??? note` 折叠块，兼顾有经验者 |
| 范围 | GoF 23 个模式（第一期） |
| 语言 | 仅中文 |
| 代码 | Go 1.23+，只用标准库，每个模式可运行并附测试 |
| 站点 | MkDocs Material，人工通过本地 `mkdocs serve` 审查 |
| 生产方式 | **全部由 AI 生成**；人工只负责审查和提建议 |
| 吉祥物 | Gopher + 黄色安全帽 + 工具腰带，已定稿（否决 Octocat：商标；否决二次元角色：授权不明、跨章不一致） |
| 原则 | 简要清晰、图先行、不长篇大论；只讲不过时的原理 |

## 详细规范（动手前必读）

| 文件 | 内容 |
|---|---|
| `spec/writing.md` | 写作规范、章节模板、自检清单 |
| `spec/code.md` | 代码目录与约束 |
| `spec/image.md` | 生图模型、Prompt 模板、Mermaid 分工 |

## 目录结构

```text
docs/                 # 正文（MkDocs）
  assets/images/      # 生成的插图
  creational/ structural/ behavioral/
codes/                # Go 代码，正文通过 snippets 引用
spec/                 # 项目规范
mkdocs.yml
```

## 工作流程

1. 按规范生成：正文 → 代码与测试 → Mermaid → 插图
2. **AI 先自审**：跑 `go vet ./...`、`go test ./...`、`mkdocs build --strict`、`scripts/playground.py`（代码有改动：先推送，再 `--update`，再提交链接），对照 `spec/writing.md` 自检清单逐条确认，查看插图
3. 自审通过后再交人工审查；按意见修改

## Git

- 远程：`git@github.com:caticat/design-pattern.git`（**公共仓库**），分支 `master`
- 禁止提交任何密码、密钥、Token、内部地址、个人/公司邮箱等敏感信息；提交前检查 `git diff --cached`
- 提交信息用中文，格式：`<类型>: <说明>`，类型取 `docs` / `code` / `chore` / `fix`

## 本地命令

```bash
python -m venv .venv && .venv/Scripts/pip install -r requirements.txt
.venv/Scripts/mkdocs serve        # 本地预览 http://127.0.0.1:8000
go vet ./... && go test ./...
```
