docker build --platform amd64 . --target prod -t asia.gcr.io/run-app-341001/goapi:latest -f docker/go/Dockerfile
docker push asia.gcr.io/run-app-341001/goapi:latest