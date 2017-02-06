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
	case errhtp.InternalServerError:
		//@TODO
	
	case errhtp.ClientError:
		//@TODO
	case errhtp.ServerError:
		//@TODO
	
	default:
		//@TODO
	}
*/
package errhttp
