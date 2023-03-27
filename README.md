# PTPT - Prompt To Plain Text

## 简介

> 低代码不如无代码

最近各种 ChatGPT 帮写代码的应用层出不穷，但与其让 ChatGPT 帮你写一个能够运行的程序，不如让 ChatGPT 直接承担各种文本生成、格式转换的工作。
对于没有编程基础的人群而言，就算拿到了能用的代码，如何让代码跑起来还需要一番折腾；对于程序员而言，重复做应用不如投入更多精力去开发 prompt。因此，我做了 PTPT，
让 ChatGPT 帮助我完成一些纯文本文件的处理工作，比如 Markdown 翻译、格式转换等。

在 PTPT 之前，我开发了一个名为 C3PO 的项目，在 C3PO 中，我需要手动去处理返回的 csv，如果想要支持 GNU po还需要写代码适配。同时，在 v2ex 论坛上，有朋友提出了很好的意见：
根本不需要做一个专门的软件来实现某个功能，开发 prompt 就足够了。这也是 PTPT 希望达到的效果。

至于为什么不使用现成的xx项目 / 使用一些 web 版的 ChatGPT 套壳工具，首先很多这些工具已经围绕 prompt 开始收费了，对他们来说 prompt 是核心资产，而我觉得 prompt 也应该是开源共享的。
另外，命令行工具能保持这个项目的操作尽可能简单，而且可以直接输出成文件，不需要再复制粘贴。最后就是，做项目带来的无限的成就感。

## 功能

- 让 ChatGPT 替你处理纯文本文件！
- 预定义 Prompt
- 方便分享和扩展的 Prompt 格式

已经支持的 prompt 一览

- [x] 🧸角色扮演（仅供娱乐）
- [x] 🧸问候语（仅供娱乐）
- [x] 📝Markdown 翻译 - translate-markdown (e.g. [Hello_translated.md](example/Hello_translated.md))
- [x] 📝csv 翻译成 csv - translate-csv (e.g. [example_translated.csv](example/example_translated.csv))
- [ ] 📝csv 翻译成 GNU po (WIP)

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