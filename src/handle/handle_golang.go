package handle

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/AVENTER-UG/vmm-agent/src/types"
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

func GoLangHandler(c echo.Context, req *types.RunReq) error {
	// TODO handle variant

	err := copy("/tmp/"+req.ID, "/tmp/"+req.ID+".go")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.RunCRes{
			Message: "Failed to copy file",
			Error:   err.Error(),
		})
	}
	// Compile code
	var compileStdOut, compileStdErr bytes.Buffer
	compileCmd := exec.Command("go", "build", "-o", "/tmp/"+req.ID+".out", "/tmp/"+req.ID+".go")
	compileCmd.Stdout = &compileStdOut
	compileCmd.Stderr = &compileStdErr
	err = compileCmd.Run()

	if err != nil {
		return c.JSON(http.StatusBadRequest, types.RunCRes{
			Message: "Failed to compile",
			Error:   err.Error(),
			Stdout:  compileStdOut.String(),
			Stderr:  compileStdErr.String(),
		})
	}

	// Run executable
	return ExecCmd(c, "/tmp/"+req.ID+".out")
}
