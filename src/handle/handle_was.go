package handle

import (
	"github.com/AVENTER-UG/vmm-agent/src/types"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func WASMHandler(c echo.Context, req *types.RunReq) error {
	logrus.WithField("func", "main.wasmHandler").Info("Execute wasm: ", req.ID)
	return ExecCmd(c, "wasmtime", "/tmp/"+req.ID)
}
