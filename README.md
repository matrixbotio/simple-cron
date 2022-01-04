# simple-cron
Library to run func by cron. Also to limit func runtime. Written in Golang

### usage example

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

---

![image](https://github.com/Sagleft/Sagleft/raw/master/image.png)

### :globe_with_meridians: [Telegram канал](https://t.me/+VIvd8j6xvm9iMzhi)
