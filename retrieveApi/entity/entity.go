package entity

type ResUser struct {
	Id                      int              `json:"id"`
	Uid                     string           `json:"uid"`
	Password                string           `json:"password"`
	First_name              string           `json:"first_name"`
	Last_name               string           `json:"last_name"`
	Username                string           `json:"username"`
	Email                   string           `json:"email"`
	Avatar                  string           `json:"avatar"`
	Gender                  string           `json:"gender"`
	Phone_number            string           `json:"phone_number"`
	Social_insurance_number string           `json:"social_insurance_number"`
	Date_of_birth           string           `json:"date_of_birth"`
	Employment              DataEmployment   `json:"employment"`
	Address                 DataAddress      `json:"address"`
	Credit_card             DataCreditCard   `json:"credit_card"`
	Subscription            DataSubscription `json:"subscription"`
}

type DataEmployment struct {
	Title     string `json:"title"`
	Key_skill string `json:"key_skill"`
}

type DataAddress struct {
	City           string          `json:"city"`
	Street_name    string          `json:"street_name"`
	Street_address string          `json:"street_address"`
	Zip_code       string          `json:"zip_code"`
	State          string          `json:"state"`
	Country        string          `json:"country"`
	Coordinates    DataCoordinates `json:"coordinates"`
}
type DataCoordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type DataCreditCard struct {
	Cc_number string `json:"cc_number"`
}
type DataSubscription struct {
	Plan           string `json:"plan"`
	Status         string `json:"status"`
	Payment_method string `json:"payment_method"`
	Term           string `json:"term"`
}

type User struct {
	Id         int         `json:"id"`
	Uid        string      `json:"uid"`
	First_name string      `json:"first_name"`
	Last_name  string      `json:"last_name"`
	Username   string      `json:"username"`
	Address    DataAddress `json:"address"`
}
