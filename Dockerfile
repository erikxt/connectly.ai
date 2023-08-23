FROM golang:1.20-alpine

ENV GOPROXY https://goproxy.cn

WORKDIR /home/connectly

COPY . .

RUN go mod tidy && go build -o app .

WORKDIR /dist

RUN cp /home/connectly/app .

# RUN cp -r /home/connectly/cert .

EXPOSE 8080

CMD [ "/dist/app" ]