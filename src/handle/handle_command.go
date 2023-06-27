package handle

import (
	"github.com/AVENTER-UG/vmm-agent/src/types"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func CommandHandler(c echo.Context, req *types.RunReq) error {
	logrus.WithField("func", "main.CommandHandler").Info("Execute Command: ", req.ID)
	return ExecCmd(c, "/bin/bash", "-c", req.Command)
}
