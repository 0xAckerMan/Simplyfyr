package data

import "time"

type Project struct{
    Id int64 `json:"id"`
    Name string `json:"name"`
    Category []string `json:"category"`
    Excerpt string `json:"excerpt"`
    Description string `json:"description"`
    Assigned_to int `json:"assigned_to"`
    Created_by int `json:"created_by"`
    Created_date time.Time `json:"-"`
    Due_date time.Time `json:"Due_date"`
    Done bool `json:"done"`
}
