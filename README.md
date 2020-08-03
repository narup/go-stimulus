# go-stimulus

Starter project with [golang](https://golang.org) backend and [Stimulus JS](https://stimulusjs.org) on the front-end. 
StimulusJS, as per their website, is suitable for the applications that have server-side rendered HTML at their core,
and want to sprinkle Javascript to make them sparkle. This project also uses Webpack to bundle Javascript and CSS, 
and Babel to transpile ES6 code back to vanilla JavaScript so that every environment (e.g. browser) can interpret it.

Golang backend is using [Fiber - Express inspired web framework written in Go](https://github.com/gofiber) to serve HTML 
and other static content such as Javascript and CSS. HTML can be rendered in server-side using 
[HTML Template Engine ](https://github.com/gofiber/template) supported by Fiber.

### Structure of the project
Front-end: inside a folder 'webapp' with subdirectories
   - **css** - contains the stylesheet files before bundling
   - **html** - contains all the - mostly HTML files for pages, layouts, and partials. Go backend can return server-side
    rendered HTML using html template engine
   - **public** - contains the bundled javascript and styles and used as a output directory of webpack build. The HTML 
   files loads the bundled static content via static URL as shown: 
   ```
   <script src="static/public/bundle.js" async></script>
   ```
   - **src** - contains StimulusJS controllers and front-end Javascript. index.js is the entry point 


#### To run the project

```
$ git clone git@github.com:narup/go-stimulus.git
$ cd go-stimulus
$ go get 
$ cd webapp
$ yarn install
$ yarn start
```
go to http://localhost:3000

#### For hot reloading of static content run the watch command inside webapp folder
```
$ yarn watch 
```

# References

### Quick go module commands

- go mod init - creates a new module, initializing the go.mod file that describes it.
- go build - go test, and other package-building commands add new dependencies to go.mod as needed.
- go list -m all - prints the current module’s dependencies.
- go get - changes the required version of a dependency (or adds a new dependency).
- go mod tidy - removes unused dependencies.

### Useful links followed for this project
- https://github.com/stimulusjs/stimulus-starter
- https://github.com/bnkamalesh/goapp
- https://stimulusjs.org
- https://blog.golang.org/using-go-modules
- https://github.com/gofiber/fiber
- https://www.reddit.com/r/golang/comments/a4kzqk/has_anyone_deployed_a_createreactapp_with_go/
- https://www.sitepoint.com/bundle-static-site-webpack/
- https://blog.jakoblind.no/css-modules-webpack/
- http://getskeleton.com/
- https://flaviocopes.com/package-json-watch/
