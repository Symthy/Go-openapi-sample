package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/Symthy/golang-practices/go-openapi-sample/autogen/client"
	"github.com/Symthy/golang-practices/go-openapi-sample/autogen/server"
	"github.com/labstack/echo/v4"
)

const clientBaseURL string = "https://qiita.com/api/v2"

type QiitaProxyHandler struct {
	Lock sync.Mutex
}

func NewQiitaProxyHandler() *QiitaProxyHandler {
	return &QiitaProxyHandler{}
}

// (GET /user/items)
func (h *QiitaProxyHandler) GetUserItems(ctx echo.Context, params server.GetUserItemsParams) error {
	h.Lock.Lock()
	defer h.Lock.Unlock()

	c, err := client.NewClient(clientBaseURL)
	if err != nil {
		return err
	}

	cParams := client.GetUserItemsParams{
		Page:          params.Page,
		PerPage:       params.PerPage,
		Authorization: params.Authorization,
	}

	clientCtx := context.Background()
	res, err := c.GetUserItems(clientCtx, &cParams)
	if err != nil {
		return err
	}
	bytes, err := convertResponseToBytes(res)
	if err != nil {
		return err
	}

	if res.StatusCode == 200 {
		var items server.Items
		json.Unmarshal(bytes, &items)
		return ctx.JSON(res.StatusCode, items)
	} else {
		return parseErrorResponse(ctx, bytes, res.StatusCode)
	}
}

// (GET /authenticated_user)
func (h *QiitaProxyHandler) GetAuthenticatedUser(ctx echo.Context, params server.GetAuthenticatedUserParams) error {
	h.Lock.Lock()
	defer h.Lock.Unlock()

	// http にすると301返ってきて、httpsで再送するが、headerが引き継がれず401
	c, err := client.NewClient(clientBaseURL)
	if err != nil {
		return err
	}

	cParams := &client.GetAuthenticatedUserParams{
		Authorization: params.Authorization,
	}

	clientCtx := context.Background()
	res, err := c.GetAuthenticatedUser(clientCtx, cParams)
	if err != nil {
		return err
	}
	bytes, err := convertResponseToBytes(res)
	if err != nil {
		return err
	}

	if res.StatusCode == 200 {
		var user server.User
		json.Unmarshal(bytes, &user)
		return ctx.JSON(res.StatusCode, user)
	} else {
		return parseErrorResponse(ctx, bytes, res.StatusCode)
	}
}

func convertResponseToBytes(res *http.Response) (bytes []byte, err error) {
	body, err := ioutil.ReadAll(res.Body)
	if err == nil {
		bytes = []byte(body)
	}
	return bytes, err
}

func parseErrorResponse(ctx echo.Context, bytes []byte, statusCode int) (err error) {
	var errRes server.Error
	json.Unmarshal(bytes, &errRes)
	return ctx.JSON(statusCode, errRes)
}
