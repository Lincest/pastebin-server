FROM alpine:latest
COPY ./main .
RUN apk add bash
RUN chmod +x main
EXPOSE 9989
CMD [ "./main" ]