package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToGeneral(rows *sql.Rows) (*General, error) {
	defer rows.Close()
	general := General{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&general.Product,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &general, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &general, nil
	}

	return &general, nil
}

func ConvertToBusinessPartner(rows *sql.Rows) (*[]BusinessPartner, error) {
	defer rows.Close()
	businessPartners := make([]BusinessPartner, 0)
	i := 0

	for rows.Next() {
		i++
		businessPartner := BusinessPartner{}
		err := rows.Scan(
			&businessPartner.Product,
			&businessPartner.BusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &businessPartners, err
		}

		businessPartners = append(businessPartners, businessPartner)
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &businessPartners, nil
	}

	return &businessPartners, nil
}
