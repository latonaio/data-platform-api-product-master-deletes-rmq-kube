package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToGeneral(rows *sql.Rows) (*[]General, error) {
	defer rows.Close()
	generals := make([]General, 0)
	i := 0

	for rows.Next() {
		i++
		general := General{}
		err := rows.Scan(
			&general.Product,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &generals, err
		}

		generals = append(generals, general)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &generals, nil
	}

	return &generals, nil
}
