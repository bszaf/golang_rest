This project is simple HTTP API server written in go-lang.

Building
-------
```
go build
```

Usage
-----

First server needs to be started:

```
./golang_rest server [-p listen_port]
```

When server is started, questions should be uploaded:

```
/golang_rest client post \
    --body "2+2= ?" \
    --valid "4" \
    --wrong "5" \
    --wrong "44" \
    --server "http://localhost:8000"

Question id = 2
```

Finally, quiz can be taken:

```
./golang_rest client quiz --server "http://localhost:8000"
Question: 2+2= ?
	 0) 5
	 1) 44
	 2) 4
> 2

Total valid answers: 1
You where better than: 100.00 %
```
