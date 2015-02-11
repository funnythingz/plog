# GoGoGo

This is web application in the Golang project template.

## use package

- Go Manager - bundle for go ... https://github.com/mattn/gom
- Web framework ... https://github.com/zenazn/goji
- Asset Management ... https://github.com/shaoshing/train
- ORMapper ... https://github.com/jinzhu/gorm
- Markdown to HTML ... https://github.com/russross/blackfriday
- Sanitizer ... https://github.com/microcosm-cc/bluemonday
- live reload ... https://github.com/codegangsta/gin
- template ... https://github.com/yosssi/ace
- AWS ... https://github.com/mitchellh/goamz
- validater ... https://github.com/asaskevich/govalidator

and more...

## tree

```
.
├── app.go
├── controller.go
├── database.yml
├── dbmap.go
├── assets
│   ├── javascripts
│   │   └── main.js
│   └── stylesheets
│       └── main.scss
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

## Getting started

```
% git clone git@github.com:funnythingz/gogogo.git
% cd gogogo/
```

### database setting

Must database setting in MySQL before try `GoGoGo`

```
% cp database.yml.sample database.yml
```

create database and user/password.

### run apprication

```
% go get github.com/mattn/gom
% gom install
% gom migration/migrate.go
% gom exec gin
```

## Command

### LiveReload server

```
% gom exec gin
```

access to `http://localhost:3000`

### migration

```
% gom run migrate/migration.go
```

### test

```
% gom test -v
```

&copy; funnythingz
