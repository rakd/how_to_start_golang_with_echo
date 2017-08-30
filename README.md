

# how to make website with echo

This repo is to introduce Go to other staffs. 

## why am i using golang?

- simple & easy 
- parallel computation 
- single binary
- friendly with docker
- good performance

ref: https://golang.org/doc/faq

---- 

# prepare Go environments

## install Go

You can install Go by Go official website or homebrew.
You might want to manage some Go versions by gvm. It's up to you.

ref: http://golang-jp.org/doc/install


## setup GOPATH

If you'd like to put repository on

> ~/src/github.com/XXX

you should setup GOPATH in .zshrc ( or .profile or .bashrc ) as below.

```
export GOPATH=$HOME
PATH=$PATH:$GOPATH/bin/
```

The path is very important to import libraries.

ref: https://golang.org/cmd/go/#hdr-GOPATH_environment_variable



## setup glide
`go get` enable you to manage go packages though, it's used by all go projects. 
You must want to use specific version of libraries for your project. 

To use specific version, you can prepare `vendor` dir under your project dir and can put the libraries unde r the dir. 

ref: https://golang.org/cmd/go/#hdr-Vendor_Directories says

You don't need to manage the dir manually, you should prepare `glide.yaml` and use `glide` to manage it. it's like `package.json`. 

To install glide, 
```
curl https://glide.sh/get | sh
```
OR
```
brew install glide
```

Please check the official repo of glide. ref: https://github.com/Masterminds/glide


If you want to add more libraries, you can just do
```
glide get LIBRARY
```
instead of
```
go get LIBRARY
```



## setup ATOM with go-plus

I'm using ATOM editor with go-plus.



## install fresh

Go is compile language but you can use it as LL as follow.

```
go run *.go
```

You still need to re-run after rewrite your code.

To avoid re-run, you can use fresh. It enable you to live-reloading.

To install fresh, you can do
```
go get github.com/pilu/fresh
```

## install mysql

```
brew install mysql
```



## install docker

please install docker.
ref: https://docs.docker.com/docker-for-mac/install/




-----

# 01Hello Echo 

## move to hello dir

```
$ mkdir 01hello
$ cd 01hello
```

## prepare glide.yaml

To prepare glide.yaml, I did 
```
$ glide init
```
and import some libraries by 
```
$ glide get XXXX
```

e.g.
```
$ glide get github.com/labstack/echo
$ glide get github.com/labstack/echo/middleware
```

## main.go


## contorllers/welcome.go


## tips
### logging

`log.Print` & `log.Printf` are useful. and adding below, you can see line where you show the log.

```main.go
func init() {
	log.SetFlags(log.Lshortfile)
}
```

## run

```
$ fresh
```

or 

```
$ go run main.go
```

or 

```
$ go build main.go
$ ./main
```


------

#02 form sample with csrf/session/template/bindata.

ref: https://echo.labstack.com/guide/templates

## install go-bindata

```
$ go get -u github.com/jteeuwen/go-bindata/...
```
ref: https://github.com/jteeuwen/go-bindata

## preapre formapp dir

- using session but no auth.
- using template 
- using bindata for template
- dummy login form 
- no Database

## prepare dir
```
$ mkdir 02formapp
$ cd 02formapp
```

## glide get 

```
$ glide get github.com/ipfans/echo-session
$ glide get "github.com/elazarl/go-bindata-assetfs"
$ glide get "github.com/itsjamie/go-bindata-templates"
$ glide get "github.com/gorilla/sessions"
```
## prepare template files 

## make bindata file
```
go-bindata -pkg=main -prefix=data -o=bindata.go data/...
```
## load bindata file. 


## prepare Makefile



## middleware for csrf/session in main.go

```
		//csrf
		engine.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
			TokenLookup: "form:csrf_token",
		}))
		//session
		store := session.NewCookieStore([]byte("secret-key-kaskdfkedknjjkelaasdfjkajkeasd"))
		store.MaxAge(86400)
		engine.Use(session.Sessions("GSESSION", store))
```

## prepare login/logout/dashboard handler 


## run 
you can just use `make`, then you can access http://localhost:3000/



----

# 03 registration 

it includes 
- flush message 
- database 
etc


## prepare dir 

```
$ mkdir 03registration
$ cd 03registration
```

```
$ glide get "github.com/jinzhu/gorm/dialects/mysql"
$ glide get "github.com/jinzhu/gorm"

```
## prepare commong logic in controllers/init-common.go

## prepare flash message template in data/templates/partial_flash.html

## prepare models 

## prepare signup/login controllers 


etcetc...



------


----

# 04 basic auth for admin pages 

it includes 
- basic auth 
- cool logging

## basic auth

```

	adminGroup := engine.Group("/admin")
	//basic auth for admin page
	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "test" && password == "test" {
			return true, nil
		}
		return false, nil
	}))

	adminGroup.GET("", controllers.AdminIndexHandler())
	adminGroup.GET("/", controllers.AdminIndexHandler())

```

## cool logging

```
	// cool custom logging
	engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${method} | ${status} | ${uri} -> ${latency_human}` + "\n",
	}))
```
------

# 05 google auth for admin pages



## glide get

```
$ glide get "golang.org/x/oauth2"
$ glide get "github.com/satori/go.uuid"
```

## controllers/google.go
## handler in main.go

```
	engine.GET("/auth/google/login", controllers.AuthGoogleLoginHandler())
	engine.GET("/auth/google/callback", controllers.AuthGoogleCallbackHandler())
	engine.GET("/auth/google/logout", controllers.AuthGoogleLogoutHandler())
```
## middleware in main.go

```

	// google auth for admin
	adminGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			adminEmails := map[string]bool{}
			adminEmails["rakd0930@gmail.com"] = true

			session := session.Default(c)
			googleEmailTmp := session.Get("google_oauth_email")
			googleEmail, ok := googleEmailTmp.(string)
			if !ok || googleEmailTmp == nil || googleEmail == "" {
				log.Print("need to login via google oauth")
				return c.Redirect(302, "/auth/google/login")
			}
			if !adminEmails[googleEmail] {
				//NG, non admin google email
				return c.Redirect(302, "/auth/google/login")
			}
			// OK
			//before

			err := next(c)
			//after

			return err
		}
	})

```


