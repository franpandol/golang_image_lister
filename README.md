# Image Server Enhancer
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview
This is a simple Go-based server that serves images from a local directory. It provides two main API endpoints: one for retrieving a specific image by its filename and another for listing all available images in the directory.

## Features
* Image Retrieval: Fetch a specific image by its filename.
* Image Listing: Get a list of all available images in the directory.
* CORS Enabled: The server is CORS-enabled, allowing requests from any origin.

## Dependencies
* github.com/gorilla/mux: A powerful URL router and dispatcher for golang.
* github.com/rs/cors: A package to handle CORS requests in Go.

## API Endpoints

### Fetch a Specific Image
* Endpoint: /api/images/{filename}
* Method: GET
* Description: Fetches a specific image by its filename.
* Parameters:
    * filename: The name of the image file to retrieve.

### List All Images
* Endpoint: /api/images
* Method: GET
* Description: Returns a JSON list of all available images in the directory.

## Running the Server

* Ensure you have Go installed on your machine.
* Clone this repository.
* Navigate to the repository directory.
* Run  `go run main.go`  to start the server.
The server will start on port 8086.

## Directory Structure
The images to be served should be placed in the ./images directory relative to the server's location.
The server currently supports .jpg and .png image formats.

## Notes
Ensure the ./images directory exists and contains the images you want to serve.
For security reasons, it's recommended to limit the allowed origins in the CORS settings when deploying to production.

## Contributing
Feel free to submit pull requests or raise issues if you find any bugs or have suggestions for improvements.