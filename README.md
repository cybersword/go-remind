# go-remind
知行合一，不忘初心

## todo

1. main.go作为应用入口，初始化基本资源，配置信息，监听端口。
2. POST请求支持 `/controller/action/` + JSON请求体。
3. GET请求支持 `/controller/action/p1/v1/p2/v2/..` 格式。
4. API查询接口 `/api/list/` 、 `/api/controller/`。
5. 关键日志存入 MySQL，关键配置存入 Redis（例如，工艺流程）。
6. 支持 SQLite 小数据处理。
7. 支持复杂流程控制。

## 目录结构
目录结构：

	|——main.go         入口文件
	|——conf            配置文件和处理模块
	|——controllers     控制器入口
	|——actions         业务逻辑入口 
	|——models          数据库处理模块
	|——utils           辅助函数库
