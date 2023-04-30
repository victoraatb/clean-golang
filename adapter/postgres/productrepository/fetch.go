package productrepository

import (
	"context"

	"github.com/booscaaa/go-paginate/paginate"
	"github.com/victoraatb/clean-go/core/domain"
	"github.com/victoraatb/clean-go/core/dto"
)

func (repository repository) Fetch(pagination *dto.PaginationRequestParms) (*domain.Pagination, error) {
	ctx := context.Background()
	products := []domain.Product{}
	total := int32(0)

	pagin := paginate.Instance(domain.Product{})
	query, queryCount := pagin.
		Query("SELECT * FROM product").
		Page(pagination.Page).
		Desc(pagination.Descending).
		Sort(pagination.Sort).
		RowsPerPage(pagination.ItemsPerPage).
		SearchBy(pagination.Search, "name", "description").
		Select()
	{
		rows, err := repository.db.Query(
			ctx,
			*query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			product := domain.Product{}

			rows.Scan(
				&product.ID,
				&product.Name,
				&product.Price,
				&product.Description,
			)

			products = append(products, product)
		}
	}
	{
		err := repository.db.QueryRow(ctx, *queryCount).Scan(&total)
		if err != nil {
			return nil, err
		}
	}
	return &domain.Pagination{
		Items: products,
		Total: total,
	}, nil
}
