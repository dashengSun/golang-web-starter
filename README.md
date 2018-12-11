# golang-web-starter


## 背景
Web应用长期以来是Ruby、Java、PHP等开发语言的战场。
- Ruby可以实现快速原型开发，Ruby On Rails “全能”框架实现“全栈”开发，缺点有大型应用性能差、调试困难；
- Java 20多年的发展历程，各种第三方库、框架健全，运行效率高，但是随着应用的功能膨胀，臃肿的get/set方法，JVM占用大量计算机资源、性能调试困难，函数式编程不友好。
- PHP，TL;DR

本文实现了一个最小化web应用，以此来了解Golang web的生态，通过使用Docker隔离开发环境，使用Posgres持久化数据，源代码请参考[这里](https://github.com/dashengSun/golang-web-starter)

## Why Go?
- 性能优越
- 部署简单，只需要将打包好的二进制文件部署到服务器上
- 内置丰富的标准库，让程序员的生活变得简单美好
- 静态语言，类型检查
- duck typing
- goroutine将开发人员从并发编程中解放出来
- 函数作为“一等公民”
- ...


## Golang第三方框架选择
- Web框架: [Gin](https://github.com/gin-gonic/gin)，性能卓越，API友好，功能完善
- ORM: [GORM](https://github.com/jinzhu/gorm),支持多种主流数据库方言，文档清晰
- 包管理工具: [Glide](https://github.com/Masterminds/glide),类似于Ruby的bundler或者NodeJS中的npm
- 测试工具:
 - [GoConvey](https://github.com/smartystreets/goconvey),符合BDD测试风格,支持浏览器测试结果的可视化
 - [Testify](https://github.com/stretchr/testify),提供丰富的断言和Mock功能
- 数据库migration: [migrate](github.com/mattes/migrate)
- 日志工具: [Logrus](https://github.com/Sirupsen/logrus),结构化日志输出，完全兼容标准库的logger

## Dockerize开发环境

### Release base app image and db image
```
auto/release-base-image
```
```
auto/release-base-db
```

### Start development environment

```
auto/dev
```

### Run test
```
auto/test
```

## 业务实现

## 总结
Go生态之活跃令我大开眼界,著名的应用如docker, Ethereum都是使用Go编写的。使用Go进行web开发的过程，感觉和搭积木一样，一个合适的第三方库需要在多个候选库中精心筛选,众多的开源作者共同构建了一个“模块”王国。在这样的环境中,编程变成了一件很自由的事情。由于Go的标准库提供了很多内置的实用命令如`go fmt`,`go test`,让编程变得异常轻松,简直是强迫型程序员的“天堂”。
当然Go语言还处在发展过程中,也有许多不完善的地方,比如
- 缺少标准的依赖管理工具（正在开发的`dep`）
- 非中心化的依赖仓库会出现由于某个依赖被删除导致应用不可用等。
