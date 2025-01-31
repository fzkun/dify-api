package dify_api

import (
	"bufio"
	"fmt"
	"github.com/fzkun/goutil/jsonutil"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"strings"
)

type DifyApi struct {
	ctx Context
}

func NewDifyApi(cfg Config) *DifyApi {
	return &DifyApi{ctx: Context{cfg: cfg}}
}

// Generate /api/generate
func (d *DifyApi) Generate(body DifyGenerateRequest) (data DifyGenerateResponse, err error) {
	var (
		httpResp *resty.Response
	)
	if httpResp, err = resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", d.ctx.cfg.ApiKey).
		SetBody(body).Post(fmt.Sprintf(d.ctx.cfg.Url + "/v1/workflows/run")); err != nil {
		return
	}
	respJson := httpResp.String()
	if err = jsonutil.JsonStrToStruct(respJson, &data); err != nil {
		err = errors.New(fmt.Sprintf("解析json失败,err=%s,json=%s", err.Error(), respJson))
		return
	}
	return
}

// GenerateSSE /api/generate sse方式对接
func (d *DifyApi) GenerateSSE(body DifyGenerateRequest, callback func(data DifySSEResponse)) (err error) {
	var (
		httpResp *resty.Response
	)
	httpResp, err = resty.New().
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", d.ctx.cfg.ApiKey).
		SetBody(body).
		SetDoNotParseResponse(true).
		Post(fmt.Sprintf(d.ctx.cfg.Url + "/v1/workflows/run"))
	if err != nil {
		return err
	}
	defer httpResp.RawResponse.Body.Close()

	scanner := bufio.NewScanner(httpResp.RawResponse.Body)
	//reply := ""
	for scanner.Scan() {
		_res := scanner.Text()
		if _res == "" {
			continue
		}
		var data DifySSEResponse
		err = jsonutil.JsonStrToStruct(strings.Replace(_res, "data:", "", 1), &data)
		if err != nil {
			err = errors.Wrap(err, "解析json失败")
			return
		}
		callback(data)
		//reply += _res
		//fmt.Println(_res)
	}
	return err
}
