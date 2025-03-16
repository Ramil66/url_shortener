package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	urlshortener "github.com/ramil66/url-shortener"
)

func (h *Handler) ShorteningUrl(c *gin.Context) {
	var input urlshortener.Url
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Printf("Ошибка BindJSON: %v\n", err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Url == "" {
		newErrorResponse(c, http.StatusBadRequest, "Поле 'url' не может быть пустым")
		return
	}
	fmt.Println(input.Url)

	alias, err := h.services.Url.Shortening(input.Url)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"url": alias,
	})
}

type getAllUrlResponse struct {
	Data []urlshortener.Url `json:"data"`
}

func (h *Handler) GetAllUrls(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	urls, err := h.services.Url.GetAllUrls(userId)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, urls)
}

func (h *Handler) RedirectUrl(c *gin.Context) {
	alias, err := getAlias(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	url, err := h.services.Url.GetUrl(alias)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Url.IncrementCounter(alias); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "failed to increment counter")
		return
	}
	if h.services.Url.CheckLinkUrlUser(alias) {
		c.Redirect(http.StatusFound, url)
		return
	}
	userAgent := c.GetHeader("User-Agent")
	ua := user_agent.New(userAgent)
	urlId, err := h.services.Url.GetIdUrl(alias)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(time.Now().Format("02.01.2006"))
	err = h.services.Statistic.SaveStatistic(urlshortener.Statistic{
		UrlId:    urlId,
		Ip:       c.ClientIP(),
		Device:   ua.OSInfo().Name,
		LastDate: time.Now().Format("02.01.2006"),
	})

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Redirect(http.StatusFound, url)
}

func (h *Handler) DeleteUrl(c *gin.Context) {
	alias, err := getAlias(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Url.DeleteUrl(alias)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

type customInput struct {
	Url   string `json:"url" binding:"required"`
	Alias string `json:"alias" binding:"required"`
}

func (h *Handler) CustomUrl(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input customInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("handler: " + input.Alias)
	fmt.Println("handler: " + input.Url)
	if !h.services.Url.CheckAlias(input.Alias) {
		newErrorResponse(c, http.StatusBadRequest, "dieser Link existiert bereits")
		return
	}
	url, err := h.services.Url.CustomUrl(userId, input.Url, input.Alias)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"url": url,
	})
}

func (h *Handler) ShorteningUrlUsers(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input urlshortener.Url
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Printf("Ошибка BindJSON: %v\n", err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Url == "" {
		newErrorResponse(c, http.StatusBadRequest, "Поле 'url' не может быть пустым")
		return
	}
	fmt.Println(input.Url)

	alias, err := h.services.Url.ShorteningUrlUsers(userId, input.Url)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"url": alias,
	})

}
