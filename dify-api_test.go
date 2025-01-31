package dify_api

import (
	"fmt"
	"testing"
)

func Init() *DifyApi {
	o := NewDifyApi(Config{
		Url:    "http://10.8.0.22",
		ApiKey: "Bearer app-Gxq0i0leXKMFFN3duYrgXU2d",
	})
	return o
}

func TestGenerate(t *testing.T) {
	sdk := Init()
	data, err := sdk.Generate(DifyGenerateRequest{
		Inputs: map[string]any{
			"Multisentiment": "True",
			"Categories":     "候诊服务,诊疗环境,医院技术,服务态度,医生沟通,服务设施,停车服务,标识清晰,楼层分布,挂号体验",
			"input_text":     "每次来都是换药，还收挂号费，医生就开个单而已，这挂号费收得真轻松",
		},
		User:         "demo-go",
		ResponseMode: "blocking",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Print(data.Data.Outputs.Text)
}

func TestSSE(t *testing.T) {
	sdk := Init()
	req := DifyGenerateRequest{
		Inputs: map[string]any{
			"Multisentiment": "True",
			"Categories":     "候诊服务,诊疗环境,医院技术,服务态度,医生沟通,服务设施,停车服务,标识清晰,楼层分布,挂号体验",
			"input_text":     "每次来都是换药，还收挂号费，医生就开个单而已，这挂号费收得真轻松",
		},
		User:         "demo-go",
		ResponseMode: "streaming",
	}
	err := sdk.GenerateSSE(req, func(data DifySSEResponse) {
		fmt.Println(data)
		if data.Event == "workflow_finished" {
			fmt.Println(data.Data.Outputs.Text)
			fmt.Println("finish: \r")
		}
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
