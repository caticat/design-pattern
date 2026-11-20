"""MkDocs 钩子：给图注自动编号。

正文中每张图（插图或 Mermaid）下方写一行图注：
    <p class="caption">图注文字</p>
所在页面属于导航中「第 N 章」时，图注按章内顺序编号为「图 N-k」，跨节连续计数（同 hello-algo）。
不在任何章节内的页面（首页、分类页）图注不编号。
"""

import re

CHAPTER_RE = re.compile(r"第\s*(\d+)\s*章")
CAPTION_RE = re.compile(r'<p class="caption"( markdown)?>')
FENCE_RE = re.compile(r"^(```|~~~).*?^\1", re.M | re.S)

_start: dict[str, tuple[int, int]] = {}  # src_uri -> (章号, 本页之前的图数)


def _chapter(page):
    node = page.parent
    while node is not None:
        m = CHAPTER_RE.search(node.title or "")
        if m:
            return int(m.group(1))
        node = node.parent
    return None


def on_nav(nav, config, files):
    _start.clear()
    counts: dict[int, int] = {}
    for page in nav.pages:  # 按导航顺序
        ch = _chapter(page)
        if ch is None:
            continue
        with open(page.file.abs_src_path, encoding="utf-8") as f:
            src = FENCE_RE.sub("", f.read())
        _start[page.file.src_uri] = (ch, counts.get(ch, 0))
        counts[ch] = counts.get(ch, 0) + len(CAPTION_RE.findall(src))
    return nav


def on_page_content(html, page, config, files):
    if page.file.src_uri not in _start:
        return html
    ch, n = _start[page.file.src_uri]

    def number(m):
        nonlocal n
        n += 1
        return f'{m.group(0)}<span class="fig-no">图 {ch}-{n}</span>'

    return re.sub(r'<p class="caption">', number, html)
