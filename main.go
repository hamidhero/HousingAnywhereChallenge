package main

import (
	"housingAnywherChallenge/routers"
)

func main()  {
	r := routers.GetRouter()
	r.Run("0.0.0.0:3030")
}


