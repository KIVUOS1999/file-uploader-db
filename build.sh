set -e

echo "Building file-uploader-db binary"

go build -o file-uploader-db

echo "Login and pusing to docker-hub"

docker login
docker build -t kivuos1999/file-uploader-db .
docker push kivuos1999/file-uploader-db

echo "cleanup"
rm file-uploader-db

echo "build succeed"