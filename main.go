// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		log.Fatal("$PORT must be set")
// 	}

// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":"+port, nil)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %q", r.URL.Path[1:])
// }

package main
import (
	"net/http"
	"github.com/labstack/echo/v4"
)
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
func main() {
	// サーバー用のインスタンスの取得
	e := echo.New()
	// ユーザー
	u := User{
		Email:    "me@example.com",
		Password: "password",
	}
	// ルーティング設定
	e.GET("/helloworld", helloWorld)
	e.POST("/login", func(c echo.Context) error {
		r := new(User)
		if err := c.Bind(r); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if r.Email != u.Email || r.Password != u.Password {
			return c.String(http.StatusUnauthorized, "login fail")
		}
		// 暫定
		token := "sampletoken"
		return c.String(http.StatusOK, "{\"token\":\""+token+"\"}")
	})
	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}
func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "hello world!!")
}
