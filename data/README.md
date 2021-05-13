# Rippled Data

Example files are stored in this folder.

Rebuild `bindata.go` with the following command:

```
$ go-bindata -pkg data -ignore '(go|md|mod|sum)$' .
```