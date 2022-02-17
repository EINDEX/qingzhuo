#!bash

hugo --minify --enableGitInfo
rsync -arvz --delete public qcloud:~/projects/blog/
scp docker-compose.yml qcloud:~/projects/blog/docker-compose.yml
scp blog.conf qcloud:~/projects/blog/blog.conf