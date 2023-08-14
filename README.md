# go-errhttp

Package **errhttp** provides errors and types that make dealing with HTTP response errors easier, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-errhttp

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-errhttp?status.svg)](https://godoc.org/sourcecode.social/reiver/go-errhttp)

## Example

Here is an example of wrapping an error:

```go
	import "sourcecode.social/reiver/go-errhttp"
	
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

Here is an example of using one of the package global variable errors:

```go
	import "sourcecode.social/reiver/go-errhttp"
	
	// ...
	
	return errhttp.ErrBadRequest
	
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
