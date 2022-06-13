

# tiktok


<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />

<p align="center">
  <a href="https://github.com/mrxuexi/tiktok/">
    <img src="https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/a2931789a5fe4bfb9c1dfa8775c9e970~tplv-k3u1fbpfcp-zoom-in-crop-mark:1304:0:0:0.awebp?" alt="Logo" width="200" height="200">
  </a>
  <h3 align="center">抖音 from 字节后端青训营</h3>
  <p align="center">
    基于架构 ver 0.0.1
    <br />
    <a href="https://github.com/mrxuexi/tiktok"><strong>探索本项目的文档 »</strong></a>
    <br />
    <br />
    <a href="https://xzi09smrpn.feishu.cn/docx/doxcnVfEWSyncDiFXVVze0Gx0Vg">查看Demo</a>
    ·
    <a href="https://github.com/mrxuexi/tiktok/issues">报告Bug</a>
    ·
    <a href="https://github.com/mrxuexi/tiktok/issues">提出新特性</a>
  </p>
### 里程碑一    <-2022年6月12日
项目功能的基本实现

### 里程碑二
引入缓存，性能优化，代码重构优化

### 里程碑三
引入 RPC 框架，进行服务分离与治理

### 里程碑四
审核系统的开发、消息推送系统的开发

### 里程碑五
请求链路追踪、日志平台引入

## 目录

- [上手指南](#上手指南)
  - [开发前的配置要求](#开发前的配置要求)
  - [安装步骤](#安装步骤)
- [文件目录说明](#文件目录说明)
- [开发的架构](#开发的架构)
- [部署](#部署)
- [使用到的框架](#使用到的框架)
- [贡献者](#贡献者)
  - [如何参与开源项目](#如何参与开源项目)
- [版本控制](#版本控制)
- [作者](#作者)
- [鸣谢](#鸣谢)

### 上手指南



###### 开发前的配置要求


###### **安装步骤**

1. Clone the repo <p>
2. go build <p>
3. ./app config.yaml`s path

```sh
git clone https://github.com/Mrxuexi/tiktok.git
cd gateway
go build -o xxxx main.go
./xxxx config.yam
```
##### DB DESIGN
![image](https://user-images.githubusercontent.com/56754549/173394004-3cd278a6-0b75-470e-8c81-2c1a1f07642f.png)

### 文件目录说明
eg:

```shell
filetree
├─base（公共基础库，封装一些通用的逻辑）
│  ├─logger（zap 日志）
│  ├─common（一些公共响应码和错误的封装）
│  ├─jwt（jwt操作的封装）
│  ├─io（请求和响应的封装，未来将抽离）
│  ├─middlewares（中间件）
│  └─mykafka（封装 Kafka 操作）
│  └─mymysql（封装 MySQL 公共 CRUD 操作）
│  └─myredis（封装 Redis 公共操作）
│  └─snowflake（分布式 id 生成器）
├─bff
│  ├─accountbff（账号相关 BFF，预计 v0.0.3）
│  ├─vediobff（视频相关 BFF，预计 v0.0.3）
├─deploy（部署）
├─design（设计文档）
├─gateway（API 网关,v0.0.1 的项目入口）
├─images 
├─service（各个服务，v0.0.1 的 logic 层实现在此处）
│  ├─commentsrv
│  ├─favoritesrv
│  ├─publishsrv
│  ├─relationsrv
│  └─usersrv
├─setting（viper 配置）
```





### 开发的架构 

![tiktok 项目架构](https://user-images.githubusercontent.com/56754549/173394272-6bef5aca-fd7d-41ca-ab49-1fb3efa7dd91.png)


请阅读[ARCHITECTURE.md](https://github.com/mrxuexi/tiktok/) 查阅为该项目的架构。

### 部署

暂无

### 使用到的技术

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Redis](https://redis.io/)
- [MySQL](https://www.mysql.com/)
- [Kubernetes](https://kubernetes.io/)
- [gRPC](https://grpc.io/)
- [Nacos](https://nacos.io)
- ......

### 贡献者

请阅读 **CONTRIBUTING.md** 查阅为该项目做出贡献的开发者。

#### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



### 版本控制

该项目使用Git进行版本管理。您可以在 repository 参看当前可用版本。

### 作者

mdowellrlph@gmail.com

博客:[Mrxuexi](https://mrxuexi.com)  &ensp; qq: 916516214

949151128@qq.com

博客:[空月](https://konyue.site/)  &ensp; qq: 949151128

xxh@xxxx

博客:[name](https://example.com)  &ensp; qq: 212222222    

 *剩余内容由目前 7 人核心开发小组填写*

### 版权说明

该项目签署了MIT 授权许可，详情请参阅 [LICENSE.txt](https://github.com/mrxuexi/tiktok/LICENSE.txt)

### 鸣谢


- [字节跳动后端青训营](https://youthcamp.bytedance.com/)

<!-- links -->
[your-project-path]:mrxuexi/tiktok
[contributors-shield]: https://img.shields.io/github/contributors/mrxuexi/tiktok.svg?style=flat-square
[contributors-url]: https://github.com/mrxuexi/tiktok/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/mrxuexi/tiktok.svg?style=flat-square
[forks-url]: https://github.com/mrxuexi/tiktok/network/members
[stars-shield]: https://img.shields.io/github/stars/mrxuexi/tiktok.svg?style=flat-square
[stars-url]: https://github.com/mrxuexi/tiktok/stargazers
[issues-shield]: https://img.shields.io/github/issues/mrxuexi/tiktoksvg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/mrxuexi/tiktok.svg
[license-shield]: https://img.shields.io/github/license/mrxuexi/tiktok.svg?style=flat-square
[license-url]: https://github.com/mrxuexi/tiktok/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/xxxx



