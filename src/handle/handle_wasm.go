package handle

import (
	"github.com/labstack/echo/v4"
)

func wasmHandler(c echo.Context, req *runReq) error {
	logrus.WithField("func", "main.wasmHandler").Infof("Execute wasm: ", req.ID)
	return execCmd(c, "wasmtime", "/tmp/"+req.ID)
}
