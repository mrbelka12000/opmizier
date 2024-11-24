# opmizier

### Project to learn how to optimize big data queries

## Part 1, read intensive:

## Me - Interviewer 1:
    
    Int) Our database started to work slowly. What will you do with it ?

    Me) Do we know what type of problem it is , read or write intensive ? (Identifying bottlenecks)

    Int) Yes we do. It is read intensive, and we have a lot of slow readers.

    Me) Okay, let's go to the next. We need to identify the root causes of the problem. After that
    I can use various tools and methods to monitor and analyze our database performance, 
    such as query execution plans, performance counters, profiling tools, and logs.

    Int) What metrics will you use ? 

    Me) It will be pg_stats, SLA(service level agreement).

    Int) What you can do with it ? 

    Me) We defined which queries is working slowly, and now we can start to investigate and optimize queiries. 
    by filtering, preprocessing. After all that i will use indexes to speed up data access and avoid table scans. 
    Batch operations to reduce networks calls

    Int) Okay, we optimized queries but still have performance issues

    Me) If optimizing the queries didn’t fully resolve the performance issues, I would move on to 
    database-level and infrastructure-level strategies. Here’s what I would consider next:
        - Caching, (TODO you can implement it in code):
            Implement a caching layer, such as Redis or Memcached, to store frequently 
            accessed data. This reduces the load on the database by serving repeat queries from the cache.

        - Read replicas:
            Introduce read replicas to distribute the read workload. Applications can be 
            configured to send read queries to replicas while the primary database handles writes.

        - Sharding:
            If the dataset is very large, consider sharding the database. Split the data across 
            multiple databases based on a shard key, such as user IDs, to balance the load.

        - Connection pooling:
            Ensure the application is using a connection pool manager, such as PgBouncer for PostgreSQL, to 
            reduce overhead from creating and tearing down database connections.
        
    Int) Good points! What would you do if we still see contention issues despite adding replicas?

    Me) If contention issues persist even after adding read replicas, I would focus on resolving 
    database locking and concurrency problems. Here's what I would do:
        - Analyze Lock Contention:
            Investigate long-running queries or frequent updates that might be holding locks for extended periods.

        - Optimize Locking Behavior:
            Reduce Lock Scope: Use more granular locks where possible (e.g., row-level locks instead of table-level locks).

        - Separate Read and Write Workloads:
            Use eventual consistency models where appropriate to reduce locking between reads and writes.

        - Queue Writes:
            Use a queuing system like RabbitMQ or Kafka to manage write operations. 
            This can help smooth spikes in write traffic and reduce contention on the primary database.
        
    Int) Excellent! What would you prioritize among these if resources are limited?

    Me) If resources are limited, I would prioritize the solutions that provide the most immediate 
    and cost-effective impact while minimizing disruption to the system. Here's how I would approach prioritization:
        1) Caching (Quick, impactful, minimal cost).

        2) Query and Index Optimization (Straightforward, direct database benefit).

        3) Read Replicas Utilization (Leverage existing or easily deployable infrastructure).

        4) Connection Pooling (Helps manage contention with minimal infrastructure).

        5) Partitioning (Improves performance for specific datasets).

        6) Scaling or Sharding (More costly but addresses systemic bottlenecks).

    Int) Makes sense. How would you assess when to stop optimizing and declare success?

    Me) Deciding when to stop optimizing and declare success requires balancing performance 
    improvements against diminishing returns and business priorities. Here’s how I would approach this decision:
        1) Performance metrics meet or exceed SLA requirements.

        2) System remains stable under peak loads.

        3) User complaints about slow queries or timeouts are resolved.

        4) Cost and complexity of further optimizations outweigh the expected benefits.

        5) A monitoring and alerting system is in place for future issues.



