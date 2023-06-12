## Init

- mkdir go-biz-admin
- cd go-biz-admin
- go mod init github.com/mousepotato/go-biz-admin

- How to update Git and tag, with new changes:

```
git add .
git commit -m "update"
git push
git tag xxx (e.g., 3_create_user_and_live_reload)
git push origin --tags
git tag (list all tags)
```

## Gin hello world

- go get -u github.com/gin-gonic/gin

## Basic Go Syntax

## GORM

- add DB

- git tag: 1_init_project

## Now create package for database and routes

- git tag: 2_refactor_with_package

## Now create and register user

- create users (dummy)
- live reloading using https://github.com/cosmtrek/air

```
# binary will be $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

- then start with `air`

- git tag: 3_create_user_and_live_reload

## register/login user with db and password encrypt

- git tag: 4_register_login_password_encrypt

## jwt, cookie, and logout

- git tag: 5_jwt_cookie_logout

## user CRUD

- git tag: 6_user_crud