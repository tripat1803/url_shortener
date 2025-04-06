package controllers

import (
	"net/http"
	"tripat3k2/url_shortner/config"
	"tripat3k2/url_shortner/models"
	"tripat3k2/url_shortner/utils"

	"github.com/gin-gonic/gin"
)

type CreateShortUrlRequest struct {
	PlainText uint64 `json:"plainText"`
	LongUrl   string `json:"longUrl"`
}

type RedirectUrlRequest struct {
	Cipher string `uri:"cipher" binding:"required"`
}

type UrlDTO struct {
	LongUrl string
}

func CreateShortUrl(ctx *gin.Context) {
	req := &CreateShortUrlRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	value, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	userId, typeOk := value.(uint)
	if !typeOk {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	cipher := utils.Encode(req.PlainText)
	newEntry := &models.Url{
		UserID:  uint(userId),
		ID:      req.PlainText,
		Cipher:  cipher,
		LongUrl: req.LongUrl,
	}

	result := config.DB.Create(newEntry)
	if result.Error != nil {
		result.Rollback()
		ctx.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	result.Commit()
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Added long url"})
}

func RedirectUrl(ctx *gin.Context) {
	req := &RedirectUrlRequest{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	plainText := utils.Decode(req.Cipher)
	url := &UrlDTO{}
	result := config.DB.Where("id = ?", plainText).Limit(1).Model(&models.Url{}).Find(url)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, result.Error.Error())
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, url.LongUrl)
}
