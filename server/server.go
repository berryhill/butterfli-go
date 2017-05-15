
package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/standard"
)


func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Restricted Access
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://butterfli.io"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// ROUTES
	e.GET("/", accessible)
	r.GET("", restricted)
	e.GET("/:username", GetUser)
	e.GET("/users", GetAllUsers)
	e.GET("/:username/accounts", GetAllAccountsByUsername)
	e.GET("/:username/accounts/:accountId/account-creds", GetAccountCreds)
	e.GET("/:username/accounts/:account_id/posts", FindAllAccountPosts)
	e.GET("/:username/accounts/:accountId/search/:socialNetwork/:searchTerm", SearchController)

	e.POST("/:username/accounts/:accountId/search/:socialNetwork/:searchTerm", SearchAndFavorite)


	e.POST("/login", Login)
	e.POST("/user", CreateUser)
	e.POST("/post/edit/:postId/title/:title", EditPost)
	e.POST("/post/approve/:postId", ApprovePost)
	e.POST("/post/disapprove/:postId", DisapprovePost)
	e.POST("/:username/accounts/:account_id/post/delete/:postId", RemovePost)
	e.POST("/:username/accounts/:account_id/post/:postId/upload/twitter/:tweetText", PostTweet)
	e.POST("/:username/accounts/:accountId/twitter/creds", CreateAccountCreds)

	e.POST("/:username/botnet/favorite/:tweetId/accounts/:accountsArray", BotnetFavoriteTweet)
	e.POST("/:username/botnet/follow-account/:accountId/accounts/:accountsArray", BotnetFollowAccountId)
	e.POST("/:username/botnet/follow/:accountName/accounts/:accountsArray", BotnetFollowAccountName)

	e.GET("/:username/accounts/:accountId/search-terms", GetAllSearchTerms)



	// NOT TESTED
	e.POST("/:username/accounts/create/:title", CreateAccount)


	fmt.Println("Server now running on port: 1323")
	e.Run(standard.New(":1323"))
}
