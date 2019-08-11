package main

import (
	"github.com/labstack/echo"
    "os"
    "io"
    "net/http"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "time"

    // "github.com/json-iterator/go"
)

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
  	// User ID from path `users/:id`
  	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func displayHello(c echo.Context) error {
    id := "Hello, World!"
    return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

// e.POST("/save", save)
// func save(c echo.Context) error {
// 	// Get name and email
// 	name := c.FormValue("name")
// 	email := c.FormValue("email")
// 	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
// }

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

type sunsetSunrise struct {
    sunrise string
    sunset string
    solar_noon string
    day_length string
    civil_twilight_begin string
    civil_twilight_end string
    nautical_twilight_begin string
    nautical_twilight_end string
    astronomical_twilight_begin string
    astronomical_twilight_end string
}

//     "sunrise":"7:27:02 AM",
//     "sunset":"5:05:55 PM",
//     "solar_noon":"12:16:28 PM",
//     "day_length":"9:38:53",
//     "civil_twilight_begin":"6:58:14 AM",
//     "civil_twilight_end":"5:34:43 PM",
//     "nautical_twilight_begin":"6:25:47 AM",
//     "nautical_twilight_end":"6:07:10 PM",
//     "astronomical_twilight_begin":"5:54:14 AM",
//     "astronomical_twilight_end":"6:38:43 PM"

func saveUsers(c echo.Context) error {
    u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func save(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
  	avatar, err := c.FormFile("avatar")
  	if err != nil {
 		return err
 	}

 	// Source
 	src, err := avatar.Open()
 	if err != nil {
 		return err
 	}
 	defer src.Close()

 	// Destination
 	dst, err := os.Create(avatar.Filename)
 	if err != nil {
 		return err
 	}
 	defer dst.Close()

 	// Copy
 	if _, err = io.Copy(dst, src); err != nil {
  		return err
  	}

	return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
}

// https://sunrise-sunset.org/api

func main() {

    url := "https://api.sunrise-sunset.org/json?lat=36.7201600&lng=-4.4203400"

    client := http.Client{
    // If you have a large team in Wakatime, it might take a while to return the data.
    // Set the timeout higher for these requests.
    Timeout: time.Second * 60,
    }

    // Use http.NewRequest when you need to specify attributes of the request. Example: setting custom headers. ??
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.Fatal(err)
    }

    // Execute the GET request.
    res, getErr := client.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
    }

    // Response data is available in response.Body; read until error or EOF.
    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }

    // Create a new sunsetSunrise to store the data.
     sunUpSunDown := sunsetSunrise{}

    // Unmarshal the JSON data to populate the team object.
    // the whole point of this exercise is "Unmarshalling" json data, but requires a module that won't load
    parseErr := json.Unmarshal([]byte(body), &sunUpSunDown)

    if parseErr != nil {
        fmt.Println(parseErr)
        return
    } else {
        fmt.Println("nil", parseErr)
    }

    fmt.Println(sunUpSunDown)

	e := echo.New()
    e.GET("/show", show)
    // e.GET("/", displayHello)
    e.GET("/users/:id", getUser)

    e.POST("/save", save) // doesn't work.
    e.POST("/users", saveUsers) // doesn't work.

    e.Static("/public", "/static")

    // e.GET("/", "static/index.html")
    e.File("/", "index.html")
	e.Logger.Fatal(e.Start(":1323"))

}
