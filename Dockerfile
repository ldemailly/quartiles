
FROM scratch
COPY quartiles /usr/bin/quartiles
ENTRYPOINT ["/usr/bin/quartiles"]
CMD ["-logger-force-color", "help"]
