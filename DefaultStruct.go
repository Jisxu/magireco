package main

type DefaultStruct struct {
	TableName  string      `json:"table_name"`
	FlowId     string      `json:"flow_id"`
	TableDatas []interface{} `json:"table_data"`
}

type TableDatas struct {
	DaqType     string `json:"daq_type"`
	MachineId   int    `json:"machine_id"`
}