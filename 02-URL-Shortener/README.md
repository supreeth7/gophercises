# URL Shortener

A http.Handler that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

- For instance, if we have a redirect setup for ```/dinos``` to ```https://www.somesite.com/extinction-of-dinosaurs```.
- It would look for any incoming web requests with the path ```/dinos``` and redirect them.

### Instructions

- Use the ```go build``` command to generate the binary file in the current directory.
- Execute the binary in your terminal which will start a server listening on the port ```8080```.
- Open your browser and go to ```localhost:8080/${path}``` to invoke the URL shortener handler functions.

