package dispatcher

import "net/http"

func main() {

	panic(http.ListenAndServe(":7002", nil))
}
