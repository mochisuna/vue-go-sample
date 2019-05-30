# vue-go-sample
This is simple web application with vue.js and golang.

These applications supports hot reload.

# USAGE
Run `make dstart` command.

```bash
$ make dstart
$ curl localhost:38080/monsters -vvv | jq .
$ open http://localhost:39000
```

# Clean up
If you clean all containers, RUN `make dclean` command

```bash
$ make dclean
docker clean
a475fb1796ff
bf5eab629891
570c3ded5f0d
54111388d61f
```