/*
Package errhttp provides types errors that make dealing with HTTP response errors easier.

Example

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
*/
package errhttp
