# Klammeräffchen

This is a simple discord soundboard bot that has a simple web UI to add new sounds with just logging in into the interface 
via discord. It is built to run only on one single server at a time, or multiple servers, that all have the same users, because all sounds are sharded among all guilds.
So you will have to self-host this application in order to use it on your own.

## Project status

This project is currently done. Due to the unknown recommendations in the beginning, there are some technical decisions 
that does not make any sense, like using websockets. Simple REST would have been better in this case, but because the implementation was already done, it was easier to keep it in place. 
So please be patient about the questionable decisions made in this project.

## Installation 

Just modify the `docker-compose.yml` for your needs.