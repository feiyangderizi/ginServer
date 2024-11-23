// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/callback/account/query": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "对账"
                ],
                "summary": "查询对账明细",
                "parameters": [
                    {
                        "type": "string",
                        "description": "查询日期",
                        "name": "queryDate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/callback/claim/notice": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "回执"
                ],
                "summary": "理赔成功回执",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.CallbackClaimCompleteReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/callback/claim/replenish": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "回执"
                ],
                "summary": "理赔补充文件回执",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.CallbackClaimReplenishReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/callback/invoice": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "回执"
                ],
                "summary": "发票回执",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.CallbackInvoiceReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/callback/surrend": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "回执"
                ],
                "summary": "退保回执",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.CallbackSurrendReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "req.CallbackClaimCompleteReq": {
            "type": "object",
            "properties": {
                "claimTotalAmount": {
                    "description": "理赔总金额",
                    "type": "number"
                },
                "payeeList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/req.Payee"
                    }
                },
                "policyNo": {
                    "type": "string"
                }
            }
        },
        "req.CallbackClaimReplenishReq": {
            "type": "object",
            "properties": {
                "claimTotalAmount": {
                    "description": "理赔总金额",
                    "type": "number"
                },
                "fileName": {
                    "description": "需要用户补传的文件名称。",
                    "type": "string"
                },
                "policyNo": {
                    "type": "string"
                }
            }
        },
        "req.CallbackInvoiceReq": {
            "type": "object",
            "properties": {
                "downLoadUrl": {
                    "description": "电子发票下下载地址",
                    "type": "string"
                },
                "policyNo": {
                    "description": "保单号",
                    "type": "string"
                }
            }
        },
        "req.CallbackSurrendReq": {
            "type": "object",
            "properties": {
                "policyNo": {
                    "description": "保单号",
                    "type": "string"
                },
                "resultFlag": {
                    "description": "true 成功 false 失败",
                    "type": "string"
                },
                "resultMessage": {
                    "description": "处理结果",
                    "type": "string"
                }
            }
        },
        "req.Payee": {
            "type": "object",
            "properties": {
                "claimAccountBankName": {
                    "description": "收款开户银行名称",
                    "type": "string"
                },
                "claimAccountNo": {
                    "description": "收款人账号",
                    "type": "string"
                },
                "claimAccountUser": {
                    "description": "收款人名称",
                    "type": "string"
                },
                "claimAmount": {
                    "description": "赔付金额",
                    "type": "number"
                }
            }
        },
        "result.Page": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "index": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "result.Result": {
            "type": "object",
            "properties": {
                "data": {},
                "msg": {
                    "type": "string"
                },
                "page": {
                    "$ref": "#/definitions/result.Page"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0(bid-dinghan-cpic)",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "鼎函太保对接系统",
	Description:      "鼎函太保对接系统",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
