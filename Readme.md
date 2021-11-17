Just a basic streamer that streams from files
`docker build  -t my_streamer .`
`docker run -p 5000:5000 --rm --name my-go-docker-run-1 my_streamer -fname cmd/wheels.txt -port 5000`
`docker run -p 3000:3000 --rm --name my-go-docker-run-2 my_streamer -fname cmd/baby.txt -port 3000`