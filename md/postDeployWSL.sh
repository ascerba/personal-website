#!/bin/sh
# converts md to .tmpl.html with pandoc
# Usage: ./postDeploy.sh title_of_post.md type
# type can be 'blog' or 'projects' in my case

if [ -z $3 ]
then
  base_dir="/home/ascerba/go/alexscerba.com"
else
  base_dir="$3"
fi

readonly posts_dir="$base_dir/html/$2"
readonly post_temp="$base_dir/md/POST_TEMPLATE.tmpl.html"

post=$(basename "$1" .md) 
echo "converting $1 to HTML"
echo "Outputting at '$posts_dir/$post.tmpl.html'"

# pandoc for md to html
pandoc -f markdown-auto_identifiers -t html --template=$post_temp --wrap=none -o "$posts_dir/$post.tmpl.html" "$1"

# # adds in filename to path no longer needed
sed -i "s;src=\";src=\"/static/media/$post/;" "$posts_dir/$post.tmpl.html"

# Pull images from <p> tags
sed -i "s;<p><img;<img;" "$posts_dir/$post.tmpl.html"
sed -i "s;/></p>;/>;" "$posts_dir/$post.tmpl.html"

# Same but for video
sed -i "s;<p><video;<video;" "$posts_dir/$post.tmpl.html"
sed -i "s;</video></p>;</video>;" "$posts_dir/$post.tmpl.html"

# makes http & https links have target _blank and rel noopener noreferrer
perl -i -0pe 's/(<\W*a\W*[^>]*href=)(["'"'"']http[s]?:\/\/[^"'"'"'>]*["'"'"'])([^>]*>)/$1$2 target="_blank" rel="noopener noreferrer"$3/g' "$posts_dir/$post.tmpl.html"

media_dir="$base_dir/static/media/$post"

if [ -d "$base_dir/md/staged/images" ]
then
  echo "Converting and deploying images:"
  if [ ! -d "$base_dir/static/media/$post" ]
  then
    mkdir "$base_dir/static/media/$post"
  fi
  for image in $(ls "$base_dir/md/staged/images")
  do
    echo "Converting $image and deploying at $media_dir/$image"
    convert $base_dir/md/staged/images/$image -resize 1800x1800\> "$media_dir/$image"
    #/home/ascerba/.local/bin/cwebp -q 70 "$media_dir/$image" "$media_dir/${image%.*}.webp" <------- ignore for now
  done
fi

if [ -d "$base_dir/md/staged/videos" ]
then
  echo "Deploying videos:"
  if [ ! -d "$base_dir/static/media/$post" ]
  then
    mkdir "$base_dir/static/media/$post"
  fi
  for video in $(ls "$base_dir/md/staged/videos")
  do
    echo "Copying $video to $media_dir/$video"
    cp $base_dir/md/staged/videos/$video $media_dir/$video
  done
fi

#echo "Contents of 'staged' will not be moved or deleted"
echo "Moving staged contents to 'deployed' folder"
mv -vT "$base_dir/md/staged" "$base_dir/md/deployed/$post"
mkdir "$base_dir/md/staged"
cp "$base_dir/md/template.md" "$base_dir/md/staged/NewPost.md"
echo "Post deploy complete!"
echo "Dont forget to edit $post_dir/$post.tmpl.html if any videos were added...!"