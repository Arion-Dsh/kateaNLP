## KateaNLP a chinese NLP project in go

#### for now 
1. segment
2. similar
3. summary
4. keywords

#### usage
first of all. you need train the model and save it or download from my website, cuz the github limit data size 100M. that will
take about a minute.

    docs :=`你们世界! Hi, 你好。`
    nlp := KateaNLP(docs)
    // download train model from my website
    nlp.Download()
    // load the train data
    nlp.Load()
    // segment
    nlp.Cut()
    



