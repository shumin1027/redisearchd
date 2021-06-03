package redis

//func TestInit(t *testing.T) {
//	addr := "127.0.0.1:6379"
//	pool := Init(addr)
//	conn := pool.Get()
//	defer conn.Close()
//	res, _ := conn.Do("PING")
//	println(res.(string))
//	println(redisearch.EscapeTextFileString("this-is-title"))
//}
//func TestCreateIndex(t *testing.T) {
//	addr := "127.0.0.1:6379"
//	Init(addr)
//	cli := Client("test")
//	sc := redisearch.NewSchema(redisearch.DefaultOptions).
//		AddField(redisearch.NewTextField("body")).
//		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
//		AddField(redisearch.NewNumericField("date"))
//	err := cli.CreateIndex(sc)
//	if err != nil {
//		println("create index error")
//	}
//}
//
//func TestCreateIndexWithIndexDefinition(t *testing.T) {
//	addr := "127.0.0.1:6379"
//	Init(addr)
//	cli := Client("test")
//	sc := redisearch.NewSchema(redisearch.DefaultOptions).
//		AddField(redisearch.NewTextField("body")).
//		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
//		AddField(redisearch.NewNumericField("date"))
//	def := redisearch.NewIndexDefinition().AddPrefix("test:")
//	err := cli.CreateIndexWithIndexDefinition(sc, def)
//	if err != nil {
//		println("create index error")
//	}
//}
//
//func TestClient(t *testing.T) {
//	addr := "127.0.0.1:6379"
//	Init(addr)
//	cli := Client("test")
//	info, _ := cli.Info()
//	data, _ := json.Marshal(info)
//	println(string(data))
//
//	list, _ := cli.List()
//	data, _ = json.Marshal(list)
//	println(string(data))
//	cli.DropIndex(false)
//}
//
//func TestHMset(t *testing.T) {
//	addr := "127.0.0.1:6379"
//	pool := Init(addr)
//	conn := pool.Get()
//	defer conn.Close()
//	data := []interface{}{"test:0"}
//	data = append(data, "title")
//	data = append(data, "this_is_title")
//	data = append(data, "body")
//	data = append(data, "this is body")
//	data = append(data, "date")
//	data = append(data, 2)
//
//	error := conn.Send("HMSET", data...)
//
//	if error != nil {
//		println(error.Error())
//	}
//	error = conn.Flush()
//	if error != nil {
//		println(error.Error())
//	}
//}
//
//func TestIndex(t *testing.T) {
//	addr := "127.0.0.1:6379"
//	pool := Init(addr)
//	conn := pool.Get()
//	defer conn.Close()
//	cli := Client("test")
//
//	properties := make(map[string]interface{})
//	properties["title"] = redisearch.EscapeTextFileString("this-is-title")
//	properties["body"] = "this is body"
//	properties["date"] = 0
//
//	doc := redisearch.Document{
//		Id:         "test:1",
//		Score:      0,
//		Payload:    nil,
//		Properties: properties,
//	}
//
//	error := cli.Index(doc)
//	if error != nil {
//		println(error.Error())
//	}
//}
//
//func TestDoc(t *testing.T) {
//	addr := "127.0.0.1:6379"
//	pool := Init(addr)
//	conn := pool.Get()
//	defer conn.Close()
//
//	properties := make(map[string]interface{})
//	properties["title"] = "this is title"
//	properties["body"] = "this is body"
//	properties["date"] = 0
//
//	doc := redisearch.Document{
//		Id:         "test:3",
//		Properties: properties,
//	}
//
//	args := make(redis.Args, 0, 1+len(doc.Properties))
//	args = append(args, doc.Id)
//	for k, f := range doc.Properties {
//		args = append(args, k, f)
//	}
//
//	error := conn.Send("HMSET", args...)
//
//	if error != nil {
//		println(error.Error())
//	}
//	error = conn.Flush()
//	if error != nil {
//		println(error.Error())
//	}
//}
