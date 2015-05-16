# GO HTTP2 Web Server
------------
Docker image for a simple http2 web server implemented by golang

Usage
------------

To start an instance:

	docker run -d --name http2server -v <your-web-dir>:/www -p 443:443 gohttp2server

To login to bash:

	docker exec -it http2server bash