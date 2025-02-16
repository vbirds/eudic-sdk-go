# Eudic SDK for Go (中文说明)

这是一个用于欧路词典 API 的 Go 客户端库。它提供了一个简单的接口来管理生词本和单词。

## 安装

```bash
go get github.com/vbirds/eudic-sdk-go
```

## 使用方法

### 初始化客户端

```go
client := eudic.NewClient("YOUR_API_KEY")
```

### 生词本操作

```go
// 获取所有生词本
lists, err := client.GetWordLists()

// 创建新生词本
newList, err := client.CreateWordList("我的生词本")

// 向生词本添加单词
err := client.AddWordToList(listID, "hello")
```

### 单词管理

```go
// 查询单词
word, err := client.LookupWord("hello")

// 获取单词发音
audio, err := client.GetWordAudio("hello")
```

## 注意事项

- API 密钥需要在欧路词典官网申请
- 请合理使用 API，遵守使用频率限制
- 建议在生产环境中使用错误处理机制

## 贡献指南

欢迎提交 Pull Request 和 Issue。在提交代码前，请确保：

1. 代码经过格式化
2. 单元测试通过
3. 符合 Go 编码规范

## 许可证

MIT License
