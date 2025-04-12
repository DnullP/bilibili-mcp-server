# 简介

这是基于go开发的一款MCP Server，目前支持的sse和stdio的方式运行

目前支持的B站相关工具包括：
- 查询个人信息（可能需要登录）
- 查询个人最新关注 (需要登录)
- 搜索视频

目前正在适配个人关注和动态相关内容，希望MCP为AI提供个性化选择视频内容的能力

# 使用方式

如果你需要查看登录才能浏览的私有信息，你可以编写一个yaml格式的配置文件，并在启动mcp服务时，第一个参数传递为该配置文件的路径，如：

```shell
D:\bili\Bilimcp.exe D:\bili\config.yml
```

配置文件需要包含你的Bilibili cookie，下面是一个示例：
```yaml
Cookie: xxxxx
```

# Contributor

欢迎代码贡献，有问题请留issue