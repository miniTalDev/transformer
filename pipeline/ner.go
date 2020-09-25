package pipeline

// Named Entity Recognition pipeline
// Extracts entities (Person, Location, Organization, Miscellaneous) from text.
// Pretrained models are available for the following languages:
// - English
// - German
// - Spanish
// - Dutch
//
// The default NER mode is an English BERT cased large model finetuned on CoNNL03, contributed by the [MDZ Digital Library team at the Bavarian State Library](https://github.com/dbmdz)

// Entity holds entity data generated by NERModel
type Entity struct {
	// String representation of the Entity
	Word string
	// Confidence score
	Score float64
	// Entity label (e.g. ORG, LOC...)
	Label string
}

// NERModel is a model to extract entities
type NERModel struct {
	tokenClassificationModel TokenClassificationModel
}

// NewNERModel creates a NERModel from input config
func NewNERModel(config TokenClassificationModel) *NERModel {
	return &NERModel{
		tokenClassificationModel: config,
	}
}

// Predict extracts entities from input text and returns slice of entities with score
func (nm *NERModel) Predict(input []string) []Entity {
	tokens := nm.tokenClassificationModel.Predict(input, true, false)

	var entities []Entity
	for _, tok := range tokens {
		if tok.Label != "0" {
			entities = append(entities, Entity{
				Word:  tok.Text,
				Score: tok.Score,
				Label: tok.Label,
			})
		}
	}

	return entities
}
