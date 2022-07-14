package main

import (
	"github.com/labstack/echo/v4"
)

func pythonHandler(c echo.Context, req *runReq) error {
	return execCmd(c, "python3", "/tmp/"+req.ID)
}
