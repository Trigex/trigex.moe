package main

import (
	"embed"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/trigex/trigex.moe/views"
	"io/fs"
	"net/http"
	"time"
)

const (
	SiteName = "trigex.moe"
	Port     = 8080
)

//go:embed assets
var embeddedAssets embed.FS

func main() {
	fmt.Println("Starting server...")
	// Create echo instance
	e := echo.New()

	// Add middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Create sub-filesystem to serve assets with correct pathing
	staticFS, err := fs.Sub(embeddedAssets, "assets")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.HTTPErrorHandler = customErrorHandler

	// -- Routes
	// Serve assets at /static
	e.StaticFS("/static", staticFS)

	// Pages
	e.GET("/", serveHomePage)
	e.GET("/music", serveMusicPage)
	e.GET("/projects", serveProjectsPage)

	portString := fmt.Sprintf(":%d", Port)
	e.Logger.Info("Started server on http://localhost:" + portString)
	e.Logger.Info("Access static assets at http://localhost:" + portString + "/static/")
	if err := e.Start(fmt.Sprintf(":%d", Port)); err != nil {
		e.Logger.Fatal(err)
	}
}

func customErrorHandler(err error, c echo.Context) {
	// Cast the error to an echo.HTTPError to get the status code
	code := http.StatusInternalServerError
	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code
	}

	// If the error is a 404 Not Found, render our custom page
	if code == http.StatusNotFound {
		// It's important to set the status code on the response
		c.Response().WriteHeader(code)
		// Render the 404 page inside our main layout
		page := views.Layout("trigex.moe | Page Not Found", views.NotFoundPage())
		err := page.Render(c.Request().Context(), c.Response().Writer)
		if err != nil {
			c.Logger().Error(err)
		}
		return
	}

	// For all other errors, fall back to Echo's default handler
	c.Echo().DefaultHTTPErrorHandler(err, c)
}

func serveHomePage(c echo.Context) error {
	data := views.PageData{
		Title: SiteName,
		Name:  SiteName,
		Bio:   "Hi I'm Trigex! Welcome to my corner of the internet! I'm a software developer, music producer, DJ, and sysadmin. (BSDs preferred ;)). I unfortunately reside in California, but I hope that'll change one day. Always looking for work!",
		Links: []views.Link{
			{Name: "GitHub", URL: "https://github.com/Trigex"},
			{Name: "Soundcloud", URL: "https://soundcloud.com/trigex"},
			{Name: "Instagram", URL: "https://www.instagram.com/seth.stokley/"},
			{Name: "Youtube", URL: "https://www.youtube.com/Trigex"},
			{Name: "Email", URL: "mailto:trigex@trigex.moe"},
		},
	}

	// Create instance of component page and call it's Render method
	// Writes the HTML directly to the response
	return views.Layout("trigex.moe | Home", views.HomePage(data)).Render(c.Request().Context(), c.Response().Writer)
}

func serveMusicPage(c echo.Context) error {
	data := []views.Track{
		{Title: "Misantrophic Drunken Terror", FlacURL: "https://static.termer.net/download/9nr7rzpvyg/Misantrophic%20Drunken%20Terror.flac", Mp3URL: "https://static.termer.net/download/qym8owcsce/Misantrophic%20Drunken%20Terror.mp3", YoutubeURL: "https://www.youtube.com/watch?v=KZY7tkzlsys", SoundcloudURL: "https://soundcloud.com/trigex/trigex-misantrophic-drunken", ReleaseDate: time.Date(2025, time.April, 18, 0, 0, 0, 0, time.UTC), CoverImage: "misanthropic.png"},
		{Title: "Alice in Psytrance Land", FlacURL: "https://static.termer.net/download/l7d9qfnker/Alice%20in%20Psytrance%20Land.flac", Mp3URL: "https://static.termer.net/download/utp0nbm3sr/Alice%20in%20Psytrance%20Land.mp3", YoutubeURL: "https://www.youtube.com/watch?v=8oRJMt3x6iw", SoundcloudURL: "https://soundcloud.com/trigex/alice-in-psytrance-land", ReleaseDate: time.Date(2025, time.April, 8, 0, 0, 0, 0, time.UTC), CoverImage: "alice.png"},
		{Title: "Guru Guru", FlacURL: "https://static.termer.net/download/of1cgo07e6/Trigex%20-%20Guru%20Guru.flac", Mp3URL: "https://static.termer.net/download/gnacpph7b1/Trigex%20-%20Guru%20Guru.mp3", YoutubeURL: "https://www.youtube.com/watch?v=YRtf9nKzbEw", SoundcloudURL: "https://soundcloud.com/trigex/guru-guru", ReleaseDate: time.Date(2025, time.February, 22, 0, 0, 0, 0, time.UTC), CoverImage: "guru.png"},
		{Title: "I Can Say Whatever I Want", Mp3URL: "https://static.termer.net/download/ql1swhlino/whatever.mp3", YoutubeURL: "https://www.youtube.com/watch?v=TLtzR51fdpk", SoundcloudURL: "https://soundcloud.com/trigex/i-can-say-whatever-i-want", ReleaseDate: time.Date(2025, time.January, 21, 0, 0, 0, 0, time.UTC), CoverImage: "whatever.png"},
		{Title: "S3RL - Fan Service (Trigex Kick Edit)", Mp3URL: "https://static.termer.net/download/m1mfeg7q5e/S3RL%20-%20Fan%20Service%20(Trigex%20Kick%20Edit).mp3", FlacURL: "https://static.termer.net/download/6qlpxygqjs/S3RL%20-%20Fan%20Service%20(Trigex%20Kick%20Edit).flac", YoutubeURL: "https://youtu.be/g2J7a9OynnA", SoundcloudURL: "https://soundcloud.com/trigex/s3rl-fan-service-trigex-kick-edit", ReleaseDate: time.Date(2024, time.June, 9, 0, 0, 0, 0, time.UTC), CoverImage: "fan.png"},
		{Title: "World's Smallest Violin (Trigex Happy Hardcore Bootleg)", FlacURL: "https://drive.google.com/file/d/1VH0CVeGEhf23RIW5qBRarELAjG7plH74/view?usp=sharing", Mp3URL: "https://static.termer.net/download/s6br7edtwi/AJR%20-%20World's%20Smallest%20Violin%20(Final%20Probably).mp3", YoutubeURL: "https://www.youtube.com/watch?v=5NXa6Egecug", SoundcloudURL: "https://soundcloud.com/trigex/worlds-smallest-violin-trigex-happy-hardcore-bootleg", ReleaseDate: time.Date(2024, time.May, 31, 0, 0, 0, 0, time.UTC), CoverImage: "violin.png"},
		{Title: "Aubz Sneeze", FlacURL: "https://static.termer.net/download/ktxke53zfx/Trigex%20-%20Aubz%20Sneeze.flac", Mp3URL: "https://static.termer.net/download/y0ka6o8ygh/Trigex%20-%20Aubz%20Sneeze.mp3", YoutubeURL: "https://www.youtube.com/watch?v=EOQjTTWu4qc", SoundcloudURL: "https://soundcloud.com/trigex/aubz-sneeze", ReleaseDate: time.Date(2024, time.May, 23, 0, 0, 0, 0, time.UTC), CoverImage: "aubz.png"},
		{Title: "Pedro (Trigex Uptempo Remix)", FlacURL: "https://static.termer.net/download/eyt6p0ne7v/Raffaella%20Carra%CC%80%20-%20Pedro%20(Trigex%20Uptempo%20Remix).flac", Mp3URL: "https://static.termer.net/download/l04jk9fdll/Raffaella%20Carra%CC%80%20-%20Pedro%20(Trigex%20Uptempo%20Remix).mp3", YoutubeURL: "https://www.youtube.com/watch?v=ZeeZyzhe-Xk", ReleaseDate: time.Date(2024, time.May, 23, 0, 0, 0, 0, time.UTC), CoverImage: "pedro.png"},
		{Title: "Smiling Friends (Makina Mix)", FlacURL: "https://static.termer.net/download/qmw1lqiz1o/Trigex%20-%20Smiling%20Friends%20(Makina%20Mix).flac", Mp3URL: "https://static.termer.net/download/z1expkcm8e/Trigex%20-%20Smiling%20Friends%20(Makina%20Mix).mp3", YoutubeURL: "https://www.youtube.com/watch?v=JyhCesdnIS0", SoundcloudURL: "https://soundcloud.com/trigex/smiling-friends-makina-mix", ReleaseDate: time.Date(2024, time.April, 9, 0, 0, 0, 0, time.UTC), CoverImage: "smiling.png"},
		{Title: "Super Idol 的笑容 (Trigex Makinatempo Remix)", FlacURL: "https://static.termer.net/download/csn3vodpep/Super%20Idol%20(Trigex%20Makinatempo%20Remix).flac", Mp3URL: "https://static.termer.net/download/j7ycis9yd5/Super%20Idol%20(Trigex%20Makinatempo%20Remix).mp3", YoutubeURL: "https://www.youtube.com/watch?v=TSM6cz_yOpo", SoundcloudURL: "https://soundcloud.com/trigex/super-idol-trigex-makinatempo-remix", ReleaseDate: time.Date(2024, time.January, 16, 0, 0, 0, 0, time.UTC), CoverImage: "idol.png"},
		{Title: "That's So Gay", FlacURL: "https://static.termer.net/download/iir6r19aop/thatssogay.flac", Mp3URL: "https://static.termer.net/download/ah1zohpxex/thatssogay.mp3", YoutubeURL: "https://www.youtube.com/watch?v=ibLjholRJgQ", SoundcloudURL: "https://soundcloud.com/trigex/thats-so-gay", ReleaseDate: time.Date(2023, time.September, 26, 0, 0, 0, 0, time.UTC), CoverImage: "gay.png"},
		{Title: "Kill Me, Baby", FlacURL: "https://static.termer.net/download/or5unjghec/Trigex%20-%20Kill%20Me%20%20Baby.flac", Mp3URL: "https://static.termer.net/download/x77kg5ew64/Trigex%20-%20Kill%20Me%20%20Baby.mp3", YoutubeURL: "https://www.youtube.com/watch?v=mT2og6R1XtI", SoundcloudURL: "https://soundcloud.com/trigex/kill-me-baby", ReleaseDate: time.Date(2023, time.August, 28, 0, 0, 0, 0, time.UTC), CoverImage: "baby.png"},
		{Title: "Creeds - Push Up (Trigex Uptempo Bootleg)", FlacURL: "https://static.termer.net/download/ednyc7j4pe/Creeds%20-%20Push%20Up%20(Trigex%20Uptempo%20Bootleg).flac", Mp3URL: "https://static.termer.net/download/u91dsw7hhm/Creeds%20-%20Push%20Up%20(Trigex%20Uptempo%20Bootleg).mp3", SoundcloudURL: "https://soundcloud.com/trigex/push-up-trigex-uptempo-bootleg", YoutubeURL: "https://www.youtube.com/watch?v=o5JnASJYGNI", ReleaseDate: time.Date(2023, time.April, 6, 0, 0, 0, 0, time.UTC), CoverImage: "pushup.png"},
		{Title: "Pill Provider", FlacURL: "https://static.termer.net/download/i1n1g6fpug/Pill%20Provider.flac", Mp3URL: "https://static.termer.net/download/w9i6gtuu7y/Pill%20Provider.mp3", YoutubeURL: "https://www.youtube.com/watch?v=UjYQ-CN6SNU", SoundcloudURL: "https://soundcloud.com/trigex/pill-provider", ReleaseDate: time.Date(2023, time.January, 4, 0, 0, 0, 0, time.UTC), CoverImage: "pill.png"},
	}

	return views.Layout("trigex.moe | Music", views.MusicPage(data)).Render(c.Request().Context(), c.Response().Writer)
}

func serveProjectsPage(c echo.Context) error {
	data := []views.Project{
		{Name: "convert-muh-music", Description: "A bulk audio library transcoder with sane defaults", RepoURL: "https://github.com/Trigex/convert-muh-music", TechStack: "Go, Python (For the working script, the Go version is a wip)"},
		{Name: "AlphaNET", Description: "AlphaNET was going to be a hacking & cracking style MMO with a whole virtual operating system and scripting, but never ended up finished...", RepoURL: "https://github.com/Trigex/AlphaNET", TechStack: "C#, .NET Standard"},
		{Name: "TextchBlazor", Description: "TextchBlazor was a 2channel-style front-end for my friend's Textbin project, but neither are active anymore", RepoURL: "https://github.com/Trigex/TextchBlazor", TechStack: "C#, Blazor"},
	}
	return views.Layout("trigex.moe | Projects", views.ProjectsPage(data)).Render(c.Request().Context(), c.Response().Writer)
}
