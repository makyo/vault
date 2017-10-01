package vault

import (
	"github.com/juju/gocharm/charmbits/httpservice"
	"github.com/juju/gocharm/hook"
)

// RegisterHooks registers each of the hooks necessary to run the vault service.
func RegisterHooks(r *hook.Registry) {
	var vc vaultCharm
	vc.svc.Register(r, "vault", "vaultserver", vc.handler)
	r.RegisterHook("start", vc.start)
	r.RegisterHook("stop", vc.stop)
}

type vaultCharm struct {
	svc httpservice.Service
}

func (vc *vaultCharm) start() error {
	return vc.svc.Start("Starting server...")
}

func (vc *vaultCharm) stop() error {
	return vc.svc.Stop()
}

func (vc *vaultCharm) handler(s string) (httpservice.Handler, error) {
	return NewServer()
}
