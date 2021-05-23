package model

type Course struct {
    Id  string `json:"id"`
    Name  string `json:"name"`
    ProfName  string `json:"prof_name"`
    Description  string `json:"description"`
    Price  string `json:"price"`
    Online  bool `json:"online"`
    Dates []int `json:"dates"`
}
