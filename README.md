<p align="center">
  <a href="" rel="noopener">
 <img width=600px height=400px src="https://github.com/honyshyota/http-rest-api/blob/master/images/rest_api.png" alt="Project logo"></a>
</p>

<h3 align="center">Http rest api</h3>

---

<p align="center"> My study project with gopher school.
    <br> 
</p>

## 🧐 About <a name = "about"></a>

my task was to write a simple rest api service, with coockies,
authentication, registration, storage, and also mock storage and tests

## 🏁 Getting Started <a name = "getting_started"></a>

Use
```
git clone https://github.com/honyshyota/http-rest-api
```
for copy this repository in your local machine
and you must have postgresql installed and given trust permissions for local users

## 🔧 Running the tests <a name = "tests"></a>

```
make test
```
to run all tests via makefile.

individual package tests can be accessed via files:
```
server_internal_test.go
user_test.go
store_test.go
userrepository_test.go
```

## 🎈 Usage <a name="usage"></a>

```
make
```
Use make to build the project, then you need to enter the name of the resulting file into the console to run

```
./apiserver
```

I recommend using [httpie](https://httpie.io/) for http requests

## ⛏️ Built Using <a name = "built_using"></a>

- [PostgreSQL](https://www.postgresql.org/) - Database
- [Toml](github.com/BurntSushi/toml) - Config
- [Gorilla Mux](github.com/gorilla/mux) - Router

## 🎉 And look

![alt text](https://github.com/honyshyota/http-rest-api/blob/master/images/http_request.png)
![alt text](https://github.com/honyshyota/http-rest-api/blob/master/images/logger.png)
