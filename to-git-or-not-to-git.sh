#!/bin/bash 
curl -s https://learn.zone01dakar.sn/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"ahmadiallo\"}}){id}}"}'>>diallo
cut -b 24-27 diallo