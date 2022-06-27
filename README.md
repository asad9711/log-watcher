# log-watcher
Tail like utility to watch files

Basic idea is to use websockets to implement continuous file watch/follow solution

Functionalities:
1. cmd line switches
    1. -f i.e the file which is to be read
    2. -n number of lines to fetch in one go  
    3. fileName i.e the file which is to be read
    4. -w to watch/follow a file for changes  (pending)


Components:
1. a http server to serve the file contents
2. the http server should expose a websocket endpoint, so that if user runs with follow option, the webserver should maintain 
    a continuous connection(ws) to the client and keep on relaying file data to client 
3. a cli client (using cobra) which will accept user commands and accordingly make request to http server
