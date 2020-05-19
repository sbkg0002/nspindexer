# NSP Indexer
An alternative to a nut server for Switch

## Installation

To run a simple webserver with a `index.tpl` of the current directory (and its subdirectories), run:
```
docker run --rm -v ${PWD}:/srv/http -p 8043:8043 sbkg0002/nspindexer -serverip 192.168.178.7
```
> Note: `-serverip` is mandatory since this IP/FQDN is used within the indexfile that is served.

Open Tinfoil, add a `http` enpoint on *serverip*:8043 and restart Tinfoil. All updates/new games should appear.

## Todo
- [ ] Auto regenerate whever there is a new file detected. (https://github.com/radovskyb/watcher)
- [ ] Try not to use a root container


## Thanks
Thanks to [PierreZ/goStatic](https://github.com/PierreZ/goStatic) who created the webserver part.