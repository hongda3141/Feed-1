package consumer

import (
	"Feed/consumer/service"
	"net/http"
)

func main() {

	service.StartUp()

	panic(http.ListenAndServe(":7003", nil))
}
