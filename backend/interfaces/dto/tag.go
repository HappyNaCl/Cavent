package dto

type TagTypeDTO struct {
    Id  string `json:"id"`
    Name string `json:"name"`
    Tags []TagDTO `json:"tags"`
}

type TagDTO struct {
    Id   string `json:"id"`
    Name string `json:"name"`
}