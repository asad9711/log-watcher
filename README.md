# log-watcher
- Tail like utility to watch files

- Idea is to use websockets to implement continuous file watch/follow solution

## Functionalities:
- cmd line switches
    1. `-f` i.e the file which is to be read
    2. `-n` number of lines to fetch in one go 
    4. `-w` to watch/follow a file for changes  (pending)


## Components:
-  An http server to serve the file contents
-  the http server should expose a websocket endpoint, so that if user runs with follow option, the webserver should maintain 
    a continuous connection(ws) to the client and keep on relaying file data to client 
-  a cli client (using cobra) which will accept user commands and accordingly make request to http server
