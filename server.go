// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package vault

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/adjspecies/vault/api"
)

type Server struct {
	r *mux.Router
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}

func (s Server) Close() error {
	return nil
}

func NewServer() (Server, error) {
	r := mux.NewRouter()
	if err := api.Register(r); err != nil {
		return Server{}, err
	}
	s := Server{r: r}
	return s, nil
}
