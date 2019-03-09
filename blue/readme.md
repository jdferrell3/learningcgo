# Red vs Blue: HTTP Showdown

Another challenge from Baltimore GoLang Meetup ([https://gist.github.com/jboursiquot/c3b8bec7dcf7589e2107dd7d72eb22e0](https://gist.github.com/jboursiquot/ec954fd2d39ff5b758b2d3342de5c3e9))

## Red Team: Build a temperamental HTTP server

Your mission, should you choose to accept it, is to build an HTTP server in Go that behaves erratically. Yes, your job is to write an HTTP server that misbehaves by doing the following in a non-deterministic fashion:

- Delays responses to clients randomly between 10 and 30 seconds sometimes
- Randomly errors out with HTTP 5XX responses
- Rate-limits client requests adaptively in real time (for example, a client that makes 10 consecutive requests within 500 ms of each other should get throttled and receive an HTTP 429 status code)

You can route all incoming requests to the same handler (i.e. no need to separate routes, just `/` will suffice).

Success Criteria:

- Keep a counter for each HTTP status you send back and how long you took to send back the response on average
- Expose a `/status` endpoint to return the counts in a `JSON` blob

## Blue Team: Build a resilient HTTP client

Your mission, if you choose to accept it, is to write a resilient http client that can keep up with Red Team's shenanigans. In other words, your Go HTTP client should be built in such a way that it handles being rate-limited and receiving errors and handling them gracefully.

You task is to successfully send 10,000 HTTP GET requests and handling every single one whether you receive an error or not and whether you get rate-limited or not.

Success Criteria:

- Keep a counter for each HTTP status code you receive and how long they took on average
- Expose a `/status` endpoint to return the counts in a `JSON` blob