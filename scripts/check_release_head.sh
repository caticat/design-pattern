#!/usr/bin/env bash
set -euo pipefail

# 在发布作业拿到并发锁后，跳过已被 master 新提交取代的构建。
head=$(git ls-remote --exit-code origin refs/heads/master)
if [[ ${head%%$'\t'*} == "$GITHUB_SHA" ]]; then
  printf 'current=true\n' >> "$GITHUB_OUTPUT"
else
  printf 'current=false\n' >> "$GITHUB_OUTPUT"
  printf '%s\n' 'master 已有更新，跳过本次旧版本发布。'
fi
