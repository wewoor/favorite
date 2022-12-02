FROM node:16

VOLUME [ "/home" ]

COPY . /home/workspace/docusaurus

WORKDIR /home/workspace/docusaurus

#RUN groupadd -r docusaurus && useradd -r -g docusaurus docusaurus \
#    && npm install && npm run start

RUN yarn install

#USER docusaurus

EXPOSE 3000

CMD [ "yarn", "start" ]

