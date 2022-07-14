package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
)

func copy(src string, dst string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, data, 0644)
	return err
}

func wasmHandler(c echo.Context, req *runReq) error {
	return execCmd(c, "wasmtime", "/tmp/"+req.ID)
}
