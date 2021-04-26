# RediSearchd

Redisearch Restful API , like [meilisearch](https://docs.meilisearch.com/references/)

apis: `docs/redisearchd.http`
```
GET     /indexes
GET     /indexes/{{index}}
POST    /indexes/{{index}}
DELETE  /indexes/{{index}}

GET     /indexes/{{index}}/search
POST    /indexes/{{index}}/search

GET     /docs/{{docid}}
POST    /docs
DELETE  /docs/{{docid}}
DELETE  /docs

```
---
