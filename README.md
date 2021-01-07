# About

A simple Go web application.

> Note:
>
> This repository is link to Docker Hub and is automatically build after `git push`

### To run the container using `docker container run`

```
 4:47:59 361  docker container run --name snippetbox -p 4000:4000 choonsiong/snippetbox
Unable to find image 'choonsiong/snippetbox:latest' locally
latest: Pulling from choonsiong/snippetbox
da7391352a9b: Already exists 
14428a6d4bcd: Already exists 
2c2d948710f2: Already exists 
732d84bc9610: Pull complete 
d920ed83a910: Pull complete 
7e4abbd28673: Pull complete 
78bc90a51009: Pull complete 
888ad530200e: Pull complete 
cfb5496f1177: Pull complete 
e42820fb090c: Pull complete 
6b95770397ea: Pull complete 
Digest: sha256:f39ffac7226dada503932ecaa8fc498a82bda869cc2f674f8f3a0115e073ed95
Status: Downloaded newer image for choonsiong/snippetbox:latest
2021/01/03 20:51:10 Starting server on :4000
```

### To run the container using `docker compose`

```
 20:35:29 358  docker-compose up
Creating network "snippetbox_my_net" with driver "bridge"
Creating snippetbox_snippetbox_1 ... done
Creating snippetbox_mysql_1      ... done
Attaching to snippetbox_mysql_1, snippetbox_snippetbox_1
snippetbox_1  | INFO	2021/01/07 12:35:34 Environment variable $PORT is not defined, set http listening address to :4000
snippetbox_1  | INFO	2021/01/07 12:35:34 Starting server on :4000
mysql_1       | 2021-01-07 12:35:34+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.7.32-1debian10 started.
mysql_1       | 2021-01-07 12:35:34+00:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
mysql_1       | 2021-01-07 12:35:34+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.7.32-1debian10 started.
mysql_1       | 2021-01-07T12:35:34.276590Z 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
mysql_1       | 2021-01-07T12:35:34.277584Z 0 [Note] mysqld (mysqld 5.7.32) starting as process 1 ...
mysql_1       | 2021-01-07T12:35:34.279741Z 0 [Note] InnoDB: PUNCH HOLE support available
mysql_1       | 2021-01-07T12:35:34.279783Z 0 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
mysql_1       | 2021-01-07T12:35:34.279788Z 0 [Note] InnoDB: Uses event mutexes
mysql_1       | 2021-01-07T12:35:34.279789Z 0 [Note] InnoDB: GCC builtin __atomic_thread_fence() is used for memory barrier
mysql_1       | 2021-01-07T12:35:34.279791Z 0 [Note] InnoDB: Compressed tables use zlib 1.2.11
mysql_1       | 2021-01-07T12:35:34.279792Z 0 [Note] InnoDB: Using Linux native AIO
mysql_1       | 2021-01-07T12:35:34.279953Z 0 [Note] InnoDB: Number of pools: 1
mysql_1       | 2021-01-07T12:35:34.280079Z 0 [Note] InnoDB: Using CPU crc32 instructions
mysql_1       | 2021-01-07T12:35:34.281401Z 0 [Note] InnoDB: Initializing buffer pool, total size = 128M, instances = 1, chunk size = 128M
mysql_1       | 2021-01-07T12:35:34.286323Z 0 [Note] InnoDB: Completed initialization of buffer pool
mysql_1       | 2021-01-07T12:35:34.287951Z 0 [Note] InnoDB: If the mysqld execution user is authorized, page cleaner thread priority can be changed. See the man page of setpriority().
mysql_1       | 2021-01-07T12:35:34.298996Z 0 [Note] InnoDB: Highest supported file format is Barracuda.
mysql_1       | 2021-01-07T12:35:34.304177Z 0 [Note] InnoDB: Creating shared tablespace for temporary tables
mysql_1       | 2021-01-07T12:35:34.304226Z 0 [Note] InnoDB: Setting file './ibtmp1' size to 12 MB. Physically writing the file full; Please wait ...
mysql_1       | 2021-01-07T12:35:34.317290Z 0 [Note] InnoDB: File './ibtmp1' size is now 12 MB.
mysql_1       | 2021-01-07T12:35:34.318115Z 0 [Note] InnoDB: 96 redo rollback segment(s) found. 96 redo rollback segment(s) are active.
mysql_1       | 2021-01-07T12:35:34.318137Z 0 [Note] InnoDB: 32 non-redo rollback segment(s) are active.
mysql_1       | 2021-01-07T12:35:34.318925Z 0 [Note] InnoDB: Waiting for purge to start
mysql_1       | 2021-01-07T12:35:34.370737Z 0 [Note] InnoDB: 5.7.32 started; log sequence number 12617603
mysql_1       | 2021-01-07T12:35:34.371085Z 0 [Note] InnoDB: Loading buffer pool(s) from /var/lib/mysql/ib_buffer_pool
mysql_1       | 2021-01-07T12:35:34.371208Z 0 [Note] Plugin 'FEDERATED' is disabled.
mysql_1       | 2021-01-07T12:35:34.373383Z 0 [Note] InnoDB: Buffer pool(s) load completed at 210107 12:35:34
mysql_1       | 2021-01-07T12:35:34.376218Z 0 [Note] Found ca.pem, server-cert.pem and server-key.pem in data directory. Trying to enable SSL support using them.
mysql_1       | 2021-01-07T12:35:34.376252Z 0 [Note] Skipping generation of SSL certificates as certificate files are present in data directory.
mysql_1       | 2021-01-07T12:35:34.376752Z 0 [Warning] CA certificate ca.pem is self signed.
mysql_1       | 2021-01-07T12:35:34.376826Z 0 [Note] Skipping generation of RSA key pair as key files are present in data directory.
mysql_1       | 2021-01-07T12:35:34.377258Z 0 [Note] Server hostname (bind-address): '*'; port: 3306
mysql_1       | 2021-01-07T12:35:34.377325Z 0 [Note] IPv6 is available.
mysql_1       | 2021-01-07T12:35:34.377338Z 0 [Note]   - '::' resolves to '::';
mysql_1       | 2021-01-07T12:35:34.377355Z 0 [Note] Server socket created on IP: '::'.
mysql_1       | 2021-01-07T12:35:34.378399Z 0 [Warning] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.
mysql_1       | 2021-01-07T12:35:34.383644Z 0 [Note] Event Scheduler: Loaded 0 events
mysql_1       | 2021-01-07T12:35:34.383889Z 0 [Note] mysqld: ready for connections.
mysql_1       | Version: '5.7.32'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server (GPL)

 20:36:01 358  
 20:36:01 358  docker container ls
CONTAINER ID   IMAGE           COMMAND                  CREATED          STATUS          PORTS                 NAMES
9206a0d469c7   mysql:5.7       "docker-entrypoint.s…"   31 seconds ago   Up 30 seconds   3306/tcp, 33060/tcp   snippetbox_mysql_1
b1c26705a453   my-snippetbox   "./snippetbox"           31 seconds ago   Up 30 seconds   4000/tcp              snippetbox_snippetbox_1
 20:36:04 358 
```
### To deploy the app to Heroku

[https://devcenter.heroku.com/articles/getting-started-with-go](https://devcenter.heroku.com/articles/getting-started-with-go)

[https://lcs-snippetbox.herokuapp.com/](https://lcs-snippetbox.herokuapp.com/)

Run commands below to have Heroku generate your app url automatically, you can also create a new app manually in Heroku and follow the instructions there to push your app to Heroku.

```
$ heroku login
$ heroku create
$ git push heroku master
```

![image](image.png)