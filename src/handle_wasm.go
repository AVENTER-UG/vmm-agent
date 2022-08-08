package main

import (
	"github.com/labstack/echo/v4"
)

func wasmHandler(c echo.Context, req *runReq) error {
	return execCmd(c, "wasmtime", "/tmp/"+req.ID)
}
