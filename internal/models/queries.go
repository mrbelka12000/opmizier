package models

const (
	// Query1 : Pass arguments to filter results
	Query1 = `
SELECT
    country,
    city,
    COUNT("Customer Id") AS customer_count
FROM
    customers
WHERE
    country = $1 -- Pass the country as a parameter
GROUP BY
    country, city
HAVING
    COUNT("Customer Id") >= $2 -- Pass the minimum customers as a parameter
ORDER BY
    customer_count DESC;
`
)
