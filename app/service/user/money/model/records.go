package model

import "gorm.io/datatypes"

type RecordsDB struct {
	UID int64 `gorm:"primary_key" json:"uid"`
	// []Record{}
	Records datatypes.JSON `json:"records"`
}

// TableName return table name
func (*RecordsDB) TableName() string {
	return "np_money_records"
}

// IsValid .
func (r *RecordsDB) IsValid() bool {
	return r.UID > 0
}

func (r *RecordsDB) Add(amount int32, reason string) error {
	records, err := GetRecordsFromJSON(r.Records)
	if err != nil {
		return err
	}
	records.Add(amount, reason)
	records.RemoveExpr()
	r.Records, err = records.ToJSON()
	return err
}

func (r *RecordsDB) Get() (*Records, error) {
	return GetRecordsFromJSON(r.Records)
}
