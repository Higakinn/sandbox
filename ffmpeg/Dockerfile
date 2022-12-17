FROM jrottenberg/ffmpeg

RUN apt-get update && apt-get install -y curl unzip
RUN curl -O https://ns1.gmkfreefonts.com/fuentes/m/monofonto.zip

RUN unzip monofonto.zip

ENTRYPOINT [ "sleep", "86400" ]