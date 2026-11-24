# 参与维护

写作、代码与插图规范见 `spec/`；项目约定见 `CLAUDE.md`。

## 更新在线书站

- **自动发布**：编辑 `docs/index.md`（首页）、`docs/patterns.md`（目录）或相应章节，将改动推送或合并到 `master`。GitHub Actions 检查并构建通过后，自动更新 [在线书站](https://caticat.github.io/design-pattern/)；容器检查通过后也会独立发布 Docker 镜像。拉取请求只检查，不发布。
- **手动重发**：在 [Actions → 检查与发布](https://github.com/caticat/design-pattern/actions/workflows/site.yml) 点击 **Run workflow**，选择 `master` 并运行。不需要空提交；手动重发只更新 Pages，不发布 Docker 镜像。

在运行记录中确认「发布到 GitHub Pages」成功，再刷新书站。日常更新无需重新设置 Pages 来源。新增页面时，在 `mkdocs.yml` 的 `nav` 里加入口；只改正文不需要刷新 Playground 链接。

如果修改了带 Playground 链接的 Go 示例包，先把代码提交推到临时分支，再按 `scripts/playground.py` 的说明运行 `python -I scripts/playground.py --update`。提交更新后的链接，再合入 `master`；保留原代码提交，不要压缩或变基，否则链接绑定的版本会变。

外部链接由 [Actions → links](https://github.com/caticat/design-pattern/actions/workflows/links.yml) 每周一检查一次，结果在运行摘要里，只报告、不阻塞；也可手动运行。

## 个人主页入口

**个人主页的项目入口**统一维护在 [`caticat/caticat.github.io` 的链接配置](https://github.com/caticat/caticat.github.io/blob/master/src/config/site.config.js)。添加其他项目的链接时，只改这份配置；更新本书内容时，不用改个人主页。
