FROM chromedp/headless-shell:latest
# Install dumb-init or tini
RUN apt-get update && apt-get install -y dumb-init
ADD main /usr/local/bin/main
# or RUN apt install tini
ENTRYPOINT ["dumb-init", "--"]
# or ENTRYPOINT ["tini", "--"]
CMD ["/usr/local/bin/main"]
