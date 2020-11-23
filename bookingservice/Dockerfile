FROM debian:jessie

COPY bookingservice /bookingservice
RUN useradd bookingservice
USER bookingservice

ENV LISTEN_URL=0.0.0.0:8181
EXPOSE 8181
CMD ["/bookingservice"]