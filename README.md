<div align="center">
  <h1>Go Weather App CLI</h1>
<p>This is a simple Command Line Interface (CLI) real-time weather application in Go that fetches and displays current weather information using the WeatherAPI.</p>
</div><br>

## Table of Contents:
- [Features](#features)
- [Installation](#installation)

## Features:
- Fetches current weather information for any location.
- Displays temperature, humidity, wind speed, and weather conditions.
- Supports multiple units (Celsius and Fahrenheit).
- Easy-to-use command-line interface.

## Installation:
1. **Clone the repository:**
   ```sh
   git clone https://github.com/Gepzuu/go-weather-app.git
   cd go-weather-app
  
## Install dependencies:
Make sure you have Go installed. 

**Then run:**
```sh 
  go mod tidy
```

Usage
To get the current weather information for a specific location, run:

```sh 
go run main.go -location="YourCity"
```

Replace "YourCity" with the name of the city you want to get the weather for.

## Using Docker:
Alternatively, you can use the provided Docker image since this project is already dockerized. Here are the instructions to build and run the Docker container:

Ensure you have Docker installed on your system.
Open a terminal and navigate to the project directory.
Build the Docker image using the following command:
```
docker build -t go-weather-app .
```
Run the Docker container using the following command:
```
docker run go-weather-app
```

This way, you can run the application without needing to install GoLang directly on your system.




## Configuration:
To use this application, you need to get an API key from WeatherAPI. Once you have the key, you can set it as an environment variable:
```sh
 export WEATHER_API_KEY="your_api_key_here"
```

Dependencies:
Go Modules: Dependency management.
WeatherAPI: Used for fetching weather data.
Additional dependencies listed in go.mod.



# Credits:
WeatherAPI: Used for fetching weather data.<br>
Go Modules: For dependency management.
