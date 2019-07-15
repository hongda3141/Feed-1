package consumer

import "net/http"

func main() {

	panic(http.ListenAndServe(":7003", nil))
}
