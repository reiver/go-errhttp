# go-errhttp

Package **errhttp** provides types errors that make dealing with HTTP response errors easier, for the Go programming language.


## Example
```go
	import "github.com/reiver/go-errhttp"
	
	// ...
	
	if err := something(); nil != err {
		return errhttp.BadRequestWrap(err)
	}
	
	// ...
	
	switch err.(type) {
	case errhttp.BadRequest:
		//@TODO
	case errhttp.NotFound:
		//@TODO
	case errhttp.InternalServerError:
		//@TODO
	
	case errhttp.ClientError:
		//@TODO
	case errhttp.ServerError:
		//@TODO
	
	default:
		//@TODO
	}
```

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-errhttp

[![GoDoc](https://godoc.org/github.com/reiver/go-errhttp?status.svg)](https://godoc.org/github.com/reiver/go-errhttp)
