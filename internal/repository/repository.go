package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/mrbelka12000/optimizer/internal/models"
)

const (
	defaultDBLoad = 10_000_000
)

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) List(ctx context.Context, req models.Request) error {

	queryWhere := " WHERE "
	var args []any
	if req.ID != 0 {
		args = append(args, req.ID)
		queryWhere += fmt.Sprintf(" index = $%v AND", len(args))
	}

	if req.FirstName != "" {
		args = append(args, req.FirstName)
		queryWhere += fmt.Sprintf(` "First Name" = $%v AND`, len(args))
	}

	if req.LastName != "" {
		args = append(args, req.LastName)
		queryWhere += fmt.Sprintf(` "Last Name" = $%v AND`, len(args))
	}

	if req.Company != "" {
		args = append(args, req.Company)
		queryWhere += fmt.Sprintf(" company = $%v AND", len(args))
	}

	if req.City != "" {
		args = append(args, req.City)
		queryWhere += fmt.Sprintf(" city = $%v AND", len(args))
	}

	if req.Country != "" {
		args = append(args, req.Country)
		queryWhere += fmt.Sprintf(" country = $%v AND", len(args))
	}

	if req.Phone1 != "" {
		args = append(args, req.Phone1)
		queryWhere += fmt.Sprintf(` "Phone 1" = $%v AND`, len(args))
	}

	if req.Phone2 != "" {
		args = append(args, req.Phone2)
		queryWhere += fmt.Sprintf(` "Phone 2" = $%v AND`, len(args))
	}

	if req.Email != "" {
		args = append(args, req.Email)
		queryWhere += fmt.Sprintf(" email = $%v AND", len(args))
	}

	if req.SubscriptionDate != "" {
		args = append(args, req.SubscriptionDate)
		queryWhere += fmt.Sprintf(` "Subscription Date" = $%v AND`, len(args))
	}

	if req.Website != "" {
		args = append(args, req.Website)
		queryWhere += fmt.Sprintf(" website = $%v AND", len(args))
	}
	queryWhere = queryWhere[:len(queryWhere)-4] // Remove the trailing " AND"

	query := `SELECT index,
	       "Customer Id",
	       "First Name",
	       "Last Name",
	       company,
	       city,
	       country,
	       "Phone 1",
	       "Phone 2",
	       email,
	       "Subscription Date",
	       website FROM customers
` + queryWhere

	if req.IsOrEnabled {
		query = strings.Replace(query, "AND", "OR", -1)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("repository.List: %w", err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		count++
	}
	fmt.Println(count)
	return nil
}
