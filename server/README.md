# server

This is golang project(backend).  
This is `super-simple-serverside-program`. So, this program's mission is only return a simple response.  
And, I know this program is very bad structure and solution.

# Setup
This package is build by `Makefile` in parent directory.  
If you want to run in host machine.

```
$ cd server
$ go build -o ./bin/server .
$ ./bin/server
```

But, this project guarantees HotReload.
If you want to run in local environment, I prefer to run `make dstart` in parent directory.
