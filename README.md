# Simple Rest API

I build rest api with golang based on concept clean architecture from Unclebob (https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

This app require mysql for running this api.

First copy and change env file for setting app config

```shell
cp .env.example .env
```

Next, run testing code in root program for make sure this program running properly:

```shell
go test ./... -v
```

And then build and run `Dockerfile` this app:

```shell
docker build --tag go-clean-api:0.0.1 .
docker run -d --rm -p 5000:5000 --name golang-api go-clean-api:0.0.1
```

If you want check swagger documentation this api, you can open it via browser (`http://localhost:5000/swagger/index.html`).


