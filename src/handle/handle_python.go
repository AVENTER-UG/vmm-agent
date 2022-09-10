package handle

import (
	"github.com/AVENTER-UG/vmm-agent/src/types"
	"github.com/labstack/echo/v4"
)

func PythonHandler(c echo.Context, req *types.RunReq) error {
	return ExecCmd(c, "python3", "/tmp/"+req.ID)
}
