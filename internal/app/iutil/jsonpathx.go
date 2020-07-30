package iutil

import (
	"encoding/json"
	"fmt"
	jsonx "github.com/iancoleman/orderedmap"
	"github.com/jmespath/go-jmespath"
	"regexp"
	"strings"
)

type JsonPathx struct {
}

type FieldWithOrder struct {
	Fields *jsonx.OrderedMap `json:"fields"`
}

// 根据语法解析出完整的json对象
func (j *JsonPathx) PathParse(data []byte) ([]byte, error) {
	// 原始json数据
	var originJsonMap map[string]interface{}
	// fields的Json数据
	var fieldsOrderMap FieldWithOrder
	// 重新组装的情景相关json数据
	sterilizeSceneFieldJsonMap := make(map[string][]interface{})
	// 情景差异字段json数据
	sceneDiffFieldMap := make(map[string]*jsonx.OrderedMap)
	err := json.Unmarshal(data, &originJsonMap)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &fieldsOrderMap)
	if err != nil {
		return nil, err
	}

	if fieldsOrderMap.Fields != nil {
		for _, fieldKey := range fieldsOrderMap.Fields.Keys() {
			if attributeMap, ok := fieldsOrderMap.Fields.Get(fieldKey); ok {
				if fieldValueMap, ok := attributeMap.(jsonx.OrderedMap); ok {
					for _, attributeKey := range fieldValueMap.Keys() {
						if attributeValue, ok := fieldValueMap.Get(attributeKey); ok {
							if strings.HasPrefix(attributeKey, "@") {
								sceneNames := strings.Split(attributeKey, "|")
								for _, sceneName := range sceneNames {
									sceneName := strings.Replace(sceneName, "@", "", 1)
									if sceneDiffFieldMap[sceneName] == nil {
										sceneDiffFieldMap[sceneName] = jsonx.New()
									}
									sceneDiffFieldMap[sceneName].Set(fieldKey, attributeValue)
								}
							}
						}
					}
				}
			}
		}

		//组装情景字段
		for sceneName, fieldMap := range sceneDiffFieldMap {
			for _, fieldKey := range fieldMap.Keys() {

				if diffAttribute, ok := fieldMap.Get(fieldKey); ok {
					if attributeMap, ok := fieldsOrderMap.Fields.Get(fieldKey); ok {
						if fieldValueMap, ok := attributeMap.(jsonx.OrderedMap); ok {
							var attributeMap = make(map[string]interface{})
							for _, attributeKey := range fieldValueMap.Keys() {
								if attributeValue, ok := fieldValueMap.Get(attributeKey); ok {
									if strings.HasPrefix(attributeKey, "@") { //忽略情景字段
										continue
									}

									if diffAttributeMap, ok := diffAttribute.(jsonx.OrderedMap); ok {
										if diffAttributeValue, ok := diffAttributeMap.Get(attributeKey); ok {
											attributeMap[attributeKey] = diffAttributeValue
										} else {
											for _, diffFieldKey := range diffAttributeMap.Keys() {
												if diffFieldValue, ok := diffAttributeMap.Get(diffFieldKey); ok {
													//判断字段处理方法
													filedDealAction := diffFieldKey[:1]

													if filedDealAction == "-" { //字段以-开头的表示去掉某个字段
														continue
													}

													attributeMap[diffFieldKey] = diffFieldValue
												}

											}
										}
									}
									attributeMap[attributeKey] = attributeValue
								}
							}
							sterilizeSceneFieldJsonMap[sceneName] = append(sterilizeSceneFieldJsonMap[sceneName], attributeMap)
							//sterilizeSceneFieldJsonMap[sceneName].Set(fieldKey, attributeMap)
						}
					}
				}
			}
		}
	}

	originJsonMap["scene"] = sterilizeSceneFieldJsonMap
	/*s,_ := json.Marshal(originJsonMap)
	fmt.Println(string(s))*/
	return json.Marshal(originJsonMap)
}

// 根据元信息和源码信息替换相关内容
func (j *JsonPathx) Search(meta, source string) (string, error) {
	if len(meta) < 0 {
		return source, nil
	}
	newMeta, err := j.PathParse([]byte(meta))

	r, err := regexp.Compile("@meta([^\"]+)")
	if err != nil {
		return source, err
	}

	all := r.FindAll([]byte(source), -1)
	var metaData interface{}
	err = json.Unmarshal(newMeta, &metaData)
	if err != nil {
		return source, err
	}

	for _, path := range all {
		pathStr := string(path)
		pathSegments := strings.Split(pathStr, "|")
		pathGrammar := pathSegments[0]
		pathPipes := pathSegments[1:]

		var customPipes []string   //自定义的pipe方法
		var jmesPathPipes []string //jmes的pipe方法
		for _, pathPipe := range pathPipes {
			if strings.HasPrefix(pathPipe, "@") {
				customPipes = append(customPipes, pathPipe[1:])
			} else {
				jmesPathPipes = append(jmesPathPipes, pathPipe)
			}
		}

		jmespathExpression := strings.Replace(pathGrammar, "@meta.", "", -1)
		if len(jmesPathPipes) > 0 {
			jmespathExpression += strings.Join(jmesPathPipes, "|")
		}

		result, err := jmespath.Search(jmespathExpression, metaData)
		if err != nil {
			return source, err
		}
		isString := false //当前值是否字符串
		resultStr, ok := result.(string)
		if !ok {
			resultByte, err := json.Marshal(result)
			if err != nil {
				return source, err
			} else {
				resultStr = string(resultByte)
			}
		} else {
			isString = true
		}

		if isString { //字符串类似需要加上双引号
			resultStr = fmt.Sprintf("\"%s\"", resultStr)
		}

		customPipeObj := customPipe{data: resultStr}
		for _, cPath := range customPipes {
			resultStr = customPipeObj.exec(cPath)
		}

		source = strings.Replace(source, "\""+pathStr+"\"", resultStr, -1)

		_ = pathPipes
	}

	return source, nil
}

type customPipe struct {
	data string
}

func (p *customPipe) exec(pipeName string) string {
	switch pipeName {
	case "expand":
		p.expand()
	}
	return p.data
}

func (p *customPipe) expand() {
	p.data = strings.Replace(p.data, "[", "", -1)
	p.data = strings.Replace(p.data, "]", "", -1)
}
