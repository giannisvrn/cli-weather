# Weather CLI Tool

**Weather CLI Tool** is a simple command-line application written in Go that fetches and displays the current weather conditions for a specified city. By default, it retrieves the weather for **Athens** if no city is provided.

## Features

- Fetches current weather conditions for any city.
- Uses **Athens** as the default city if no city is specified.
- Lightweight and easy-to-use CLI application.
- Can be installed globally for convenient access from anywhere in the terminal.

## Prerequisites

To use this tool, ensure you have:
- [Go](https://golang.org/) installed (v1.16 or later recommended).
- An internet connection (to fetch weather data).

The application uses the [WeatherAPI](https://www.weatherapi.com/) to retrieve weather data.

## Installation

### Running Locally
1. Clone the repository or download the source code.
2. Open your terminal and navigate to the folder containing `weather-cli.go`.
3. Run the application with:
   ```
   go run weather-cli.go <city>
   ```

    #### Example
      ```
       go run weather-cli.go London
      ```

## Global Installation
To make the tool accessible from any directory:

  
1.```go build```
2. ```sudo mv weather-cli /usr/local/bin```
3. ```weather-cli <city>```
