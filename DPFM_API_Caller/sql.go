package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-product-master-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-master-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strings"
)

func (c *DPFMAPICaller) GeneralRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	where := strings.Join([]string{
		fmt.Sprintf("WHERE general.Product = \"%s\" ", input.General.Product),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	general.Product
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_general_data as general 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGeneral(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) BusinessPartnersRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.BusinessPartner {
	where := strings.Join([]string{
		fmt.Sprintf("WHERE businessPartner.Product = \"%s\" ", input.General.Product),
		fmt.Sprintf("AND businessPartner.BusinessPartner = %d ", input.General.BusinessPartner[0].BusinessPartner),
		fmt.Sprintf("AND validityStartDate = \"%s\" ", input.General.BusinessPartner[0].ValidityStartDate),
		fmt.Sprintf("AND validityEndDate = \"%s\" ", input.General.BusinessPartner[0].ValidityEndDate),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	businessPartner.Product,
    	businessPartner.BusinessPartner,
    	businessPartner.ValidityStartDate,
    	businessPartner.ValidityEndDate
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_product_master_business_partner_data as businessPartner 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToBusinessPartner(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
