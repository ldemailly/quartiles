FROM scratch
COPY quartiles /usr/bin/quartiles
COPY /usr/share/dict/words /usr/share/dict/words
ENTRYPOINT ["/usr/bin/lll-fixer"]
CMD ["help"]
