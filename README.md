# optimizer

## Overview
This project consists of **three different SQL queries** to optimize. Your task is to improve the performance of each query while maintaining the correctness of results. Tests are provided to validate your optimizations.

- **Query 1**: Focus on optimizing a high volume of requests with a relatively simple query.
- **Queries 2 and 3**: Focus on optimizing and analyzing large, complex SQL queries. These queries have minimal request loads but strict timeout requirements.

### Goals:
1. Learn how to optimize a large number of requests with a small query.
2. Understand how to analyze and optimize complex SQL queries with tight performance constraints.

**Note:** After passing the tests for a query, you may remove the previous code and proceed to the next one.

---

## Usage

### Start the environment:
```bash
docker-compose up --build
```

---

## Testing Your Implementation
 
- To verify your implementation, you can run the unit tests located in the `tests/unit` folder.  
Each test file corresponds to a specific query.

---


## After completing the task:
1) Shut down your environment and remove local volumes:
```bash
docker-compose down
docker volume rm prom_data redis_data postgres_data 
```
2) Proceed to the next query.

---

## Additional Information:
### Grafana Dashboard
- A pre-configured Grafana instance is available with two dashboards.
- Access Grafana at: http://localhost:3000
- You may add or replace information on the dashboards.
- **Important**: Grafana is not persistent by default. Ensure you make it persistent or proceed cautiously.

---

## PostgreSQL Debugging Tools
- Use EXPLAIN and EXPLAIN ANALYZE to understand and debug your query plans effectively.

---

## Best Practices:

- Implement caching mechanisms, metrics, and logging where necessary.
- If you change the structure of a query, ensure that the responses remain consistent with the original implementation.
- Do not modify the provided tests, as they are essential for tracking your progress.

---


## Database Performance Optimization Guide

This guide provides a step-by-step process to identify and resolve database performance issues.

## Steps to Solve the Problem

1. **Examine `pg_stat_statements`**  
   Start by analyzing query performance using `pg_stat_statements`. Run diagnostic queries in the database to gather insights into slow or resource-intensive queries. You can find them in db/1.queries.sql

2. **Investigate Problematic Queries**  
   Identify queries causing performance bottlenecks, such as high memory usage or long execution times.

3. **Analyze Slow Queries**
   - Define metrics such as SLA thresholds or log detailed request-response data for further analysis.
   - Use tools like `EXPLAIN` and `EXPLAIN ANALYZE` to inspect query execution plans and pinpoint inefficiencies.

4. **Optimize Queries**  
   Use the insights gained from query analysis to optimize slow or resource-intensive queries.

5. **Evaluate Metrics**  
   Reassess the metrics defined earlier to measure the effectiveness of the optimizations and identify any remaining bottlenecks.

6. **Implement Caching**  
   If performance issues persist, consider implementing caching for request-response pairs to reduce database load and improve response times.

7. **Reevaluate Metrics**  
   After introducing cache, revisit the defined metrics to ensure optimal performance and stability.

8. **Explore Advanced Optimization Techniques**  
   If necessary, implement advanced strategies such as:
   - **Replication**: For high availability and read scalability.
   - **Sharding**: To distribute data horizontally across multiple nodes.
   - **Partitioning**: For managing large datasets more efficiently.

---


## Part 1, read intensive:

## Me - Interviewer:
    
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
        - Caching:
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

    Int) Excellent! What would you prioritize among these if resources are limited?

    Me) If resources are limited, I would prioritize the solutions that provide the most immediate 
    and cost-effective impact while minimizing disruption to the system. Here's how I would approach prioritization:
        1) Query and Index Optimization (Straightforward, direct database benefit).

        2) Caching (Quick, impactful, minimal cost).

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

---

