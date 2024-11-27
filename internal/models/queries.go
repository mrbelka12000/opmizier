package models

const (
	// Query1 : Retrieve the country, cities, and customers with at least n customers.
	//Level1
	Query1 = `
SELECT
    country,
    city,
    COUNT("Customer Id") AS customer_count
FROM
    customers
WHERE
    country = $1 
GROUP BY
    country, city
HAVING
    COUNT("Customer Id") >= $2
ORDER BY
    customer_count DESC;
`

	// Query2 : Retrieve the count of customers grouped by country, city, and company, showing combinations with at least n customers.
	//Level2
	Query2 = `
SELECT 
    country, 
    city, 
    company, 
    COUNT(*) AS customer_count
FROM customers
WHERE country = $1 AND city = $2  AND company = $3
GROUP BY country, city, company
HAVING COUNT(*) > $4
ORDER BY customer_count DESC;
`

	// Query3 : Retrieve the top n companies in country with the most customers, analyzing their subscription trends over the past m years.
	//Level3
	Query3 = `
WITH subscription_data AS (
    SELECT 
        country,
        company,
        EXTRACT(YEAR FROM TO_DATE("Subscription Date", 'YYYY-MM-DD')) AS subscription_year,
        COUNT(*) AS customer_count
    FROM customers
    WHERE country = $1
      AND "Subscription Date" ~ '^\d{4}-\d{2}-\d{2}$'
      AND EXTRACT(YEAR FROM TO_DATE("Subscription Date", 'YYYY-MM-DD')) >= EXTRACT(YEAR FROM CURRENT_DATE) - $2
    GROUP BY country, company, EXTRACT(YEAR FROM TO_DATE("Subscription Date", 'YYYY-MM-DD'))
),

ranked_companies AS (
    SELECT 
        country,
        company,
        SUM(customer_count) AS total_customers,
        RANK() OVER (PARTITION BY country ORDER BY SUM(customer_count) DESC) AS rank
    FROM subscription_data
    GROUP BY country, company
)

SELECT 
    country,
    company,
    subscription_year,
    customer_count
FROM subscription_data
WHERE company IN (
    SELECT company 
    FROM ranked_companies 
    WHERE rank <= $3
)
ORDER BY country, company, subscription_year;
`
)
