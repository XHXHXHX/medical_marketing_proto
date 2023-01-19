- [demo 说明](#demo-说明)
- [Get started](#get-started)
  - [Requirements](#requirements)
  - [目录说明](#目录说明)
  - [编写 .proto 文件](#编写-proto-文件)
  - [.proto 文件常见问题](#proto-文件常见问题)
    - [关于 package](#关于-package)
    - [关于 go_package](#关于-go_package)
    - [关于 proto 文件中的注释](#关于-proto-文件中的注释)
  - [关于 v1api, v2api](#关于-v1api-v2api)
  - [google.api.http 格式](#googleapihttp-格式)
  - [关于 v1iapi](#关于-v1iapi)
  - [生成代码](#生成代码)
  - [使用该项目的代码](#使用该项目的代码)
    - [说明](#说明)
    - [服务端 golang](#服务端-golang)
    - [客户端 golang](#客户端-golang)
    - [客户端 web](#客户端-web)
  - [openAPI (TODO)](#openapi-todo)
- [从该项目创建新的项目](#从该项目创建新的项目)
- [使用其他项目的 proto 文件](#使用其他项目的-proto-文件)
- [其他](#其他)
  - [参考资料](#参考资料)

# demo 说明

演示 gRPC + gRPC gateway 的 demo,同时内置一些脚本,可以基于该 demo创建新的项目.

该项目只包含 proto 文件以及生成的代码部分.


# Get started

## Requirements

本项目需要安装 `protoc` 程序, 用于自动生成代码. 一次性.

```sh
# 确认安装的版本
protoc --version
libprotoc 3.14.0
# 只要是和该版本差距不大理论上都可以.
```

如果没有安装, 尝试使用下面的命令进行安装.

```sh
./misc/install-protoc.sh
```

## 目录说明

```sh
proto/   # 编写的 .proto 文件, 和具体语言无关
gen/     # 放置生成的代码部分.
   /go/pb  # 生成的 golang 部分的代码
   # TODO 未来支持更多的语言,例如 .js, 可以直接用于前端, 进一步简化前端的负担.
tools/   # golang 依赖的一些工具链
misc/*   # 放置一些杂项
```

使用者需要重点关心的文件(夹)为 `proto` 文件夹, 所有的接口定义, 数据定义均在该目录中.

其次为 `gen`目录,下面会生成多个语言的代码, 例如 `gen/go/pb` 为生成的 golang代码.

## 编写 .proto 文件

`.proto`文件是项目的核心, 里面定义了 gRPC 接口的方法, 入参 以及 返回值.

编写完可以生产对应的代码, 既可以用于 server 端来实现, 又可以直接作为 gRPC 的 client 端直接使用.

同时,我们使用 gRPC gateway, 支持 server 端自动将 gRPC 接口转换为 HTTP 接口.做到同一套代码既可以用于 gRPC 又可以使用 HTTP 接口.

使用方可以根据自己团队的接入能力,自行选择接入方式是 gRPC 或 HTTP.

**参考资料:**

- protobuf 的书写风格, 按照[官方的建议](https://developers.google.com/protocol-buffers/docs/style)即可.


**proto 目录组织**:

所有的 `.proto` 文件均放置在 `./proto`目录下. 按照模块自己划分,见示例.  
不同的模块可以引用, 方便抽象一些公共的数据结构使用.


如下:
1. 将一些公共的数据结构,放置在了 `./proto/common`目录下 (可以按照模块放置其他目录名字中).
2. 将接口层的定义放置在了 `./proto/v1api`下.

    - 建议 接口的输入和输出接口定义放在接口定义文件中.  
    - 同时建议不同的模块的接口可以放在不同的文件中以区别.  

```sh
proto/common               # common 目录, 放置一些共享的数据接口
           /product.proto
proto/v1api                # v1api接口的定义
          /product.proto   #   v1api模块中的 product 模块接口定义
          /common.proto    #   v1api中的一些共享数据
proto/v2api                # v2api接口定义(通常不需要 v2,这里使用 v2 来演示自定义风格接口)
          /product.proto
```

## .proto 文件常见问题

### 关于 package

package 是 proto 文件在被其他的 proto 文件引用时的包名.

建议格式: `项目名字(短名字).目录1.目录2.文件名`

例如 `proto/common/product.proto` 文件, 他的 package 名字为 `package demo.common.product;`

当该文件被其他文件引用时候, 需要导入全路径, 然后使用该 package 名字.

例如  `proto/v1api/pruduct.proto`文件中

```protobuf
// 首先导入需要使用的 proto 文件
import "proto/common/product.proto";

// 使用该文件中的数据结构 demo.common.product.Product
message ListProductsResponse {
    repeated demo.common.product.Product products = 1;
}
```

### 关于 go_package

go_package 选项控制的是最终生成的 go 代码的包名. 由于 golang 是以目录作为一个 package, 而且现在的代码生成器要求 go_package 必须是全路径, 所以, go_package 的格式定义为:

`<go module 名字>/gen/go/<proto 所在的文件夹全路径>`

例如, `proto/common/*.proto`, 即 `proto/common`目录下的所有 proto 文件,他们的 go_package 都是一致的:

`option go_package = "gitlab.aiforward.cn/proto/demo/gen/go/proto/common";`

- 其中 `gitlab.aiforward.cn/proto/demo` 是 go module 的名字.
- `gen/go` 是 golang 代码的固定前缀.
- `proto/common` 是 proto 文件所在的文件夹路径.

### 关于 proto 文件中的注释

再 proto 文件中,所有的 `message` 和 rpc 的方法,都尽量加上注释, 在代码生成时会自动生成接口以及接口文档. 其中文档内容大部分来自于注释.

## 关于 v1api, v2api

正常的项目, 可以仅仅规划 `api` 即可,默认是 v1 的,再需要的时候,再创建 v2api. 因为大部分情况的项目是不需要 v2 版本的 API 的.

该项目的 v2api 是为了演示自定义风格接口来使用的.

## google.api.http 格式

- [官方的例子: a_bit_of_everything.proto](https://github.com/grpc-ecosystem/grpc-gateway/blob/master/examples/internal/proto/examplepb/a_bit_of_everything.proto) 上面有很多常用的场景.
- [google.api.http.proto 的原始定义](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto)

## 关于 v1iapi

提供给内部的（如服务端之间）API调用使用，内部API调用会对进行简单的签名验证。

## 生成代码

运行如下脚本生成代码, 如果不报错则代表生成无误.

```sh
./generate.sh
```

## 使用该项目的代码

### 说明

该项目的使用,分为多端, 例如典型的如下:

1. 服务端. 服务端需要实现这些接口.一般我们公司使用 golang 作为服务端的实现语言.
2. 客户端. 客户端指调用接口方, 一般是其他的后台服务或者前端, 直接使用该语言支持的依赖方式即可.

### 服务端 golang

直接使用 go get 来添加依赖.
```sh
go get -u gitlab.aiforward.cn/proto/demo
```

代码中需要实现 `gitlab.aiforward.cn/proto/demo/gen/go/pb` 中的server 方法. 例如

```golang
type Server struct{}

// V1CreateProduct implements gitlab.aiforward.cn/proto/demo/gen/go/pb/v1api.V1APIProductServer 
func (s *Server)V1CreateProduct(context.Context, *CreateProductRequest) (*common.Product, error) {
    ...
}
```

具体见 `http://gitlab.aiforward.cn/inf/grpcgateway-demo` 项目.

### 客户端 golang

```sh
mkdir my_test && cd my_test
go mod init
go get -u gitlab.aiforward.cn/proto/demo

# 之后便可以使用其中的 Client 方法.
```

client-demo 见 `TODO`

### 客户端 web

TODO 暂不支持, 未来可以支持.


## openAPI (TODO)

TODO

# 从该项目创建新的项目

如果要创建一个项目,可以基于该项目并且修改部分配置即可.
下面以 创建一个项目 `gitlab.aiforward.cn/proto/iot_bus` 为例.

1. gitlab 创建项目

选择 公开项目

2. 本地 clone 该项目
   
```sh
git clone http://gitlab.aiforward.cn/proto/iot_bus.git
```

3. 运行脚本

```sh
# 进入到项目目录
cd iot_bus
curl -o- -s  http://gitlab.aiforward.cn/proto/demo/raw/master/misc/new_project.sh | bash
```

4. 重新运行 generate.sh

```sh
./generate.sh
```

5. 提交代码

```sh
git add .
git commit -m 'init'
```

6. 修改 proto 文件正常开发.

见 README.md

# 使用其他项目的 proto 文件

参考 `proto/demo/shared/shared.proto` 文件.

1. 在 `tools/tools.go`  import 依赖的项目的任意一个 go package 目录.
2. go mod tidy 确保 go.mod 包含了依赖.
3. `./generate.sh` 中, 找到 EXTRA_DEPS 变量,添加依赖的 go module name.
4. 运行 `./generate.sh`

# 其他

## 参考资料

- [gRPC gateway](https://github.com/grpc-ecosystem/grpc-gateway)
