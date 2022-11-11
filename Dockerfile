#FROM golang:1.19
#MAINTAINER xuxinnan
#ADD . /nitr
#WORKDIR /nitr
#RUN go build
#CMD ./nitr 60 root sb1234567890SB 49.235.101.180:3306 nitr

FROM centos:7.9.2009
MAINTAINER xuxinnan
ADD ./nitr /nitr
WORKDIR /nitr
CMD ./nitr 60 root sb1234567890SB 49.235.101.180:3306 nitr
