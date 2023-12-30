# URL Shortener

This repository contains a Golang implementation of a URL shortener that uses Redis as its database for storing URLs. The project is designed to provide a simple and efficient solution for shortening URLs and managing their corresponding redirects.

## Features

- **Golang**: The application is written in Go, providing a fast and reliable solution for URL shortening and management.
- **Redis**: The Redis database is used for storing URLs and their corresponding redirects, allowing for efficient and scalable storage.
- **Short UUIDs**: The system generates short UUIDs for each shortened URL, ensuring uniqueness and simplicity.
- **Redirect Management**: The application handles redirects efficiently, ensuring that users are redirected to the correct destination when they access a shortened URL.

## Installation

1. Clone the repository from GitHub: `git clone https://github.com/99pouria/url-shortener.git`.
2. Install the required dependencies, such as Redis.
3. Run the application: `go run cmd/url_shortener/main.go`.

## Usage

To use the URL shortener, simply enter the long URL you want to shorten in the provided input field, and the application will generate a short URL for you. Once you have the short URL, you can use it to access the original content.

## Contributing

We welcome contributions and suggestions to improve the URL shortener. Please submit your ideas or code changes through a pull request.

## License

This project is licensed under the MIT License.

Citations:
[1] https://pkg.go.dev/github.com/mxschmitt/golang-url-shortener
[2] https://pkg.go.dev/github.com/thitiph0n/go-url-shortener
[3] https://pkg.go.dev/github.com/pantrif/url-shortener
[4] https://pkg.go.dev/github.com/mattkelly/url-shortener-go
[5] https://www.reddit.com/r/golang/comments/vahlui/code_review_url_shortener_w_the_clean_architecture/?rdt=45620
