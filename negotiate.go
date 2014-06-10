package negotiate

import (
  "net/http"

  "github.com/K-Phoen/negotiation"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/encoder"
)

func NegotiateFormat(negotiators map[string](encoder.Encoder)) martini.Handler {
  alternatives := make([]string, 0, len(negotiators))

  for key := range negotiators {
    alternatives = append(alternatives, key)
  }

  return func(r *http.Request, c martini.Context, w http.ResponseWriter) {
    if len(r.Header["Accept"]) == 0 {
      return
    }

    alternative, err := negotiation.NegotiateAccept(r.Header["Accept"][0], alternatives)

    if err != nil || negotiators[alternative.Value] == nil {
      return
    }

    w.Header().Set("Content-Type", alternative.Value)
    c.MapTo(negotiators[alternative.Value], (*encoder.Encoder)(nil))
  }
}
