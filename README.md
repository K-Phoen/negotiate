negotiate
=========

Content negotiation middleware for Martini.

It's a simple wrapper to the [negotiation library](https://github.com/K-Phoen/negotiation) and
[martini-contrib/encoder](https://github.com/martini-contrib/encoder) service.

## Usage

Here is a ready to use example:

```go
package main

import (
  "github.com/K-Phoen/negotiate"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/encoder"
  "log"
  "net/http"
)

type Some struct {
  Login    string `json:"login"`
  Password string `json:"password" out:"false"`
}

func main() {
  m := martini.New()
  route := martini.NewRouter()

  // create a format -> encoder map
  negotiators := make(map[string]encoder.Encoder)
  negotiators["application/xml"] = encoder.XmlEncoder{}
  negotiators["application/json"] = encoder.JsonEncoder{}

  // use the middleware
  m.Use(negotiate.NegotiateFormat(negotiators))

  // and the right encoder will be automatically injected
  route.Get("/test", func(enc encoder.Encoder) (int, []byte) {
    result := &Some{"awesome", "hidden"}
    return http.StatusOK, encoder.Must(enc.Encode(result))
  })

  m.Action(route.Handle)

  log.Println("Waiting for connections...")

  if err := http.ListenAndServe(":8000", m); err != nil {
    log.Fatal(err)
  }
}
```

## ToDo

  * provide tools to negotiate other things (language for instance)
  * write tests

## License

This library is released under the MIT License. See the bundled LICENSE file for
details.
