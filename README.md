https://golang.org/pkg/net/http/#Response

Reader Interface:

* https://golang.org/pkg/io/#ReadCloser
* https://golang.org/pkg/io/#Reader
* https://golang.org/pkg/io/#Closer

The Reader interface is an interface that allows you to easily work with a lot of different types of data. Since Go cares about what type you pass, the Reader interface allows you to convert any data into a slice of bytes.

```go

...

bs := make([]byte, 99999) // create a byte slice with an arbitrary large number of spaces
resp.Body.Read(bs)        // take the response body and use the Reader .Read interface to feed the bs the Body data
fmt.Println(string(bs))   // convert bytes to a string and print

...

```

Writer interface describes something that can take info and send it out of ourr program

https://golang.org/pkg/io/#Writer