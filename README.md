# Qing Zhuo

A clean, modern blog theme for hugo. [Demo](https://eindex.github.io/qingzhuo)

[![Status](https://github.com/EINDEX/qingzhuo/actions/workflows/build.yaml/badge.svg)](https://github.com/EINDEX/qingzhuo/actions/workflows/build.yaml)


## Quick Start

### AS Theme (recommended)
```shell
git submodule add https://github.com/eindex/qingzhuo themes/qingzhuo
cp themes/qingzhuo/config.toml .
```

```shell
hugo mod npm pack
npm install
```

### As Project

```
git clone https://github.com/eindex/qingzhuo blog && cd blog
npm install
```

### Before Build

Edit `config.toml` content to adjust the theme and enable functions.

(AS Theme) Uncomment theme in `config.toml` file.

run `hugo` to gerenate the static files or `hugo server` to start server.

## How to update

### As Theme (recommended)

```shell
git submodule update --init --remote
```

### As Project 

Merge changes form [qingzhuo](https://github.com/eindex/qingzhuo) via git.

## Functions

- Latex
- Code Hightlight

- Giscus
- Outdateing Warning
- NavBar menu customizing
- RSS
- TOC Support

- Analytics
    - Google Analytics
    - Microsoft Clarify
    - Baidu Analytics

- Heti style
- Chinese ICP Number
- BuSuanZi Visit Counter

### TODO
- i18n support
- gallery page for photos
- enhancement 404 page

## Who using this theme?

- [EINDEX's Blog](https://eindex.me)

If you are using this theme, please create pull request adding your site here.

## Screen Shots

### Homepage
![](/docs/imgs/index-page.jpg)

### Archive
![](/docs/imgs/archive-page.jpg)

## Thanks
- hugo
- tailwindcss
- heti

## License
GPL v3 License
