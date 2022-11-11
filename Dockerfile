FROM golang:1.19
MAINTAINER xuxinnan
ADD . /nitr
WORKDIR /nitr
RUN npm install && go build
CMD ./nitr 60 
