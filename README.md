# RediSearchd

Redisearch API

`redisearch-go` 以 `submodule` 的形式被引入到当前工程，便于同步维护
```shell
git submodule -b develop git@gitlab.xtc.home:xtc/redisearch-go.git libs/redisearch-go
```

1. clone project

    ```shell
    git clone git@gitlab.xtc.home:xtc/redisearchd.git
    git submodule init && git submodule update
    ```
    or 
    ```shell
    git clone --recursive git@gitlab.xtc.home:xtc/redisearchd.git
    ```

2. update submodule
    ```shell
    cd libs/redisearch-go
    git checkout develop
    git pull
    ```
