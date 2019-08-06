package dispatcher

import (
	"net/http"

	"github.com/tsuna/gohbase"
)

func main() {
	client := gohbase.NewClient("localhost")
	panic(http.ListenAndServe(":7002", nil))
}
