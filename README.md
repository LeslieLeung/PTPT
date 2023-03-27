# PTPT - Prompt To Plain Text

## 简介

> 低代码不如无代码

最近各种 ChatGPT 帮写代码的应用层出不穷，但与其让 ChatGPT 帮你写一个能够运行的程序，不如让 ChatGPT 直接承担各种文本生成、格式转换的工作。
对于没有编程基础的人群而言，就算拿到了能用的代码，如何让代码跑起来还需要一番折腾；对于程序员而言，重复做应用不如投入更多精力去开发 prompt。因此，我做了 PTPT，
让 ChatGPT 帮助我完成一些纯文本文件的处理工作，比如 Markdown 翻译、格式转换等。

## 功能

- 让 ChatGPT 替你处理纯文本文件！
- 预定义 Prompt
- 方便分享和扩展的 Prompt 格式

已经支持的 prompt 一览

- [x] 🧸角色扮演（仅供娱乐）
- [x] 🧸问候语（仅供娱乐）
- [x] 📝Markdown 翻译
- [x] 📝csv 翻译成 csv
- [x] 📝csv 翻译成 GNU po

## 安装

```bash
go install github.com/leslieleung/ptpt
```

## 使用

首先需要初始化 `OPENAI_API_KEY`，可以通过以下方式设置。

`export OPENAI_API_KEY="sk-xxxxxx"` 或 `echo "sk-xxxxxx" > 〜/.ptptcfg`。

### 交互式

目前已经预置了几个好用的 prompt，后续会继续增加。同时也可以通过 PromptHub (WIP) 获取更多的 prompt。

```bash
> ptpt run
```
![](docs/screenshots/interactive.gif)

### 通过命令行参数
```bash
ptpt run [prompt] [inFile] [outFile]

# 使用重定向
> ptpt run translate-markdown Hello.md > Hello_tranlsated.md
# 或直接指定输出文件
> ptpt run translate-markdown Hello.md Hello_tranlsated.md
```

## 创造你自己的 prompt

### 通过交互式创建(WIP)
```bash
> ptpt prompt create
```

### 格式说明

```yaml
version: v0 # 版本号，暂时为v0
prompts: # 定义的 prompt
  - name: role-yoda # prompt 名称
    description: "Role Play as Yoda" # prompt 描述
    system: You are Yoda master from Star Wars, speak in his tongue you must. # system 指令
  - name: role-spock
    description: "Role Play as Spock"
    system: You are Spock from Star Trek, you must speak in his tongue.
```

通过下载别人分享的 prompt，保存在 `~/.ptpt/prompt` 目录下后，即可使用更多的 prompt。

## Roadmap
- [ ] 支持代理配置
- [ ] 支持ChatGPT参数配置
- [ ] PromptHub - 通过 yaml 文件分享 prompt
- [ ] 支持更多的 prompt

本项目暂时不会专注于：
- 连续对话、聊天记录
- 复杂花哨的命令行交互

## Credits
本项目灵感来源于 [sigoden/aichat](https://github.com/sigoden/aichat)，该项目使用 Rust 语言，由于能力有限，我想用自己熟悉的技术栈做一个自己使用的版本。