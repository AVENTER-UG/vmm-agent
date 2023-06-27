package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/AVENTER-UG/vmm-agent/src/handle"
	"github.com/AVENTER-UG/vmm-agent/src/types"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type Language int

type CustomValidator struct {
	validator *validator.Validate
}

// BuildVersion
var BuildVersion string

// GitVersion is the revision and commit number
var GitVersion string

const (
	Python = iota + 1
	C
	Cpp
	Golang
	WASM
	Bash
	Command
)

func (s Language) String() string {
	return toString[s]
}

var toString = map[Language]string{
	Python:  "python",
	C:       "c",
	Cpp:     "cpp",
	Golang:  "go",
	WASM:    "wasm",
	Bash:    "bash",
	Command: "command",
}

var toID = map[string]Language{
	"python":  Python,
	"c":       C,
	"cpp":     Cpp,
	"go":      Golang,
	"wasm":    WASM,
	"bash":    Bash,
	"command": Command,
}

// MarshalJSON marshals the enum as a quoted json string
func (s Language) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *Language) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	if toID[j] == 0 {
		return errors.New("invalid language")
	}
	*s = toID[j]
	return nil
}

func main() {

	logrus.WithField("func", "main").Println("vmm-agent build " + BuildVersion + " git " + GitVersion)
	logrus.SetLevel(logrus.DebugLevel)

	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/run", handleCodeRun)
	e.GET("/health", health)
	e.GET("/shutdown", shutdown)

	logrus.WithField("func", "main").Info(e.Start(":8085"))
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func shutdown(c echo.Context) error {
	c.Echo().Server.Close()
	return c.String(http.StatusOK, "OK")
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func handleCodeRun(c echo.Context) error {
	req := new(types.RunReq)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	err = c.Validate(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Write code to file
	f, err := os.Create("/tmp/" + req.ID)

	if err != nil {
		logrus.WithField("func", "main.handleCodeRun").WithError(err).Error()
		return c.JSON(http.StatusInternalServerError, types.RunCRes{
			Stdout: "",
			Stderr: err.Error(),
		})
	}

	defer f.Close()

	// use code
	if req.Code != "" {
		_, err = f.WriteString(req.Code)

		if err != nil {
			logrus.WithError(err).Error()
			return c.JSON(http.StatusInternalServerError, types.RunCRes{
				Stdout: "",
				Stderr: err.Error(),
			})
		}
	}

	// use precompiles binary
	req.File, err = c.FormFile("file")
	if err == nil {

		logrus.WithField("func", "main.handleCodeRun").Infof("Uploaded File: %+v\n", req.File.Filename)
		logrus.WithField("func", "main.handleCodeRun").Infof("File Size: %+v\n", req.File.Size)
		logrus.WithField("func", "main.handleCodeRun").Infof("MIME Header: %+v\n", req.File.Header)

		src, err := req.File.Open()
		if err != nil {
			logrus.WithField("func", "main.handleCodeRun").Error("Error during open receive file: ", err.Error())
			return c.JSON(http.StatusInternalServerError, types.RunCRes{
				Stdout: "",
				Stderr: err.Error(),
			})
		}

		defer src.Close()

		if _, err = io.Copy(f, src); err != nil {
			logrus.WithField("func", "main.handleCodeRun").Error("Error during write receive file: ", err.Error())
			return c.JSON(http.StatusInternalServerError, types.RunCRes{
				Stdout: "",
				Stderr: err.Error(),
			})
		}
	}

	// Call language handler
	switch toID[req.Language] {
	case WASM:
		return handle.WASMHandler(c, req)
	case Python:
		return handle.PythonHandler(c, req)
	case C:
		return handle.CHandler(c, req)
	case Cpp:
		return handle.CPPHandler(c, req)
	case Golang:
		return handle.GoLangHandler(c, req)
	case Bash:
		return handle.BashHandler(c, req)
	case Command:
		return handle.CommandHandler(c, req)
	default:
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid language")
	}
}
