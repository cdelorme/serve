
# [serve](https://github.com/cdelorme/serve)

A utility that makes testing static web files stupidly simple.

Written in [go](https://golang.org/), it produces a cli that when run from any folder will host that folder as the root path of a webserver.

It defaults to [http://localhost:3000/](http://localhost:3000/), but you can override the port via the `PORT` environment variable or the first parameter supplied to the app.
