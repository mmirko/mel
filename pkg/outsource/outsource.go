package outsource

import (
	"io"
	"log"
	"net/http"
)

type OutSource interface {
	SendGenotype() ([]byte, error)
	ReceiveGenotype([]byte) error
}

type outSourceServer struct {
	waitList chan []byte
}

func (m *outSourceServer) newFenoHandler(w http.ResponseWriter, r *http.Request) {
	result := ""
	result += `
<!DOCTYPE html>
<meta charset="utf-8">
<title>m</title>
<body>
</body>
</html>
`
	io.WriteString(w, result)
}

func (m *outSourceServer) Serve() {
	http.HandleFunc("/newfenotype", m.newFenoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
