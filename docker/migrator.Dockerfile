FROM neo4j:community-trixie

WORKDIR /migrator

RUN mkdir migrations

COPY ./migrations ./migrations

CMD [ "cypher-shell", "--help" ]
