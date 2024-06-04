
FROM scratch
COPY quartiles /usr/bin/quartiles
ENTRYPOINT ["/usr/bin/quartiles"]
CMD ["help"]
