package transformer_test

import (
	"fmt"
	"log"

	"github.com/sugarme/transformer"
	"github.com/sugarme/transformer/bert"
)

func ExampleLoadConfig() {
	// url := "https://s3.amazonaws.com/models.huggingface.co/bert/bert-base-uncased-config.json"
	localFile := "bert-base-uncased-config.json"
	var config bert.BertConfig
	err := transformer.LoadConfig(&config, localFile, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.VocabSize)

	// Output:
	// 30522
}
