#!/bin/bash
# converts md to .tmpl.html with pandoc
# Usage: ./postDeploy.sh title_of_post.md type
# type can be 'blog' or 'projects' in my case

readonly POSTS_DIR="../html/$2"
readonly POST_TEMP="POST_TEMPLATE.tmpl.html"

post=$(basename "$1" .md) 
echo "converting $post to HTML"
echo "Outputting at '$POSTS_DIR/$post.tmpl.html'"

# pandoc for md to html
pandoc -f markdown-auto_identifiers -t html --template=$POST_TEMP --wrap=none -o "$POSTS_DIR/$post.tmpl.html" "$post".md

# # adds in filename to path no longer needed
# sed -i "s/posts\//posts\/$post/" "$POSTS_DIR/$post.tmpl.html"

# makes http & https links have target _blank and rel noopener noreferrer
perl -i -0pe 's/(<\W*a\W*[^>]*href=)(["'"'"']http[s]?:\/\/[^"'"'"'>]*["'"'"'])([^>]*>)/$1$2 target="_blank" rel="noopener noreferrer"$3/g' "$POSTS_DIR/$post.tmpl.html"

echo "moving $1 to 'deployed' folder"
mkdir -p "deployed/$post"
mv $1 "deployed/$post/$1"