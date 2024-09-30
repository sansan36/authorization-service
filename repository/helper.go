package repository

import (
	"fmt"
	"strings"

	commonv1 "github.com/sansan36/authorization-service/gen/common/v1"
	"gorm.io/gorm"
)

func searchQuery(statement *gorm.DB, query *commonv1.StandardQuery, columns []string) *gorm.DB {
	if query != nil {
		if query.Search != "" {
			query.Search = strings.ToLower(query.Search)
			searchQuery := ""
			arguments := []any{}
			for _, column := range columns {
				if searchQuery != "" {
					searchQuery += " OR "
				}
				searchQuery += fmt.Sprintf("lower(%s) LIKE ?", column)
				arguments = append(arguments, "%"+query.Search+"%")
			}

			fmt.Println(searchQuery + ", " + fmt.Sprint(arguments))
			statement = statement.Where(searchQuery, arguments...)
		}
	}

	return statement
}

func PaginationCounterQuery(statement *gorm.DB, query *commonv1.StandardQuery, pagination *commonv1.StandardPaginationResponse) *gorm.DB {
	if query != nil {
		if query.Page < 1 {
			query.Page = 1
		}
		pagination.Page = query.Page

		statement.Count(&pagination.Total)
		if query.Page > 0 && query.PageSize > 0 {
			statement = statement.Offset(int(query.Page-1) * int(query.PageSize)).Limit(int(query.PageSize))
		}
	}
	return statement
}
