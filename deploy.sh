time {
# Will only work after docker login!
eval $(docker-machine env eqb-bot)

# In case we run this from a different directory
cd $GOPATH/src/github.com/yamamushi/EQB

# Build and tag
docker build -t eqb-bot .
docker tag eqb-bot yamamushi/eqb-bot:latest

# Cleanup remote
docker stop remote-eqbbot
yes | docker container prune

# We use the name remote-eqbbot to avoid confusion with localhost
docker run --name remote-eqbbot -d yamamushi/eqb-bot

# Announce deployment has completed
say deployment completed
}