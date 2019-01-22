package indexer

const (
	schoolIndexName = "schools"
	schoolTypeName  = "school"
	schoolMap       = `
    "tweet":{
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
	UUID        string
	Name        string `json:"name"`
	Description string `json:"description"`
}
