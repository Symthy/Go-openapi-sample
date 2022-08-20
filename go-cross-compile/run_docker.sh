docker run \
  --name gotest -it -v ".:/work" \
  1.19.0-windowsservercore-1809 \
  sh -c 'cd /work; go test -v -tags=windows ./...'
result="$?"
docker rm gotest
exit ${result}