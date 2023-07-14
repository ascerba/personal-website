#!/bin/bash
# converts md to .tmpl.html with pandoc

readonly POSTS_DIR="../../html/posts"
readonly POST_TEMP="POST_TEMPLATE.tmpl.html"

post=$(basename "$1" .md) 
echo "converting $post to HTML"

# pandoc for md to html
pandoc -f markdown-auto_identifiers -t html --template=$POST_TEMP --wrap=none -o "$POSTS_DIR/$post.tmpl.html" "$post".md

# # adds in filename to path no longer needed
# sed -i "s/posts\//posts\/$post/" "$POSTS_DIR/$post.tmpl.html"

# makes http & https links have target _blank and rel noopener noreferrer
perl -i -0pe 's/(<\W*a\W*[^>]*href=)(["'"'"']http[s]?:\/\/[^"'"'"'>]*["'"'"'])([^>]*>)/$1$2 target="_blank" rel="noopener noreferrer"$3/g' "$POSTS_DIR/$post.tmpl.html"