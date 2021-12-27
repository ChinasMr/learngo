# Dependency Injection

We are using the dependency injection design principle. In practice, that means we pass in whatever each component needs. This style of design lends itself to writing easily tested code and makes it easy to swap out one dependency with another.
One downside to dependency injection is the need for so many initialization steps, so we got wire
Wire: Automated Initialization in Go
Installing 
```shell
go get github.com/google/wire/cmd/wire
wire
```