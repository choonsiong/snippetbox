### Running `docker-compose up` for the first time

The application will failed because the database required is not created yet.

```
WARNING: Image for service snippetbox was built because it did not already exist. To rebuild this image you must use `docker-compose build` or `docker-compose up --build`.
Creating snippetbox_mysql_1 ... done
Creating snippetbox_snippetbox_1 ... done
Attaching to snippetbox_mysql_1, snippetbox_snippetbox_1
snippetbox_1  | Wait 10 sec before starting up application...
mysql_1       | 2021-01-08 12:27:12+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.7.32-1debian10 started.
mysql_1       | 2021-01-08 12:27:12+00:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
mysql_1       | 2021-01-08 12:27:12+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.7.32-1debian10 started.
mysql_1       | 2021-01-08T12:27:13.161581Z 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
mysql_1       | 2021-01-08T12:27:13.162580Z 0 [Note] mysqld (mysqld 5.7.32) starting as process 1 ...
mysql_1       | 2021-01-08T12:27:13.164844Z 0 [Note] InnoDB: PUNCH HOLE support available
mysql_1       | 2021-01-08T12:27:13.164879Z 0 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
mysql_1       | 2021-01-08T12:27:13.164886Z 0 [Note] InnoDB: Uses event mutexes
mysql_1       | 2021-01-08T12:27:13.164892Z 0 [Note] InnoDB: GCC builtin __atomic_thread_fence() is used for memory barrier
mysql_1       | 2021-01-08T12:27:13.164896Z 0 [Note] InnoDB: Compressed tables use zlib 1.2.11
mysql_1       | 2021-01-08T12:27:13.164904Z 0 [Note] InnoDB: Using Linux native AIO
mysql_1       | 2021-01-08T12:27:13.165081Z 0 [Note] InnoDB: Number of pools: 1
mysql_1       | 2021-01-08T12:27:13.165238Z 0 [Note] InnoDB: Using CPU crc32 instructions
mysql_1       | 2021-01-08T12:27:13.166333Z 0 [Note] InnoDB: Initializing buffer pool, total size = 128M, instances = 1, chunk size = 128M
mysql_1       | 2021-01-08T12:27:13.171515Z 0 [Note] InnoDB: Completed initialization of buffer pool
mysql_1       | 2021-01-08T12:27:13.173067Z 0 [Note] InnoDB: If the mysqld execution user is authorized, page cleaner thread priority can be changed. See the man page of setpriority().
mysql_1       | 2021-01-08T12:27:13.184432Z 0 [Note] InnoDB: Highest supported file format is Barracuda.
mysql_1       | 2021-01-08T12:27:13.191356Z 0 [Note] InnoDB: Creating shared tablespace for temporary tables
mysql_1       | 2021-01-08T12:27:13.191403Z 0 [Note] InnoDB: Setting file './ibtmp1' size to 12 MB. Physically writing the file full; Please wait ...
mysql_1       | 2021-01-08T12:27:13.204668Z 0 [Note] InnoDB: File './ibtmp1' size is now 12 MB.
mysql_1       | 2021-01-08T12:27:13.205660Z 0 [Note] InnoDB: 96 redo rollback segment(s) found. 96 redo rollback segment(s) are active.
mysql_1       | 2021-01-08T12:27:13.205688Z 0 [Note] InnoDB: 32 non-redo rollback segment(s) are active.
mysql_1       | 2021-01-08T12:27:13.206327Z 0 [Note] InnoDB: 5.7.32 started; log sequence number 12632004
mysql_1       | 2021-01-08T12:27:13.206456Z 0 [Note] InnoDB: Loading buffer pool(s) from /var/lib/mysql/ib_buffer_pool
mysql_1       | 2021-01-08T12:27:13.206716Z 0 [Note] Plugin 'FEDERATED' is disabled.
mysql_1       | 2021-01-08T12:27:13.207935Z 0 [Note] InnoDB: Buffer pool(s) load completed at 210108 12:27:13
mysql_1       | 2021-01-08T12:27:13.210244Z 0 [Note] Found ca.pem, server-cert.pem and server-key.pem in data directory. Trying to enable SSL support using them.
mysql_1       | 2021-01-08T12:27:13.210270Z 0 [Note] Skipping generation of SSL certificates as certificate files are present in data directory.
mysql_1       | 2021-01-08T12:27:13.210695Z 0 [Warning] CA certificate ca.pem is self signed.
mysql_1       | 2021-01-08T12:27:13.210735Z 0 [Note] Skipping generation of RSA key pair as key files are present in data directory.
mysql_1       | 2021-01-08T12:27:13.211060Z 0 [Note] Server hostname (bind-address): '*'; port: 3306
mysql_1       | 2021-01-08T12:27:13.211146Z 0 [Note] IPv6 is available.
mysql_1       | 2021-01-08T12:27:13.211160Z 0 [Note]   - '::' resolves to '::';
mysql_1       | 2021-01-08T12:27:13.211177Z 0 [Note] Server socket created on IP: '::'.
mysql_1       | 2021-01-08T12:27:13.212299Z 0 [Warning] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.
mysql_1       | 2021-01-08T12:27:13.220310Z 0 [Note] Event Scheduler: Loaded 0 events
mysql_1       | 2021-01-08T12:27:13.220693Z 0 [Note] mysqld: ready for connections.
mysql_1       | Version: '5.7.32'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server (GPL)
snippetbox_1  | INFO	2021/01/08 12:27:23 Environment variable $PORT is not defined, set http listening address to :4000
mysql_1       | 2021-01-08T12:27:23.203388Z 2 [Note] Access denied for user 'admin'@'172.28.0.3' (using password: YES)
snippetbox_1  | ERROR	2021/01/08 12:27:23 main.go:57: Error 1045: Access denied for user 'admin'@'172.28.0.3' (using password: YES)
snippetbox_snippetbox_1 exited with code 1
```

Login to the MySQL container and import the `setup_mysql.sql` script to create the database. After that exit the contaner and run `docker compose down` to shutdown the running containers.

```
 20:29:54 357  docker container ls
CONTAINER ID   IMAGE      COMMAND                  CREATED         STATUS         PORTS                               NAMES
ce7210cb71a2   my-mysql   "docker-entrypoint.s…"   2 minutes ago   Up 2 minutes   0.0.0.0:3306->3306/tcp, 33060/tcp   snippetbox_mysql_1
 20:29:56 357 
 
  20:29:57 357  docker container exec -it ce7 bash
root@ce7210cb71a2:~# 
root@ce7210cb71a2:~# ls -l
total 4
-rw-r--r-- 1 root root 1282 Jan  8 12:25 setup_mysql.sql
root@ce7210cb71a2:~#
root@ce7210cb71a2:~# mysql -uroot -p < setup_mysql.sql 
Enter password: 
root@ce7210cb71a2:~# exit
exit
 20:30:47 357  docker compose down
The new 'docker compose' command is currently experimental. To provide feedback or request new features please open issues at https://github.com/docker/compose-cli
[+] Running 3/2
 ⠿ Container snippetbox_snippetbox_1  Removed                                                                                                                                                                                                            0.0s
 ⠿ Container snippetbox_mysql_1       Removed                                                                                                                                                                                                            1.2s
 ⠿ Network "snippetbox_my_net"        Removed                                                                                                                                                                                                            0.1s
 20:30:55 357 
```

Now if you run `docker-compose up` again, you should be able to open your web browser and access `http://localhost:4000`

```
 20:30:55 357  docker-compose up
Creating network "snippetbox_my_net" with driver "bridge"
Creating snippetbox_mysql_1 ... done
Creating snippetbox_snippetbox_1 ... done
Attaching to snippetbox_mysql_1, snippetbox_snippetbox_1
mysql_1       | 2021-01-08 12:30:59+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.7.32-1debian10 started.
snippetbox_1  | Wait 10 sec before starting up application...
mysql_1       | 2021-01-08 12:30:59+00:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
mysql_1       | 2021-01-08 12:30:59+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 5.7.32-1debian10 started.
mysql_1       | 2021-01-08T12:30:59.974359Z 0 [Warning] TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details).
mysql_1       | 2021-01-08T12:30:59.975336Z 0 [Note] mysqld (mysqld 5.7.32) starting as process 1 ...
mysql_1       | 2021-01-08T12:30:59.977414Z 0 [Note] InnoDB: PUNCH HOLE support available
mysql_1       | 2021-01-08T12:30:59.977450Z 0 [Note] InnoDB: Mutexes and rw_locks use GCC atomic builtins
mysql_1       | 2021-01-08T12:30:59.977483Z 0 [Note] InnoDB: Uses event mutexes
mysql_1       | 2021-01-08T12:30:59.977492Z 0 [Note] InnoDB: GCC builtin __atomic_thread_fence() is used for memory barrier
mysql_1       | 2021-01-08T12:30:59.977513Z 0 [Note] InnoDB: Compressed tables use zlib 1.2.11
mysql_1       | 2021-01-08T12:30:59.977518Z 0 [Note] InnoDB: Using Linux native AIO
mysql_1       | 2021-01-08T12:30:59.977651Z 0 [Note] InnoDB: Number of pools: 1
mysql_1       | 2021-01-08T12:30:59.977719Z 0 [Note] InnoDB: Using CPU crc32 instructions
mysql_1       | 2021-01-08T12:30:59.978852Z 0 [Note] InnoDB: Initializing buffer pool, total size = 128M, instances = 1, chunk size = 128M
mysql_1       | 2021-01-08T12:30:59.984163Z 0 [Note] InnoDB: Completed initialization of buffer pool
mysql_1       | 2021-01-08T12:30:59.985762Z 0 [Note] InnoDB: If the mysqld execution user is authorized, page cleaner thread priority can be changed. See the man page of setpriority().
mysql_1       | 2021-01-08T12:30:59.996909Z 0 [Note] InnoDB: Highest supported file format is Barracuda.
mysql_1       | 2021-01-08T12:31:00.003613Z 0 [Note] InnoDB: Creating shared tablespace for temporary tables
mysql_1       | 2021-01-08T12:31:00.003662Z 0 [Note] InnoDB: Setting file './ibtmp1' size to 12 MB. Physically writing the file full; Please wait ...
mysql_1       | 2021-01-08T12:31:00.016029Z 0 [Note] InnoDB: File './ibtmp1' size is now 12 MB.
mysql_1       | 2021-01-08T12:31:00.016932Z 0 [Note] InnoDB: 96 redo rollback segment(s) found. 96 redo rollback segment(s) are active.
mysql_1       | 2021-01-08T12:31:00.016956Z 0 [Note] InnoDB: 32 non-redo rollback segment(s) are active.
mysql_1       | 2021-01-08T12:31:00.018050Z 0 [Note] InnoDB: 5.7.32 started; log sequence number 12649593
mysql_1       | 2021-01-08T12:31:00.018566Z 0 [Note] InnoDB: Loading buffer pool(s) from /var/lib/mysql/ib_buffer_pool
mysql_1       | 2021-01-08T12:31:00.018653Z 0 [Note] Plugin 'FEDERATED' is disabled.
mysql_1       | 2021-01-08T12:31:00.019757Z 0 [Note] InnoDB: Buffer pool(s) load completed at 210108 12:31:00
mysql_1       | 2021-01-08T12:31:00.022130Z 0 [Note] Found ca.pem, server-cert.pem and server-key.pem in data directory. Trying to enable SSL support using them.
mysql_1       | 2021-01-08T12:31:00.022154Z 0 [Note] Skipping generation of SSL certificates as certificate files are present in data directory.
mysql_1       | 2021-01-08T12:31:00.022560Z 0 [Warning] CA certificate ca.pem is self signed.
mysql_1       | 2021-01-08T12:31:00.022599Z 0 [Note] Skipping generation of RSA key pair as key files are present in data directory.
mysql_1       | 2021-01-08T12:31:00.022903Z 0 [Note] Server hostname (bind-address): '*'; port: 3306
mysql_1       | 2021-01-08T12:31:00.022989Z 0 [Note] IPv6 is available.
mysql_1       | 2021-01-08T12:31:00.023023Z 0 [Note]   - '::' resolves to '::';
mysql_1       | 2021-01-08T12:31:00.023041Z 0 [Note] Server socket created on IP: '::'.
mysql_1       | 2021-01-08T12:31:00.024198Z 0 [Warning] Insecure configuration for --pid-file: Location '/var/run/mysqld' in the path is accessible to all OS users. Consider choosing a different directory.
mysql_1       | 2021-01-08T12:31:00.032553Z 0 [Note] Event Scheduler: Loaded 0 events
mysql_1       | 2021-01-08T12:31:00.032982Z 0 [Note] mysqld: ready for connections.
mysql_1       | Version: '5.7.32'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server (GPL)
snippetbox_1  | INFO	2021/01/08 12:31:10 Environment variable $PORT is not defined, set http listening address to :4000
snippetbox_1  | INFO	2021/01/08 12:31:10 Starting server on :4000
```