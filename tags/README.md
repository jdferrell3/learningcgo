# Experimenting with Build Tags

```
$ go build -tags "debug get"
$ ./tags
This is main()
DEBUG = true
Performing GET action
$ go build -tags "put"
$ ./tags
This is main()
DEBUG = false
Performing PUT action
```