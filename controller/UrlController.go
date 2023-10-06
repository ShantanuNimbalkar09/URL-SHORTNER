package controller

import (
	"app/constant"
	"app/database"
	"app/helper"
	"app/types"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ShortTheUrl(c *gin.Context) {
	var shortUrlBody types.ShortUrlBody
	err := c.BindJSON(&shortUrlBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": constant.BindError})
		return
	}
	code := helper.GetRnadomString(6)

	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

	if record.UrlCode != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "this code is already in use"})
		return
	}

	var url types.UrlDb

	url.CreateAt = time.Now().Unix()
	url.ExpireAt = time.Now().Unix()
	url.UrlCode = code
	url.LongUrl = shortUrlBody.LongUrl
	url.ShortUrl = constant.BaseUrl + code

	resp, err := database.Mgr.Insert(url, constant.UrlCollection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": false, "data": resp, "short_url": url.ShortUrl})

}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")

	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

	if record.UrlCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "there is no url found"})
		return
	}
	fmt.Println(record.LongUrl)

	c.Redirect(http.StatusPermanentRedirect, record.LongUrl)
}
