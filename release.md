# 将go的项目发到github作成开源公用包

> 版本：v1.0.0，需要是3位的版本号，不然会报错

# git标签
```
# 添加标签
git tag -a v1.0.0 -m "v1.0.0"
git push origin v1.0.0

# 删除
git tag -d v1.0.0
git push origin :refs/tags/v1.0.0
```

# go mod版本化
`go.sum`文件是Go模块的一个关键文件，用于记录项目依赖模块的校验和信息，以确保构建时使用的依赖模块版本的一致性和完整性。`go.sum`文件是由Go工具自动生成和维护的，其生成过程如下：

1. **初始化项目**：
    - 当您在一个新的Go项目中使用Go模块时，可以通过运行`go mod init <module-name>`来初始化项目，并生成`go.mod`文件。

2. **添加或更新依赖**：
    - 当您使用`go get`、`go build`、`go run`等命令引入新的依赖模块或更新已有依赖模块时，Go工具会自动更新`go.mod`文件中的依赖信息，并生成或更新`go.sum`文件。

3. **下载依赖**：
    - 当执行`go build`、`go run`等命令时，Go工具会根据`go.mod`文件中的依赖信息下载相应的依赖模块，并计算每个模块的校验和。

4. **生成`go.sum`文件**：
    - 在下载依赖模块的过程中，Go工具会计算每个依赖模块的内容的哈希值，并将这些哈希值记录在`go.sum`文件中。

5. **校验依赖完整性**：
    - 每次构建项目时，Go工具会读取`go.sum`文件中记录的依赖模块的哈希值，与实际下载的依赖模块进行比对，以确保依赖模块的完整性和一致性。

总之，`go.sum`文件是通过计算依赖模块的哈希值生成的，用于记录项目依赖的校验和信息，以确保依赖模块的安全性和完整性。在日常开发中，无需手动编辑`go.sum`文件，Go工具会自动维护和更新该文件。