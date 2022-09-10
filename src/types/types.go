package types

import (
	"mime/multipart"
)

type (
	runReq struct {
		ID       string                `json:"id" validate:"required" form:"id"`
		Code     string                `json:"code"`
		File     *multipart.FileHeader `json:"file" form:"file"`
		Language string                `json:"language" validate:"required" form:"language"`
		Variant  string                `json:"variant" validate:"required" form:"variant"`
	}

	runCRes struct {
		Message      string `json:"message"`
		Error        string `json:"error"`
		Stdout       string `json:"stdout"`
		Stderr       string `json:"stderr"`
		ExecDuration int64  `json:"exec_duration"`
		MemUsage     int64  `json:"mem_usage"`
	}
)
