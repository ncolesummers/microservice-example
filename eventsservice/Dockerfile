FROM byrnedo/alpine-curl

COPY eventsservice /eventsservice

ENV LISTEN_URL=0.0.0.0:8181
EXPOSE 8181
EXPOSE 9100
CMD [ "/eventsservice" ]