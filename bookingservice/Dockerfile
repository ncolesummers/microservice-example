FROM byrnedo/alpine-curl

COPY bookingservice /bookingservice

ENV LISTEN_URL=0.0.0.0:8181
EXPOSE 8181
EXPOSE 9100
CMD ["/bookingservice"]