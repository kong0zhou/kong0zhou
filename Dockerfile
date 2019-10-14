FROM golang:kz as back
COPY ./src /home/src
RUN cd /home/src && go build

FROM alpine

EXPOSE 8083

COPY ./src/conf.yaml /home/
COPY ./client/dist /home/dist
COPY --from=back /home/src/src /home/

WORKDIR /home

RUN mkdir logfiles
ENV DIRPATH=./logfiles
ENV UID=admin
ENV PASSWORD=newtest
CMD ["sh","-c","./src"]