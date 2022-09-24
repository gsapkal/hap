package hap

import (
	"github.com/gsapkal/hap/log"

	"net/http"
)

func (srv *Server) identify(res http.ResponseWriter, req *http.Request) {
	if srv.isPaired() {
		log.Info.Printf("request only valid if unpaired")
		JsonError(res, JsonStatusInsufficientPrivileges)
		return
	}

	if srv.a.IdentifyFunc != nil {
		srv.a.IdentifyFunc(req)
	}

	res.WriteHeader(http.StatusNoContent)
}
