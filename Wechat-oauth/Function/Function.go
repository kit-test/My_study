package Function

import (
	"net/http"
	"io"
)

type Myhander interface {
	Hellhander()
	Nihanhander()
}
type HANDER struct{}   //(s *HANDER)

func (s *HANDER)Hellohander(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"Hello!")
}
func (s *HANDER)Nihaohander(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"nihao!")
}
