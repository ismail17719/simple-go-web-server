This is a project that demonstrates how to create a simple Go web server. There are three routes on the server that can accept user's requets and return an HTTP response.

These are the routes:
- / This route demonstrates how to return a static html file to a user request. In our case its index.html
- /greet This route demonstrates how to call a function in response to a user request
- /form.html to return form.html file in the static files directory. This file has basic login form and "admin" is the default username and password
- /form This route shows basic form handling by first calling a function and then return an html file that contains the form.
