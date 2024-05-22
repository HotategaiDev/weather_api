# JHA Golang Coding Exercise

This project is a simple HTTP server written in Go that utilizes the Open Weather API to provide weather information based on latitude and longitude coordinates.

## Features

- Exposes an endpoint that takes latitude and longitude coordinates as input.
- Retrieves weather condition and temperature information from the Open Weather API.
- Returns the current weather condition (e.g., rain, snow) and temperature category (hot, cold, moderate) for the specified location.


## Configuration

Before running the application, make sure to configure your OpenWeather API key in the `config.dev.yaml` file. You can obtain your API key from the [Open Weather website](https://openweathermap.org/api). Once obtained, replace the `YOUR_API_KEY_HERE` placeholder in the `config.dev.yaml` file with your actual API key.

Example `config.dev.yaml`:

```yaml
...
openWeather:
    apiKey: YOUR_API_KEY_HERE
...
```

## Installation

1. Clone this repository to your local machine:

```bash
git clone https://github.com/urimeba/codeChallenge-jha.git
```

2. Navigate to the project directory:

```bash
cd codeChallenge-jha
```

3. Build the project:

```bash
go build
```

4. Run the executable:

```bash
./codeChallenge-jha
```


## Usage

Make a GET request to the following endpoint with latitude and longitude coordinates:

```
GET /?lat={latitude}&long={longitude}
```

Replace `{latitude}` and `{longitude}` with the actual latitude and longitude coordinates.

Example:

```
GET /?lat=37.7749&long=-122.4194
```

You can also use curl:

```bash
curl "http://localhost:port/?lat=37.7749&long=-122.4194"
```

If the port wasn't changed, and the service is running on localhost, you may use:

```bash
curl "http://localhost:8075/?lat=37.7749&long=-122.4194"
```
