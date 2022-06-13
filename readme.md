
在开始之前首先在命令行运行如下脚本
---

```python
python -m venv venv
.\venv\Scripts\activate.bat
.\venv\Scripts\python.exe -m pip install -r requirements.txt
```

修改`get_data.py`文件中的 `id_info`中的字段值

```go
go mod tidy
set GOOS=windows GOARCH=amd64
go build -o cron-gen-excel.exe main.go
```

执行`cron-gen-excel.exe`，运行结束后，excel文件位于当前目录下的`output_file`中。