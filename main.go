package main

import ("github.com/AndrewAhr51/JumpCloud/controllers"
"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
