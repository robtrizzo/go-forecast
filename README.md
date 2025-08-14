# go-forecast

This is a simple api to get a weather forecast summary based on coordinate input.

## Development

To run the application in development mode, run `make run/api`. Optional flags include:

 - `port`: Port the api will run on. Defaults to `4000`
 - `weather-url`: URL of the weather API this service fetches data from. Defaults to `https://api.weather.gov`
 - `cors-trusted-origins`: Space-separated whitelist of origins permitted to make requests to this API. Defaults to an empty list, which permits all origins.

## Build

To build the application, run `make build/api`. This will create an executable binary in the `bin/` directory. You can run it with `bin/api`.