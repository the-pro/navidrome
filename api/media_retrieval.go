package api

import (
	"io"
	"os"

	"github.com/astaxie/beego"
	"github.com/deluan/gosonic/api/responses"
	"github.com/deluan/gosonic/engine"
	"github.com/deluan/gosonic/utils"
)

type MediaRetrievalController struct {
	BaseAPIController
	cover engine.Cover
}

func (c *MediaRetrievalController) Prepare() {
	utils.ResolveDependencies(&c.cover)
}

func (c *MediaRetrievalController) GetAvatar() {
	var f *os.File
	f, err := os.Open("static/itunes.png")
	if err != nil {
		beego.Error(err, "Image not found")
		c.SendError(responses.ERROR_DATA_NOT_FOUND, "Avatar image not found")
	}
	defer f.Close()
	io.Copy(c.Ctx.ResponseWriter, f)
}

func (c *MediaRetrievalController) GetCover() {
	id := c.RequiredParamString("id", "id parameter required")
	size := c.ParamInt("size", 0)

	err := c.cover.Get(id, size, c.Ctx.ResponseWriter)

	switch {
	case err == engine.ErrDataNotFound:
		beego.Error(err, "Id:", id)
		c.SendError(responses.ERROR_DATA_NOT_FOUND, "Directory not found")
	case err != nil:
		beego.Error(err)
		c.SendError(responses.ERROR_GENERIC, "Internal Error")
	}
}