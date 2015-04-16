# youcode

This is the code behind [YouCode](http://youcode.io/), a plateform to find talks and tutorials for developers

## Setup
### The back-end
The whole project is powered by Go and App Engine, so you need [Go](https://golang.org/), a working [workspace](https://golang.org/doc/code.html) and [Google App Engine SDK for Go](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go).
To test, run on the root directory:
```bash
cd backend
goapp get github.com/GoogleCloudPlatform/go-endpoints/endpoints
cd ../
goapp serve dispatch.yaml backend/app.yaml frontend/app.yaml
```
### The frontend
We are using [Bower](http://bower.io/) to handle Polymer's dependencies. So you need to run
```bash
cd frontend/static
bower install
```