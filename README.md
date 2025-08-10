# go-errhttp

Package **errhttp** provides errors and types that make dealing with HTTP response errors easier, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-errhttp

[![GoDoc](https://godoc.org/github.com/reiver/go-errhttp?status.svg)](https://godoc.org/github.com/reiver/go-errhttp)

## Example

Here is an example of wrapping an error:

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

Here is an example of using one of the package global variable errors:

```go
	import "github.com/reiver/go-errhttp"
	
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

Here is another example, where it used the `.ErrHTTP()` method to get the HTTP response status code:

```go
	import "github.com/reiver/go-errhttp"
	
	// ...
	
	return errhttp.ErrBadRequest
	
	// ...
	
	switch casted := err.(type) {
	case errhttp.ClientError:
		statuscode := casted.ErrHTTP()
		
		http.Error(responsewriter, http.StatusText(statuscode), statuscode)
		return
	case errhttp.ServerError:
		statuscode := casted.ErrHTTP()

		http.Error(responsewriter, http.StatusText(statuscode), statuscode)
		return
	default:
		//@TODO
	}
```

Here is an example of returning an error based on an HTTP status-code:

```golang

	import "github.com/reiver/go-errhttp"
	
	// ...

	err := errhttp.Return(resp.StatusCode)
```

## Import

To import package **errhttp** use import code like the follownig:

```
import "github.com/reiver/go-errhttp"
```

## Installation

To install package **errhttp** do the following:

```
GOPROXY=direct go get github.com/reiver/go-errhttp
```

## Author

Package **errhttp** was written by [Charles Iliya Krempeaux](http://reiver.link/)
