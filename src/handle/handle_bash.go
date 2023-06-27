package handle

import (
	"github.com/AVENTER-UG/vmm-agent/src/types"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func BashHandler(c echo.Context, req *types.RunReq) error {
	logrus.WithField("func", "main.BashHandler").Info("Execute bash: ", req.ID)
	return ExecCmd(c, "/bin/bash", "/tmp/"+req.ID)
}
