FROM debian:jessie

COPY eventsservice /eventsservice
RUN useradd eventsservice
USER eventsservice

ENV LISTEN_URL=0.0.0.0:8181
EXPOSE 8181
CMD [ "/eventsservice" ]