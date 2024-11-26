-- Sort database queries based on total_exec_time:

-- The following query, uses the pg_stat_statements view,
-- shows the running queries sorted descending by total_exec_time, re-formats
-- the calls column, and deduces the prop_exec_time and sync_io_time:

SELECT interval '1 millisecond' * total_exec_time AS total_exec_time,
       to_char((total_exec_time/sum(total_exec_time) OVER()) * 100, 'FM90D0') || '%'  AS prop_exec_time,
       to_char(calls, 'FM999G999G999G990') AS calls,
       interval '1 millisecond' * (blk_read_time + blk_write_time) AS sync_io_time,
       query AS query
FROM pg_stat_statements
WHERE userid =
      (
          SELECT usesysid
          FROM pg_user
          WHERE usename = current_user
          LIMIT 1
      )
ORDER BY total_exec_time DESC
LIMIT 10;


-- Find top queries with high I/O activity:

-- The following SQL shows queries with their id and mean time in seconds.
-- The result set is ordered based on the sum of blk_read_time and
-- blk_write_time so that queries with the highest read/write are shown at the top.

SELECT userid::regrole,
       dbid,
       query,
       queryid,
       mean_exec_time/1000 as mean_time_seconds
FROM pg_stat_statements
ORDER by (blk_read_time+blk_write_time) DESC
LIMIT 10;

-- See top time-consuming queries:

-- Aside from relevant information about the database, the following SQL retrieves:
-- Number of calls
-- Consumption time as total_time_seconds (in milliseconds)
-- Minimum time (in milliseconds)
-- Maximum time (in milliseconds)
-- Mean times (in milliseconds) The result set is ordered in descending order by mean_time, showing the queries with the longest consumption time first.

SELECT userid::regrole,
       dbid,
       query,
       calls,
       total_exec_time/1000 as total_time_seconds,
       min_exec_time/1000 as min_time_seconds,
       max_exec_time/1000 as max_time_seconds,
       mean_exec_time/1000 as mean_time_seconds
FROM pg_stat_statements
ORDER by mean_exec_time desc
LIMIT 10;


-- Check queries with high memory usage:

-- The following SQL retrieves the query, its id, and relevant information about the database.
-- The result set is ordered by showing the queries with the highest memory usage at the top,
-- summing the number of shared memory blocks returned from the cache (shared_blks_hit) and the number of shared memory
-- blocks marked as "dirty" during a request needed to be written to the disk (shared_blks_dirtied).


SELECT userid::regrole,
       dbid,
       queryid,
       query
FROM pg_stat_statements
ORDER by (shared_blks_hit+shared_blks_dirtied) DESC limit 10;