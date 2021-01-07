package handler

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
)

type execCommandRequest struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}

type execCommandResponse struct {
	Error          string `json:"error"`
	CombinedOutput string `json:"combined_output"`
}

func execCommand(c *gin.Context) (out []byte, err error) {
	var req execCommandRequest
	err = c.Bind(&req)
	if err != nil {
		return
	}
	out, err = exec.Command(req.Name, req.Args...).CombinedOutput()
	fmt.Println("exec")
	fmt.Println(req.Name)
	fmt.Println(req.Args)
	fmt.Println("result")
	fmt.Printf("result: %v, %v", err, DecodeGBK(string(out)))
	return
}

func ExecCommand(c *gin.Context) {
	out, err := execCommand(c)

	var resp execCommandResponse
	if err != nil {
		resp.Error = err.Error()
	}
	if out != nil {
		resp.CombinedOutput = DecodeGBK(string(out))
	}

	c.JSON(http.StatusOK, &resp)
}

func DecodeGBK(s string) string {
	dec := mahonia.NewDecoder("GBK")
	return dec.ConvertString(s)
}
