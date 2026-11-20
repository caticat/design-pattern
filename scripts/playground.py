"""生成并校验各章的 Go Playground「在线运行」链接。

每章正文中放一行按钮（ID 由本脚本维护）：
    [:material-play: 在 Go Playground 运行](https://go.dev/play/p/<ID>){ .md-button target=_blank }

脚本把该章引用的 Go 包打包成 Playground 的多文件格式（txtar）：
    - 主文件：example_test.go 改成 package main，并追加 main() 依次调用各 Example
    - go.mod + 包内所有非测试 .go 文件（保持原目录与模块路径，import 无需改写）

用法：
    python -I scripts/playground.py           # 离线校验链接是否与当前代码一致（自审用）
    python -I scripts/playground.py --update  # 上传到 go.dev 并改写正文中的链接
"""

import base64
import hashlib
import pathlib
import re
import sys
import urllib.request

ROOT = pathlib.Path(__file__).resolve().parent.parent
MODULE = "github.com/caticat/design-pattern"
SHARE_URL = "https://go.dev/_/share"
# 与 golang.org/x/playground 的 snippet ID 算法一致：sha256(salt + body) 的 base64url 前 11 位
SALT = b"Go playground salt\n"

LINK_RE = re.compile(r"https://go\.dev/play/p/[\w-]+")
SNIPPET_RE = re.compile(r'--8<-- "(codes/[^/"]+/[^/"]+)/')
EXAMPLE_RE = re.compile(r"^func (Example\w*)\(\) \{", re.M)
MARKER_RE = re.compile(r"^\s*// --8<-- .*\n", re.M)


def read(path: pathlib.Path) -> str:
    return MARKER_RE.sub("", path.read_text(encoding="utf-8"))


def build(pkg: str) -> bytes:
    """把一个模式包打包成 Playground 可运行的 txtar。"""
    pkg_dir = ROOT / pkg
    main = read(pkg_dir / "example_test.go")
    if '"fmt"' not in main:
        sys.exit(f"{pkg}/example_test.go 必须导入 fmt")
    main = re.sub(r"^package \w+", "package main", main, count=1, flags=re.M)
    names = EXAMPLE_RE.findall(main)
    if not names:
        sys.exit(f"{pkg}/example_test.go 中没有 Example 函数")
    calls = []
    for name in names:
        if len(names) > 1:
            calls.append(f'\tfmt.Println("=== {name} ===")')
        calls.append(f"\t{name}()")
    main = (
        "// 由 scripts/playground.py 从 example_test.go 生成；期望输出见各函数的 // Output: 注释。\n"
        + main.rstrip()
        + "\n\nfunc main() {\n"
        + "\n".join(calls)
        + "\n}\n"
    )

    parts = [main, f"-- go.mod --\nmodule {MODULE}\n\ngo 1.23\n"]
    for f in sorted(pkg_dir.glob("*.go")):
        if not f.name.endswith("_test.go"):
            parts.append(f"-- {pkg}/{f.name} --\n{read(f)}")
    return "".join(parts).encode("utf-8")


def snippet_id(body: bytes) -> str:
    b = base64.urlsafe_b64encode(hashlib.sha256(SALT + body).digest()).decode()
    n = 11
    while n <= len(b) and b[n - 1] == "_":
        n += 1
    return b[:n]


def share(body: bytes) -> str:
    req = urllib.request.Request(SHARE_URL, data=body, method="POST")
    # 必须声明纯文本；默认的表单类型会让服务端把代码按 URL 编码存储
    req.add_header("Content-Type", "text/plain; charset=utf-8")
    with urllib.request.urlopen(req, timeout=30) as resp:
        return resp.read().decode().strip()


def main() -> None:
    update = "--update" in sys.argv[1:]
    stale = 0
    for md in sorted((ROOT / "docs").rglob("*.md")):
        text = md.read_text(encoding="utf-8")
        links = LINK_RE.findall(text)
        if not links:
            continue
        pkgs = set(SNIPPET_RE.findall(text))
        if len(pkgs) != 1:
            sys.exit(f"{md}: 需恰好引用一个代码包，实际 {sorted(pkgs)}")
        pkg = pkgs.pop()
        body = build(pkg)
        want = f"https://go.dev/play/p/{snippet_id(body)}"
        rel = md.relative_to(ROOT).as_posix()
        if all(link == want for link in links):
            print(f"ok      {rel}")
            continue
        if not update:
            print(f"STALE   {rel}  (运行 --update 刷新)")
            stale += 1
            continue
        got = f"https://go.dev/play/p/{share(body)}"
        if got != want:
            sys.exit(f"{rel}: 服务端 ID {got} 与本地算法 {want} 不一致")
        md.write_text(LINK_RE.sub(want, text), encoding="utf-8", newline="\n")
        print(f"updated {rel}  {want}")
    sys.exit(1 if stale else 0)


if __name__ == "__main__":
    main()
