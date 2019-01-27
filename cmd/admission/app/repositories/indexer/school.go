package indexer

const (
	schoolIndexName = "schools"
	schoolTypeName  = "school"
	schoolMap       = `
    "school":{
      "properties":{
        "name":{
          "type":"keyword"
        },
        "description":{
          "type":"text",
          "store": true,
          "fielddata": true
        }
      }
    }
    `
)

type School struct {
	UUID string `json:"uuid"`

	Name        string `json:"name"`
	Description string `json:"description"`
}
