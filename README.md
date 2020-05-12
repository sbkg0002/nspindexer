# NSP Indexer

To generate a `index.tpl` of the current directory (and its subdirectories), run:
```
docker run --rm -v ${PWD}:/nsp sbkg0002/nspindexer:latest "."
```