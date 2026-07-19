# trigex.moe

My (Trigex) personal website, that links to my profiles on other websites, lists downloads and links to most of the music I make (and let's you listen on the site itself), has a blog post section with RSS feed support, and displays select GitHub repos and general programming projects I wanna show off.
It is written in Go, and uses Echo for the web framework, templ as the template rendering engine, HTMX for some dynamic page stuff, Tailwind CSS + daisyUI for styling, and SQLite via sqlc for database stuff.
There's also a nice admin panel for managing all the database data through the site.

## Building

I have no idea why you'd want to build the binary serving my personal website, but do the following on any Unix-y system that has Go installed:

``` sh
git clone https://github.com/Trigex/trigex.moe
cd trigex.moe
npm install
export ADMIN_USER=your-user
export ADMIN_PASSWORD=your-password
make build
# All done, should produce this binary in the same folder
./trigexmoe
```
The app will create its SQLite database in `data/trigexmoe.sqlite` on first run.

The admin panel lives at `/admin/` and uses HTTP Basic Auth with those env vars. Blog previews are live and render Markdown through HTMX. Blog images and music cover uploads are stored under `data/uploads/` and served at `/uploads/...`. Existing blog posts can be edited from the admin post list.

To install it to the system, which is only supported on FreeBSD currently (may change if I decide to change my server OS. Looking at you, OpenBSD!), you would do:
```sh
# while still in the trigex.moe/ directory
make install
# if successfully installed, you can enable the service so it's ran with the proper user and logging, and automatically on boot
sysrc trigexmoe_enable="YES"
# start the service
service trigexmoe start
```
I like having it installed in a jail, they make nice application containers. Then you can reverse proxy to it from a web server, I like Caddy personally.