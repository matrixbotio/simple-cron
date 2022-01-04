# simple-cron

usage example

```go
var myfunc = func() {
	// func example
	time.Sleep(time.Second * 10)
}

isTimeIsUP := simplecron.NewRuntimeLimitHandler(
	time.Second * 2,
	myfunc,
).Run()
if isTimeIsUP {
	// handle timeout
}
```

![image](image.png)

### :globe_with_meridians: [Telegram канал](https://t.me/+VIvd8j6xvm9iMzhi)
