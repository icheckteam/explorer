FROM centos

COPY ./build/explorercli /explorercli
COPY ./dashboard/build /build

WORKDIR /

ENTRYPOINT ["/explorercli"]