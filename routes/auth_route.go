package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/peewlaom/testgo/services"
)

type Cat struct {
	NameCat string `json:"name"`
	TypeCat string `json:"type"`
}
type JWTClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func AuthRoutes(route *echo.Echo) {
	adminGroup := route.Group("/admin")
	adminGroup.Use(middleware.Logger())
	// adminGroup.Use(middleware.JWT("Ws0844038001"))
	// adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	// subtle.ConstantTimeCompare()
	// 	if username == "peewlaom" && password == "Ws0844038001" {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))
	adminGroup.GET("/main", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is main")
	})
	route.GET("/cats", testParams)
	route.GET("/cats/:data", testReturnType)
	route.POST("/cats", AddCat)
	route.GET("/login", func(c echo.Context) error {
		if c.QueryParam("username") == "peewlaom" && c.QueryParam("password") == "Ws0844038001" {
			cookie := &http.Cookie{}
			cookie.Name = "sesionID"
			cookie.Value = "some_ting"
			cookie.Expires = time.Now().Add(time.Hour * 24)
			token, err := CreateToken()
			if err != nil {
				return c.String(http.StatusBadRequest, "error JWT")
			}
			return c.JSON(http.StatusOK, map[string]string{
				"massage": "token value",
				"token":   token,
			})
		}
		return c.String(http.StatusUnauthorized, "Please check username and password")
	})
}

//middleware///
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "tesgo")
		return next(c)
	}
}
func testParams(c echo.Context) error {
	nameCat := c.QueryParam("name")
	typeCat := c.QueryParam("type")
	return c.String(http.StatusOK, fmt.Sprintf("cat name: %s\nand cat type:%s\n", nameCat, typeCat))
}
func testReturnType(c echo.Context) error {
	nameCat := c.QueryParam("name")
	typeCat := c.QueryParam("type")
	if c.Param("data") == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("cat name: %s\nand cat type:%s\n", nameCat, typeCat))
	} else if c.Param("data") == "json" {
		return c.JSON(http.StatusOK, Cat{nameCat, typeCat})
	} else {
		return c.String(http.StatusBadRequest, "not codition")
	}
}

func AddCat(c echo.Context) error {
	ser := &services.ManagerSql{}
	ser.ConnectSql()
	//c.Request().URL
	cat := &Cat{}
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error_massage": "data empty"})
	}
	err = json.Unmarshal(b, cat)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error_massage": "server can not covert data to json"})
	}
	log.Println("DataCat:", cat)
	return c.String(http.StatusOK, "I got cat")
}

func CreateToken() (string, error) {
	claims := JWTClaims{
		"testgo",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawtoken := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	token, err := rawtoken.SignedString([]byte("mysecret"))
	if err != nil {
		return "", err
	}
	return token, nil

}
