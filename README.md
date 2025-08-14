# go-forecast

This is a simple api to get a weather forecast summary based on coordinate input.

## Development

To run the application in development mode, run `make run/api`. Optional flags include:

 - `port`: Port the api will run on. Defaults to `4000`
 - `weather-url`: URL of the weather API this service fetches data from. Defaults to `https://api.weather.gov`
 - `cors-trusted-origins`: Space-separated whitelist of origins permitted to make requests to this API. Defaults to an empty list, which permits all origins.

## Build

To build the application, run `make build/api`. This will create an executable binary in the `bin/` directory. You can run it with `./bin/api`.

## Example Usage

There are two endpoints in this application

`/v1/healthcheck` reports application status.

Example:

```shell
$ curl "http://localhost:4000/v1/healthcheck"
{
        "status": "available"
}
```

`/v1/forecast` returns a summary of today's forecast

```shell
$ curl "http://localhost:4000/v1/forecast"
Today's forecast is Mostly Clear and moderate.
```

This endpoint defaults to Detroit, but you may provide coordinates via query parameters. For example (for Anchorage, Alaska):

```shell
$ curl "http://localhost:4000/v1/forecast?lat=61.218&lon=-149.899"
Today's forecast is Scattered Rain Showers and moderate.
```