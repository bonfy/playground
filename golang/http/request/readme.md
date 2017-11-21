# http request

## 01 基本的 http request 

[01-request.go]()

设置 timeout
```Golang
url := "xxx"
timeout := time.Duration(5 * time.Second)
client := http.Client{
    Timeout: timeout,
}
resp, _ := client.Get(url)
```

## 02 parse xml

