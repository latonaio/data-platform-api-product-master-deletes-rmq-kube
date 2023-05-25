package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey    string   `json:"connection_key"`
	Result           bool     `json:"result"`
	RedisKey         string   `json:"redis_key"`
	Filepath         string   `json:"filepath"`
	APIStatusCode    int      `json:"api_status_code"`
	RuntimeSessionID string   `json:"runtime_session_id"`
	BusinessPartner  *int     `json:"business_partner"`
	ServiceLabel     string   `json:"service_label"`
	APIType          string   `json:"api_type"`
	General          General  `json:"ProductMaster"`
	APISchema        string   `json:"api_schema"`
	Accepter         []string `json:"accepter"`
	Deleted          bool     `json:"deleted"`
}

type General struct {
	Product             string            `json:"Product"`
	IsMarkedForDeletion *bool             `json:"IsMarkedForDeletion"`
	BusinessPartner     []BusinessPartner `json:"BusinessPartner"`
}

type BusinessPartner struct {
	Product                string    `json:"Product"`
	BusinessPartner        int       `json:"BusinessPartner"`
	BusinessPartnerProduct *string   `json:"BusinessPartnerProduct"`
	IsMarkedForDeletion    *bool     `json:"IsMarkedForDeletion"`
	BPPlant                []BPPlant `json:"BPPlant"`
}

type BPPlant struct {
	Product               string            `json:"Product"`
	BusinessPartner       int               `json:"BusinessPartner"`
	Plant                 string            `json:"Plant"`
	AvailabilityCheckType *string           `json:"AvailabilityCheckType"`
	IsMarkedForDeletion   *bool             `json:"IsMarkedForDeletion"`
	StorageLocation       []StorageLocation `json:"StorageLocation"`
	MRPArea               []MRPArea         `json:"MRPArea"`
	WorkScheduling        []WorkScheduling  `json:"WorkScheduling"`
	Accounting            []Accounting      `json:"Accounting"`
}

type StorageLocation struct {
	Product              string `json:"Product"`
	BusinessPartner      int    `json:"BusinessPartner"`
	Plant                string `json:"Plant"`
	StorageLocation      string `json:"StorageLocation"`
	InventoryBlockStatus *bool  `json:"InventoryBlockStatus"`
	IsMarkedForDeletion  *bool  `json:"IsMarkedForDeletion"`
}

type MRPArea struct {
	Product               string  `json:"Product"`
	BusinessPartner       int     `json:"BusinessPartner"`
	Plant                 string  `json:"Plant"`
	MRPArea               string  `json:"MRPArea"`
	StorageLocationForMRP *string `json:"StorageLocationForMRP"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
}

type WorkScheduling struct {
	Product                       string  `json:"Product"`
	BusinessPartner               int     `json:"BusinessPartner"`
	Plant                         string  `json:"Plant"`
	ProductionInvtryManagedLoc    *string `json:"ProductionInvtryManagedLoc"`
	ProductProcessingTime         *int    `json:"ProductProcessingTime"`
	ProductionSupervisor          *string `json:"ProductionSupervisor"`
	ProductProductionQuantityUnit *string `json:"ProductProductionQuantityUnit"`
	ProdnOrderIsBatchRequired     *bool   `json:"ProdnOrderIsBatchRequired"`
	PDTCompIsMarkedForBackflush   *bool   `json:"PDTCompIsMarkedForBackflush"`
	ProductionSchedulingProfile   *string `json:"ProductionSchedulingProfile"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}

type Accounting struct {
	Product             string  `json:"Product"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Plant               string  `json:"Plant"`
	ValuationClass      *string `json:"ValuationClass"`
	PriceLastChangeDate *string `json:"PriceLastChangeDate"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
