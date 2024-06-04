
FROM scratch
COPY quartiles /usr/bin/quartiles
COPY words /usr/share/dict/words
ENTRYPOINT ["/usr/bin/quartiles"]
CMD ["-logger-force-color", "help"]
