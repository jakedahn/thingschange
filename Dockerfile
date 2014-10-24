FROM scratch
MAINTAINER Jake Dahn <jake@markupisart.com>
ADD bin/worker /worker
ENTRYPOINT ["/worker"]
