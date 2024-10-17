# An example GOTTH WebApp application

A simple Go WebApp using Go, Templ, TailwindCSS, and HTMX! 

Tutorial Video on Youtube: [Companion Video](https://www.youtube.com/watch?v=qd1c8LnuWjw)

The structure and tooling is based on https://github.com/TomDoesTech/GOTTH with thanks!

A spooky screenshot of the demo application, which allows user to sign up for a Halloween party:
![image](https://github.com/user-attachments/assets/5cc315d1-13ef-43ce-a4e1-95107643f8a1)

**Note**: This is not production ready code - it is simply an example application to show off how the GOTTH stack works.

### Running

The project uses:
- [Air](https://github.com/air-verse/air) for golang hot reload compiling
- [Templ](https://templ.guide/) for templating
- [TailwindCSS](https://tailwindcss.com/blog/standalone-cli) for styling
- Make for building and tooling
- Go 1.22.x or newer

It has several make targets to help you get started, run `make help` to display them:

```
build                          compile tailwindcss and templ files and build the project
get-install-air                Installs the air build reload system using cUrl
get-install-tailwindcss        Installs the tailwindcss cli
go-install-air                 Installs the air build reload system using 'go install'
go-install-templ               Installs the templ Templating system for Go
help                           print make targets
tailwind-build                 one-time compile tailwindcss styles
tailwind-watch                 compile tailwindcss and watch for changes
watch                          build and watch the project with air
```
