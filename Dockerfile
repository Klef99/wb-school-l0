FROM ubuntu:latest
LABEL authors="egor"

ENTRYPOINT ["top", "-b"]