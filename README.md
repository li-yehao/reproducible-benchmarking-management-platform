
<div align=center>
<img src="https://img.shields.io/badge/golang-1.16-blue"/>
<img src="https://img.shields.io/badge/gin-1.7.0-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.2.25-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.0.1-green"/>
<img src="https://img.shields.io/badge/gorm-1.22.5-red"/>
</div>

English | [简体中文](./README.md)

[gitee](https://gitee.com/pixelmax/gin-vue-admin): https://gitee.com/pixelmax/gin-vue-admin

[github](https://github.com/flipped-aurora/gin-vue-admin): https://github.com/flipped-aurora/gin-vue-admin

[Vue3 version branch address](https://github.com/flipped-aurora/gin-vue-admin/tree/vue3Develop): https://github.com/flipped-aurora/gin-vue-admin/tree/vue3Develop

[Approval flow branch](https://github.com/flipped-aurora/gin-vue-admin/tree/gva_workflow): https://github.com/flipped-aurora/gin-vue-admin/tree/gva_workflow

# Project Guidelines
[Online Documentation](https://www.gin-vue-admin.com/) : https://www.gin-vue-admin.com/

[From the environment to the deployment of teaching videos](https://www.bilibili.com/video/BV1fV411y7dT)

[Development Steps](https://www.gin-vue-admin.com/guide/start-quickly/env.html) (Contributor:  <a href="https://github.com/LLemonGreen">LLemonGreen</a> And <a href="https://github.com/fkk0509">Fann</a>)

## 1. Basic Introduction

### 1.1 Project Introduction

> Gin-vue-admin is a backstage management system based on [vue](https://vuejs.org) and [gin](https://gin-gonic.com), which separates the front and rear of the full stack. It integrates jwt authentication, dynamic routing, dynamic menu, casbin authentication, form generator, code generator and other functions. It provides a variety of sample files, allowing you to focus more time on business development.

[Online Demo](http://demo.gin-vue-admin.com): http://demo.gin-vue-admin.com

username：admin

password：123456

## 2. Getting started

```
- node version > v8.6.0
- golang version >= v1.14
- initialization project: different versions of the database are not initialized. See synonyms at initialization https://www.gin-vue-admin.com/docs/first
- Replace the Qiniuyun public key, private key, warehouse name and default url address in the project to avoid data confusion in the test file.
```


### 2.1 web project

```bash
# enter the project directory
cd web

# develop
npm run serve
```

### 2.2 Server

```bash
# enter the project directory
cd server

# develop
go run main.go
```
