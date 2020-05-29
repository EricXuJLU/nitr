<p align="center">
    <img alt="Nitr" height="125" src="https://raw.githubusercontent.com/juanhuttemann/nitr-agent/master/app/assets/images/logo.png" style="max-width:100%;">
    <br>
</p>

[![Build Status](https://travis-ci.org/juanhuttemann/nitr-agent.svg?branch=master)](https://travis-ci.org/juanhuttemann/nitr-agent)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/juanhuttemann/nitr-agent/blob/master/LICENSE)

# nitr agent
Nitr is a webserver that collects System and Hardware information and makes it accessible through an JSON API

### Building from source
Note: go version 1.13 or higher is required building it from the source.

#### Clone
```
git clone git@github.com:juanhuttemann/nitr-agent.git
```
#### Build
```
cd nitr-agent
go build
```

### Using Nitr Agent

**Linux**
```
./nitr-agent
```

**Windows**
You can double click the .exe file or type in cmd
```
nitr-agent.exe
```
the server will start listening on port 8000 by default

### Accessing web panel
Go to [http://localhost:8000](http://localhost:8000) in your web browser

![preview](https://raw.githubusercontent.com/juanhuttemann/nitr-agent/master/images/login-web.png)

Access with default **username** and **password**: **Admin Admin**

![preview](https://raw.githubusercontent.com/juanhuttemann/nitr-agent/master/images/panel-web.png)

# API endpoints

These endpoints allow you to get system and hardware information about your host.

## Available Endpoints

| Verb   | Endpoint                      | Action                       |
|--------|-------------------------------|------------------------------|
|GET     |'/cpu'                         | CPU Information              |

## How to Use

Call the above endpoints with ?key=secret in the URL or pass the x-api-key header with value secret you will get success response.

Example:
```
http://localhost:8000/api/v1/cpu?key=secret
```


## Powered by

* [Fiber](https://gofiber.io/) - The web framework used
* [bbolt](https://github.com/etcd-io/bbolt) - Database
* [UIKit](https://getuikit.com/) - Styling
