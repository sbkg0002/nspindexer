# NSP Indexer

To run a simple webserver with a `index.tpl` of the current directory (and its subdirectories), run:
```
docker run --rm -v ${PWD}:/srv/http -p 8043:8043 sbkg0002/nspindexer -serverip 192.168.178.7
```
> Note: `-serverip` is mandatory since this IP/FQDN is used within the indexfile that is served.

Open Tinfoil, add a `http` enpoint on *serverip*:8043 and restart Tinfoil. All updates/new games should appear.