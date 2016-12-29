# Home

API for home management stuff.

## Running Locally

Setup the local PG database:

```sh
make docker
make seed
```

Run with gin auto-reload:

```
make run
```

```
curl localhost:3000/health
```

### Get a PSQL session

`make psql`

### Make sure Heroku is installed
Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ go get -u github.com/heroku/go-getting-started
$ cd $GOPATH/src/github.com/heroku/go-getting-started
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

You should also install [Godep](https://github.com/tools/godep) if you are going to add any dependencies to the sample app.

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku master
$ heroku open
```

or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## Refreshing Heroku DB for development

`make clean-heroku`


## Documentation

For more information about using Go on Heroku, see these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)
