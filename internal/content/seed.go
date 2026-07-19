package content

import (
	"context"
	"time"
)

func (s *Store) seed(ctx context.Context) error {
	for _, track := range []CreateMusicTrackParams{
		{ID: "kick-drum-bass", Title: "Kick Drum Bass", Genre: "Rawtempo", FlacUrl: "https://static.termer.net/download/t389z78gnv/Trigex%20-%20Kick%20Drum%20Bass.flac", Mp3Url: "https://static.termer.net/download/8lw5i03jly/Trigex%20-%20Kick%20Drum%20Bass.mp3", YoutubeUrl: "https://www.youtube.com/watch?v=RflfNsFCHU8", SoundcloudUrl: "https://soundcloud.com/trigex/kick-drum-bass", ReleaseDate: time.Date(2026, time.April, 28, 0, 0, 0, 0, time.UTC), CoverImage: "kick-drum-bass.jpg"},
		{ID: "go-down-deh-trigex-dnb-remix", Title: "Spice - Go Down Deh (Trigex DnB Remix)", Genre: "Drum and Bass", FlacUrl: "https://static.termer.net/download/v5beoc38kk/Spice%20-%20Go%20Down%20Deh%20(Trigex%20DnB%20Remix).flac", Mp3Url: "https://static.termer.net/download/oy8efe5xsk/Spice%20-%20Go%20Down%20Deh%20(Trigex%20DnB%20Remix).mp3", YoutubeUrl: "https://www.youtube.com/watch?v=DLysBaQr31o", ReleaseDate: time.Date(2026, time.March, 31, 0, 0, 0, 0, time.UTC), CoverImage: "go-down-deh-trigex-dnb-remix.jpg"},
		{ID: "aubz-sneeze-2026-refix", Title: "Aubz Sneeze (2026 Refix)", Genre: "Uptempo Hardcore", FlacUrl: "https://drive.google.com/file/d/1m7GAEhQxcm_6CnlHsvmVlwMZ_2f1tnKG/view", Mp3Url: "https://static.termer.net/download/4z34o9pl8m/01%20%20-%20Aubz%20Sneeze%20(2026%20Refix).mp3", YoutubeUrl: "https://www.youtube.com/watch?v=4PZ0GETATiQ", SoundcloudUrl: "https://soundcloud.com/trigex/aubz-sneeze-2026-refix", ReleaseDate: time.Date(2026, time.March, 16, 0, 0, 0, 0, time.UTC), CoverImage: "aubz-sneeze-2026-refix.jpg"},
		{ID: "jews-in-a-battle", Title: "Jews In A Battle", Genre: "Makina / Happy Hardcore / Makinatempo", FlacUrl: "https://static.termer.net/download/vey3d4sw0g/Trigex%20-%20Jews%20In%20A%20Battle.flac", Mp3Url: "https://static.termer.net/download/cnu6i8z5ox/Trigex%20-%20Jews%20In%20A%20Battle.mp3", YoutubeUrl: "https://www.youtube.com/watch?v=D-msUMZtaoU", SoundcloudUrl: "https://soundcloud.com/trigex/jews-in-a-battle", ReleaseDate: time.Date(2026, time.March, 11, 0, 0, 0, 0, time.UTC), CoverImage: "jews-in-a-battle.jpg"},
		{ID: "misantrophic-drunken-terror", Title: "Misantrophic Drunken Terror", Genre: "Terror / Hardcore Techno", FlacUrl: "https://static.termer.net/download/9nr7rzpvyg/Misantrophic%20Drunken%20Terror.flac", Mp3Url: "https://static.termer.net/download/qym8owcsce/Misantrophic%20Drunken%20Terror.mp3", YoutubeUrl: "https://www.youtube.com/watch?v=KZY7tkzlsys", SoundcloudUrl: "https://soundcloud.com/trigex/trigex-misantrophic-drunken", ReleaseDate: time.Date(2025, time.April, 18, 0, 0, 0, 0, time.UTC), CoverImage: "misanthropic.jpg"},
		{ID: "alice-in-psytrance-land", Title: "Alice in Psytrance Land", Genre: "Psytrance", FlacUrl: "https://static.termer.net/download/l7d9qfnker/Alice%20in%20Psytrance%20Land.flac", Mp3Url: "https://static.termer.net/download/utp0nbm3sr/Alice%20in%20Psytrance%20Land.mp3", YoutubeUrl: "https://www.youtube.com/watch?v=8oRJMt3x6iw", SoundcloudUrl: "https://soundcloud.com/trigex/alice-in-psytrance-land", ReleaseDate: time.Date(2025, time.April, 8, 0, 0, 0, 0, time.UTC), CoverImage: "alice.jpg"},
		{ID: "guru-guru", Title: "Guru Guru", Genre: "Uptempo Hardcore / Lolicore", FlacUrl: "https://static.termer.net/download/of1cgo07e6/Trigex%20-%20Guru%20Guru.flac", Mp3Url: "https://static.termer.net/download/gnacpph7b1/Trigex%20-%20Guru%20Guru.mp3", YoutubeUrl: "https://www.youtube.com/watch?v=YRtf9nKzbEw", SoundcloudUrl: "https://soundcloud.com/trigex/guru-guru", ReleaseDate: time.Date(2025, time.February, 22, 0, 0, 0, 0, time.UTC), CoverImage: "guru.jpg"},
		{ID: "i-can-say-whatever-i-want", Title: "I Can Say Whatever I Want", Genre: "Hardtek / Hardcore Techno", FlacUrl: "", Mp3Url: "https://static.termer.net/download/ql1swhlino/whatever.mp3", YoutubeUrl: "https://www.youtube.com/watch?v=TLtzR51fdpk", SoundcloudUrl: "https://soundcloud.com/trigex/i-can-say-whatever-i-want", ReleaseDate: time.Date(2025, time.January, 21, 0, 0, 0, 0, time.UTC), CoverImage: "whatever.jpg"},
		{ID: "fan-service-kick-edit", Title: "S3RL - Fan Service (Trigex Kick Edit)", Genre: "Hardcore Techno / Happy Hardcore", FlacUrl: "https://static.termer.net/download/6qlpxygqjs/S3RL%20-%20Fan%20Service%20(Trigex%20Kick%20Edit).flac", Mp3Url: "https://static.termer.net/download/m1mfeg7q5e/S3RL%20-%20Fan%20Service%20(Trigex%20Kick%20Edit).mp3", YoutubeUrl: "https://youtu.be/g2J7a9OynnA", SoundcloudUrl: "https://soundcloud.com/trigex/s3rl-fan-service-trigex-kick-edit", ReleaseDate: time.Date(2024, time.June, 9, 0, 0, 0, 0, time.UTC), CoverImage: "fan.jpg"},
		{ID: "worlds-smallest-violin-bootleg", Title: "World's Smallest Violin (Trigex Happy Hardcore Bootleg)", Genre: "Makina / Happy Hardcore", FlacUrl: "https://drive.google.com/file/d/1VH0CVeGEhf23RIW5qBRarELAjG7plH74/view?usp=sharing", Mp3Url: "https://static.termer.net/download/s6br7edtwi/AJR%20-%20World's%20Smallest%20Violin%20(Final%20Probably).mp3", YoutubeUrl: "https://www.youtube.com/watch?v=5NXa6Egecug", SoundcloudUrl: "https://soundcloud.com/trigex/worlds-smallest-violin-trigex-happy-hardcore-bootleg", ReleaseDate: time.Date(2024, time.May, 31, 0, 0, 0, 0, time.UTC), CoverImage: "violin.jpg"},
		{ID: "aubz-sneeze", Title: "Aubz Sneeze", Genre: "Uptempo Hardcore", FlacUrl: "https://static.termer.net/download/ktxke53zfx/Trigex%20-%20Aubz%20Sneeze.flac", Mp3Url: "https://static.termer.net/download/y0ka6o8ygh/Trigex%20-%20Aubz%20Sneeze.mp3", YoutubeUrl: "https://www.youtube.com/watch?v=EOQjTTWu4qc", SoundcloudUrl: "https://soundcloud.com/trigex/aubz-sneeze", ReleaseDate: time.Date(2024, time.May, 23, 0, 0, 0, 0, time.UTC), CoverImage: "aubz.jpg"},
		{ID: "pedro-uptempo-remix", Title: "Pedro (Trigex Uptempo Remix)", Genre: "Uptempo Hardcore", FlacUrl: "https://static.termer.net/download/eyt6p0ne7v/Raffaella%20Carra%CC%80%20-%20Pedro%20(Trigex%20Uptempo%20Remix).flac", Mp3Url: "https://static.termer.net/download/l04jk9fdll/Raffaella%20Carra%CC%80%20-%20Pedro%20(Trigex%20Uptempo%20Remix).mp3", YoutubeUrl: "https://www.youtube.com/watch?v=ZeeZyzhe-Xk", SoundcloudUrl: "", ReleaseDate: time.Date(2024, time.May, 23, 0, 0, 0, 0, time.UTC), CoverImage: "pedro.jpg"},
		{ID: "smiling-friends-makina-mix", Title: "Smiling Friends (Makina Mix)", Genre: "Makina", FlacUrl: "https://static.termer.net/download/qmw1lqiz1o/Trigex%20-%20Smiling%20Friends%20(Makina%20Mix).flac", Mp3Url: "https://static.termer.net/download/z1expkcm8e/Trigex%20-%20Smiling%20Friends%20(Makina%20Mix).mp3", YoutubeUrl: "https://www.youtube.com/watch?v=JyhCesdnIS0", SoundcloudUrl: "https://soundcloud.com/trigex/smiling-friends-makina-mix", ReleaseDate: time.Date(2024, time.April, 9, 0, 0, 0, 0, time.UTC), CoverImage: "smiling.jpg"},
		{ID: "super-idol-makinatempo-remix", Title: "Super Idol 的笑容 (Trigex Makinatempo Remix)", Genre: "Makina / Makinatempo / Happy Hardcore", FlacUrl: "https://static.termer.net/download/csn3vodpep/Super%20Idol%20(Trigex%20Makinatempo%20Remix).flac", Mp3Url: "https://static.termer.net/download/j7ycis9yd5/Super%20Idol%20(Trigex%20Makinatempo%20Remix).mp3", YoutubeUrl: "https://www.youtube.com/watch?v=TSM6cz_yOpo", SoundcloudUrl: "https://soundcloud.com/trigex/super-idol-trigex-makinatempo-remix", ReleaseDate: time.Date(2024, time.January, 16, 0, 0, 0, 0, time.UTC), CoverImage: "idol.jpg"},
		{ID: "thats-so-gay", Title: "That's So Gay", Genre: "Uptempo Hardcore", FlacUrl: "https://static.termer.net/download/iir6r19aop/thatssogay.flac", Mp3Url: "https://static.termer.net/download/ah1zohpxex/thatssogay.mp3", YoutubeUrl: "https://www.youtube.com/watch?v=ibLjholRJgQ", SoundcloudUrl: "https://soundcloud.com/trigex/thats-so-gay", ReleaseDate: time.Date(2023, time.September, 26, 0, 0, 0, 0, time.UTC), CoverImage: "gay.jpg"},
		{ID: "kill-me-baby", Title: "Kill Me, Baby", Genre: "Happy Hardcore / Reverse Bass", FlacUrl: "https://static.termer.net/download/or5unjghec/Trigex%20-%20Kill%20Me%20%20Baby.flac", Mp3Url: "https://static.termer.net/download/x77kg5ew64/Trigex%20-%20Kill%20Me%20%20Baby.mp3", YoutubeUrl: "https://www.youtube.com/watch?v=mT2og6R1XtI", SoundcloudUrl: "https://soundcloud.com/trigex/kill-me-baby", ReleaseDate: time.Date(2023, time.August, 28, 0, 0, 0, 0, time.UTC), CoverImage: "baby.jpg"},
		{ID: "push-up-bootleg", Title: "Creeds - Push Up (Trigex Uptempo Bootleg)", Genre: "Uptempo Hardcore", FlacUrl: "https://static.termer.net/download/ednyc7j4pe/Creeds%20-%20Push%20Up%20(Trigex%20Uptempo%20Bootleg).flac", Mp3Url: "https://static.termer.net/download/u91dsw7hhm/Creeds%20-%20Push%20Up%20(Trigex%20Uptempo%20Bootleg).mp3", YoutubeUrl: "https://www.youtube.com/watch?v=o5JnASJYGNI", SoundcloudUrl: "https://soundcloud.com/trigex/push-up-trigex-uptempo-bootleg", ReleaseDate: time.Date(2023, time.April, 6, 0, 0, 0, 0, time.UTC), CoverImage: "pushup.jpg"},
		{ID: "pill-provider", Title: "Pill Provider", Genre: "Uptempo Hardcore / Frenchcore", FlacUrl: "https://static.termer.net/download/i1n1g6fpug/Pill%20Provider.flac", Mp3Url: "https://static.termer.net/download/w9i6gtuu7y/Pill%20Provider.mp3", YoutubeUrl: "https://www.youtube.com/watch?v=UjYQ-CN6SNU", SoundcloudUrl: "https://soundcloud.com/trigex/pill-provider", ReleaseDate: time.Date(2023, time.January, 4, 0, 0, 0, 0, time.UTC), CoverImage: "pill.jpg"},
	} {
		if err := s.CreateMusicTrack(ctx, track); err != nil {
			return err
		}
		if _, err := s.db.ExecContext(ctx, `
			UPDATE music_tracks
			SET genre = ?
			WHERE id = ? AND COALESCE(genre, '') = ''
		`, track.Genre, track.ID); err != nil {
			return err
		}
	}

	for _, project := range []CreateProjectParams{
		{ID: "netchat", Name: "Netchat", Description: "A pretty unremarkable chat client & server written in C++, as a programming exercise", RepoUrl: "https://github.com/Trigex/Netchat", TechStack: "C++"},
		{ID: "trigex.moe", Name: "trigex.moe", Description: "My personal website you're currently on! Features a blog system (with RSS feed generation), my music releases, my projects, and my social links, all editable through the admin panel. Basically just a personal ghetto CMS for me.", RepoUrl: "https://github.com/Trigex/trigex.moe", TechStack: "Go, Echo, templ, HTMX, sqlc, TailwindCSS+daisyUI"},
		{ID: "convert-muh-music", Name: "convert-muh-music", Description: "A bulk audio library transcoder with sane defaults", RepoUrl: "https://github.com/Trigex/convert-muh-music", TechStack: "Go, Python"},
		{ID: "alphanet", Name: "AlphaNET", Description: "AlphaNET was going to be a hacking & cracking style MMO with a whole virtual operating system and scripting, but never ended up finished...", RepoUrl: "https://github.com/Trigex/AlphaNET", TechStack: "C#, .NET Standard"},
		{ID: "textchblazor", Name: "TextchBlazor", Description: "TextchBlazor was a 2channel-style front-end for my friend's Textbin project, but neither are active anymore", RepoUrl: "https://github.com/Trigex/TextchBlazor", TechStack: "C#, Blazor"},
	} {
		if err := s.CreateProject(ctx, project); err != nil {
			return err
		}
	}

	for _, post := range []CreateBlogPostParams{
		{
			ID:          "welcome-to-the-blog",
			Slug:        "welcome-to-the-blog",
			Title:       "Welcome to the Blog",
			Excerpt:     "A starter post for the new blog section.",
			Body:        "This is the first post in the new blog section.\n\nFrom here I can keep site updates, notes, and longer thoughts in the database instead of hardcoding them into templates.",
			PublishedAt: time.Date(2026, time.July, 18, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:          "building-the-site",
			Slug:        "building-the-site",
			Title:       "Building the Site",
			Excerpt:     "How the site is wired together with Go, Echo, templ, DaisyUI, and SQLite.",
			Body:        "This post is a placeholder for the blog system.\n\nIt gives the blog page something real to render while the content model is still being fleshed out.",
			PublishedAt: time.Date(2026, time.July, 18, 0, 0, 0, 0, time.UTC),
		},
	} {
		if err := s.CreateBlogPost(ctx, post); err != nil {
			return err
		}
	}

	for _, link := range []CreateSocialLinkParams{
		{ID: "email", Name: "Email", Url: "mailto:trigex@trigex.moe"},
		{ID: "github", Name: "GitHub", Url: "https://github.com/Trigex"},
		{ID: "instagram", Name: "Instagram", Url: "https://www.instagram.com/seth.stokley/"},
		{ID: "soundcloud", Name: "Soundcloud", Url: "https://soundcloud.com/trigex"},
		{ID: "youtube", Name: "Youtube", Url: "https://www.youtube.com/Trigex"},
		{ID: "spotify", Name: "Spotify", Url: "https://open.spotify.com/artist/5bROX0acYXYaCU8z5Rns6z?si=IZbjaWXhTReYSiklc2gQMg"},
	} {
		if err := s.CreateSocialLink(ctx, link); err != nil {
			return err
		}
	}

	return nil
}
