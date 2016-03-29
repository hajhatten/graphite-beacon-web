# Build the image (--rm writes over previous versions)
docker build --rm -t graphite-beacon-web .

# Run container from current directory
docker run -p 3000:3000 -v $PWD:/go/src/github.com/hajhatten/graphite-beacon-web -it --name="graphite-beacon-web" graphite-beacon-web
