package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/mrbelka12000/optimizer/internal/models"
)

const (
	defaultDBLoad = 2_000_000
)

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) List(ctx context.Context, pars models.Data) error {

	queryWhere := " WHERE "
	var args []any
	if pars.ID != 0 {
		args = append(args, pars.ID)
		queryWhere += fmt.Sprintf(" index = $%v AND", len(args))
	}

	if pars.FirstName != "" {
		args = append(args, pars.FirstName)
		queryWhere += fmt.Sprintf(" 'First Name' = $%v AND", len(args))
	}

	if pars.LastName != "" {
		args = append(args, pars.LastName)
		queryWhere += fmt.Sprintf(" 'Last Name' = $%v AND", len(args))
	}

	if pars.Company != "" {
		args = append(args, pars.Company)
		queryWhere += fmt.Sprintf(" company = $%v AND", len(args))
	}

	if pars.City != "" {
		args = append(args, pars.City)
		queryWhere += fmt.Sprintf(" city = $%v AND", len(args))
	}

	if pars.Country != "" {
		args = append(args, pars.Country)
		queryWhere += fmt.Sprintf(" country = $%v AND", len(args))
	}

	if pars.Phone1 != "" {
		args = append(args, pars.Phone1)
		queryWhere += fmt.Sprintf(" 'Phone 1' = $%v AND", len(args))
	}

	if pars.Phone2 != "" {
		args = append(args, pars.Phone2)
		queryWhere += fmt.Sprintf(" 'Phone 2' = $%v AND", len(args))
	}

	if pars.Email != "" {
		args = append(args, pars.Email)
		queryWhere += fmt.Sprintf(" email = $%v AND", len(args))
	}

	if pars.SubscriptionDate != "" {
		args = append(args, pars.SubscriptionDate)
		queryWhere += fmt.Sprintf(" 'Subscription Date' = $%v AND", len(args))
	}

	if pars.Website != "" {
		args = append(args, pars.Website)
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

	if pars.IsOrEnabled {
		query = strings.Replace(query, "AND", "OR", -1)
	}
	start := time.Now()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("repository.List: %w", err)
	}
	defer rows.Close()

	fmt.Println(time.Since(start).Seconds())
	return nil
}
