docker rmi my-golang-app
docker build --force-rm -t my-golang-app .

if [[ $? -eq 0 ]]; then
    clear
    docker run -t --rm --name my-running-app my-golang-app
fi
