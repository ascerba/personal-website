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
    +-- project1.tmpl.html
    +-- project2.tmpl.html
    +-- ...
\-- static
    \-- media
        \-- projectN
            \-- hq-main
                +-- img1.jpg
                +-- img2.jpg
                +-- ...
            \-- sq-main
                +-- img1_1000.jpg
                +-- img2_1000.jpg
                +-- ...
            \-- lowres-main
                +-- img1_400.jpg
                +-- img2_400.jpg
                +-- ...
            \-- sq-thumb
                +-- img1_600.jpg
                +-- img2_600.jpg
                +-- ...
            \-- lowres-thumb
                +-- img1_200.jpg
                +-- img2_200.jpg
                +-- ...
    \-- assets
        \-- fonts
            \-- <font-folder>
        +-- style.css
        +-- favicon.ico
        +-- logo.svg
        +-- profile-pic.jpg
        +-- sitemap.xml
```
