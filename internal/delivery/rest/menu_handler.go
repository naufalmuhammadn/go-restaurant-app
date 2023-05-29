package rest

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"fmt"
)

func (h *handler) getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	menuData, err := h.restoUsecase.GetMenuList(menuType)
	if err != nil {
		fmt.Printf("got error %s\n", err.Error())

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}