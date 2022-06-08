package service

import (
	"etweb/internal/models"
	"etweb/utils"
	"fmt"
	"net/http"
	"path"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	// 创建seesion
	session := sessions.Default(c)

	// session过期
	if session.Get("currentUser") == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":     utils.GetErrMsg(utils.UPLOAD_PICTRUE_FAILD),
			"retcode": utils.UPLOAD_PICTRUE_FAILD,
		})
		return
	}

	file, err := c.FormFile("file")
	// 上传图片失败
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":     utils.GetErrMsg(utils.SAVE_PICTRUE_FAILD),
			"retcode": utils.SAVE_PICTRUE_FAILD,
		})
		return
	}
	// 图片存放地址
	filePath := path.Join(utils.PicFileAddr, file.Filename)
	// 保存图片
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":     utils.GetErrMsg(utils.UPLOAD_PICTRUE_FAILD),
			"retcode": utils.UPLOAD_PICTRUE_FAILD,
		})
		return
	}

	//fileType := c.PostForm("type")

	// 更新user
	oldUser := session.Get("currentUser").(models.User)
	newUser := models.User{
		Username:   oldUser.Username,
		Password:   oldUser.Password,
		Nickname:   oldUser.Nickname,
		Profilepic: filePath,
	}
	models.UpdateUser(oldUser.Username, &newUser)

	// 返回数据
	returndata := fmt.Sprintf("%+v", struct {
		fileName string
		fileUrl  string
	}{fileName: file.Filename, fileUrl: filePath})

	c.JSON(http.StatusOK, gin.H{
		"data":    returndata,
		"msg":     utils.GetErrMsg(utils.SUCCSE),
		"retcode": utils.SUCCSE,
	})

}
