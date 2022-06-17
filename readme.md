
在开始之前首先在命令行运行如下脚本
---
```powershell
Set-ExecutionPolicy -ExecutionPolicy Bypass -Scope CurrentUser
```

```python
python -m pip install virtualenv -i https://pypi.tuna.tsinghua.edu.cn/simple
virtualenv --version
virtualenv venv
.\venv\Scripts\activate.ps1
.\venv\Scripts\python.exe -m pip install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple
```

修改`config.toml`文件中的字段

```go
go mod tidy
set GOOS=windows GOARCH=amd64
go build -o cron-gen-excel.exe main.go
```

执行`cron-gen-excel.exe`，运行结束后，excel文件位于当前目录下的`output_file`中。