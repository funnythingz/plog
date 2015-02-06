# GoGoGo

This is web application in the Golang project template.

## use package

- Web framework ... https://github.com/zenazn/goji
- ORMapper ... https://github.com/jinzhu/gorm
- Markdown to HTML ... https://github.com/russross/blackfriday
- Sanitizer ... https://github.com/microcosm-cc/bluemonday
- live reload ... https://github.com/codegangsta/gin
- template ... https://github.com/yosssi/ace
- AWS ... https://github.com/mitchellh/goamz
- Image Resize ... https://github.com/nfnt/resize

and more...

## tree

```
.
├── app.go
├── controller.go
├── database.yml
├── dbmap.go
├── helper
│   └── helper.go
├── migrate
│   └── migration.go
├── models
│   ├── entry.go
│   └── foo.go
├── entry_test.go
└── views
    ├── layouts
    │   └── layout.ace
    ├── index.ace
    └── show.ace
```

## run

```
% cd /path/to/gogogo
% gin
```

access to `http://localhost:3000`

## migration

```
% go run migrate/migration.go
```

## test

```
% go test -v
```

&copy; funnythingz
