# alexscerba.com

## Overview

* Developed in Go
* Use of builtin text/template system
* [See the site in action](https://alexscerba.com/)

## Structure Planning

### Templates
* base - _useful constants: head, header/nav, footer_
* index - _homepage: displays all projects and short 'about' in TBD layout_
* projectN - _individual project pages_

### Go Structure
* main.go - _high level redirects, handle function and http server calls_
* middle.go - _gzip + any other intermediary_
* handle.go - _handle function definitions_
* load.go - _reads and loads post into a struct + other relevant functions_
* render.go - _combines templates and renders the final output_
* errors.go - _functions for error handling and logging_

### Folder Structure
```
\-- cmd/http
    +-- main.go
    +-- middle.go
    +-- handle.go
    +-- load.go
    +-- render.go
    +-- errors.go
\-- html
\-- static
    \-- media
    \-- fonts
        \-- <font-folder>
    +-- style.css
    +-- favicon.ico
    +-- logo.svg
    +-- sitemap.xml
```