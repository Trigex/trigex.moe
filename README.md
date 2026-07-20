# trigex.moe

My (Trigex) personal website, that links to my profiles on other websites, lists downloads and links to most of the music I make (and lets you listen on the site itself), has a blog post section with RSS feed support, and displays select GitHub repos and general programming projects I wanna show off.
It is written in Go, and uses Echo for the web framework, templ as the template rendering engine, HTMX for some dynamic page stuff, Tailwind CSS + daisyUI for styling, and SQLite via sqlc for database stuff.
There's also a nice admin panel for managing all the database entries through the site.

## Building

I have no idea why you'd want to build the binary serving my personal website, but do the following on any Unix-y system that has Go installed:

``` sh
git clone https://github.com/Trigex/trigex.moe
cd trigex.moe
npm install
# edit config.yaml and set server.port, admin.username, and admin.password
make build
# All done, should produce this binary in the same folder
./trigexmoe
```
The admin panel lives at `/admin/` and uses HTTP Basic Auth with credentials from config. The server port is also configured there (`server.port`). Blog previews are live and render Markdown through HTMX. Blog images and music cover uploads are served at `/uploads/...`. Existing blog posts can be edited from the admin post list.

To install it to the system, which is only supported on FreeBSD in the Makefile currently, you would do:
```sh
# while still in the trigex.moe/ directory
make install
# if successfully installed, you can enable the service so it's ran with the proper user and logging, and automatically on boot
sysrc trigexmoe_enable="YES"
# start the service
service trigexmoe start
```
FreeBSD installs use:
- config: `/usr/local/etc/trigexmoe.yaml`
- db: `/var/db/trigexmoe.sqlite`
- data/uploads root: `/usr/local/share/trigexmoe/data`

`make install` creates the data directory and db file with `trigexmoe` ownership, and installs `config.freebsd.yaml` as `/usr/local/etc/trigexmoe.yaml` if no config exists yet.
On FreeBSD, database/data paths are expected to be absolute; if relative paths are configured, the app now falls back to the defaults above.

I like having it installed in a jail, they make nice application containers. Then you can reverse proxy to it from a web server, I like Caddy personally.