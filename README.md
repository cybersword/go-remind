# go-remind
知行合一，不忘初心

## 摘要
1. main.go作为应用入口，初始化基本资源，配置信息，监听端口，组装请求参数，路由到 app。
2. 支持 RESTful API
3. query 格式：`/app/controller/action/version/p1/v1/p2/v2` 形式。
    * `app`,`controller`,`action`,`version` 为预置的层级，用于路由，`version`可省略。
    * 其中 `/p1/v1 ... /pn/vn` 必须成对出现。
    * query 中的参数`?k1=v1&k2=v2`，会解析到`params["FORM"]`中。
    * `POST`、`PUT`、`PATCH`请求，`Content-Type: application/x-www-form-urlencoded`的内容会一并解析到`params["FORM"]`中。
    * `POST`、`PUT`、`PATCH`请求，`Content-Type: application/json`的内容会解析到`params["JSON"]`中。
4. API查询接口 `/app` 、 `/app/controller`。
5. 关键日志存入 MySQL，关键配置存入 Redis（例如，工艺流程）。
6. 支持 SQLite 小数据处理。
7. 支持复杂流程控制。
8. 避免硬编码。
9. 支持热上线。

## 目录结构
目录结构：

	|——main.go         入口文件
	|——conf            配置文件和处理模块
	|——controllers     控制器入口
	|——actions         业务逻辑入口 
	|——models          数据库处理模块
	|——utils           辅助函数库
