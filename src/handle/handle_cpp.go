package handle

import (
	"bytes"
	"net/http"
	"os/exec"

	"github.com/AVENTER-UG/vmm-agent/src/types"
	"github.com/labstack/echo/v4"
)

func CPPHandler(c echo.Context, req *types.RunReq) error {
	// TODO handle variant

	// Compile code
	var compileStdOut, compileStdErr bytes.Buffer
	compileCmd := exec.Command("g++", "-x", "c++", "/tmp/"+req.ID, "-o", "/tmp/"+req.ID+".out")
	compileCmd.Stdout = &compileStdOut
	compileCmd.Stderr = &compileStdErr
	err := compileCmd.Run()

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
