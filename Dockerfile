FROM scratch
COPY ./gus ./cmd
RUN ./cmd
