package main

var aws = []byte(`
{
	"format_version": "0.1",
	"provider_schemas": {
	  "aws": {
		"provider": {
		  "version": 0,
		  "block": {
			"attributes": {
			  "access_key": {
				"type": "string",
				"description": "The access key for API operations. You can retrieve this\nfrom the 'Security & Credentials' section of the AWS console.",
				"optional": true
			  },
			  "allowed_account_ids": {
				"type": [
				  "set",
				  "string"
				],
				"optional": true
			  },
			  "forbidden_account_ids": {
				"type": [
				  "set",
				  "string"
				],
				"optional": true
			  },
			  "insecure": {
				"type": "bool",
				"description": "Explicitly allow the provider to perform \"insecure\" SSL requests. If omitted,default value is \"false\"",
				"optional": true
			  },
			  "max_retries": {
				"type": "number",
				"description": "The maximum number of times an AWS API request is\nbeing executed. If the API request still fails, an error is\nthrown.",
				"optional": true
			  },
			  "profile": {
				"type": "string",
				"description": "The profile for API operations. If not set, the default profile\ncreated with \"aws configure\" will be used.",
				"optional": true
			  },
			  "region": {
				"type": "string",
				"description": "The region where AWS operations will take place. Examples\nare us-east-1, us-west-2, etc.",
				"required": true
			  },
			  "s3_force_path_style": {
				"type": "bool",
				"description": "Set this to true to force the request to use path-style addressing,\ni.e., http://s3.amazonaws.com/BUCKET/KEY. By default, the S3 client will\nuse virtual hosted bucket addressing when possible\n(http://BUCKET.s3.amazonaws.com/KEY). Specific to the Amazon S3 service.",
				"optional": true
			  },
			  "secret_key": {
				"type": "string",
				"description": "The secret key for API operations. You can retrieve this\nfrom the 'Security & Credentials' section of the AWS console.",
				"optional": true
			  },
			  "shared_credentials_file": {
				"type": "string",
				"description": "The path to the shared credentials file. If not set\nthis defaults to ~/.aws/credentials.",
				"optional": true
			  },
			  "skip_credentials_validation": {
				"type": "bool",
				"description": "Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.",
				"optional": true
			  },
			  "skip_get_ec2_platforms": {
				"type": "bool",
				"description": "Skip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes permissions.",
				"optional": true
			  },
			  "skip_metadata_api_check": {
				"type": "bool",
				"optional": true
			  },
			  "skip_region_validation": {
				"type": "bool",
				"description": "Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to regions that are not public (yet).",
				"optional": true
			  },
			  "skip_requesting_account_id": {
				"type": "bool",
				"description": "Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.",
				"optional": true
			  },
			  "token": {
				"type": "string",
				"description": "session token. A session token is only required if you are\nusing temporary security credentials.",
				"optional": true
			  }
			},
			"block_types": {
			  "assume_role": {
				"nesting_mode": "list",
				"block": {
				  "attributes": {
					"duration_seconds": {
					  "type": "number",
					  "description": "Seconds to restrict the assume role session duration.",
					  "optional": true
					},
					"external_id": {
					  "type": "string",
					  "description": "Unique identifier that might be required for assuming a role in another account.",
					  "optional": true
					},
					"policy": {
					  "type": "string",
					  "description": "IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.",
					  "optional": true
					},
					"policy_arns": {
					  "type": [
						"set",
						"string"
					  ],
					  "description": "Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed.",
					  "optional": true
					},
					"role_arn": {
					  "type": "string",
					  "description": "Amazon Resource Name of an IAM Role to assume prior to making API calls.",
					  "optional": true
					},
					"session_name": {
					  "type": "string",
					  "description": "Identifier for the assumed role session.",
					  "optional": true
					},
					"tags": {
					  "type": [
						"map",
						"string"
					  ],
					  "description": "Assume role session tags.",
					  "optional": true
					},
					"transitive_tag_keys": {
					  "type": [
						"set",
						"string"
					  ],
					  "description": "Assume role session tag keys to pass to any subsequent sessions.",
					  "optional": true
					}
				  }
				},
				"max_items": 1
			  },
			  "endpoints": {
				"nesting_mode": "set",
				"block": {
				  "attributes": {
					"accessanalyzer": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"acm": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"acmpca": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"amplify": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"apigateway": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"applicationautoscaling": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"applicationinsights": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"appmesh": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"appstream": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"appsync": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"athena": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"autoscaling": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"autoscalingplans": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"backup": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"batch": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"budgets": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"cloud9": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"cloudformation": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"cloudfront": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"cloudhsm": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"cloudsearch": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"cloudtrail": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"cloudwatch": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"cloudwatchevents": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"cloudwatchlogs": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"codeartifact": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"codebuild": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"codecommit": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"codedeploy": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"codepipeline": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"cognitoidentity": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"cognitoidp": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"configservice": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"cur": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"dataexchange": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"datapipeline": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"datasync": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"dax": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"devicefarm": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"directconnect": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"dlm": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"dms": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"docdb": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"ds": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"dynamodb": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"ec2": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"ecr": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"ecs": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"efs": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"eks": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"elasticache": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"elasticbeanstalk": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"elastictranscoder": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"elb": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"emr": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"es": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"firehose": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"fms": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"forecast": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"fsx": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"gamelift": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"glacier": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"globalaccelerator": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"glue": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"greengrass": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"guardduty": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"iam": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"imagebuilder": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"inspector": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"iot": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"iotanalytics": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"iotevents": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"kafka": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"kinesis": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"kinesisanalytics": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"kinesisanalyticsv2": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"kinesisvideo": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"kms": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"lakeformation": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"lambda": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"lexmodels": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"licensemanager": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"lightsail": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"macie": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"managedblockchain": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"marketplacecatalog": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"mediaconnect": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"mediaconvert": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"medialive": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"mediapackage": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"mediastore": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"mediastoredata": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"mq": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"neptune": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"networkmanager": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"opsworks": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"organizations": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"outposts": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"personalize": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"pinpoint": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"pricing": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"qldb": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"quicksight": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"ram": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"rds": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"redshift": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"resourcegroups": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"resourcegroupstaggingapi": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"route53": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"route53domains": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"route53resolver": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"s3": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"s3control": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"sagemaker": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"sdb": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"secretsmanager": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"securityhub": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"serverlessrepo": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"servicecatalog": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"servicediscovery": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"servicequotas": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"ses": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"shield": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"sns": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"sqs": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"ssm": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"stepfunctions": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"storagegateway": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"sts": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"swf": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"synthetics": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"transfer": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"waf": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"wafregional": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"wafv2": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"worklink": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"workmail": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"workspaces": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					},
					"xray": {
					  "type": "string",
					  "description": "Use this to override the default service endpoint URL",
					  "optional": true
					}
				  }
				}
			  },
			  "ignore_tags": {
				"nesting_mode": "list",
				"block": {
				  "attributes": {
					"key_prefixes": {
					  "type": [
						"set",
						"string"
					  ],
					  "description": "Resource tag key prefixes to ignore across all resources.",
					  "optional": true
					},
					"keys": {
					  "type": [
						"set",
						"string"
					  ],
					  "description": "Resource tag keys to ignore across all resources.",
					  "optional": true
					}
				  }
				},
				"max_items": 1
			  }
			}
		  }
		},
		"resource_schemas": {
		  "aws_accessanalyzer_analyzer": {
			"version": 0,
			"block": {
			  "attributes": {
				"analyzer_name": {
				  "type": "string",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"type": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_acm_certificate": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"certificate_authority_arn": {
				  "type": "string",
				  "optional": true
				},
				"certificate_body": {
				  "type": "string",
				  "optional": true
				},
				"certificate_chain": {
				  "type": "string",
				  "optional": true
				},
				"domain_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"domain_validation_options": {
				  "type": [
					"set",
					[
					  "object",
					  {
						"domain_name": "string",
						"resource_record_name": "string",
						"resource_record_type": "string",
						"resource_record_value": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"private_key": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"subject_alternative_names": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"validation_emails": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"validation_method": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "certificate_transparency_logging_preference": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_acm_certificate_validation": {
			"version": 0,
			"block": {
			  "attributes": {
				"certificate_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"validation_record_fqdns": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_acmpca_certificate_authority": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"certificate": {
				  "type": "string",
				  "computed": true
				},
				"certificate_chain": {
				  "type": "string",
				  "computed": true
				},
				"certificate_signing_request": {
				  "type": "string",
				  "computed": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"not_after": {
				  "type": "string",
				  "computed": true
				},
				"not_before": {
				  "type": "string",
				  "computed": true
				},
				"permanent_deletion_time_in_days": {
				  "type": "number",
				  "optional": true
				},
				"serial": {
				  "type": "string",
				  "computed": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"type": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"certificate_authority_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "key_algorithm": {
						"type": "string",
						"required": true
					  },
					  "signing_algorithm": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "subject": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"common_name": {
							  "type": "string",
							  "optional": true
							},
							"country": {
							  "type": "string",
							  "optional": true
							},
							"distinguished_name_qualifier": {
							  "type": "string",
							  "optional": true
							},
							"generation_qualifier": {
							  "type": "string",
							  "optional": true
							},
							"given_name": {
							  "type": "string",
							  "optional": true
							},
							"initials": {
							  "type": "string",
							  "optional": true
							},
							"locality": {
							  "type": "string",
							  "optional": true
							},
							"organization": {
							  "type": "string",
							  "optional": true
							},
							"organizational_unit": {
							  "type": "string",
							  "optional": true
							},
							"pseudonym": {
							  "type": "string",
							  "optional": true
							},
							"state": {
							  "type": "string",
							  "optional": true
							},
							"surname": {
							  "type": "string",
							  "optional": true
							},
							"title": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"revocation_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "crl_configuration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"custom_cname": {
							  "type": "string",
							  "optional": true
							},
							"enabled": {
							  "type": "bool",
							  "optional": true
							},
							"expiration_in_days": {
							  "type": "number",
							  "required": true
							},
							"s3_bucket_name": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_alb": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"arn_suffix": {
				  "type": "string",
				  "computed": true
				},
				"dns_name": {
				  "type": "string",
				  "computed": true
				},
				"drop_invalid_header_fields": {
				  "type": "bool",
				  "optional": true
				},
				"enable_cross_zone_load_balancing": {
				  "type": "bool",
				  "optional": true
				},
				"enable_deletion_protection": {
				  "type": "bool",
				  "optional": true
				},
				"enable_http2": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"idle_timeout": {
				  "type": "number",
				  "optional": true
				},
				"internal": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"ip_address_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"load_balancer_type": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"security_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"subnets": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				},
				"zone_id": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"access_logs": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket": {
						"type": "string",
						"required": true
					  },
					  "enabled": {
						"type": "bool",
						"optional": true
					  },
					  "prefix": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"subnet_mapping": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "allocation_id": {
						"type": "string",
						"optional": true
					  },
					  "private_ipv4_address": {
						"type": "string",
						"optional": true
					  },
					  "subnet_id": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_alb_listener": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"certificate_arn": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"load_balancer_arn": {
				  "type": "string",
				  "required": true
				},
				"port": {
				  "type": "number",
				  "required": true
				},
				"protocol": {
				  "type": "string",
				  "optional": true
				},
				"ssl_policy": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"default_action": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "order": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "target_group_arn": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "authenticate_cognito": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"authentication_request_extra_params": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"on_unauthenticated_request": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"scope": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_cookie_name": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_timeout": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"user_pool_arn": {
							  "type": "string",
							  "required": true
							},
							"user_pool_client_id": {
							  "type": "string",
							  "required": true
							},
							"user_pool_domain": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "authenticate_oidc": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"authentication_request_extra_params": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"authorization_endpoint": {
							  "type": "string",
							  "required": true
							},
							"client_id": {
							  "type": "string",
							  "required": true
							},
							"client_secret": {
							  "type": "string",
							  "required": true,
							  "sensitive": true
							},
							"issuer": {
							  "type": "string",
							  "required": true
							},
							"on_unauthenticated_request": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"scope": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_cookie_name": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_timeout": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"token_endpoint": {
							  "type": "string",
							  "required": true
							},
							"user_info_endpoint": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "fixed_response": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"content_type": {
							  "type": "string",
							  "required": true
							},
							"message_body": {
							  "type": "string",
							  "optional": true
							},
							"status_code": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							}
						  }
						},
						"max_items": 1
					  },
					  "forward": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"stickiness": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "duration": {
									"type": "number",
									"required": true
								  },
								  "enabled": {
									"type": "bool",
									"optional": true
								  }
								}
							  },
							  "max_items": 1
							},
							"target_group": {
							  "nesting_mode": "set",
							  "block": {
								"attributes": {
								  "arn": {
									"type": "string",
									"required": true
								  },
								  "weight": {
									"type": "number",
									"optional": true
								  }
								}
							  },
							  "min_items": 2,
							  "max_items": 5
							}
						  }
						},
						"max_items": 1
					  },
					  "redirect": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"host": {
							  "type": "string",
							  "optional": true
							},
							"path": {
							  "type": "string",
							  "optional": true
							},
							"port": {
							  "type": "string",
							  "optional": true
							},
							"protocol": {
							  "type": "string",
							  "optional": true
							},
							"query": {
							  "type": "string",
							  "optional": true
							},
							"status_code": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "read": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_alb_listener_certificate": {
			"version": 0,
			"block": {
			  "attributes": {
				"certificate_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"listener_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_alb_listener_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"listener_arn": {
				  "type": "string",
				  "required": true
				},
				"priority": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"action": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "order": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "target_group_arn": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "authenticate_cognito": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"authentication_request_extra_params": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"on_unauthenticated_request": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"scope": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_cookie_name": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_timeout": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"user_pool_arn": {
							  "type": "string",
							  "required": true
							},
							"user_pool_client_id": {
							  "type": "string",
							  "required": true
							},
							"user_pool_domain": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "authenticate_oidc": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"authentication_request_extra_params": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"authorization_endpoint": {
							  "type": "string",
							  "required": true
							},
							"client_id": {
							  "type": "string",
							  "required": true
							},
							"client_secret": {
							  "type": "string",
							  "required": true,
							  "sensitive": true
							},
							"issuer": {
							  "type": "string",
							  "required": true
							},
							"on_unauthenticated_request": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"scope": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_cookie_name": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_timeout": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"token_endpoint": {
							  "type": "string",
							  "required": true
							},
							"user_info_endpoint": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "fixed_response": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"content_type": {
							  "type": "string",
							  "required": true
							},
							"message_body": {
							  "type": "string",
							  "optional": true
							},
							"status_code": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							}
						  }
						},
						"max_items": 1
					  },
					  "forward": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"stickiness": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "duration": {
									"type": "number",
									"required": true
								  },
								  "enabled": {
									"type": "bool",
									"optional": true
								  }
								}
							  },
							  "max_items": 1
							},
							"target_group": {
							  "nesting_mode": "set",
							  "block": {
								"attributes": {
								  "arn": {
									"type": "string",
									"required": true
								  },
								  "weight": {
									"type": "number",
									"optional": true
								  }
								}
							  },
							  "min_items": 2,
							  "max_items": 5
							}
						  }
						},
						"max_items": 1
					  },
					  "redirect": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"host": {
							  "type": "string",
							  "optional": true
							},
							"path": {
							  "type": "string",
							  "optional": true
							},
							"port": {
							  "type": "string",
							  "optional": true
							},
							"protocol": {
							  "type": "string",
							  "optional": true
							},
							"query": {
							  "type": "string",
							  "optional": true
							},
							"status_code": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1
				},
				"condition": {
				  "nesting_mode": "set",
				  "block": {
					"block_types": {
					  "host_header": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"values": {
							  "type": [
								"set",
								"string"
							  ],
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "http_header": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"http_header_name": {
							  "type": "string",
							  "required": true
							},
							"values": {
							  "type": [
								"set",
								"string"
							  ],
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "http_request_method": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"values": {
							  "type": [
								"set",
								"string"
							  ],
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "path_pattern": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"values": {
							  "type": [
								"set",
								"string"
							  ],
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "query_string": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"key": {
							  "type": "string",
							  "optional": true
							},
							"value": {
							  "type": "string",
							  "required": true
							}
						  }
						}
					  },
					  "source_ip": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"values": {
							  "type": [
								"set",
								"string"
							  ],
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1
				}
			  }
			}
		  },
		  "aws_alb_target_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"arn_suffix": {
				  "type": "string",
				  "computed": true
				},
				"deregistration_delay": {
				  "type": "number",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"lambda_multi_value_headers_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"load_balancing_algorithm_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"port": {
				  "type": "number",
				  "optional": true
				},
				"protocol": {
				  "type": "string",
				  "optional": true
				},
				"proxy_protocol_v2": {
				  "type": "bool",
				  "optional": true
				},
				"slow_start": {
				  "type": "number",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"target_type": {
				  "type": "string",
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"health_check": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"optional": true
					  },
					  "healthy_threshold": {
						"type": "number",
						"optional": true
					  },
					  "interval": {
						"type": "number",
						"optional": true
					  },
					  "matcher": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "path": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "port": {
						"type": "string",
						"optional": true
					  },
					  "protocol": {
						"type": "string",
						"optional": true
					  },
					  "timeout": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "unhealthy_threshold": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"stickiness": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "cookie_duration": {
						"type": "number",
						"optional": true
					  },
					  "enabled": {
						"type": "bool",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_alb_target_group_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"availability_zone": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "optional": true
				},
				"target_group_arn": {
				  "type": "string",
				  "required": true
				},
				"target_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ami": {
			"version": 0,
			"block": {
			  "attributes": {
				"architecture": {
				  "type": "string",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"ena_support": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"image_location": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kernel_id": {
				  "type": "string",
				  "optional": true
				},
				"manage_ebs_snapshots": {
				  "type": "bool",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"ramdisk_id": {
				  "type": "string",
				  "optional": true
				},
				"root_device_name": {
				  "type": "string",
				  "optional": true
				},
				"root_snapshot_id": {
				  "type": "string",
				  "computed": true
				},
				"sriov_net_support": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"virtualization_type": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"ebs_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "delete_on_termination": {
						"type": "bool",
						"optional": true
					  },
					  "device_name": {
						"type": "string",
						"required": true
					  },
					  "encrypted": {
						"type": "bool",
						"optional": true
					  },
					  "iops": {
						"type": "number",
						"optional": true
					  },
					  "snapshot_id": {
						"type": "string",
						"optional": true
					  },
					  "volume_size": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "volume_type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"ephemeral_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "device_name": {
						"type": "string",
						"required": true
					  },
					  "virtual_name": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_ami_copy": {
			"version": 0,
			"block": {
			  "attributes": {
				"architecture": {
				  "type": "string",
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"ena_support": {
				  "type": "bool",
				  "computed": true
				},
				"encrypted": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"image_location": {
				  "type": "string",
				  "computed": true
				},
				"kernel_id": {
				  "type": "string",
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"manage_ebs_snapshots": {
				  "type": "bool",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"ramdisk_id": {
				  "type": "string",
				  "computed": true
				},
				"root_device_name": {
				  "type": "string",
				  "computed": true
				},
				"root_snapshot_id": {
				  "type": "string",
				  "computed": true
				},
				"source_ami_id": {
				  "type": "string",
				  "required": true
				},
				"source_ami_region": {
				  "type": "string",
				  "required": true
				},
				"sriov_net_support": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"virtualization_type": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"ebs_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "delete_on_termination": {
						"type": "bool",
						"computed": true
					  },
					  "device_name": {
						"type": "string",
						"computed": true
					  },
					  "encrypted": {
						"type": "bool",
						"computed": true
					  },
					  "iops": {
						"type": "number",
						"computed": true
					  },
					  "snapshot_id": {
						"type": "string",
						"computed": true
					  },
					  "volume_size": {
						"type": "number",
						"computed": true
					  },
					  "volume_type": {
						"type": "string",
						"computed": true
					  }
					}
				  }
				},
				"ephemeral_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "device_name": {
						"type": "string",
						"computed": true
					  },
					  "virtual_name": {
						"type": "string",
						"computed": true
					  }
					}
				  }
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_ami_from_instance": {
			"version": 0,
			"block": {
			  "attributes": {
				"architecture": {
				  "type": "string",
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"ena_support": {
				  "type": "bool",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"image_location": {
				  "type": "string",
				  "computed": true
				},
				"kernel_id": {
				  "type": "string",
				  "computed": true
				},
				"manage_ebs_snapshots": {
				  "type": "bool",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"ramdisk_id": {
				  "type": "string",
				  "computed": true
				},
				"root_device_name": {
				  "type": "string",
				  "computed": true
				},
				"root_snapshot_id": {
				  "type": "string",
				  "computed": true
				},
				"snapshot_without_reboot": {
				  "type": "bool",
				  "optional": true
				},
				"source_instance_id": {
				  "type": "string",
				  "required": true
				},
				"sriov_net_support": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"virtualization_type": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"ebs_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "delete_on_termination": {
						"type": "bool",
						"computed": true
					  },
					  "device_name": {
						"type": "string",
						"computed": true
					  },
					  "encrypted": {
						"type": "bool",
						"computed": true
					  },
					  "iops": {
						"type": "number",
						"computed": true
					  },
					  "snapshot_id": {
						"type": "string",
						"computed": true
					  },
					  "volume_size": {
						"type": "number",
						"computed": true
					  },
					  "volume_type": {
						"type": "string",
						"computed": true
					  }
					}
				  }
				},
				"ephemeral_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "device_name": {
						"type": "string",
						"computed": true
					  },
					  "virtual_name": {
						"type": "string",
						"computed": true
					  }
					}
				  }
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_ami_launch_permission": {
			"version": 0,
			"block": {
			  "attributes": {
				"account_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"image_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_api_gateway_account": {
			"version": 0,
			"block": {
			  "attributes": {
				"cloudwatch_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"throttle_settings": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"burst_limit": "number",
						"rate_limit": "number"
					  }
					]
				  ],
				  "computed": true
				}
			  }
			}
		  },
		  "aws_api_gateway_api_key": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"created_date": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"last_updated_date": {
				  "type": "string",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"value": {
				  "type": "string",
				  "optional": true,
				  "computed": true,
				  "sensitive": true
				}
			  }
			}
		  },
		  "aws_api_gateway_authorizer": {
			"version": 0,
			"block": {
			  "attributes": {
				"authorizer_credentials": {
				  "type": "string",
				  "optional": true
				},
				"authorizer_result_ttl_in_seconds": {
				  "type": "number",
				  "optional": true
				},
				"authorizer_uri": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identity_source": {
				  "type": "string",
				  "optional": true
				},
				"identity_validation_expression": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"provider_arns": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				},
				"type": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_api_gateway_base_path_mapping": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"base_path": {
				  "type": "string",
				  "optional": true
				},
				"domain_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"stage_name": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_api_gateway_client_certificate": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"created_date": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"expiration_date": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"pem_encoded_certificate": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_api_gateway_deployment": {
			"version": 0,
			"block": {
			  "attributes": {
				"created_date": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"execution_arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"invoke_url": {
				  "type": "string",
				  "computed": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				},
				"stage_description": {
				  "type": "string",
				  "optional": true
				},
				"stage_name": {
				  "type": "string",
				  "optional": true
				},
				"triggers": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"variables": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_api_gateway_documentation_part": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"properties": {
				  "type": "string",
				  "required": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"location": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "method": {
						"type": "string",
						"optional": true
					  },
					  "name": {
						"type": "string",
						"optional": true
					  },
					  "path": {
						"type": "string",
						"optional": true
					  },
					  "status_code": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_api_gateway_documentation_version": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				},
				"version": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_api_gateway_domain_name": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"certificate_arn": {
				  "type": "string",
				  "optional": true
				},
				"certificate_body": {
				  "type": "string",
				  "optional": true
				},
				"certificate_chain": {
				  "type": "string",
				  "optional": true
				},
				"certificate_name": {
				  "type": "string",
				  "optional": true
				},
				"certificate_private_key": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"certificate_upload_date": {
				  "type": "string",
				  "computed": true
				},
				"cloudfront_domain_name": {
				  "type": "string",
				  "computed": true
				},
				"cloudfront_zone_id": {
				  "type": "string",
				  "computed": true
				},
				"domain_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"regional_certificate_arn": {
				  "type": "string",
				  "optional": true
				},
				"regional_certificate_name": {
				  "type": "string",
				  "optional": true
				},
				"regional_domain_name": {
				  "type": "string",
				  "computed": true
				},
				"regional_zone_id": {
				  "type": "string",
				  "computed": true
				},
				"security_policy": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"endpoint_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "types": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_api_gateway_gateway_response": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"response_parameters": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"response_templates": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"response_type": {
				  "type": "string",
				  "required": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				},
				"status_code": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_api_gateway_integration": {
			"version": 0,
			"block": {
			  "attributes": {
				"cache_key_parameters": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"cache_namespace": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"connection_id": {
				  "type": "string",
				  "optional": true
				},
				"connection_type": {
				  "type": "string",
				  "optional": true
				},
				"content_handling": {
				  "type": "string",
				  "optional": true
				},
				"credentials": {
				  "type": "string",
				  "optional": true
				},
				"http_method": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"integration_http_method": {
				  "type": "string",
				  "optional": true
				},
				"passthrough_behavior": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"request_parameters": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"request_templates": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"resource_id": {
				  "type": "string",
				  "required": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				},
				"timeout_milliseconds": {
				  "type": "number",
				  "optional": true
				},
				"type": {
				  "type": "string",
				  "required": true
				},
				"uri": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_api_gateway_integration_response": {
			"version": 0,
			"block": {
			  "attributes": {
				"content_handling": {
				  "type": "string",
				  "optional": true
				},
				"http_method": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"resource_id": {
				  "type": "string",
				  "required": true
				},
				"response_parameters": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"response_templates": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				},
				"selection_pattern": {
				  "type": "string",
				  "optional": true
				},
				"status_code": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_api_gateway_method": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_key_required": {
				  "type": "bool",
				  "optional": true
				},
				"authorization": {
				  "type": "string",
				  "required": true
				},
				"authorization_scopes": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"authorizer_id": {
				  "type": "string",
				  "optional": true
				},
				"http_method": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"request_models": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"request_parameters": {
				  "type": [
					"map",
					"bool"
				  ],
				  "optional": true
				},
				"request_validator_id": {
				  "type": "string",
				  "optional": true
				},
				"resource_id": {
				  "type": "string",
				  "required": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_api_gateway_method_response": {
			"version": 0,
			"block": {
			  "attributes": {
				"http_method": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"resource_id": {
				  "type": "string",
				  "required": true
				},
				"response_models": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"response_parameters": {
				  "type": [
					"map",
					"bool"
				  ],
				  "optional": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				},
				"status_code": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_api_gateway_method_settings": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"method_path": {
				  "type": "string",
				  "required": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				},
				"stage_name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "cache_data_encrypted": {
						"type": "bool",
						"optional": true,
						"computed": true
					  },
					  "cache_ttl_in_seconds": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "caching_enabled": {
						"type": "bool",
						"optional": true,
						"computed": true
					  },
					  "data_trace_enabled": {
						"type": "bool",
						"optional": true,
						"computed": true
					  },
					  "logging_level": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "metrics_enabled": {
						"type": "bool",
						"optional": true,
						"computed": true
					  },
					  "require_authorization_for_cache_control": {
						"type": "bool",
						"optional": true,
						"computed": true
					  },
					  "throttling_burst_limit": {
						"type": "number",
						"optional": true
					  },
					  "throttling_rate_limit": {
						"type": "number",
						"optional": true
					  },
					  "unauthorized_cache_control_header_strategy": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_api_gateway_model": {
			"version": 0,
			"block": {
			  "attributes": {
				"content_type": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				},
				"schema": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_api_gateway_request_validator": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				},
				"validate_request_body": {
				  "type": "bool",
				  "optional": true
				},
				"validate_request_parameters": {
				  "type": "bool",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_api_gateway_resource": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"parent_id": {
				  "type": "string",
				  "required": true
				},
				"path": {
				  "type": "string",
				  "computed": true
				},
				"path_part": {
				  "type": "string",
				  "required": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_api_gateway_rest_api": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_key_source": {
				  "type": "string",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"binary_media_types": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"body": {
				  "type": "string",
				  "optional": true
				},
				"created_date": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"execution_arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"minimum_compression_size": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"policy": {
				  "type": "string",
				  "optional": true
				},
				"root_resource_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"endpoint_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "types": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  },
					  "vpc_endpoint_ids": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_api_gateway_stage": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"cache_cluster_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"cache_cluster_size": {
				  "type": "string",
				  "optional": true
				},
				"client_certificate_id": {
				  "type": "string",
				  "optional": true
				},
				"deployment_id": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"documentation_version": {
				  "type": "string",
				  "optional": true
				},
				"execution_arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"invoke_url": {
				  "type": "string",
				  "computed": true
				},
				"rest_api_id": {
				  "type": "string",
				  "required": true
				},
				"stage_name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"variables": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"xray_tracing_enabled": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"access_log_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "destination_arn": {
						"type": "string",
						"required": true
					  },
					  "format": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_api_gateway_usage_plan": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"product_code": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"api_stages": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "api_id": {
						"type": "string",
						"required": true
					  },
					  "stage": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"quota_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "limit": {
						"type": "number",
						"required": true
					  },
					  "offset": {
						"type": "number",
						"optional": true
					  },
					  "period": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"throttle_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "burst_limit": {
						"type": "number",
						"optional": true
					  },
					  "rate_limit": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_api_gateway_usage_plan_key": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key_id": {
				  "type": "string",
				  "required": true
				},
				"key_type": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "computed": true
				},
				"usage_plan_id": {
				  "type": "string",
				  "required": true
				},
				"value": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_api_gateway_vpc_link": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"target_arns": {
				  "type": [
					"list",
					"string"
				  ],
				  "required": true
				}
			  }
			}
		  },
		  "aws_apigatewayv2_api": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_endpoint": {
				  "type": "string",
				  "computed": true
				},
				"api_key_selection_expression": {
				  "type": "string",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"credentials_arn": {
				  "type": "string",
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"execution_arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"protocol_type": {
				  "type": "string",
				  "required": true
				},
				"route_key": {
				  "type": "string",
				  "optional": true
				},
				"route_selection_expression": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"target": {
				  "type": "string",
				  "optional": true
				},
				"version": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"cors_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allow_credentials": {
						"type": "bool",
						"optional": true
					  },
					  "allow_headers": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "allow_methods": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "allow_origins": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "expose_headers": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "max_age": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_apigatewayv2_api_mapping": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"api_mapping_key": {
				  "type": "string",
				  "optional": true
				},
				"domain_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"stage": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_apigatewayv2_authorizer": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"authorizer_credentials_arn": {
				  "type": "string",
				  "optional": true
				},
				"authorizer_type": {
				  "type": "string",
				  "required": true
				},
				"authorizer_uri": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identity_sources": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"jwt_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "audience": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "issuer": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_apigatewayv2_deployment": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"auto_deployed": {
				  "type": "bool",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"triggers": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_apigatewayv2_domain_name": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_mapping_selection_expression": {
				  "type": "string",
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"domain_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"domain_name_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "certificate_arn": {
						"type": "string",
						"required": true
					  },
					  "endpoint_type": {
						"type": "string",
						"required": true
					  },
					  "hosted_zone_id": {
						"type": "string",
						"computed": true
					  },
					  "security_policy": {
						"type": "string",
						"required": true
					  },
					  "target_domain_name": {
						"type": "string",
						"computed": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_apigatewayv2_integration": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"connection_id": {
				  "type": "string",
				  "optional": true
				},
				"connection_type": {
				  "type": "string",
				  "optional": true
				},
				"content_handling_strategy": {
				  "type": "string",
				  "optional": true
				},
				"credentials_arn": {
				  "type": "string",
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"integration_method": {
				  "type": "string",
				  "optional": true
				},
				"integration_response_selection_expression": {
				  "type": "string",
				  "computed": true
				},
				"integration_type": {
				  "type": "string",
				  "required": true
				},
				"integration_uri": {
				  "type": "string",
				  "optional": true
				},
				"passthrough_behavior": {
				  "type": "string",
				  "optional": true
				},
				"payload_format_version": {
				  "type": "string",
				  "optional": true
				},
				"request_parameters": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"request_templates": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"template_selection_expression": {
				  "type": "string",
				  "optional": true
				},
				"timeout_milliseconds": {
				  "type": "number",
				  "optional": true
				}
			  },
			  "block_types": {
				"tls_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "server_name_to_verify": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_apigatewayv2_integration_response": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"content_handling_strategy": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"integration_id": {
				  "type": "string",
				  "required": true
				},
				"integration_response_key": {
				  "type": "string",
				  "required": true
				},
				"response_templates": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"template_selection_expression": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_apigatewayv2_model": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"content_type": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"schema": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_apigatewayv2_route": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"api_key_required": {
				  "type": "bool",
				  "optional": true
				},
				"authorization_scopes": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"authorization_type": {
				  "type": "string",
				  "optional": true
				},
				"authorizer_id": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"model_selection_expression": {
				  "type": "string",
				  "optional": true
				},
				"operation_name": {
				  "type": "string",
				  "optional": true
				},
				"request_models": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"route_key": {
				  "type": "string",
				  "required": true
				},
				"route_response_selection_expression": {
				  "type": "string",
				  "optional": true
				},
				"target": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_apigatewayv2_route_response": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"model_selection_expression": {
				  "type": "string",
				  "optional": true
				},
				"response_models": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"route_id": {
				  "type": "string",
				  "required": true
				},
				"route_response_key": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_apigatewayv2_stage": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_deploy": {
				  "type": "bool",
				  "optional": true
				},
				"client_certificate_id": {
				  "type": "string",
				  "optional": true
				},
				"deployment_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"execution_arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"invoke_url": {
				  "type": "string",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"stage_variables": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"access_log_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "destination_arn": {
						"type": "string",
						"required": true
					  },
					  "format": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"default_route_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "data_trace_enabled": {
						"type": "bool",
						"optional": true
					  },
					  "detailed_metrics_enabled": {
						"type": "bool",
						"optional": true
					  },
					  "logging_level": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "throttling_burst_limit": {
						"type": "number",
						"optional": true
					  },
					  "throttling_rate_limit": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"route_settings": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "data_trace_enabled": {
						"type": "bool",
						"optional": true
					  },
					  "detailed_metrics_enabled": {
						"type": "bool",
						"optional": true
					  },
					  "logging_level": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "route_key": {
						"type": "string",
						"required": true
					  },
					  "throttling_burst_limit": {
						"type": "number",
						"optional": true
					  },
					  "throttling_rate_limit": {
						"type": "number",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_apigatewayv2_vpc_link": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_app_cookie_stickiness_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"cookie_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"lb_port": {
				  "type": "number",
				  "required": true
				},
				"load_balancer": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_appautoscaling_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"policy_type": {
				  "type": "string",
				  "optional": true
				},
				"resource_id": {
				  "type": "string",
				  "required": true
				},
				"scalable_dimension": {
				  "type": "string",
				  "required": true
				},
				"service_namespace": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"step_scaling_policy_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "adjustment_type": {
						"type": "string",
						"optional": true
					  },
					  "cooldown": {
						"type": "number",
						"optional": true
					  },
					  "metric_aggregation_type": {
						"type": "string",
						"optional": true
					  },
					  "min_adjustment_magnitude": {
						"type": "number",
						"optional": true
					  }
					},
					"block_types": {
					  "step_adjustment": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"metric_interval_lower_bound": {
							  "type": "string",
							  "optional": true
							},
							"metric_interval_upper_bound": {
							  "type": "string",
							  "optional": true
							},
							"scaling_adjustment": {
							  "type": "number",
							  "required": true
							}
						  }
						}
					  }
					}
				  },
				  "max_items": 1
				},
				"target_tracking_scaling_policy_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "disable_scale_in": {
						"type": "bool",
						"optional": true
					  },
					  "scale_in_cooldown": {
						"type": "number",
						"optional": true
					  },
					  "scale_out_cooldown": {
						"type": "number",
						"optional": true
					  },
					  "target_value": {
						"type": "number",
						"required": true
					  }
					},
					"block_types": {
					  "customized_metric_specification": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"metric_name": {
							  "type": "string",
							  "required": true
							},
							"namespace": {
							  "type": "string",
							  "required": true
							},
							"statistic": {
							  "type": "string",
							  "required": true
							},
							"unit": {
							  "type": "string",
							  "optional": true
							}
						  },
						  "block_types": {
							"dimensions": {
							  "nesting_mode": "set",
							  "block": {
								"attributes": {
								  "name": {
									"type": "string",
									"required": true
								  },
								  "value": {
									"type": "string",
									"required": true
								  }
								}
							  }
							}
						  }
						},
						"max_items": 1
					  },
					  "predefined_metric_specification": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"predefined_metric_type": {
							  "type": "string",
							  "required": true
							},
							"resource_label": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_appautoscaling_scheduled_action": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"end_time": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"resource_id": {
				  "type": "string",
				  "required": true
				},
				"scalable_dimension": {
				  "type": "string",
				  "optional": true
				},
				"schedule": {
				  "type": "string",
				  "optional": true
				},
				"service_namespace": {
				  "type": "string",
				  "required": true
				},
				"start_time": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"scalable_target_action": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "max_capacity": {
						"type": "number",
						"optional": true
					  },
					  "min_capacity": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_appautoscaling_target": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"max_capacity": {
				  "type": "number",
				  "required": true
				},
				"min_capacity": {
				  "type": "number",
				  "required": true
				},
				"resource_id": {
				  "type": "string",
				  "required": true
				},
				"role_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"scalable_dimension": {
				  "type": "string",
				  "required": true
				},
				"service_namespace": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_appmesh_mesh": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"created_date": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"last_updated_date": {
				  "type": "string",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"spec": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "egress_filter": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"type": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_appmesh_route": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"created_date": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"last_updated_date": {
				  "type": "string",
				  "computed": true
				},
				"mesh_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"virtual_router_name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"spec": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "priority": {
						"type": "number",
						"optional": true
					  }
					},
					"block_types": {
					  "http_route": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"action": {
							  "nesting_mode": "list",
							  "block": {
								"block_types": {
								  "weighted_target": {
									"nesting_mode": "set",
									"block": {
									  "attributes": {
										"virtual_node": {
										  "type": "string",
										  "required": true
										},
										"weight": {
										  "type": "number",
										  "required": true
										}
									  }
									},
									"min_items": 1,
									"max_items": 10
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							},
							"match": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "method": {
									"type": "string",
									"optional": true
								  },
								  "prefix": {
									"type": "string",
									"required": true
								  },
								  "scheme": {
									"type": "string",
									"optional": true
								  }
								},
								"block_types": {
								  "header": {
									"nesting_mode": "set",
									"block": {
									  "attributes": {
										"invert": {
										  "type": "bool",
										  "optional": true
										},
										"name": {
										  "type": "string",
										  "required": true
										}
									  },
									  "block_types": {
										"match": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "exact": {
												"type": "string",
												"optional": true
											  },
											  "prefix": {
												"type": "string",
												"optional": true
											  },
											  "regex": {
												"type": "string",
												"optional": true
											  },
											  "suffix": {
												"type": "string",
												"optional": true
											  }
											},
											"block_types": {
											  "range": {
												"nesting_mode": "list",
												"block": {
												  "attributes": {
													"end": {
													  "type": "number",
													  "required": true
													},
													"start": {
													  "type": "number",
													  "required": true
													}
												  }
												},
												"max_items": 1
											  }
											}
										  },
										  "max_items": 1
										}
									  }
									},
									"max_items": 10
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  },
					  "tcp_route": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"action": {
							  "nesting_mode": "list",
							  "block": {
								"block_types": {
								  "weighted_target": {
									"nesting_mode": "set",
									"block": {
									  "attributes": {
										"virtual_node": {
										  "type": "string",
										  "required": true
										},
										"weight": {
										  "type": "number",
										  "required": true
										}
									  }
									},
									"min_items": 1,
									"max_items": 10
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_appmesh_virtual_node": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"created_date": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"last_updated_date": {
				  "type": "string",
				  "computed": true
				},
				"mesh_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"spec": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "backend": {
						"nesting_mode": "set",
						"block": {
						  "block_types": {
							"virtual_service": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "virtual_service_name": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"max_items": 25
					  },
					  "listener": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"health_check": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "healthy_threshold": {
									"type": "number",
									"required": true
								  },
								  "interval_millis": {
									"type": "number",
									"required": true
								  },
								  "path": {
									"type": "string",
									"optional": true
								  },
								  "port": {
									"type": "number",
									"optional": true,
									"computed": true
								  },
								  "protocol": {
									"type": "string",
									"required": true
								  },
								  "timeout_millis": {
									"type": "number",
									"required": true
								  },
								  "unhealthy_threshold": {
									"type": "number",
									"required": true
								  }
								}
							  },
							  "max_items": 1
							},
							"port_mapping": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "port": {
									"type": "number",
									"required": true
								  },
								  "protocol": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  },
					  "logging": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"access_log": {
							  "nesting_mode": "list",
							  "block": {
								"block_types": {
								  "file": {
									"nesting_mode": "list",
									"block": {
									  "attributes": {
										"path": {
										  "type": "string",
										  "required": true
										}
									  }
									},
									"max_items": 1
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  },
					  "service_discovery": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"aws_cloud_map": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "attributes": {
									"type": [
									  "map",
									  "string"
									],
									"optional": true
								  },
								  "namespace_name": {
									"type": "string",
									"required": true
								  },
								  "service_name": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "max_items": 1
							},
							"dns": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "hostname": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_appmesh_virtual_router": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"created_date": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"last_updated_date": {
				  "type": "string",
				  "computed": true
				},
				"mesh_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"spec": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "listener": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"port_mapping": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "port": {
									"type": "number",
									"required": true
								  },
								  "protocol": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_appmesh_virtual_service": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"created_date": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"last_updated_date": {
				  "type": "string",
				  "computed": true
				},
				"mesh_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"spec": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "provider": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"virtual_node": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "virtual_node_name": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "max_items": 1
							},
							"virtual_router": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "virtual_router_name": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_appsync_api_key": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"expires": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key": {
				  "type": "string",
				  "computed": true,
				  "sensitive": true
				}
			  }
			}
		  },
		  "aws_appsync_datasource": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"service_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"type": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"dynamodb_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "region": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "table_name": {
						"type": "string",
						"required": true
					  },
					  "use_caller_credentials": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"elasticsearch_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "endpoint": {
						"type": "string",
						"required": true
					  },
					  "region": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				},
				"http_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "endpoint": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"lambda_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "function_arn": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_appsync_function": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"data_source": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"function_id": {
				  "type": "string",
				  "computed": true
				},
				"function_version": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"request_mapping_template": {
				  "type": "string",
				  "required": true
				},
				"response_mapping_template": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_appsync_graphql_api": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"authentication_type": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"schema": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"uris": {
				  "type": [
					"map",
					"string"
				  ],
				  "computed": true
				},
				"xray_enabled": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"additional_authentication_provider": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "authentication_type": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "openid_connect_config": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"auth_ttl": {
							  "type": "number",
							  "optional": true
							},
							"client_id": {
							  "type": "string",
							  "optional": true
							},
							"iat_ttl": {
							  "type": "number",
							  "optional": true
							},
							"issuer": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "user_pool_config": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"app_id_client_regex": {
							  "type": "string",
							  "optional": true
							},
							"aws_region": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"user_pool_id": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  }
				},
				"log_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "cloudwatch_logs_role_arn": {
						"type": "string",
						"required": true
					  },
					  "exclude_verbose_content": {
						"type": "bool",
						"optional": true
					  },
					  "field_log_level": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"openid_connect_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "auth_ttl": {
						"type": "number",
						"optional": true
					  },
					  "client_id": {
						"type": "string",
						"optional": true
					  },
					  "iat_ttl": {
						"type": "number",
						"optional": true
					  },
					  "issuer": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"user_pool_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "app_id_client_regex": {
						"type": "string",
						"optional": true
					  },
					  "aws_region": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "default_action": {
						"type": "string",
						"required": true
					  },
					  "user_pool_id": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_appsync_resolver": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_id": {
				  "type": "string",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"data_source": {
				  "type": "string",
				  "optional": true
				},
				"field": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kind": {
				  "type": "string",
				  "optional": true
				},
				"request_template": {
				  "type": "string",
				  "required": true
				},
				"response_template": {
				  "type": "string",
				  "required": true
				},
				"type": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"caching_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "caching_keys": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "ttl": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"pipeline_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "functions": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_athena_database": {
			"version": 0,
			"block": {
			  "attributes": {
				"bucket": {
				  "type": "string",
				  "required": true
				},
				"force_destroy": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"encryption_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "encryption_option": {
						"type": "string",
						"required": true
					  },
					  "kms_key": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_athena_named_query": {
			"version": 0,
			"block": {
			  "attributes": {
				"database": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"query": {
				  "type": "string",
				  "required": true
				},
				"workgroup": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_athena_workgroup": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"force_destroy": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"state": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bytes_scanned_cutoff_per_query": {
						"type": "number",
						"optional": true
					  },
					  "enforce_workgroup_configuration": {
						"type": "bool",
						"optional": true
					  },
					  "publish_cloudwatch_metrics_enabled": {
						"type": "bool",
						"optional": true
					  }
					},
					"block_types": {
					  "result_configuration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"output_location": {
							  "type": "string",
							  "optional": true
							}
						  },
						  "block_types": {
							"encryption_configuration": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "encryption_option": {
									"type": "string",
									"optional": true
								  },
								  "kms_key_arn": {
									"type": "string",
									"optional": true
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_autoscaling_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"alb_target_group_arn": {
				  "type": "string",
				  "optional": true
				},
				"autoscaling_group_name": {
				  "type": "string",
				  "required": true
				},
				"elb": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_autoscaling_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"availability_zones": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"default_cooldown": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"desired_capacity": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"enabled_metrics": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"force_delete": {
				  "type": "bool",
				  "optional": true
				},
				"health_check_grace_period": {
				  "type": "number",
				  "optional": true
				},
				"health_check_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"launch_configuration": {
				  "type": "string",
				  "optional": true
				},
				"load_balancers": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"max_instance_lifetime": {
				  "type": "number",
				  "optional": true
				},
				"max_size": {
				  "type": "number",
				  "required": true
				},
				"metrics_granularity": {
				  "type": "string",
				  "optional": true
				},
				"min_elb_capacity": {
				  "type": "number",
				  "optional": true
				},
				"min_size": {
				  "type": "number",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"placement_group": {
				  "type": "string",
				  "optional": true
				},
				"protect_from_scale_in": {
				  "type": "bool",
				  "optional": true
				},
				"service_linked_role_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"suspended_processes": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"set",
					[
					  "map",
					  "string"
					]
				  ],
				  "optional": true
				},
				"target_group_arns": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"termination_policies": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"vpc_zone_identifier": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"wait_for_capacity_timeout": {
				  "type": "string",
				  "optional": true
				},
				"wait_for_elb_capacity": {
				  "type": "number",
				  "optional": true
				}
			  },
			  "block_types": {
				"initial_lifecycle_hook": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "default_result": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "heartbeat_timeout": {
						"type": "number",
						"optional": true
					  },
					  "lifecycle_transition": {
						"type": "string",
						"required": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  },
					  "notification_metadata": {
						"type": "string",
						"optional": true
					  },
					  "notification_target_arn": {
						"type": "string",
						"optional": true
					  },
					  "role_arn": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"launch_template": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "name": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "version": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"mixed_instances_policy": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "instances_distribution": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"on_demand_allocation_strategy": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"on_demand_base_capacity": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"on_demand_percentage_above_base_capacity": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"spot_allocation_strategy": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"spot_instance_pools": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"spot_max_price": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "launch_template": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"launch_template_specification": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "launch_template_id": {
									"type": "string",
									"optional": true,
									"computed": true
								  },
								  "launch_template_name": {
									"type": "string",
									"optional": true,
									"computed": true
								  },
								  "version": {
									"type": "string",
									"optional": true
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							},
							"override": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "instance_type": {
									"type": "string",
									"optional": true
								  },
								  "weighted_capacity": {
									"type": "string",
									"optional": true
								  }
								}
							  }
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"tag": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "key": {
						"type": "string",
						"required": true
					  },
					  "propagate_at_launch": {
						"type": "bool",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_autoscaling_lifecycle_hook": {
			"version": 0,
			"block": {
			  "attributes": {
				"autoscaling_group_name": {
				  "type": "string",
				  "required": true
				},
				"default_result": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"heartbeat_timeout": {
				  "type": "number",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"lifecycle_transition": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"notification_metadata": {
				  "type": "string",
				  "optional": true
				},
				"notification_target_arn": {
				  "type": "string",
				  "optional": true
				},
				"role_arn": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_autoscaling_notification": {
			"version": 0,
			"block": {
			  "attributes": {
				"group_names": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"notifications": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"topic_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_autoscaling_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"adjustment_type": {
				  "type": "string",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"autoscaling_group_name": {
				  "type": "string",
				  "required": true
				},
				"cooldown": {
				  "type": "number",
				  "optional": true
				},
				"estimated_instance_warmup": {
				  "type": "number",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"metric_aggregation_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"min_adjustment_magnitude": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"policy_type": {
				  "type": "string",
				  "optional": true
				},
				"scaling_adjustment": {
				  "type": "number",
				  "optional": true
				}
			  },
			  "block_types": {
				"step_adjustment": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "metric_interval_lower_bound": {
						"type": "string",
						"optional": true
					  },
					  "metric_interval_upper_bound": {
						"type": "string",
						"optional": true
					  },
					  "scaling_adjustment": {
						"type": "number",
						"required": true
					  }
					}
				  }
				},
				"target_tracking_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "disable_scale_in": {
						"type": "bool",
						"optional": true
					  },
					  "target_value": {
						"type": "number",
						"required": true
					  }
					},
					"block_types": {
					  "customized_metric_specification": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"metric_name": {
							  "type": "string",
							  "required": true
							},
							"namespace": {
							  "type": "string",
							  "required": true
							},
							"statistic": {
							  "type": "string",
							  "required": true
							},
							"unit": {
							  "type": "string",
							  "optional": true
							}
						  },
						  "block_types": {
							"metric_dimension": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "name": {
									"type": "string",
									"required": true
								  },
								  "value": {
									"type": "string",
									"required": true
								  }
								}
							  }
							}
						  }
						},
						"max_items": 1
					  },
					  "predefined_metric_specification": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"predefined_metric_type": {
							  "type": "string",
							  "required": true
							},
							"resource_label": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_autoscaling_schedule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"autoscaling_group_name": {
				  "type": "string",
				  "required": true
				},
				"desired_capacity": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"end_time": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"max_size": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"min_size": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"recurrence": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"scheduled_action_name": {
				  "type": "string",
				  "required": true
				},
				"start_time": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_backup_plan": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"version": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"rule": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "completion_window": {
						"type": "number",
						"optional": true
					  },
					  "recovery_point_tags": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  },
					  "rule_name": {
						"type": "string",
						"required": true
					  },
					  "schedule": {
						"type": "string",
						"optional": true
					  },
					  "start_window": {
						"type": "number",
						"optional": true
					  },
					  "target_vault_name": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "copy_action": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"destination_vault_arn": {
							  "type": "string",
							  "required": true
							}
						  },
						  "block_types": {
							"lifecycle": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "cold_storage_after": {
									"type": "number",
									"optional": true
								  },
								  "delete_after": {
									"type": "number",
									"optional": true
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						}
					  },
					  "lifecycle": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"cold_storage_after": {
							  "type": "number",
							  "optional": true
							},
							"delete_after": {
							  "type": "number",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1
				}
			  }
			}
		  },
		  "aws_backup_selection": {
			"version": 0,
			"block": {
			  "attributes": {
				"iam_role_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"plan_id": {
				  "type": "string",
				  "required": true
				},
				"resources": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"selection_tag": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "key": {
						"type": "string",
						"required": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_backup_vault": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"recovery_points": {
				  "type": "number",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_batch_compute_environment": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"compute_environment_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"compute_environment_name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"ecs_cluster_arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"service_role": {
				  "type": "string",
				  "required": true
				},
				"state": {
				  "type": "string",
				  "optional": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"status_reason": {
				  "type": "string",
				  "computed": true
				},
				"type": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"compute_resources": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allocation_strategy": {
						"type": "string",
						"optional": true
					  },
					  "bid_percentage": {
						"type": "number",
						"optional": true
					  },
					  "desired_vcpus": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "ec2_key_pair": {
						"type": "string",
						"optional": true
					  },
					  "image_id": {
						"type": "string",
						"optional": true
					  },
					  "instance_role": {
						"type": "string",
						"required": true
					  },
					  "instance_type": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "max_vcpus": {
						"type": "number",
						"required": true
					  },
					  "min_vcpus": {
						"type": "number",
						"required": true
					  },
					  "security_group_ids": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "spot_iam_fleet_role": {
						"type": "string",
						"optional": true
					  },
					  "subnets": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "tags": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "launch_template": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"launch_template_id": {
							  "type": "string",
							  "optional": true
							},
							"launch_template_name": {
							  "type": "string",
							  "optional": true
							},
							"version": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_batch_job_definition": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"container_properties": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"parameters": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"revision": {
				  "type": "number",
				  "computed": true
				},
				"type": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"retry_strategy": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "attempts": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeout": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "attempt_duration_seconds": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_batch_job_queue": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"compute_environments": {
				  "type": [
					"list",
					"string"
				  ],
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"priority": {
				  "type": "number",
				  "required": true
				},
				"state": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_budgets_budget": {
			"version": 0,
			"block": {
			  "attributes": {
				"account_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"budget_type": {
				  "type": "string",
				  "required": true
				},
				"cost_filters": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"limit_amount": {
				  "type": "string",
				  "required": true
				},
				"limit_unit": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"time_period_end": {
				  "type": "string",
				  "optional": true
				},
				"time_period_start": {
				  "type": "string",
				  "required": true
				},
				"time_unit": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"cost_types": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "include_credit": {
						"type": "bool",
						"optional": true
					  },
					  "include_discount": {
						"type": "bool",
						"optional": true
					  },
					  "include_other_subscription": {
						"type": "bool",
						"optional": true
					  },
					  "include_recurring": {
						"type": "bool",
						"optional": true
					  },
					  "include_refund": {
						"type": "bool",
						"optional": true
					  },
					  "include_subscription": {
						"type": "bool",
						"optional": true
					  },
					  "include_support": {
						"type": "bool",
						"optional": true
					  },
					  "include_tax": {
						"type": "bool",
						"optional": true
					  },
					  "include_upfront": {
						"type": "bool",
						"optional": true
					  },
					  "use_amortized": {
						"type": "bool",
						"optional": true
					  },
					  "use_blended": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"notification": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "comparison_operator": {
						"type": "string",
						"required": true
					  },
					  "notification_type": {
						"type": "string",
						"required": true
					  },
					  "subscriber_email_addresses": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "subscriber_sns_topic_arns": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "threshold": {
						"type": "number",
						"required": true
					  },
					  "threshold_type": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_cloud9_environment_ec2": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"automatic_stop_time_minutes": {
				  "type": "number",
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_type": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"owner_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"subnet_id": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"type": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_cloudformation_stack": {
			"version": 0,
			"block": {
			  "attributes": {
				"capabilities": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"disable_rollback": {
				  "type": "bool",
				  "optional": true
				},
				"iam_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"notification_arns": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"on_failure": {
				  "type": "string",
				  "optional": true
				},
				"outputs": {
				  "type": [
					"map",
					"string"
				  ],
				  "computed": true
				},
				"parameters": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"policy_body": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy_url": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"template_body": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"template_url": {
				  "type": "string",
				  "optional": true
				},
				"timeout_in_minutes": {
				  "type": "number",
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_cloudformation_stack_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"administration_role_arn": {
				  "type": "string",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"capabilities": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"execution_role_name": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"parameters": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"stack_set_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"template_body": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"template_url": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_cloudformation_stack_set_instance": {
			"version": 0,
			"block": {
			  "attributes": {
				"account_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"parameter_overrides": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"region": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"retain_stack": {
				  "type": "bool",
				  "optional": true
				},
				"stack_id": {
				  "type": "string",
				  "computed": true
				},
				"stack_set_name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_cloudfront_distribution": {
			"version": 1,
			"block": {
			  "attributes": {
				"aliases": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"caller_reference": {
				  "type": "string",
				  "computed": true
				},
				"comment": {
				  "type": "string",
				  "optional": true
				},
				"default_root_object": {
				  "type": "string",
				  "optional": true
				},
				"domain_name": {
				  "type": "string",
				  "computed": true
				},
				"enabled": {
				  "type": "bool",
				  "required": true
				},
				"etag": {
				  "type": "string",
				  "computed": true
				},
				"hosted_zone_id": {
				  "type": "string",
				  "computed": true
				},
				"http_version": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"in_progress_validation_batches": {
				  "type": "number",
				  "computed": true
				},
				"is_ipv6_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"last_modified_time": {
				  "type": "string",
				  "computed": true
				},
				"price_class": {
				  "type": "string",
				  "optional": true
				},
				"retain_on_delete": {
				  "type": "bool",
				  "optional": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"trusted_signers": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"enabled": "bool",
						"items": [
						  "list",
						  [
							"object",
							{
							  "aws_account_number": "string",
							  "key_pair_ids": [
								"set",
								"string"
							  ]
							}
						  ]
						]
					  }
					]
				  ],
				  "computed": true
				},
				"wait_for_deployment": {
				  "type": "bool",
				  "optional": true
				},
				"web_acl_id": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"custom_error_response": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "error_caching_min_ttl": {
						"type": "number",
						"optional": true
					  },
					  "error_code": {
						"type": "number",
						"required": true
					  },
					  "response_code": {
						"type": "number",
						"optional": true
					  },
					  "response_page_path": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"default_cache_behavior": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allowed_methods": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "cached_methods": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "compress": {
						"type": "bool",
						"optional": true
					  },
					  "default_ttl": {
						"type": "number",
						"optional": true
					  },
					  "field_level_encryption_id": {
						"type": "string",
						"optional": true
					  },
					  "max_ttl": {
						"type": "number",
						"optional": true
					  },
					  "min_ttl": {
						"type": "number",
						"optional": true
					  },
					  "smooth_streaming": {
						"type": "bool",
						"optional": true
					  },
					  "target_origin_id": {
						"type": "string",
						"required": true
					  },
					  "trusted_signers": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "viewer_protocol_policy": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "forwarded_values": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"headers": {
							  "type": [
								"set",
								"string"
							  ],
							  "optional": true
							},
							"query_string": {
							  "type": "bool",
							  "required": true
							},
							"query_string_cache_keys": {
							  "type": [
								"list",
								"string"
							  ],
							  "optional": true
							}
						  },
						  "block_types": {
							"cookies": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "forward": {
									"type": "string",
									"required": true
								  },
								  "whitelisted_names": {
									"type": [
									  "set",
									  "string"
									],
									"optional": true
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  },
					  "lambda_function_association": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"event_type": {
							  "type": "string",
							  "required": true
							},
							"include_body": {
							  "type": "bool",
							  "optional": true
							},
							"lambda_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 4
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"logging_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket": {
						"type": "string",
						"required": true
					  },
					  "include_cookies": {
						"type": "bool",
						"optional": true
					  },
					  "prefix": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"ordered_cache_behavior": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allowed_methods": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "cached_methods": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "compress": {
						"type": "bool",
						"optional": true
					  },
					  "default_ttl": {
						"type": "number",
						"optional": true
					  },
					  "field_level_encryption_id": {
						"type": "string",
						"optional": true
					  },
					  "max_ttl": {
						"type": "number",
						"optional": true
					  },
					  "min_ttl": {
						"type": "number",
						"optional": true
					  },
					  "path_pattern": {
						"type": "string",
						"required": true
					  },
					  "smooth_streaming": {
						"type": "bool",
						"optional": true
					  },
					  "target_origin_id": {
						"type": "string",
						"required": true
					  },
					  "trusted_signers": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "viewer_protocol_policy": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "forwarded_values": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"headers": {
							  "type": [
								"set",
								"string"
							  ],
							  "optional": true
							},
							"query_string": {
							  "type": "bool",
							  "required": true
							},
							"query_string_cache_keys": {
							  "type": [
								"list",
								"string"
							  ],
							  "optional": true
							}
						  },
						  "block_types": {
							"cookies": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "forward": {
									"type": "string",
									"required": true
								  },
								  "whitelisted_names": {
									"type": [
									  "set",
									  "string"
									],
									"optional": true
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  },
					  "lambda_function_association": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"event_type": {
							  "type": "string",
							  "required": true
							},
							"include_body": {
							  "type": "bool",
							  "optional": true
							},
							"lambda_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 4
					  }
					}
				  }
				},
				"origin": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "domain_name": {
						"type": "string",
						"required": true
					  },
					  "origin_id": {
						"type": "string",
						"required": true
					  },
					  "origin_path": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "custom_header": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"name": {
							  "type": "string",
							  "required": true
							},
							"value": {
							  "type": "string",
							  "required": true
							}
						  }
						}
					  },
					  "custom_origin_config": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"http_port": {
							  "type": "number",
							  "required": true
							},
							"https_port": {
							  "type": "number",
							  "required": true
							},
							"origin_keepalive_timeout": {
							  "type": "number",
							  "optional": true
							},
							"origin_protocol_policy": {
							  "type": "string",
							  "required": true
							},
							"origin_read_timeout": {
							  "type": "number",
							  "optional": true
							},
							"origin_ssl_protocols": {
							  "type": [
								"set",
								"string"
							  ],
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "s3_origin_config": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"origin_access_identity": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1
				},
				"origin_group": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "origin_id": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "failover_criteria": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"status_codes": {
							  "type": [
								"set",
								"number"
							  ],
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  },
					  "member": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"origin_id": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 2,
						"max_items": 2
					  }
					}
				  }
				},
				"restrictions": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "geo_restriction": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"locations": {
							  "type": [
								"set",
								"string"
							  ],
							  "optional": true
							},
							"restriction_type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"viewer_certificate": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "acm_certificate_arn": {
						"type": "string",
						"optional": true
					  },
					  "cloudfront_default_certificate": {
						"type": "bool",
						"optional": true
					  },
					  "iam_certificate_id": {
						"type": "string",
						"optional": true
					  },
					  "minimum_protocol_version": {
						"type": "string",
						"optional": true
					  },
					  "ssl_support_method": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_cloudfront_origin_access_identity": {
			"version": 0,
			"block": {
			  "attributes": {
				"caller_reference": {
				  "type": "string",
				  "computed": true
				},
				"cloudfront_access_identity_path": {
				  "type": "string",
				  "computed": true
				},
				"comment": {
				  "type": "string",
				  "optional": true
				},
				"etag": {
				  "type": "string",
				  "computed": true
				},
				"iam_arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"s3_canonical_user_id": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_cloudfront_public_key": {
			"version": 0,
			"block": {
			  "attributes": {
				"caller_reference": {
				  "type": "string",
				  "computed": true
				},
				"comment": {
				  "type": "string",
				  "optional": true
				},
				"encoded_key": {
				  "type": "string",
				  "required": true
				},
				"etag": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_cloudhsm_v2_cluster": {
			"version": 0,
			"block": {
			  "attributes": {
				"cluster_certificates": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"aws_hardware_certificate": "string",
						"cluster_certificate": "string",
						"cluster_csr": "string",
						"hsm_certificate": "string",
						"manufacturer_hardware_certificate": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"cluster_id": {
				  "type": "string",
				  "computed": true
				},
				"cluster_state": {
				  "type": "string",
				  "computed": true
				},
				"hsm_type": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"security_group_id": {
				  "type": "string",
				  "computed": true
				},
				"source_backup_identifier": {
				  "type": "string",
				  "optional": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_cloudhsm_v2_hsm": {
			"version": 0,
			"block": {
			  "attributes": {
				"availability_zone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_id": {
				  "type": "string",
				  "required": true
				},
				"hsm_eni_id": {
				  "type": "string",
				  "computed": true
				},
				"hsm_id": {
				  "type": "string",
				  "computed": true
				},
				"hsm_state": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ip_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"subnet_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_cloudtrail": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"cloud_watch_logs_group_arn": {
				  "type": "string",
				  "optional": true
				},
				"cloud_watch_logs_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"enable_log_file_validation": {
				  "type": "bool",
				  "optional": true
				},
				"enable_logging": {
				  "type": "bool",
				  "optional": true
				},
				"home_region": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"include_global_service_events": {
				  "type": "bool",
				  "optional": true
				},
				"is_multi_region_trail": {
				  "type": "bool",
				  "optional": true
				},
				"is_organization_trail": {
				  "type": "bool",
				  "optional": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"s3_bucket_name": {
				  "type": "string",
				  "required": true
				},
				"s3_key_prefix": {
				  "type": "string",
				  "optional": true
				},
				"sns_topic_name": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"event_selector": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "include_management_events": {
						"type": "bool",
						"optional": true
					  },
					  "read_write_type": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "data_resource": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"type": {
							  "type": "string",
							  "required": true
							},
							"values": {
							  "type": [
								"list",
								"string"
							  ],
							  "required": true
							}
						  }
						}
					  }
					}
				  },
				  "max_items": 5
				}
			  }
			}
		  },
		  "aws_cloudwatch_dashboard": {
			"version": 0,
			"block": {
			  "attributes": {
				"dashboard_arn": {
				  "type": "string",
				  "computed": true
				},
				"dashboard_body": {
				  "type": "string",
				  "required": true
				},
				"dashboard_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_cloudwatch_event_permission": {
			"version": 0,
			"block": {
			  "attributes": {
				"action": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"principal": {
				  "type": "string",
				  "required": true
				},
				"statement_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"condition": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "key": {
						"type": "string",
						"required": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_cloudwatch_event_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"event_pattern": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"is_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"role_arn": {
				  "type": "string",
				  "optional": true
				},
				"schedule_expression": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_cloudwatch_event_target": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"input": {
				  "type": "string",
				  "optional": true
				},
				"input_path": {
				  "type": "string",
				  "optional": true
				},
				"role_arn": {
				  "type": "string",
				  "optional": true
				},
				"rule": {
				  "type": "string",
				  "required": true
				},
				"target_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"batch_target": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "array_size": {
						"type": "number",
						"optional": true
					  },
					  "job_attempts": {
						"type": "number",
						"optional": true
					  },
					  "job_definition": {
						"type": "string",
						"required": true
					  },
					  "job_name": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"ecs_target": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "group": {
						"type": "string",
						"optional": true
					  },
					  "launch_type": {
						"type": "string",
						"optional": true
					  },
					  "platform_version": {
						"type": "string",
						"optional": true
					  },
					  "task_count": {
						"type": "number",
						"optional": true
					  },
					  "task_definition_arn": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "network_configuration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"assign_public_ip": {
							  "type": "bool",
							  "optional": true
							},
							"security_groups": {
							  "type": [
								"set",
								"string"
							  ],
							  "optional": true
							},
							"subnets": {
							  "type": [
								"set",
								"string"
							  ],
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"input_transformer": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "input_paths": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  },
					  "input_template": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"kinesis_target": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "partition_key_path": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"run_command_targets": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "key": {
						"type": "string",
						"required": true
					  },
					  "values": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  }
					}
				  },
				  "max_items": 5
				},
				"sqs_target": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "message_group_id": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_cloudwatch_log_destination": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"role_arn": {
				  "type": "string",
				  "required": true
				},
				"target_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_cloudwatch_log_destination_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"access_policy": {
				  "type": "string",
				  "required": true
				},
				"destination_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_cloudwatch_log_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"retention_in_days": {
				  "type": "number",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_cloudwatch_log_metric_filter": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"log_group_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"pattern": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"metric_transformation": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "default_value": {
						"type": "string",
						"optional": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  },
					  "namespace": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_cloudwatch_log_resource_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy_document": {
				  "type": "string",
				  "required": true
				},
				"policy_name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_cloudwatch_log_stream": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"log_group_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_cloudwatch_log_subscription_filter": {
			"version": 0,
			"block": {
			  "attributes": {
				"destination_arn": {
				  "type": "string",
				  "required": true
				},
				"distribution": {
				  "type": "string",
				  "optional": true
				},
				"filter_pattern": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"log_group_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"role_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_cloudwatch_metric_alarm": {
			"version": 1,
			"block": {
			  "attributes": {
				"actions_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"alarm_actions": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"alarm_description": {
				  "type": "string",
				  "optional": true
				},
				"alarm_name": {
				  "type": "string",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"comparison_operator": {
				  "type": "string",
				  "required": true
				},
				"datapoints_to_alarm": {
				  "type": "number",
				  "optional": true
				},
				"dimensions": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"evaluate_low_sample_count_percentiles": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"evaluation_periods": {
				  "type": "number",
				  "required": true
				},
				"extended_statistic": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"insufficient_data_actions": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"metric_name": {
				  "type": "string",
				  "optional": true
				},
				"namespace": {
				  "type": "string",
				  "optional": true
				},
				"ok_actions": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"period": {
				  "type": "number",
				  "optional": true
				},
				"statistic": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"threshold": {
				  "type": "number",
				  "optional": true
				},
				"threshold_metric_id": {
				  "type": "string",
				  "optional": true
				},
				"treat_missing_data": {
				  "type": "string",
				  "optional": true
				},
				"unit": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"metric_query": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "expression": {
						"type": "string",
						"optional": true
					  },
					  "id": {
						"type": "string",
						"required": true
					  },
					  "label": {
						"type": "string",
						"optional": true
					  },
					  "return_data": {
						"type": "bool",
						"optional": true
					  }
					},
					"block_types": {
					  "metric": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"dimensions": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"metric_name": {
							  "type": "string",
							  "required": true
							},
							"namespace": {
							  "type": "string",
							  "optional": true
							},
							"period": {
							  "type": "number",
							  "required": true
							},
							"stat": {
							  "type": "string",
							  "required": true
							},
							"unit": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_codebuild_project": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"badge_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"badge_url": {
				  "type": "string",
				  "computed": true
				},
				"build_timeout": {
				  "type": "number",
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"encryption_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"queued_timeout": {
				  "type": "number",
				  "optional": true
				},
				"service_role": {
				  "type": "string",
				  "required": true
				},
				"source_version": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"artifacts": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "artifact_identifier": {
						"type": "string",
						"optional": true
					  },
					  "encryption_disabled": {
						"type": "bool",
						"optional": true
					  },
					  "location": {
						"type": "string",
						"optional": true
					  },
					  "name": {
						"type": "string",
						"optional": true
					  },
					  "namespace_type": {
						"type": "string",
						"optional": true
					  },
					  "override_artifact_name": {
						"type": "bool",
						"optional": true
					  },
					  "packaging": {
						"type": "string",
						"optional": true
					  },
					  "path": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"cache": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "location": {
						"type": "string",
						"optional": true
					  },
					  "modes": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"environment": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "certificate": {
						"type": "string",
						"optional": true
					  },
					  "compute_type": {
						"type": "string",
						"required": true
					  },
					  "image": {
						"type": "string",
						"required": true
					  },
					  "image_pull_credentials_type": {
						"type": "string",
						"optional": true
					  },
					  "privileged_mode": {
						"type": "bool",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "environment_variable": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"name": {
							  "type": "string",
							  "required": true
							},
							"type": {
							  "type": "string",
							  "optional": true
							},
							"value": {
							  "type": "string",
							  "required": true
							}
						  }
						}
					  },
					  "registry_credential": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"credential": {
							  "type": "string",
							  "required": true
							},
							"credential_provider": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"logs_config": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "cloudwatch_logs": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"group_name": {
							  "type": "string",
							  "optional": true
							},
							"status": {
							  "type": "string",
							  "optional": true
							},
							"stream_name": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "s3_logs": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"encryption_disabled": {
							  "type": "bool",
							  "optional": true
							},
							"location": {
							  "type": "string",
							  "optional": true
							},
							"status": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"secondary_artifacts": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "artifact_identifier": {
						"type": "string",
						"required": true
					  },
					  "encryption_disabled": {
						"type": "bool",
						"optional": true
					  },
					  "location": {
						"type": "string",
						"optional": true
					  },
					  "name": {
						"type": "string",
						"optional": true
					  },
					  "namespace_type": {
						"type": "string",
						"optional": true
					  },
					  "override_artifact_name": {
						"type": "bool",
						"optional": true
					  },
					  "packaging": {
						"type": "string",
						"optional": true
					  },
					  "path": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"secondary_sources": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "buildspec": {
						"type": "string",
						"optional": true
					  },
					  "git_clone_depth": {
						"type": "number",
						"optional": true
					  },
					  "insecure_ssl": {
						"type": "bool",
						"optional": true
					  },
					  "location": {
						"type": "string",
						"optional": true
					  },
					  "report_build_status": {
						"type": "bool",
						"optional": true
					  },
					  "source_identifier": {
						"type": "string",
						"required": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "auth": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"resource": {
							  "type": "string",
							  "optional": true,
							  "sensitive": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "git_submodules_config": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"fetch_submodules": {
							  "type": "bool",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  }
				},
				"source": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "buildspec": {
						"type": "string",
						"optional": true
					  },
					  "git_clone_depth": {
						"type": "number",
						"optional": true
					  },
					  "insecure_ssl": {
						"type": "bool",
						"optional": true
					  },
					  "location": {
						"type": "string",
						"optional": true
					  },
					  "report_build_status": {
						"type": "bool",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "auth": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"resource": {
							  "type": "string",
							  "optional": true,
							  "sensitive": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "git_submodules_config": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"fetch_submodules": {
							  "type": "bool",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"vpc_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "security_group_ids": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "subnets": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "vpc_id": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_codebuild_source_credential": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auth_type": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"server_type": {
				  "type": "string",
				  "required": true
				},
				"token": {
				  "type": "string",
				  "required": true,
				  "sensitive": true
				},
				"user_name": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_codebuild_webhook": {
			"version": 0,
			"block": {
			  "attributes": {
				"branch_filter": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"payload_url": {
				  "type": "string",
				  "computed": true
				},
				"project_name": {
				  "type": "string",
				  "required": true
				},
				"secret": {
				  "type": "string",
				  "computed": true,
				  "sensitive": true
				},
				"url": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"filter_group": {
				  "nesting_mode": "set",
				  "block": {
					"block_types": {
					  "filter": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"exclude_matched_pattern": {
							  "type": "bool",
							  "optional": true
							},
							"pattern": {
							  "type": "string",
							  "required": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						}
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_codecommit_repository": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"clone_url_http": {
				  "type": "string",
				  "computed": true
				},
				"clone_url_ssh": {
				  "type": "string",
				  "computed": true
				},
				"default_branch": {
				  "type": "string",
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"repository_id": {
				  "type": "string",
				  "computed": true
				},
				"repository_name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_codecommit_trigger": {
			"version": 0,
			"block": {
			  "attributes": {
				"configuration_id": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"repository_name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"trigger": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "branches": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "custom_data": {
						"type": "string",
						"optional": true
					  },
					  "destination_arn": {
						"type": "string",
						"required": true
					  },
					  "events": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 10
				}
			  }
			}
		  },
		  "aws_codedeploy_app": {
			"version": 0,
			"block": {
			  "attributes": {
				"compute_platform": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"unique_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_codedeploy_deployment_config": {
			"version": 0,
			"block": {
			  "attributes": {
				"compute_platform": {
				  "type": "string",
				  "optional": true
				},
				"deployment_config_id": {
				  "type": "string",
				  "computed": true
				},
				"deployment_config_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"minimum_healthy_hosts": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "type": {
						"type": "string",
						"optional": true
					  },
					  "value": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"traffic_routing_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "type": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "time_based_canary": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"interval": {
							  "type": "number",
							  "optional": true
							},
							"percentage": {
							  "type": "number",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "time_based_linear": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"interval": {
							  "type": "number",
							  "optional": true
							},
							"percentage": {
							  "type": "number",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_codedeploy_deployment_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"app_name": {
				  "type": "string",
				  "required": true
				},
				"autoscaling_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"deployment_config_name": {
				  "type": "string",
				  "optional": true
				},
				"deployment_group_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"service_role_arn": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"alarm_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "alarms": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "enabled": {
						"type": "bool",
						"optional": true
					  },
					  "ignore_poll_alarm_failure": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"auto_rollback_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"optional": true
					  },
					  "events": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"blue_green_deployment_config": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "deployment_ready_option": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"action_on_timeout": {
							  "type": "string",
							  "optional": true
							},
							"wait_time_in_minutes": {
							  "type": "number",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "green_fleet_provisioning_option": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"action": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "terminate_blue_instances_on_deployment_success": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"action": {
							  "type": "string",
							  "optional": true
							},
							"termination_wait_time_in_minutes": {
							  "type": "number",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"deployment_style": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "deployment_option": {
						"type": "string",
						"optional": true
					  },
					  "deployment_type": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"ec2_tag_filter": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "key": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  },
					  "value": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"ec2_tag_set": {
				  "nesting_mode": "set",
				  "block": {
					"block_types": {
					  "ec2_tag_filter": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"key": {
							  "type": "string",
							  "optional": true
							},
							"type": {
							  "type": "string",
							  "optional": true
							},
							"value": {
							  "type": "string",
							  "optional": true
							}
						  }
						}
					  }
					}
				  }
				},
				"ecs_service": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "cluster_name": {
						"type": "string",
						"required": true
					  },
					  "service_name": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"load_balancer_info": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "elb_info": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"name": {
							  "type": "string",
							  "optional": true
							}
						  }
						}
					  },
					  "target_group_info": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"name": {
							  "type": "string",
							  "optional": true
							}
						  }
						}
					  },
					  "target_group_pair_info": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"prod_traffic_route": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "listener_arns": {
									"type": [
									  "set",
									  "string"
									],
									"required": true
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							},
							"target_group": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "name": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 2
							},
							"test_traffic_route": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "listener_arns": {
									"type": [
									  "set",
									  "string"
									],
									"required": true
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"on_premises_instance_tag_filter": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "key": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  },
					  "value": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"trigger_configuration": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "trigger_events": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "trigger_name": {
						"type": "string",
						"required": true
					  },
					  "trigger_target_arn": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_codepipeline": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"role_arn": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"artifact_store": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "location": {
						"type": "string",
						"required": true
					  },
					  "region": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "encryption_key": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"id": {
							  "type": "string",
							  "required": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1
				},
				"stage": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "name": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "action": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"category": {
							  "type": "string",
							  "required": true
							},
							"configuration": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"input_artifacts": {
							  "type": [
								"list",
								"string"
							  ],
							  "optional": true
							},
							"name": {
							  "type": "string",
							  "required": true
							},
							"namespace": {
							  "type": "string",
							  "optional": true
							},
							"output_artifacts": {
							  "type": [
								"list",
								"string"
							  ],
							  "optional": true
							},
							"owner": {
							  "type": "string",
							  "required": true
							},
							"provider": {
							  "type": "string",
							  "required": true
							},
							"region": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"role_arn": {
							  "type": "string",
							  "optional": true
							},
							"run_order": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"version": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1
					  }
					}
				  },
				  "min_items": 2
				}
			  }
			}
		  },
		  "aws_codepipeline_webhook": {
			"version": 0,
			"block": {
			  "attributes": {
				"authentication": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"target_action": {
				  "type": "string",
				  "required": true
				},
				"target_pipeline": {
				  "type": "string",
				  "required": true
				},
				"url": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"authentication_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allowed_ip_range": {
						"type": "string",
						"optional": true
					  },
					  "secret_token": {
						"type": "string",
						"optional": true,
						"sensitive": true
					  }
					}
				  },
				  "max_items": 1
				},
				"filter": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "json_path": {
						"type": "string",
						"required": true
					  },
					  "match_equals": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1
				}
			  }
			}
		  },
		  "aws_codestarnotifications_notification_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"detail_type": {
				  "type": "string",
				  "required": true
				},
				"event_type_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"resource": {
				  "type": "string",
				  "required": true
				},
				"status": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"target": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "address": {
						"type": "string",
						"required": true
					  },
					  "status": {
						"type": "string",
						"computed": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 10
				}
			  }
			}
		  },
		  "aws_cognito_identity_pool": {
			"version": 0,
			"block": {
			  "attributes": {
				"allow_unauthenticated_identities": {
				  "type": "bool",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"developer_provider_name": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identity_pool_name": {
				  "type": "string",
				  "required": true
				},
				"openid_connect_provider_arns": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"saml_provider_arns": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"supported_login_providers": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"cognito_identity_providers": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "client_id": {
						"type": "string",
						"optional": true
					  },
					  "provider_name": {
						"type": "string",
						"optional": true
					  },
					  "server_side_token_check": {
						"type": "bool",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_cognito_identity_pool_roles_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identity_pool_id": {
				  "type": "string",
				  "required": true
				},
				"roles": {
				  "type": [
					"map",
					"string"
				  ],
				  "required": true
				}
			  },
			  "block_types": {
				"role_mapping": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "ambiguous_role_resolution": {
						"type": "string",
						"optional": true
					  },
					  "identity_provider": {
						"type": "string",
						"required": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "mapping_rule": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"claim": {
							  "type": "string",
							  "required": true
							},
							"match_type": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							},
							"value": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 25
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_cognito_identity_provider": {
			"version": 0,
			"block": {
			  "attributes": {
				"attribute_mapping": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"idp_identifiers": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"provider_details": {
				  "type": [
					"map",
					"string"
				  ],
				  "required": true
				},
				"provider_name": {
				  "type": "string",
				  "required": true
				},
				"provider_type": {
				  "type": "string",
				  "required": true
				},
				"user_pool_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_cognito_resource_server": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identifier": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"scope_identifiers": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"user_pool_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"scope": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "scope_description": {
						"type": "string",
						"required": true
					  },
					  "scope_name": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 100
				}
			  }
			}
		  },
		  "aws_cognito_user_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"precedence": {
				  "type": "number",
				  "optional": true
				},
				"role_arn": {
				  "type": "string",
				  "optional": true
				},
				"user_pool_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_cognito_user_pool": {
			"version": 0,
			"block": {
			  "attributes": {
				"alias_attributes": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_verified_attributes": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"creation_date": {
				  "type": "string",
				  "computed": true
				},
				"email_verification_message": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"email_verification_subject": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"endpoint": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"last_modified_date": {
				  "type": "string",
				  "computed": true
				},
				"mfa_configuration": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"sms_authentication_message": {
				  "type": "string",
				  "optional": true
				},
				"sms_verification_message": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"username_attributes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"admin_create_user_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allow_admin_create_user_only": {
						"type": "bool",
						"optional": true
					  }
					},
					"block_types": {
					  "invite_message_template": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"email_message": {
							  "type": "string",
							  "optional": true
							},
							"email_subject": {
							  "type": "string",
							  "optional": true
							},
							"sms_message": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"device_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "challenge_required_on_new_device": {
						"type": "bool",
						"optional": true
					  },
					  "device_only_remembered_on_user_prompt": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"email_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "email_sending_account": {
						"type": "string",
						"optional": true
					  },
					  "from_email_address": {
						"type": "string",
						"optional": true
					  },
					  "reply_to_email_address": {
						"type": "string",
						"optional": true
					  },
					  "source_arn": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"lambda_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "create_auth_challenge": {
						"type": "string",
						"optional": true
					  },
					  "custom_message": {
						"type": "string",
						"optional": true
					  },
					  "define_auth_challenge": {
						"type": "string",
						"optional": true
					  },
					  "post_authentication": {
						"type": "string",
						"optional": true
					  },
					  "post_confirmation": {
						"type": "string",
						"optional": true
					  },
					  "pre_authentication": {
						"type": "string",
						"optional": true
					  },
					  "pre_sign_up": {
						"type": "string",
						"optional": true
					  },
					  "pre_token_generation": {
						"type": "string",
						"optional": true
					  },
					  "user_migration": {
						"type": "string",
						"optional": true
					  },
					  "verify_auth_challenge_response": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"password_policy": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "minimum_length": {
						"type": "number",
						"optional": true
					  },
					  "require_lowercase": {
						"type": "bool",
						"optional": true
					  },
					  "require_numbers": {
						"type": "bool",
						"optional": true
					  },
					  "require_symbols": {
						"type": "bool",
						"optional": true
					  },
					  "require_uppercase": {
						"type": "bool",
						"optional": true
					  },
					  "temporary_password_validity_days": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"schema": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "attribute_data_type": {
						"type": "string",
						"required": true
					  },
					  "developer_only_attribute": {
						"type": "bool",
						"optional": true
					  },
					  "mutable": {
						"type": "bool",
						"optional": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  },
					  "required": {
						"type": "bool",
						"optional": true
					  }
					},
					"block_types": {
					  "number_attribute_constraints": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"max_value": {
							  "type": "string",
							  "optional": true
							},
							"min_value": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "string_attribute_constraints": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"max_length": {
							  "type": "string",
							  "optional": true
							},
							"min_length": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 50
				},
				"sms_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "external_id": {
						"type": "string",
						"required": true
					  },
					  "sns_caller_arn": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"software_token_mfa_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"user_pool_add_ons": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "advanced_security_mode": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"username_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "case_sensitive": {
						"type": "bool",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"verification_message_template": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "default_email_option": {
						"type": "string",
						"optional": true
					  },
					  "email_message": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "email_message_by_link": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "email_subject": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "email_subject_by_link": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "sms_message": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_cognito_user_pool_client": {
			"version": 0,
			"block": {
			  "attributes": {
				"allowed_oauth_flows": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"allowed_oauth_flows_user_pool_client": {
				  "type": "bool",
				  "optional": true
				},
				"allowed_oauth_scopes": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"callback_urls": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"client_secret": {
				  "type": "string",
				  "computed": true,
				  "sensitive": true
				},
				"default_redirect_uri": {
				  "type": "string",
				  "optional": true
				},
				"explicit_auth_flows": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"generate_secret": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"logout_urls": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"prevent_user_existence_errors": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"read_attributes": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"refresh_token_validity": {
				  "type": "number",
				  "optional": true
				},
				"supported_identity_providers": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"user_pool_id": {
				  "type": "string",
				  "required": true
				},
				"write_attributes": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"analytics_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "application_id": {
						"type": "string",
						"required": true
					  },
					  "external_id": {
						"type": "string",
						"required": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "user_data_shared": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_cognito_user_pool_domain": {
			"version": 0,
			"block": {
			  "attributes": {
				"aws_account_id": {
				  "type": "string",
				  "computed": true
				},
				"certificate_arn": {
				  "type": "string",
				  "optional": true
				},
				"cloudfront_distribution_arn": {
				  "type": "string",
				  "computed": true
				},
				"domain": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"s3_bucket": {
				  "type": "string",
				  "computed": true
				},
				"user_pool_id": {
				  "type": "string",
				  "required": true
				},
				"version": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_config_aggregate_authorization": {
			"version": 0,
			"block": {
			  "attributes": {
				"account_id": {
				  "type": "string",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"region": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_config_config_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"input_parameters": {
				  "type": "string",
				  "optional": true
				},
				"maximum_execution_frequency": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"rule_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"scope": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "compliance_resource_id": {
						"type": "string",
						"optional": true
					  },
					  "compliance_resource_types": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "tag_key": {
						"type": "string",
						"optional": true
					  },
					  "tag_value": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"source": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "owner": {
						"type": "string",
						"required": true
					  },
					  "source_identifier": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "source_detail": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"event_source": {
							  "type": "string",
							  "optional": true
							},
							"maximum_execution_frequency": {
							  "type": "string",
							  "optional": true
							},
							"message_type": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 25
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_config_configuration_aggregator": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"account_aggregation_source": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "account_ids": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  },
					  "all_regions": {
						"type": "bool",
						"optional": true
					  },
					  "regions": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"organization_aggregation_source": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "all_regions": {
						"type": "bool",
						"optional": true
					  },
					  "regions": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_config_configuration_recorder": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"role_arn": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"recording_group": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "all_supported": {
						"type": "bool",
						"optional": true
					  },
					  "include_global_resource_types": {
						"type": "bool",
						"optional": true
					  },
					  "resource_types": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_config_configuration_recorder_status": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"is_enabled": {
				  "type": "bool",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_config_delivery_channel": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"s3_bucket_name": {
				  "type": "string",
				  "required": true
				},
				"s3_key_prefix": {
				  "type": "string",
				  "optional": true
				},
				"sns_topic_arn": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"snapshot_delivery_properties": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "delivery_frequency": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_config_organization_custom_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"excluded_accounts": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"input_parameters": {
				  "type": "string",
				  "optional": true
				},
				"lambda_function_arn": {
				  "type": "string",
				  "required": true
				},
				"maximum_execution_frequency": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"resource_id_scope": {
				  "type": "string",
				  "optional": true
				},
				"resource_types_scope": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tag_key_scope": {
				  "type": "string",
				  "optional": true
				},
				"tag_value_scope": {
				  "type": "string",
				  "optional": true
				},
				"trigger_types": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_config_organization_managed_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"excluded_accounts": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"input_parameters": {
				  "type": "string",
				  "optional": true
				},
				"maximum_execution_frequency": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"resource_id_scope": {
				  "type": "string",
				  "optional": true
				},
				"resource_types_scope": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"rule_identifier": {
				  "type": "string",
				  "required": true
				},
				"tag_key_scope": {
				  "type": "string",
				  "optional": true
				},
				"tag_value_scope": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_cur_report_definition": {
			"version": 0,
			"block": {
			  "attributes": {
				"additional_artifacts": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"additional_schema_elements": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"compression": {
				  "type": "string",
				  "required": true
				},
				"format": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"report_name": {
				  "type": "string",
				  "required": true
				},
				"s3_bucket": {
				  "type": "string",
				  "required": true
				},
				"s3_prefix": {
				  "type": "string",
				  "optional": true
				},
				"s3_region": {
				  "type": "string",
				  "required": true
				},
				"time_unit": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_customer_gateway": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"bgp_asn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ip_address": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"type": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_datapipeline_pipeline": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_datasync_agent": {
			"version": 0,
			"block": {
			  "attributes": {
				"activation_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ip_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_datasync_location_efs": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"efs_file_system_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"subdirectory": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"uri": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"ec2_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "security_group_arns": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "subnet_arn": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_datasync_location_nfs": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"server_hostname": {
				  "type": "string",
				  "required": true
				},
				"subdirectory": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"uri": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"on_prem_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "agent_arns": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_datasync_location_s3": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"s3_bucket_arn": {
				  "type": "string",
				  "required": true
				},
				"subdirectory": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"uri": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"s3_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket_access_role_arn": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_datasync_location_smb": {
			"version": 0,
			"block": {
			  "attributes": {
				"agent_arns": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"domain": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"password": {
				  "type": "string",
				  "required": true,
				  "sensitive": true
				},
				"server_hostname": {
				  "type": "string",
				  "required": true
				},
				"subdirectory": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"uri": {
				  "type": "string",
				  "computed": true
				},
				"user": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"mount_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "version": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_datasync_task": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"cloudwatch_log_group_arn": {
				  "type": "string",
				  "optional": true
				},
				"destination_location_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"source_location_arn": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "atime": {
						"type": "string",
						"optional": true
					  },
					  "bytes_per_second": {
						"type": "number",
						"optional": true
					  },
					  "gid": {
						"type": "string",
						"optional": true
					  },
					  "mtime": {
						"type": "string",
						"optional": true
					  },
					  "posix_permissions": {
						"type": "string",
						"optional": true
					  },
					  "preserve_deleted_files": {
						"type": "string",
						"optional": true
					  },
					  "preserve_devices": {
						"type": "string",
						"optional": true
					  },
					  "uid": {
						"type": "string",
						"optional": true
					  },
					  "verify_mode": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dax_cluster": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"availability_zones": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"cluster_address": {
				  "type": "string",
				  "computed": true
				},
				"cluster_name": {
				  "type": "string",
				  "required": true
				},
				"configuration_endpoint": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"iam_role_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"maintenance_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"node_type": {
				  "type": "string",
				  "required": true
				},
				"nodes": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"address": "string",
						"availability_zone": "string",
						"id": "string",
						"port": "number"
					  }
					]
				  ],
				  "computed": true
				},
				"notification_topic_arn": {
				  "type": "string",
				  "optional": true
				},
				"parameter_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "computed": true
				},
				"replication_factor": {
				  "type": "number",
				  "required": true
				},
				"security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"subnet_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"server_side_encryption": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dax_parameter_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"parameters": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "name": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dax_subnet_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_db_cluster_snapshot": {
			"version": 0,
			"block": {
			  "attributes": {
				"allocated_storage": {
				  "type": "number",
				  "computed": true
				},
				"availability_zones": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"db_cluster_identifier": {
				  "type": "string",
				  "required": true
				},
				"db_cluster_snapshot_arn": {
				  "type": "string",
				  "computed": true
				},
				"db_cluster_snapshot_identifier": {
				  "type": "string",
				  "required": true
				},
				"engine": {
				  "type": "string",
				  "computed": true
				},
				"engine_version": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "computed": true
				},
				"license_model": {
				  "type": "string",
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "computed": true
				},
				"snapshot_type": {
				  "type": "string",
				  "computed": true
				},
				"source_db_cluster_snapshot_arn": {
				  "type": "string",
				  "computed": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"storage_encrypted": {
				  "type": "bool",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_db_event_subscription": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"customer_aws_id": {
				  "type": "string",
				  "computed": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"event_categories": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"sns_topic": {
				  "type": "string",
				  "required": true
				},
				"source_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"source_type": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_db_instance": {
			"version": 1,
			"block": {
			  "attributes": {
				"address": {
				  "type": "string",
				  "computed": true
				},
				"allocated_storage": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"allow_major_version_upgrade": {
				  "type": "bool",
				  "optional": true
				},
				"apply_immediately": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_minor_version_upgrade": {
				  "type": "bool",
				  "optional": true
				},
				"availability_zone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"backup_retention_period": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"backup_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ca_cert_identifier": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"character_set_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"copy_tags_to_snapshot": {
				  "type": "bool",
				  "optional": true
				},
				"db_subnet_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"delete_automated_backups": {
				  "type": "bool",
				  "optional": true
				},
				"deletion_protection": {
				  "type": "bool",
				  "optional": true
				},
				"domain": {
				  "type": "string",
				  "optional": true
				},
				"domain_iam_role_name": {
				  "type": "string",
				  "optional": true
				},
				"enabled_cloudwatch_logs_exports": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"endpoint": {
				  "type": "string",
				  "computed": true
				},
				"engine": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"engine_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"final_snapshot_identifier": {
				  "type": "string",
				  "optional": true
				},
				"hosted_zone_id": {
				  "type": "string",
				  "computed": true
				},
				"iam_database_authentication_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identifier": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identifier_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_class": {
				  "type": "string",
				  "required": true
				},
				"iops": {
				  "type": "number",
				  "optional": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"license_model": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"maintenance_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"max_allocated_storage": {
				  "type": "number",
				  "optional": true
				},
				"monitoring_interval": {
				  "type": "number",
				  "optional": true
				},
				"monitoring_role_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"multi_az": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"option_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"parameter_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"password": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"performance_insights_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"performance_insights_kms_key_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"performance_insights_retention_period": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"publicly_accessible": {
				  "type": "bool",
				  "optional": true
				},
				"replicas": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"replicate_source_db": {
				  "type": "string",
				  "optional": true
				},
				"resource_id": {
				  "type": "string",
				  "computed": true
				},
				"security_group_names": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"skip_final_snapshot": {
				  "type": "bool",
				  "optional": true
				},
				"snapshot_identifier": {
				  "type": "string",
				  "optional": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"storage_encrypted": {
				  "type": "bool",
				  "optional": true
				},
				"storage_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"timezone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"username": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"vpc_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"s3_import": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket_name": {
						"type": "string",
						"required": true
					  },
					  "bucket_prefix": {
						"type": "string",
						"optional": true
					  },
					  "ingestion_role": {
						"type": "string",
						"required": true
					  },
					  "source_engine": {
						"type": "string",
						"required": true
					  },
					  "source_engine_version": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_db_instance_role_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"db_instance_identifier": {
				  "type": "string",
				  "required": true
				},
				"feature_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"role_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_db_option_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"engine_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"major_engine_version": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"option_group_description": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"option": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "db_security_group_memberships": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "option_name": {
						"type": "string",
						"required": true
					  },
					  "port": {
						"type": "number",
						"optional": true
					  },
					  "version": {
						"type": "string",
						"optional": true
					  },
					  "vpc_security_group_memberships": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  }
					},
					"block_types": {
					  "option_settings": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"name": {
							  "type": "string",
							  "required": true
							},
							"value": {
							  "type": "string",
							  "required": true
							}
						  }
						}
					  }
					}
				  }
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_db_parameter_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"family": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"parameter": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "apply_method": {
						"type": "string",
						"optional": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_db_security_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"ingress": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "cidr": {
						"type": "string",
						"optional": true
					  },
					  "security_group_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "security_group_name": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "security_group_owner_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "min_items": 1
				}
			  }
			}
		  },
		  "aws_db_snapshot": {
			"version": 0,
			"block": {
			  "attributes": {
				"allocated_storage": {
				  "type": "number",
				  "computed": true
				},
				"availability_zone": {
				  "type": "string",
				  "computed": true
				},
				"db_instance_identifier": {
				  "type": "string",
				  "required": true
				},
				"db_snapshot_arn": {
				  "type": "string",
				  "computed": true
				},
				"db_snapshot_identifier": {
				  "type": "string",
				  "required": true
				},
				"encrypted": {
				  "type": "bool",
				  "computed": true
				},
				"engine": {
				  "type": "string",
				  "computed": true
				},
				"engine_version": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"iops": {
				  "type": "number",
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "computed": true
				},
				"license_model": {
				  "type": "string",
				  "computed": true
				},
				"option_group_name": {
				  "type": "string",
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "computed": true
				},
				"snapshot_type": {
				  "type": "string",
				  "computed": true
				},
				"source_db_snapshot_identifier": {
				  "type": "string",
				  "computed": true
				},
				"source_region": {
				  "type": "string",
				  "computed": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"storage_type": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "read": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_db_subnet_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_default_network_acl": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"default_network_acl_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"egress": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "action": {
						"type": "string",
						"required": true
					  },
					  "cidr_block": {
						"type": "string",
						"optional": true
					  },
					  "from_port": {
						"type": "number",
						"required": true
					  },
					  "icmp_code": {
						"type": "number",
						"optional": true
					  },
					  "icmp_type": {
						"type": "number",
						"optional": true
					  },
					  "ipv6_cidr_block": {
						"type": "string",
						"optional": true
					  },
					  "protocol": {
						"type": "string",
						"required": true
					  },
					  "rule_no": {
						"type": "number",
						"required": true
					  },
					  "to_port": {
						"type": "number",
						"required": true
					  }
					}
				  }
				},
				"ingress": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "action": {
						"type": "string",
						"required": true
					  },
					  "cidr_block": {
						"type": "string",
						"optional": true
					  },
					  "from_port": {
						"type": "number",
						"required": true
					  },
					  "icmp_code": {
						"type": "number",
						"optional": true
					  },
					  "icmp_type": {
						"type": "number",
						"optional": true
					  },
					  "ipv6_cidr_block": {
						"type": "string",
						"optional": true
					  },
					  "protocol": {
						"type": "string",
						"required": true
					  },
					  "rule_no": {
						"type": "number",
						"required": true
					  },
					  "to_port": {
						"type": "number",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_default_route_table": {
			"version": 0,
			"block": {
			  "attributes": {
				"default_route_table_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"propagating_vgws": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"route": {
				  "type": [
					"set",
					[
					  "object",
					  {
						"cidr_block": "string",
						"egress_only_gateway_id": "string",
						"gateway_id": "string",
						"instance_id": "string",
						"ipv6_cidr_block": "string",
						"nat_gateway_id": "string",
						"network_interface_id": "string",
						"transit_gateway_id": "string",
						"vpc_peering_connection_id": "string"
					  }
					]
				  ],
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_default_security_group": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "computed": true
				},
				"egress": {
				  "type": [
					"set",
					[
					  "object",
					  {
						"cidr_blocks": [
						  "list",
						  "string"
						],
						"description": "string",
						"from_port": "number",
						"ipv6_cidr_blocks": [
						  "list",
						  "string"
						],
						"prefix_list_ids": [
						  "list",
						  "string"
						],
						"protocol": "string",
						"security_groups": [
						  "set",
						  "string"
						],
						"self": "bool",
						"to_port": "number"
					  }
					]
				  ],
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ingress": {
				  "type": [
					"set",
					[
					  "object",
					  {
						"cidr_blocks": [
						  "list",
						  "string"
						],
						"description": "string",
						"from_port": "number",
						"ipv6_cidr_blocks": [
						  "list",
						  "string"
						],
						"prefix_list_ids": [
						  "list",
						  "string"
						],
						"protocol": "string",
						"security_groups": [
						  "set",
						  "string"
						],
						"self": "bool",
						"to_port": "number"
					  }
					]
				  ],
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"revoke_rules_on_delete": {
				  "type": "bool",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_default_subnet": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"assign_ipv6_address_on_creation": {
				  "type": "bool",
				  "computed": true
				},
				"availability_zone": {
				  "type": "string",
				  "required": true
				},
				"availability_zone_id": {
				  "type": "string",
				  "computed": true
				},
				"cidr_block": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ipv6_cidr_block": {
				  "type": "string",
				  "computed": true
				},
				"ipv6_cidr_block_association_id": {
				  "type": "string",
				  "computed": true
				},
				"map_public_ip_on_launch": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"outpost_arn": {
				  "type": "string",
				  "optional": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_default_vpc": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"assign_generated_ipv6_cidr_block": {
				  "type": "bool",
				  "computed": true
				},
				"cidr_block": {
				  "type": "string",
				  "computed": true
				},
				"default_network_acl_id": {
				  "type": "string",
				  "computed": true
				},
				"default_route_table_id": {
				  "type": "string",
				  "computed": true
				},
				"default_security_group_id": {
				  "type": "string",
				  "computed": true
				},
				"dhcp_options_id": {
				  "type": "string",
				  "computed": true
				},
				"enable_classiclink": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"enable_classiclink_dns_support": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"enable_dns_hostnames": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"enable_dns_support": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_tenancy": {
				  "type": "string",
				  "computed": true
				},
				"ipv6_association_id": {
				  "type": "string",
				  "computed": true
				},
				"ipv6_cidr_block": {
				  "type": "string",
				  "computed": true
				},
				"main_route_table_id": {
				  "type": "string",
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_default_vpc_dhcp_options": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"domain_name": {
				  "type": "string",
				  "computed": true
				},
				"domain_name_servers": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"netbios_name_servers": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"netbios_node_type": {
				  "type": "string",
				  "optional": true
				},
				"ntp_servers": {
				  "type": "string",
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_devicefarm_project": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_directory_service_conditional_forwarder": {
			"version": 0,
			"block": {
			  "attributes": {
				"directory_id": {
				  "type": "string",
				  "required": true
				},
				"dns_ips": {
				  "type": [
					"list",
					"string"
				  ],
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"remote_domain_name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_directory_service_directory": {
			"version": 0,
			"block": {
			  "attributes": {
				"access_url": {
				  "type": "string",
				  "computed": true
				},
				"alias": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"dns_ip_addresses": {
				  "type": [
					"set",
					"string"
				  ],
				  "computed": true
				},
				"edition": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"enable_sso": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"password": {
				  "type": "string",
				  "required": true,
				  "sensitive": true
				},
				"security_group_id": {
				  "type": "string",
				  "computed": true
				},
				"short_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"size": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"type": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"connect_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "availability_zones": {
						"type": [
						  "set",
						  "string"
						],
						"computed": true
					  },
					  "connect_ips": {
						"type": [
						  "set",
						  "string"
						],
						"computed": true
					  },
					  "customer_dns_ips": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "customer_username": {
						"type": "string",
						"required": true
					  },
					  "subnet_ids": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "vpc_id": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"vpc_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "availability_zones": {
						"type": [
						  "set",
						  "string"
						],
						"computed": true
					  },
					  "subnet_ids": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "vpc_id": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_directory_service_log_subscription": {
			"version": 0,
			"block": {
			  "attributes": {
				"directory_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"log_group_name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_dlm_lifecycle_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "required": true
				},
				"execution_role_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"state": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"policy_details": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "resource_types": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  },
					  "target_tags": {
						"type": [
						  "map",
						  "string"
						],
						"required": true
					  }
					},
					"block_types": {
					  "schedule": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"copy_tags": {
							  "type": "bool",
							  "optional": true,
							  "computed": true
							},
							"name": {
							  "type": "string",
							  "required": true
							},
							"tags_to_add": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							}
						  },
						  "block_types": {
							"create_rule": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "interval": {
									"type": "number",
									"required": true
								  },
								  "interval_unit": {
									"type": "string",
									"optional": true
								  },
								  "times": {
									"type": [
									  "list",
									  "string"
									],
									"optional": true,
									"computed": true
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							},
							"retain_rule": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "count": {
									"type": "number",
									"required": true
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"min_items": 1
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_dms_certificate": {
			"version": 0,
			"block": {
			  "attributes": {
				"certificate_arn": {
				  "type": "string",
				  "computed": true
				},
				"certificate_id": {
				  "type": "string",
				  "required": true
				},
				"certificate_pem": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"certificate_wallet": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_dms_endpoint": {
			"version": 0,
			"block": {
			  "attributes": {
				"certificate_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"database_name": {
				  "type": "string",
				  "optional": true
				},
				"endpoint_arn": {
				  "type": "string",
				  "computed": true
				},
				"endpoint_id": {
				  "type": "string",
				  "required": true
				},
				"endpoint_type": {
				  "type": "string",
				  "required": true
				},
				"engine_name": {
				  "type": "string",
				  "required": true
				},
				"extra_connection_attributes": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"password": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"port": {
				  "type": "number",
				  "optional": true
				},
				"server_name": {
				  "type": "string",
				  "optional": true
				},
				"service_access_role": {
				  "type": "string",
				  "optional": true
				},
				"ssl_mode": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"username": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"elasticsearch_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "endpoint_uri": {
						"type": "string",
						"required": true
					  },
					  "error_retry_duration": {
						"type": "number",
						"optional": true
					  },
					  "full_load_error_percentage": {
						"type": "number",
						"optional": true
					  },
					  "service_access_role_arn": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"kafka_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "broker": {
						"type": "string",
						"required": true
					  },
					  "topic": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"kinesis_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "message_format": {
						"type": "string",
						"optional": true
					  },
					  "service_access_role_arn": {
						"type": "string",
						"optional": true
					  },
					  "stream_arn": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"mongodb_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "auth_mechanism": {
						"type": "string",
						"optional": true
					  },
					  "auth_source": {
						"type": "string",
						"optional": true
					  },
					  "auth_type": {
						"type": "string",
						"optional": true
					  },
					  "docs_to_investigate": {
						"type": "string",
						"optional": true
					  },
					  "extract_doc_id": {
						"type": "string",
						"optional": true
					  },
					  "nesting_level": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"s3_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket_folder": {
						"type": "string",
						"optional": true
					  },
					  "bucket_name": {
						"type": "string",
						"optional": true
					  },
					  "compression_type": {
						"type": "string",
						"optional": true
					  },
					  "csv_delimiter": {
						"type": "string",
						"optional": true
					  },
					  "csv_row_delimiter": {
						"type": "string",
						"optional": true
					  },
					  "external_table_definition": {
						"type": "string",
						"optional": true
					  },
					  "service_access_role_arn": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_dms_event_subscription": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"event_categories": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"sns_topic_arn": {
				  "type": "string",
				  "required": true
				},
				"source_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"source_type": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dms_replication_instance": {
			"version": 0,
			"block": {
			  "attributes": {
				"allocated_storage": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"apply_immediately": {
				  "type": "bool",
				  "optional": true
				},
				"auto_minor_version_upgrade": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"availability_zone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"engine_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"multi_az": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"preferred_maintenance_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"publicly_accessible": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"replication_instance_arn": {
				  "type": "string",
				  "computed": true
				},
				"replication_instance_class": {
				  "type": "string",
				  "required": true
				},
				"replication_instance_id": {
				  "type": "string",
				  "required": true
				},
				"replication_instance_private_ips": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"replication_instance_public_ips": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"replication_subnet_group_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dms_replication_subnet_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"replication_subnet_group_arn": {
				  "type": "string",
				  "computed": true
				},
				"replication_subnet_group_description": {
				  "type": "string",
				  "required": true
				},
				"replication_subnet_group_id": {
				  "type": "string",
				  "required": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_dms_replication_task": {
			"version": 0,
			"block": {
			  "attributes": {
				"cdc_start_time": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"migration_type": {
				  "type": "string",
				  "required": true
				},
				"replication_instance_arn": {
				  "type": "string",
				  "required": true
				},
				"replication_task_arn": {
				  "type": "string",
				  "computed": true
				},
				"replication_task_id": {
				  "type": "string",
				  "required": true
				},
				"replication_task_settings": {
				  "type": "string",
				  "optional": true
				},
				"source_endpoint_arn": {
				  "type": "string",
				  "required": true
				},
				"table_mappings": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"target_endpoint_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_docdb_cluster": {
			"version": 0,
			"block": {
			  "attributes": {
				"apply_immediately": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"availability_zones": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"backup_retention_period": {
				  "type": "number",
				  "optional": true
				},
				"cluster_identifier": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_identifier_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_members": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"cluster_resource_id": {
				  "type": "string",
				  "computed": true
				},
				"db_cluster_parameter_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"db_subnet_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"deletion_protection": {
				  "type": "bool",
				  "optional": true
				},
				"enabled_cloudwatch_logs_exports": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"endpoint": {
				  "type": "string",
				  "computed": true
				},
				"engine": {
				  "type": "string",
				  "optional": true
				},
				"engine_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"final_snapshot_identifier": {
				  "type": "string",
				  "optional": true
				},
				"hosted_zone_id": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"master_password": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"master_username": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "optional": true
				},
				"preferred_backup_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"preferred_maintenance_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"reader_endpoint": {
				  "type": "string",
				  "computed": true
				},
				"skip_final_snapshot": {
				  "type": "bool",
				  "optional": true
				},
				"snapshot_identifier": {
				  "type": "string",
				  "optional": true
				},
				"storage_encrypted": {
				  "type": "bool",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_docdb_cluster_instance": {
			"version": 0,
			"block": {
			  "attributes": {
				"apply_immediately": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_minor_version_upgrade": {
				  "type": "bool",
				  "optional": true
				},
				"availability_zone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ca_cert_identifier": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_identifier": {
				  "type": "string",
				  "required": true
				},
				"db_subnet_group_name": {
				  "type": "string",
				  "computed": true
				},
				"dbi_resource_id": {
				  "type": "string",
				  "computed": true
				},
				"endpoint": {
				  "type": "string",
				  "computed": true
				},
				"engine": {
				  "type": "string",
				  "optional": true
				},
				"engine_version": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identifier": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identifier_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_class": {
				  "type": "string",
				  "required": true
				},
				"kms_key_id": {
				  "type": "string",
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "computed": true
				},
				"preferred_backup_window": {
				  "type": "string",
				  "computed": true
				},
				"preferred_maintenance_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"promotion_tier": {
				  "type": "number",
				  "optional": true
				},
				"publicly_accessible": {
				  "type": "bool",
				  "computed": true
				},
				"storage_encrypted": {
				  "type": "bool",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"writer": {
				  "type": "bool",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_docdb_cluster_parameter_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"family": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"parameter": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "apply_method": {
						"type": "string",
						"optional": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_docdb_cluster_snapshot": {
			"version": 0,
			"block": {
			  "attributes": {
				"availability_zones": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"db_cluster_identifier": {
				  "type": "string",
				  "required": true
				},
				"db_cluster_snapshot_arn": {
				  "type": "string",
				  "computed": true
				},
				"db_cluster_snapshot_identifier": {
				  "type": "string",
				  "required": true
				},
				"engine": {
				  "type": "string",
				  "computed": true
				},
				"engine_version": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "computed": true
				},
				"snapshot_type": {
				  "type": "string",
				  "computed": true
				},
				"source_db_cluster_snapshot_arn": {
				  "type": "string",
				  "computed": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"storage_encrypted": {
				  "type": "bool",
				  "computed": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_docdb_subnet_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_dx_bgp_peer": {
			"version": 0,
			"block": {
			  "attributes": {
				"address_family": {
				  "type": "string",
				  "required": true
				},
				"amazon_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"aws_device": {
				  "type": "string",
				  "computed": true
				},
				"bgp_asn": {
				  "type": "number",
				  "required": true
				},
				"bgp_auth_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"bgp_peer_id": {
				  "type": "string",
				  "computed": true
				},
				"bgp_status": {
				  "type": "string",
				  "computed": true
				},
				"customer_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"virtual_interface_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dx_connection": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"aws_device": {
				  "type": "string",
				  "computed": true
				},
				"bandwidth": {
				  "type": "string",
				  "required": true
				},
				"has_logical_redundancy": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"jumbo_frame_capable": {
				  "type": "bool",
				  "computed": true
				},
				"location": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_dx_connection_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"connection_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"lag_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_dx_gateway": {
			"version": 0,
			"block": {
			  "attributes": {
				"amazon_side_asn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"owner_account_id": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dx_gateway_association": {
			"version": 1,
			"block": {
			  "attributes": {
				"allowed_prefixes": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"associated_gateway_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"associated_gateway_owner_account_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"associated_gateway_type": {
				  "type": "string",
				  "computed": true
				},
				"dx_gateway_association_id": {
				  "type": "string",
				  "computed": true
				},
				"dx_gateway_id": {
				  "type": "string",
				  "required": true
				},
				"dx_gateway_owner_account_id": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"proposal_id": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dx_gateway_association_proposal": {
			"version": 0,
			"block": {
			  "attributes": {
				"allowed_prefixes": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"associated_gateway_id": {
				  "type": "string",
				  "required": true
				},
				"associated_gateway_owner_account_id": {
				  "type": "string",
				  "computed": true
				},
				"associated_gateway_type": {
				  "type": "string",
				  "computed": true
				},
				"dx_gateway_id": {
				  "type": "string",
				  "required": true
				},
				"dx_gateway_owner_account_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_dx_hosted_private_virtual_interface": {
			"version": 0,
			"block": {
			  "attributes": {
				"address_family": {
				  "type": "string",
				  "required": true
				},
				"amazon_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"amazon_side_asn": {
				  "type": "string",
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"aws_device": {
				  "type": "string",
				  "computed": true
				},
				"bgp_asn": {
				  "type": "number",
				  "required": true
				},
				"bgp_auth_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"connection_id": {
				  "type": "string",
				  "required": true
				},
				"customer_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"jumbo_frame_capable": {
				  "type": "bool",
				  "computed": true
				},
				"mtu": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"owner_account_id": {
				  "type": "string",
				  "required": true
				},
				"vlan": {
				  "type": "number",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dx_hosted_private_virtual_interface_accepter": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"dx_gateway_id": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"virtual_interface_id": {
				  "type": "string",
				  "required": true
				},
				"vpn_gateway_id": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dx_hosted_public_virtual_interface": {
			"version": 0,
			"block": {
			  "attributes": {
				"address_family": {
				  "type": "string",
				  "required": true
				},
				"amazon_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"amazon_side_asn": {
				  "type": "string",
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"aws_device": {
				  "type": "string",
				  "computed": true
				},
				"bgp_asn": {
				  "type": "number",
				  "required": true
				},
				"bgp_auth_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"connection_id": {
				  "type": "string",
				  "required": true
				},
				"customer_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"owner_account_id": {
				  "type": "string",
				  "required": true
				},
				"route_filter_prefixes": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"vlan": {
				  "type": "number",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dx_hosted_public_virtual_interface_accepter": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"virtual_interface_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dx_hosted_transit_virtual_interface": {
			"version": 0,
			"block": {
			  "attributes": {
				"address_family": {
				  "type": "string",
				  "required": true
				},
				"amazon_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"amazon_side_asn": {
				  "type": "string",
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"aws_device": {
				  "type": "string",
				  "computed": true
				},
				"bgp_asn": {
				  "type": "number",
				  "required": true
				},
				"bgp_auth_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"connection_id": {
				  "type": "string",
				  "required": true
				},
				"customer_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"jumbo_frame_capable": {
				  "type": "bool",
				  "computed": true
				},
				"mtu": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"owner_account_id": {
				  "type": "string",
				  "required": true
				},
				"vlan": {
				  "type": "number",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dx_hosted_transit_virtual_interface_accepter": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"dx_gateway_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"virtual_interface_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dx_lag": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"connections_bandwidth": {
				  "type": "string",
				  "required": true
				},
				"force_destroy": {
				  "type": "bool",
				  "optional": true
				},
				"has_logical_redundancy": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"jumbo_frame_capable": {
				  "type": "bool",
				  "computed": true
				},
				"location": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_dx_private_virtual_interface": {
			"version": 0,
			"block": {
			  "attributes": {
				"address_family": {
				  "type": "string",
				  "required": true
				},
				"amazon_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"amazon_side_asn": {
				  "type": "string",
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"aws_device": {
				  "type": "string",
				  "computed": true
				},
				"bgp_asn": {
				  "type": "number",
				  "required": true
				},
				"bgp_auth_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"connection_id": {
				  "type": "string",
				  "required": true
				},
				"customer_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"dx_gateway_id": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"jumbo_frame_capable": {
				  "type": "bool",
				  "computed": true
				},
				"mtu": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vlan": {
				  "type": "number",
				  "required": true
				},
				"vpn_gateway_id": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dx_public_virtual_interface": {
			"version": 0,
			"block": {
			  "attributes": {
				"address_family": {
				  "type": "string",
				  "required": true
				},
				"amazon_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"amazon_side_asn": {
				  "type": "string",
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"aws_device": {
				  "type": "string",
				  "computed": true
				},
				"bgp_asn": {
				  "type": "number",
				  "required": true
				},
				"bgp_auth_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"connection_id": {
				  "type": "string",
				  "required": true
				},
				"customer_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"route_filter_prefixes": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vlan": {
				  "type": "number",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dx_transit_virtual_interface": {
			"version": 0,
			"block": {
			  "attributes": {
				"address_family": {
				  "type": "string",
				  "required": true
				},
				"amazon_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"amazon_side_asn": {
				  "type": "string",
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"aws_device": {
				  "type": "string",
				  "computed": true
				},
				"bgp_asn": {
				  "type": "number",
				  "required": true
				},
				"bgp_auth_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"connection_id": {
				  "type": "string",
				  "required": true
				},
				"customer_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"dx_gateway_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"jumbo_frame_capable": {
				  "type": "bool",
				  "computed": true
				},
				"mtu": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vlan": {
				  "type": "number",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dynamodb_global_table": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"replica": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "region_name": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_dynamodb_table": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"billing_mode": {
				  "type": "string",
				  "optional": true
				},
				"hash_key": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"range_key": {
				  "type": "string",
				  "optional": true
				},
				"read_capacity": {
				  "type": "number",
				  "optional": true
				},
				"stream_arn": {
				  "type": "string",
				  "computed": true
				},
				"stream_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"stream_label": {
				  "type": "string",
				  "computed": true
				},
				"stream_view_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"write_capacity": {
				  "type": "number",
				  "optional": true
				}
			  },
			  "block_types": {
				"attribute": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "name": {
						"type": "string",
						"required": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1
				},
				"global_secondary_index": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "hash_key": {
						"type": "string",
						"required": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  },
					  "non_key_attributes": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "projection_type": {
						"type": "string",
						"required": true
					  },
					  "range_key": {
						"type": "string",
						"optional": true
					  },
					  "read_capacity": {
						"type": "number",
						"optional": true
					  },
					  "write_capacity": {
						"type": "number",
						"optional": true
					  }
					}
				  }
				},
				"local_secondary_index": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "name": {
						"type": "string",
						"required": true
					  },
					  "non_key_attributes": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "projection_type": {
						"type": "string",
						"required": true
					  },
					  "range_key": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"point_in_time_recovery": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"replica": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "region_name": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"server_side_encryption": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"required": true
					  },
					  "kms_key_arn": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"ttl": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "attribute_name": {
						"type": "string",
						"required": true
					  },
					  "enabled": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_dynamodb_table_item": {
			"version": 0,
			"block": {
			  "attributes": {
				"hash_key": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"item": {
				  "type": "string",
				  "required": true
				},
				"range_key": {
				  "type": "string",
				  "optional": true
				},
				"table_name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ebs_default_kms_key": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ebs_encryption_by_default": {
			"version": 0,
			"block": {
			  "attributes": {
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_ebs_snapshot": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"data_encryption_key_id": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"encrypted": {
				  "type": "bool",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "computed": true
				},
				"owner_alias": {
				  "type": "string",
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"volume_id": {
				  "type": "string",
				  "required": true
				},
				"volume_size": {
				  "type": "number",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_ebs_snapshot_copy": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"data_encryption_key_id": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"encrypted": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true
				},
				"owner_alias": {
				  "type": "string",
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"source_region": {
				  "type": "string",
				  "required": true
				},
				"source_snapshot_id": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"volume_id": {
				  "type": "string",
				  "computed": true
				},
				"volume_size": {
				  "type": "number",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_ebs_volume": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"availability_zone": {
				  "type": "string",
				  "required": true
				},
				"encrypted": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"iops": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"multi_attach_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"outpost_arn": {
				  "type": "string",
				  "optional": true
				},
				"size": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"snapshot_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_ec2_availability_zone_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"group_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"opt_in_status": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ec2_capacity_reservation": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"availability_zone": {
				  "type": "string",
				  "required": true
				},
				"ebs_optimized": {
				  "type": "bool",
				  "optional": true
				},
				"end_date": {
				  "type": "string",
				  "optional": true
				},
				"end_date_type": {
				  "type": "string",
				  "optional": true
				},
				"ephemeral_storage": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_count": {
				  "type": "number",
				  "required": true
				},
				"instance_match_criteria": {
				  "type": "string",
				  "optional": true
				},
				"instance_platform": {
				  "type": "string",
				  "required": true
				},
				"instance_type": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"tenancy": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_ec2_client_vpn_authorization_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"access_group_id": {
				  "type": "string",
				  "optional": true
				},
				"authorize_all_groups": {
				  "type": "bool",
				  "optional": true
				},
				"client_vpn_endpoint_id": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"target_network_cidr": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ec2_client_vpn_endpoint": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"client_cidr_block": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"dns_name": {
				  "type": "string",
				  "computed": true
				},
				"dns_servers": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"server_certificate_arn": {
				  "type": "string",
				  "required": true
				},
				"split_tunnel": {
				  "type": "bool",
				  "optional": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"transport_protocol": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"authentication_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "active_directory_id": {
						"type": "string",
						"optional": true
					  },
					  "root_certificate_chain_arn": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 2
				},
				"connection_log_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "cloudwatch_log_group": {
						"type": "string",
						"optional": true
					  },
					  "cloudwatch_log_stream": {
						"type": "string",
						"optional": true
					  },
					  "enabled": {
						"type": "bool",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_ec2_client_vpn_network_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"client_vpn_endpoint_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"security_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "computed": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"subnet_id": {
				  "type": "string",
				  "required": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_ec2_client_vpn_route": {
			"version": 0,
			"block": {
			  "attributes": {
				"client_vpn_endpoint_id": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"destination_cidr_block": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"origin": {
				  "type": "string",
				  "computed": true
				},
				"target_vpc_subnet_id": {
				  "type": "string",
				  "required": true
				},
				"type": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_ec2_fleet": {
			"version": 0,
			"block": {
			  "attributes": {
				"excess_capacity_termination_policy": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"replace_unhealthy_instances": {
				  "type": "bool",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"terminate_instances": {
				  "type": "bool",
				  "optional": true
				},
				"terminate_instances_with_expiration": {
				  "type": "bool",
				  "optional": true
				},
				"type": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"launch_template_config": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "launch_template_specification": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"launch_template_id": {
							  "type": "string",
							  "optional": true
							},
							"launch_template_name": {
							  "type": "string",
							  "optional": true
							},
							"version": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  },
					  "override": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"availability_zone": {
							  "type": "string",
							  "optional": true
							},
							"instance_type": {
							  "type": "string",
							  "optional": true
							},
							"max_price": {
							  "type": "string",
							  "optional": true
							},
							"priority": {
							  "type": "number",
							  "optional": true
							},
							"subnet_id": {
							  "type": "string",
							  "optional": true
							},
							"weighted_capacity": {
							  "type": "number",
							  "optional": true
							}
						  }
						},
						"max_items": 50
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"on_demand_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allocation_strategy": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"spot_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allocation_strategy": {
						"type": "string",
						"optional": true
					  },
					  "instance_interruption_behavior": {
						"type": "string",
						"optional": true
					  },
					  "instance_pools_to_use_count": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"target_capacity_specification": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "default_target_capacity_type": {
						"type": "string",
						"required": true
					  },
					  "on_demand_target_capacity": {
						"type": "number",
						"optional": true
					  },
					  "spot_target_capacity": {
						"type": "number",
						"optional": true
					  },
					  "total_target_capacity": {
						"type": "number",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_ec2_local_gateway_route": {
			"version": 0,
			"block": {
			  "attributes": {
				"destination_cidr_block": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"local_gateway_route_table_id": {
				  "type": "string",
				  "required": true
				},
				"local_gateway_virtual_interface_group_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ec2_local_gateway_route_table_vpc_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"local_gateway_id": {
				  "type": "string",
				  "computed": true
				},
				"local_gateway_route_table_id": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ec2_tag": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key": {
				  "type": "string",
				  "required": true
				},
				"resource_id": {
				  "type": "string",
				  "required": true
				},
				"value": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ec2_traffic_mirror_filter": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"network_services": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_ec2_traffic_mirror_filter_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"destination_cidr_block": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"protocol": {
				  "type": "number",
				  "optional": true
				},
				"rule_action": {
				  "type": "string",
				  "required": true
				},
				"rule_number": {
				  "type": "number",
				  "required": true
				},
				"source_cidr_block": {
				  "type": "string",
				  "required": true
				},
				"traffic_direction": {
				  "type": "string",
				  "required": true
				},
				"traffic_mirror_filter_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"destination_port_range": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "from_port": {
						"type": "number",
						"optional": true
					  },
					  "to_port": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"source_port_range": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "from_port": {
						"type": "number",
						"optional": true
					  },
					  "to_port": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_ec2_traffic_mirror_session": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"network_interface_id": {
				  "type": "string",
				  "required": true
				},
				"packet_length": {
				  "type": "number",
				  "optional": true
				},
				"session_number": {
				  "type": "number",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"traffic_mirror_filter_id": {
				  "type": "string",
				  "required": true
				},
				"traffic_mirror_target_id": {
				  "type": "string",
				  "required": true
				},
				"virtual_network_id": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_ec2_traffic_mirror_target": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"network_interface_id": {
				  "type": "string",
				  "optional": true
				},
				"network_load_balancer_arn": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_ec2_transit_gateway": {
			"version": 0,
			"block": {
			  "attributes": {
				"amazon_side_asn": {
				  "type": "number",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"association_default_route_table_id": {
				  "type": "string",
				  "computed": true
				},
				"auto_accept_shared_attachments": {
				  "type": "string",
				  "optional": true
				},
				"default_route_table_association": {
				  "type": "string",
				  "optional": true
				},
				"default_route_table_propagation": {
				  "type": "string",
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"dns_support": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"propagation_default_route_table_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpn_ecmp_support": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_ec2_transit_gateway_peering_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"peer_account_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"peer_region": {
				  "type": "string",
				  "required": true
				},
				"peer_transit_gateway_id": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"transit_gateway_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ec2_transit_gateway_peering_attachment_accepter": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"peer_account_id": {
				  "type": "string",
				  "computed": true
				},
				"peer_region": {
				  "type": "string",
				  "computed": true
				},
				"peer_transit_gateway_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"transit_gateway_attachment_id": {
				  "type": "string",
				  "required": true
				},
				"transit_gateway_id": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_ec2_transit_gateway_route": {
			"version": 0,
			"block": {
			  "attributes": {
				"blackhole": {
				  "type": "bool",
				  "optional": true
				},
				"destination_cidr_block": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"transit_gateway_attachment_id": {
				  "type": "string",
				  "optional": true
				},
				"transit_gateway_route_table_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ec2_transit_gateway_route_table": {
			"version": 0,
			"block": {
			  "attributes": {
				"default_association_route_table": {
				  "type": "bool",
				  "computed": true
				},
				"default_propagation_route_table": {
				  "type": "bool",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"transit_gateway_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ec2_transit_gateway_route_table_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"resource_id": {
				  "type": "string",
				  "computed": true
				},
				"resource_type": {
				  "type": "string",
				  "computed": true
				},
				"transit_gateway_attachment_id": {
				  "type": "string",
				  "required": true
				},
				"transit_gateway_route_table_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ec2_transit_gateway_route_table_propagation": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"resource_id": {
				  "type": "string",
				  "computed": true
				},
				"resource_type": {
				  "type": "string",
				  "computed": true
				},
				"transit_gateway_attachment_id": {
				  "type": "string",
				  "required": true
				},
				"transit_gateway_route_table_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ec2_transit_gateway_vpc_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"dns_support": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ipv6_support": {
				  "type": "string",
				  "optional": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"transit_gateway_default_route_table_association": {
				  "type": "bool",
				  "optional": true
				},
				"transit_gateway_default_route_table_propagation": {
				  "type": "bool",
				  "optional": true
				},
				"transit_gateway_id": {
				  "type": "string",
				  "required": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				},
				"vpc_owner_id": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_ec2_transit_gateway_vpc_attachment_accepter": {
			"version": 0,
			"block": {
			  "attributes": {
				"dns_support": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ipv6_support": {
				  "type": "string",
				  "computed": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"transit_gateway_attachment_id": {
				  "type": "string",
				  "required": true
				},
				"transit_gateway_default_route_table_association": {
				  "type": "bool",
				  "optional": true
				},
				"transit_gateway_default_route_table_propagation": {
				  "type": "bool",
				  "optional": true
				},
				"transit_gateway_id": {
				  "type": "string",
				  "computed": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				},
				"vpc_owner_id": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_ecr_lifecycle_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				},
				"registry_id": {
				  "type": "string",
				  "computed": true
				},
				"repository": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ecr_repository": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"image_tag_mutability": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"registry_id": {
				  "type": "string",
				  "computed": true
				},
				"repository_url": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"encryption_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "encryption_type": {
						"type": "string",
						"optional": true
					  },
					  "kms_key": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  }
				},
				"image_scanning_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "scan_on_push": {
						"type": "bool",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_ecr_repository_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				},
				"registry_id": {
				  "type": "string",
				  "computed": true
				},
				"repository": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ecs_capacity_provider": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"auto_scaling_group_provider": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "auto_scaling_group_arn": {
						"type": "string",
						"required": true
					  },
					  "managed_termination_protection": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					},
					"block_types": {
					  "managed_scaling": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"maximum_scaling_step_size": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"minimum_scaling_step_size": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"status": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"target_capacity": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_ecs_cluster": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"capacity_providers": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"default_capacity_provider_strategy": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "base": {
						"type": "number",
						"optional": true
					  },
					  "capacity_provider": {
						"type": "string",
						"required": true
					  },
					  "weight": {
						"type": "number",
						"optional": true
					  }
					}
				  }
				},
				"setting": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "name": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_ecs_service": {
			"version": 0,
			"block": {
			  "attributes": {
				"cluster": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"deployment_maximum_percent": {
				  "type": "number",
				  "optional": true
				},
				"deployment_minimum_healthy_percent": {
				  "type": "number",
				  "optional": true
				},
				"desired_count": {
				  "type": "number",
				  "optional": true
				},
				"enable_ecs_managed_tags": {
				  "type": "bool",
				  "optional": true
				},
				"force_new_deployment": {
				  "type": "bool",
				  "optional": true
				},
				"health_check_grace_period_seconds": {
				  "type": "number",
				  "optional": true
				},
				"iam_role": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"launch_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"platform_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"propagate_tags": {
				  "type": "string",
				  "optional": true
				},
				"scheduling_strategy": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"task_definition": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"capacity_provider_strategy": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "base": {
						"type": "number",
						"optional": true
					  },
					  "capacity_provider": {
						"type": "string",
						"required": true
					  },
					  "weight": {
						"type": "number",
						"optional": true
					  }
					}
				  }
				},
				"deployment_controller": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"load_balancer": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "container_name": {
						"type": "string",
						"required": true
					  },
					  "container_port": {
						"type": "number",
						"required": true
					  },
					  "elb_name": {
						"type": "string",
						"optional": true
					  },
					  "target_group_arn": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"network_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "assign_public_ip": {
						"type": "bool",
						"optional": true
					  },
					  "security_groups": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "subnets": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"ordered_placement_strategy": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "field": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 5
				},
				"placement_constraints": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "expression": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 10
				},
				"service_registries": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "container_name": {
						"type": "string",
						"optional": true
					  },
					  "container_port": {
						"type": "number",
						"optional": true
					  },
					  "port": {
						"type": "number",
						"optional": true
					  },
					  "registry_arn": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_ecs_task_definition": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"container_definitions": {
				  "type": "string",
				  "required": true
				},
				"cpu": {
				  "type": "string",
				  "optional": true
				},
				"execution_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"family": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ipc_mode": {
				  "type": "string",
				  "optional": true
				},
				"memory": {
				  "type": "string",
				  "optional": true
				},
				"network_mode": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"pid_mode": {
				  "type": "string",
				  "optional": true
				},
				"requires_compatibilities": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"revision": {
				  "type": "number",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"task_role_arn": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"inference_accelerator": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "device_name": {
						"type": "string",
						"required": true
					  },
					  "device_type": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"placement_constraints": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "expression": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 10
				},
				"proxy_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "container_name": {
						"type": "string",
						"required": true
					  },
					  "properties": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"volume": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "host_path": {
						"type": "string",
						"optional": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "docker_volume_configuration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"autoprovision": {
							  "type": "bool",
							  "optional": true
							},
							"driver": {
							  "type": "string",
							  "optional": true
							},
							"driver_opts": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"labels": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"scope": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							}
						  }
						},
						"max_items": 1
					  },
					  "efs_volume_configuration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"file_system_id": {
							  "type": "string",
							  "required": true
							},
							"root_directory": {
							  "type": "string",
							  "optional": true
							},
							"transit_encryption": {
							  "type": "string",
							  "optional": true
							},
							"transit_encryption_port": {
							  "type": "number",
							  "optional": true
							}
						  },
						  "block_types": {
							"authorization_config": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "access_point_id": {
									"type": "string",
									"optional": true
								  },
								  "iam": {
									"type": "string",
									"optional": true
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_efs_access_point": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"file_system_arn": {
				  "type": "string",
				  "computed": true
				},
				"file_system_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"posix_user": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "gid": {
						"type": "number",
						"required": true
					  },
					  "secondary_gids": {
						"type": [
						  "set",
						  "number"
						],
						"optional": true
					  },
					  "uid": {
						"type": "number",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"root_directory": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "path": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					},
					"block_types": {
					  "creation_info": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"owner_gid": {
							  "type": "number",
							  "required": true
							},
							"owner_uid": {
							  "type": "number",
							  "required": true
							},
							"permissions": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_efs_file_system": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"creation_token": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"dns_name": {
				  "type": "string",
				  "computed": true
				},
				"encrypted": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"performance_mode": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"provisioned_throughput_in_mibps": {
				  "type": "number",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"throughput_mode": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"lifecycle_policy": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "transition_to_ia": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_efs_file_system_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"file_system_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_efs_mount_target": {
			"version": 0,
			"block": {
			  "attributes": {
				"availability_zone_id": {
				  "type": "string",
				  "computed": true
				},
				"availability_zone_name": {
				  "type": "string",
				  "computed": true
				},
				"dns_name": {
				  "type": "string",
				  "computed": true
				},
				"file_system_arn": {
				  "type": "string",
				  "computed": true
				},
				"file_system_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ip_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"mount_target_dns_name": {
				  "type": "string",
				  "computed": true
				},
				"network_interface_id": {
				  "type": "string",
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"security_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"subnet_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_egress_only_internet_gateway": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_eip": {
			"version": 0,
			"block": {
			  "attributes": {
				"allocation_id": {
				  "type": "string",
				  "computed": true
				},
				"associate_with_private_ip": {
				  "type": "string",
				  "optional": true
				},
				"association_id": {
				  "type": "string",
				  "computed": true
				},
				"customer_owned_ip": {
				  "type": "string",
				  "computed": true
				},
				"customer_owned_ipv4_pool": {
				  "type": "string",
				  "optional": true
				},
				"domain": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"network_interface": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"private_dns": {
				  "type": "string",
				  "computed": true
				},
				"private_ip": {
				  "type": "string",
				  "computed": true
				},
				"public_dns": {
				  "type": "string",
				  "computed": true
				},
				"public_ip": {
				  "type": "string",
				  "computed": true
				},
				"public_ipv4_pool": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "read": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_eip_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"allocation_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"allow_reassociation": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"network_interface_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"private_ip_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"public_ip": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_eks_cluster": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"certificate_authority": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"data": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"created_at": {
				  "type": "string",
				  "computed": true
				},
				"enabled_cluster_log_types": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"endpoint": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identity": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"oidc": [
						  "list",
						  [
							"object",
							{
							  "issuer": "string"
							}
						  ]
						]
					  }
					]
				  ],
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"platform_version": {
				  "type": "string",
				  "computed": true
				},
				"role_arn": {
				  "type": "string",
				  "required": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"encryption_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "resources": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  }
					},
					"block_types": {
					  "provider": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"key_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"vpc_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "cluster_security_group_id": {
						"type": "string",
						"computed": true
					  },
					  "endpoint_private_access": {
						"type": "bool",
						"optional": true
					  },
					  "endpoint_public_access": {
						"type": "bool",
						"optional": true
					  },
					  "public_access_cidrs": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true,
						"computed": true
					  },
					  "security_group_ids": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "subnet_ids": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "vpc_id": {
						"type": "string",
						"computed": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_eks_fargate_profile": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"cluster_name": {
				  "type": "string",
				  "required": true
				},
				"fargate_profile_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"pod_execution_role_arn": {
				  "type": "string",
				  "required": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"selector": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "labels": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  },
					  "namespace": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_eks_node_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"ami_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"cluster_name": {
				  "type": "string",
				  "required": true
				},
				"disk_size": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"force_update_version": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_types": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"labels": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"node_group_name": {
				  "type": "string",
				  "required": true
				},
				"node_role_arn": {
				  "type": "string",
				  "required": true
				},
				"release_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"resources": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"autoscaling_groups": [
						  "list",
						  [
							"object",
							{
							  "name": "string"
							}
						  ]
						],
						"remote_access_security_group_id": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"remote_access": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "ec2_ssh_key": {
						"type": "string",
						"optional": true
					  },
					  "source_security_group_ids": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"scaling_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "desired_size": {
						"type": "number",
						"required": true
					  },
					  "max_size": {
						"type": "number",
						"required": true
					  },
					  "min_size": {
						"type": "number",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_elastic_beanstalk_application": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"appversion_lifecycle": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "delete_source_from_s3": {
						"type": "bool",
						"optional": true
					  },
					  "max_age_in_days": {
						"type": "number",
						"optional": true
					  },
					  "max_count": {
						"type": "number",
						"optional": true
					  },
					  "service_role": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_elastic_beanstalk_application_version": {
			"version": 0,
			"block": {
			  "attributes": {
				"application": {
				  "type": "string",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"bucket": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"force_delete": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_elastic_beanstalk_configuration_template": {
			"version": 0,
			"block": {
			  "attributes": {
				"application": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"environment_id": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"solution_stack_name": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"setting": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "name": {
						"type": "string",
						"required": true
					  },
					  "namespace": {
						"type": "string",
						"required": true
					  },
					  "resource": {
						"type": "string",
						"optional": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_elastic_beanstalk_environment": {
			"version": 1,
			"block": {
			  "attributes": {
				"all_settings": {
				  "type": [
					"set",
					[
					  "object",
					  {
						"name": "string",
						"namespace": "string",
						"resource": "string",
						"value": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"application": {
				  "type": "string",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"autoscaling_groups": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"cname": {
				  "type": "string",
				  "computed": true
				},
				"cname_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"endpoint_url": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instances": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"launch_configurations": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"load_balancers": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"platform_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"poll_interval": {
				  "type": "string",
				  "optional": true
				},
				"queues": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"solution_stack_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"template_name": {
				  "type": "string",
				  "optional": true
				},
				"tier": {
				  "type": "string",
				  "optional": true
				},
				"triggers": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"version_label": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"wait_for_ready_timeout": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"setting": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "name": {
						"type": "string",
						"required": true
					  },
					  "namespace": {
						"type": "string",
						"required": true
					  },
					  "resource": {
						"type": "string",
						"optional": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_elasticache_cluster": {
			"version": 0,
			"block": {
			  "attributes": {
				"apply_immediately": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"availability_zone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"az_mode": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cache_nodes": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"address": "string",
						"availability_zone": "string",
						"id": "string",
						"port": "number"
					  }
					]
				  ],
				  "computed": true
				},
				"cluster_address": {
				  "type": "string",
				  "computed": true
				},
				"cluster_id": {
				  "type": "string",
				  "required": true
				},
				"configuration_endpoint": {
				  "type": "string",
				  "computed": true
				},
				"engine": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"engine_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"maintenance_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"node_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"notification_topic_arn": {
				  "type": "string",
				  "optional": true
				},
				"num_cache_nodes": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"parameter_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"preferred_availability_zones": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"replication_group_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"security_group_names": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"snapshot_arns": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"snapshot_name": {
				  "type": "string",
				  "optional": true
				},
				"snapshot_retention_limit": {
				  "type": "number",
				  "optional": true
				},
				"snapshot_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"subnet_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_elasticache_parameter_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"family": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"parameter": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "name": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_elasticache_replication_group": {
			"version": 1,
			"block": {
			  "attributes": {
				"apply_immediately": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"at_rest_encryption_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"auth_token": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"auto_minor_version_upgrade": {
				  "type": "bool",
				  "optional": true
				},
				"automatic_failover_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"availability_zones": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"configuration_endpoint_address": {
				  "type": "string",
				  "computed": true
				},
				"engine": {
				  "type": "string",
				  "optional": true
				},
				"engine_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true
				},
				"maintenance_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"member_clusters": {
				  "type": [
					"set",
					"string"
				  ],
				  "computed": true
				},
				"node_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"notification_topic_arn": {
				  "type": "string",
				  "optional": true
				},
				"number_cache_clusters": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"parameter_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "optional": true
				},
				"primary_endpoint_address": {
				  "type": "string",
				  "computed": true
				},
				"replication_group_description": {
				  "type": "string",
				  "required": true
				},
				"replication_group_id": {
				  "type": "string",
				  "required": true
				},
				"security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"security_group_names": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"snapshot_arns": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"snapshot_name": {
				  "type": "string",
				  "optional": true
				},
				"snapshot_retention_limit": {
				  "type": "number",
				  "optional": true
				},
				"snapshot_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"subnet_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"transit_encryption_enabled": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"cluster_mode": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "num_node_groups": {
						"type": "number",
						"required": true
					  },
					  "replicas_per_node_group": {
						"type": "number",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_elasticache_security_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"security_group_names": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				}
			  }
			}
		  },
		  "aws_elasticache_subnet_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				}
			  }
			}
		  },
		  "aws_elasticsearch_domain": {
			"version": 0,
			"block": {
			  "attributes": {
				"access_policies": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"advanced_options": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"domain_id": {
				  "type": "string",
				  "computed": true
				},
				"domain_name": {
				  "type": "string",
				  "required": true
				},
				"elasticsearch_version": {
				  "type": "string",
				  "optional": true
				},
				"endpoint": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kibana_endpoint": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"advanced_security_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"required": true
					  },
					  "internal_user_database_enabled": {
						"type": "bool",
						"optional": true
					  }
					},
					"block_types": {
					  "master_user_options": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"master_user_arn": {
							  "type": "string",
							  "optional": true
							},
							"master_user_name": {
							  "type": "string",
							  "optional": true
							},
							"master_user_password": {
							  "type": "string",
							  "optional": true,
							  "sensitive": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"cluster_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "dedicated_master_count": {
						"type": "number",
						"optional": true
					  },
					  "dedicated_master_enabled": {
						"type": "bool",
						"optional": true
					  },
					  "dedicated_master_type": {
						"type": "string",
						"optional": true
					  },
					  "instance_count": {
						"type": "number",
						"optional": true
					  },
					  "instance_type": {
						"type": "string",
						"optional": true
					  },
					  "warm_count": {
						"type": "number",
						"optional": true
					  },
					  "warm_enabled": {
						"type": "bool",
						"optional": true
					  },
					  "warm_type": {
						"type": "string",
						"optional": true
					  },
					  "zone_awareness_enabled": {
						"type": "bool",
						"optional": true
					  }
					},
					"block_types": {
					  "zone_awareness_config": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"availability_zone_count": {
							  "type": "number",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"cognito_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"optional": true
					  },
					  "identity_pool_id": {
						"type": "string",
						"required": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "user_pool_id": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"domain_endpoint_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enforce_https": {
						"type": "bool",
						"required": true
					  },
					  "tls_security_policy": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				},
				"ebs_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "ebs_enabled": {
						"type": "bool",
						"required": true
					  },
					  "iops": {
						"type": "number",
						"optional": true
					  },
					  "volume_size": {
						"type": "number",
						"optional": true
					  },
					  "volume_type": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				},
				"encrypt_at_rest": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"required": true
					  },
					  "kms_key_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				},
				"log_publishing_options": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "cloudwatch_log_group_arn": {
						"type": "string",
						"required": true
					  },
					  "enabled": {
						"type": "bool",
						"optional": true
					  },
					  "log_type": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"node_to_node_encryption": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"snapshot_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "automated_snapshot_start_hour": {
						"type": "number",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"vpc_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "availability_zones": {
						"type": [
						  "set",
						  "string"
						],
						"computed": true
					  },
					  "security_group_ids": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "subnet_ids": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "vpc_id": {
						"type": "string",
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_elasticsearch_domain_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"access_policies": {
				  "type": "string",
				  "required": true
				},
				"domain_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_elastictranscoder_pipeline": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"aws_kms_key_arn": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"input_bucket": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"output_bucket": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"role": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"content_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "storage_class": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"content_config_permissions": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "access": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "grantee": {
						"type": "string",
						"optional": true
					  },
					  "grantee_type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"notifications": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "completed": {
						"type": "string",
						"optional": true
					  },
					  "error": {
						"type": "string",
						"optional": true
					  },
					  "progressing": {
						"type": "string",
						"optional": true
					  },
					  "warning": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"thumbnail_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "storage_class": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"thumbnail_config_permissions": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "access": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "grantee": {
						"type": "string",
						"optional": true
					  },
					  "grantee_type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_elastictranscoder_preset": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"container": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"video_codec_options": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"audio": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "audio_packing_mode": {
						"type": "string",
						"optional": true
					  },
					  "bit_rate": {
						"type": "string",
						"optional": true
					  },
					  "channels": {
						"type": "string",
						"optional": true
					  },
					  "codec": {
						"type": "string",
						"optional": true
					  },
					  "sample_rate": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"audio_codec_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bit_depth": {
						"type": "string",
						"optional": true
					  },
					  "bit_order": {
						"type": "string",
						"optional": true
					  },
					  "profile": {
						"type": "string",
						"optional": true
					  },
					  "signed": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"thumbnails": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "aspect_ratio": {
						"type": "string",
						"optional": true
					  },
					  "format": {
						"type": "string",
						"optional": true
					  },
					  "interval": {
						"type": "string",
						"optional": true
					  },
					  "max_height": {
						"type": "string",
						"optional": true
					  },
					  "max_width": {
						"type": "string",
						"optional": true
					  },
					  "padding_policy": {
						"type": "string",
						"optional": true
					  },
					  "resolution": {
						"type": "string",
						"optional": true
					  },
					  "sizing_policy": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"video": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "aspect_ratio": {
						"type": "string",
						"optional": true
					  },
					  "bit_rate": {
						"type": "string",
						"optional": true
					  },
					  "codec": {
						"type": "string",
						"optional": true
					  },
					  "display_aspect_ratio": {
						"type": "string",
						"optional": true
					  },
					  "fixed_gop": {
						"type": "string",
						"optional": true
					  },
					  "frame_rate": {
						"type": "string",
						"optional": true
					  },
					  "keyframes_max_dist": {
						"type": "string",
						"optional": true
					  },
					  "max_frame_rate": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "max_height": {
						"type": "string",
						"optional": true
					  },
					  "max_width": {
						"type": "string",
						"optional": true
					  },
					  "padding_policy": {
						"type": "string",
						"optional": true
					  },
					  "resolution": {
						"type": "string",
						"optional": true
					  },
					  "sizing_policy": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"video_watermarks": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "horizontal_align": {
						"type": "string",
						"optional": true
					  },
					  "horizontal_offset": {
						"type": "string",
						"optional": true
					  },
					  "id": {
						"type": "string",
						"optional": true
					  },
					  "max_height": {
						"type": "string",
						"optional": true
					  },
					  "max_width": {
						"type": "string",
						"optional": true
					  },
					  "opacity": {
						"type": "string",
						"optional": true
					  },
					  "sizing_policy": {
						"type": "string",
						"optional": true
					  },
					  "target": {
						"type": "string",
						"optional": true
					  },
					  "vertical_align": {
						"type": "string",
						"optional": true
					  },
					  "vertical_offset": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_elb": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"availability_zones": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"connection_draining": {
				  "type": "bool",
				  "optional": true
				},
				"connection_draining_timeout": {
				  "type": "number",
				  "optional": true
				},
				"cross_zone_load_balancing": {
				  "type": "bool",
				  "optional": true
				},
				"dns_name": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"idle_timeout": {
				  "type": "number",
				  "optional": true
				},
				"instances": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"internal": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"security_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"source_security_group": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"source_security_group_id": {
				  "type": "string",
				  "computed": true
				},
				"subnets": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"zone_id": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"access_logs": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket": {
						"type": "string",
						"required": true
					  },
					  "bucket_prefix": {
						"type": "string",
						"optional": true
					  },
					  "enabled": {
						"type": "bool",
						"optional": true
					  },
					  "interval": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"health_check": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "healthy_threshold": {
						"type": "number",
						"required": true
					  },
					  "interval": {
						"type": "number",
						"required": true
					  },
					  "target": {
						"type": "string",
						"required": true
					  },
					  "timeout": {
						"type": "number",
						"required": true
					  },
					  "unhealthy_threshold": {
						"type": "number",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"listener": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "instance_port": {
						"type": "number",
						"required": true
					  },
					  "instance_protocol": {
						"type": "string",
						"required": true
					  },
					  "lb_port": {
						"type": "number",
						"required": true
					  },
					  "lb_protocol": {
						"type": "string",
						"required": true
					  },
					  "ssl_certificate_id": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "min_items": 1
				}
			  }
			}
		  },
		  "aws_elb_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"elb": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_emr_cluster": {
			"version": 0,
			"block": {
			  "attributes": {
				"additional_info": {
				  "type": "string",
				  "optional": true
				},
				"applications": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"autoscaling_role": {
				  "type": "string",
				  "optional": true
				},
				"cluster_state": {
				  "type": "string",
				  "computed": true
				},
				"configurations": {
				  "type": "string",
				  "optional": true
				},
				"configurations_json": {
				  "type": "string",
				  "optional": true
				},
				"custom_ami_id": {
				  "type": "string",
				  "optional": true
				},
				"ebs_root_volume_size": {
				  "type": "number",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"keep_job_flow_alive_when_no_steps": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"log_uri": {
				  "type": "string",
				  "optional": true
				},
				"master_public_dns": {
				  "type": "string",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"release_label": {
				  "type": "string",
				  "required": true
				},
				"scale_down_behavior": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"security_configuration": {
				  "type": "string",
				  "optional": true
				},
				"service_role": {
				  "type": "string",
				  "required": true
				},
				"step": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"action_on_failure": "string",
						"hadoop_jar_step": [
						  "list",
						  [
							"object",
							{
							  "args": [
								"list",
								"string"
							  ],
							  "jar": "string",
							  "main_class": "string",
							  "properties": [
								"map",
								"string"
							  ]
							}
						  ]
						],
						"name": "string"
					  }
					]
				  ],
				  "optional": true,
				  "computed": true
				},
				"step_concurrency_level": {
				  "type": "number",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"termination_protection": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"visible_to_all_users": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"bootstrap_action": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "args": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  },
					  "path": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"core_instance_group": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "autoscaling_policy": {
						"type": "string",
						"optional": true
					  },
					  "bid_price": {
						"type": "string",
						"optional": true
					  },
					  "id": {
						"type": "string",
						"computed": true
					  },
					  "instance_count": {
						"type": "number",
						"optional": true
					  },
					  "instance_type": {
						"type": "string",
						"required": true
					  },
					  "name": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "ebs_config": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"iops": {
							  "type": "number",
							  "optional": true
							},
							"size": {
							  "type": "number",
							  "required": true
							},
							"type": {
							  "type": "string",
							  "required": true
							},
							"volumes_per_instance": {
							  "type": "number",
							  "optional": true
							}
						  }
						}
					  }
					}
				  },
				  "max_items": 1
				},
				"ec2_attributes": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "additional_master_security_groups": {
						"type": "string",
						"optional": true
					  },
					  "additional_slave_security_groups": {
						"type": "string",
						"optional": true
					  },
					  "emr_managed_master_security_group": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "emr_managed_slave_security_group": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "instance_profile": {
						"type": "string",
						"required": true
					  },
					  "key_name": {
						"type": "string",
						"optional": true
					  },
					  "service_access_security_group": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "subnet_id": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"kerberos_attributes": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "ad_domain_join_password": {
						"type": "string",
						"optional": true,
						"sensitive": true
					  },
					  "ad_domain_join_user": {
						"type": "string",
						"optional": true
					  },
					  "cross_realm_trust_principal_password": {
						"type": "string",
						"optional": true,
						"sensitive": true
					  },
					  "kdc_admin_password": {
						"type": "string",
						"required": true,
						"sensitive": true
					  },
					  "realm": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"master_instance_group": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bid_price": {
						"type": "string",
						"optional": true
					  },
					  "id": {
						"type": "string",
						"computed": true
					  },
					  "instance_count": {
						"type": "number",
						"optional": true
					  },
					  "instance_type": {
						"type": "string",
						"required": true
					  },
					  "name": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "ebs_config": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"iops": {
							  "type": "number",
							  "optional": true
							},
							"size": {
							  "type": "number",
							  "required": true
							},
							"type": {
							  "type": "string",
							  "required": true
							},
							"volumes_per_instance": {
							  "type": "number",
							  "optional": true
							}
						  }
						}
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_emr_instance_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"autoscaling_policy": {
				  "type": "string",
				  "optional": true
				},
				"bid_price": {
				  "type": "string",
				  "optional": true
				},
				"cluster_id": {
				  "type": "string",
				  "required": true
				},
				"configurations_json": {
				  "type": "string",
				  "optional": true
				},
				"ebs_optimized": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_count": {
				  "type": "number",
				  "optional": true
				},
				"instance_type": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"running_instance_count": {
				  "type": "number",
				  "computed": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"ebs_config": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "iops": {
						"type": "number",
						"optional": true
					  },
					  "size": {
						"type": "number",
						"required": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  },
					  "volumes_per_instance": {
						"type": "number",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_emr_security_configuration": {
			"version": 0,
			"block": {
			  "attributes": {
				"configuration": {
				  "type": "string",
				  "required": true
				},
				"creation_date": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_flow_log": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"eni_id": {
				  "type": "string",
				  "optional": true
				},
				"iam_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"log_destination": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"log_destination_type": {
				  "type": "string",
				  "optional": true
				},
				"log_format": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"log_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"max_aggregation_interval": {
				  "type": "number",
				  "optional": true
				},
				"subnet_id": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"traffic_type": {
				  "type": "string",
				  "required": true
				},
				"vpc_id": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_fms_admin_account": {
			"version": 0,
			"block": {
			  "attributes": {
				"account_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_fsx_lustre_file_system": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"deployment_type": {
				  "type": "string",
				  "optional": true
				},
				"dns_name": {
				  "type": "string",
				  "computed": true
				},
				"export_path": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"import_path": {
				  "type": "string",
				  "optional": true
				},
				"imported_file_chunk_size": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"network_interface_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"per_unit_storage_throughput": {
				  "type": "number",
				  "optional": true
				},
				"security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"storage_capacity": {
				  "type": "number",
				  "required": true
				},
				"subnet_ids": {
				  "type": [
					"list",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				},
				"weekly_maintenance_start_time": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_fsx_windows_file_system": {
			"version": 0,
			"block": {
			  "attributes": {
				"active_directory_id": {
				  "type": "string",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"automatic_backup_retention_days": {
				  "type": "number",
				  "optional": true
				},
				"copy_tags_to_backups": {
				  "type": "bool",
				  "optional": true
				},
				"daily_automatic_backup_start_time": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"deployment_type": {
				  "type": "string",
				  "optional": true
				},
				"dns_name": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"network_interface_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"preferred_file_server_ip": {
				  "type": "string",
				  "computed": true
				},
				"preferred_subnet_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"remote_administration_endpoint": {
				  "type": "string",
				  "computed": true
				},
				"security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"skip_final_backup": {
				  "type": "bool",
				  "optional": true
				},
				"storage_capacity": {
				  "type": "number",
				  "required": true
				},
				"storage_type": {
				  "type": "string",
				  "optional": true
				},
				"subnet_ids": {
				  "type": [
					"list",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"throughput_capacity": {
				  "type": "number",
				  "required": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				},
				"weekly_maintenance_start_time": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"self_managed_active_directory": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "dns_ips": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "domain_name": {
						"type": "string",
						"required": true
					  },
					  "file_system_administrators_group": {
						"type": "string",
						"optional": true
					  },
					  "organizational_unit_distinguished_name": {
						"type": "string",
						"optional": true
					  },
					  "password": {
						"type": "string",
						"required": true,
						"sensitive": true
					  },
					  "username": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_gamelift_alias": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"routing_strategy": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "fleet_id": {
						"type": "string",
						"optional": true
					  },
					  "message": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_gamelift_build": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"operating_system": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"version": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"storage_location": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket": {
						"type": "string",
						"required": true
					  },
					  "key": {
						"type": "string",
						"required": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_gamelift_fleet": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"build_id": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"ec2_instance_type": {
				  "type": "string",
				  "required": true
				},
				"fleet_type": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"log_paths": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"metric_groups": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"new_game_session_protection_policy": {
				  "type": "string",
				  "optional": true
				},
				"operating_system": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"ec2_inbound_permission": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "from_port": {
						"type": "number",
						"required": true
					  },
					  "ip_range": {
						"type": "string",
						"required": true
					  },
					  "protocol": {
						"type": "string",
						"required": true
					  },
					  "to_port": {
						"type": "number",
						"required": true
					  }
					}
				  },
				  "max_items": 50
				},
				"resource_creation_limit_policy": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "new_game_sessions_per_creator": {
						"type": "number",
						"optional": true
					  },
					  "policy_period_in_minutes": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"runtime_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "game_session_activation_timeout_seconds": {
						"type": "number",
						"optional": true
					  },
					  "max_concurrent_game_session_activations": {
						"type": "number",
						"optional": true
					  }
					},
					"block_types": {
					  "server_process": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"concurrent_executions": {
							  "type": "number",
							  "required": true
							},
							"launch_path": {
							  "type": "string",
							  "required": true
							},
							"parameters": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 50
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_gamelift_game_session_queue": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"destinations": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"timeout_in_seconds": {
				  "type": "number",
				  "optional": true
				}
			  },
			  "block_types": {
				"player_latency_policy": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "maximum_individual_player_latency_milliseconds": {
						"type": "number",
						"required": true
					  },
					  "policy_duration_seconds": {
						"type": "number",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_glacier_vault": {
			"version": 0,
			"block": {
			  "attributes": {
				"access_policy": {
				  "type": "string",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"location": {
				  "type": "string",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"notification": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "events": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "sns_topic": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_glacier_vault_lock": {
			"version": 0,
			"block": {
			  "attributes": {
				"complete_lock": {
				  "type": "bool",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ignore_deletion_error": {
				  "type": "bool",
				  "optional": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				},
				"vault_name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_globalaccelerator_accelerator": {
			"version": 0,
			"block": {
			  "attributes": {
				"dns_name": {
				  "type": "string",
				  "computed": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"hosted_zone_id": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ip_address_type": {
				  "type": "string",
				  "optional": true
				},
				"ip_sets": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"ip_addresses": [
						  "list",
						  "string"
						],
						"ip_family": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"attributes": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "flow_logs_enabled": {
						"type": "bool",
						"optional": true
					  },
					  "flow_logs_s3_bucket": {
						"type": "string",
						"optional": true
					  },
					  "flow_logs_s3_prefix": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_globalaccelerator_endpoint_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"endpoint_group_region": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"health_check_interval_seconds": {
				  "type": "number",
				  "optional": true
				},
				"health_check_path": {
				  "type": "string",
				  "optional": true
				},
				"health_check_port": {
				  "type": "number",
				  "optional": true
				},
				"health_check_protocol": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"listener_arn": {
				  "type": "string",
				  "required": true
				},
				"threshold_count": {
				  "type": "number",
				  "optional": true
				},
				"traffic_dial_percentage": {
				  "type": "number",
				  "optional": true
				}
			  },
			  "block_types": {
				"endpoint_configuration": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "endpoint_id": {
						"type": "string",
						"optional": true
					  },
					  "weight": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 10
				}
			  }
			}
		  },
		  "aws_globalaccelerator_listener": {
			"version": 0,
			"block": {
			  "attributes": {
				"accelerator_arn": {
				  "type": "string",
				  "required": true
				},
				"client_affinity": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"protocol": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"port_range": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "from_port": {
						"type": "number",
						"optional": true
					  },
					  "to_port": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 10
				}
			  }
			}
		  },
		  "aws_glue_catalog_database": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"catalog_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"location_uri": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"parameters": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_glue_catalog_table": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"catalog_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"database_name": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"owner": {
				  "type": "string",
				  "optional": true
				},
				"parameters": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"retention": {
				  "type": "number",
				  "optional": true
				},
				"table_type": {
				  "type": "string",
				  "optional": true
				},
				"view_expanded_text": {
				  "type": "string",
				  "optional": true
				},
				"view_original_text": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"partition_keys": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "comment": {
						"type": "string",
						"optional": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"storage_descriptor": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket_columns": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "compressed": {
						"type": "bool",
						"optional": true
					  },
					  "input_format": {
						"type": "string",
						"optional": true
					  },
					  "location": {
						"type": "string",
						"optional": true
					  },
					  "number_of_buckets": {
						"type": "number",
						"optional": true
					  },
					  "output_format": {
						"type": "string",
						"optional": true
					  },
					  "parameters": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  },
					  "stored_as_sub_directories": {
						"type": "bool",
						"optional": true
					  }
					},
					"block_types": {
					  "columns": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"comment": {
							  "type": "string",
							  "optional": true
							},
							"name": {
							  "type": "string",
							  "required": true
							},
							"type": {
							  "type": "string",
							  "optional": true
							}
						  }
						}
					  },
					  "ser_de_info": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"name": {
							  "type": "string",
							  "optional": true
							},
							"parameters": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"serialization_library": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "skewed_info": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"skewed_column_names": {
							  "type": [
								"list",
								"string"
							  ],
							  "optional": true
							},
							"skewed_column_value_location_maps": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"skewed_column_values": {
							  "type": [
								"list",
								"string"
							  ],
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "sort_columns": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"column": {
							  "type": "string",
							  "required": true
							},
							"sort_order": {
							  "type": "number",
							  "required": true
							}
						  }
						}
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_glue_classifier": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"csv_classifier": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allow_single_column": {
						"type": "bool",
						"optional": true
					  },
					  "contains_header": {
						"type": "string",
						"optional": true
					  },
					  "delimiter": {
						"type": "string",
						"optional": true
					  },
					  "disable_value_trimming": {
						"type": "bool",
						"optional": true
					  },
					  "header": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "quote_symbol": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"grok_classifier": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "classification": {
						"type": "string",
						"required": true
					  },
					  "custom_patterns": {
						"type": "string",
						"optional": true
					  },
					  "grok_pattern": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"json_classifier": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "json_path": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"xml_classifier": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "classification": {
						"type": "string",
						"required": true
					  },
					  "row_tag": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_glue_connection": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"catalog_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"connection_properties": {
				  "type": [
					"map",
					"string"
				  ],
				  "required": true,
				  "sensitive": true
				},
				"connection_type": {
				  "type": "string",
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"match_criteria": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"physical_connection_requirements": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "availability_zone": {
						"type": "string",
						"optional": true
					  },
					  "security_group_id_list": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "subnet_id": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_glue_crawler": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"classifiers": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"configuration": {
				  "type": "string",
				  "optional": true
				},
				"database_name": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"role": {
				  "type": "string",
				  "required": true
				},
				"schedule": {
				  "type": "string",
				  "optional": true
				},
				"security_configuration": {
				  "type": "string",
				  "optional": true
				},
				"table_prefix": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"catalog_target": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "database_name": {
						"type": "string",
						"required": true
					  },
					  "tables": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  }
					}
				  }
				},
				"dynamodb_target": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "path": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"jdbc_target": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "connection_name": {
						"type": "string",
						"required": true
					  },
					  "exclusions": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "path": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"s3_target": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "exclusions": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "path": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"schema_change_policy": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "delete_behavior": {
						"type": "string",
						"optional": true
					  },
					  "update_behavior": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_glue_job": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"connections": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"default_arguments": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"glue_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"max_capacity": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"max_retries": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"number_of_workers": {
				  "type": "number",
				  "optional": true
				},
				"role_arn": {
				  "type": "string",
				  "required": true
				},
				"security_configuration": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"timeout": {
				  "type": "number",
				  "optional": true
				},
				"worker_type": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"command": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "name": {
						"type": "string",
						"optional": true
					  },
					  "python_version": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "script_location": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"execution_property": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "max_concurrent_runs": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"notification_property": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "notify_delay_after": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_glue_security_configuration": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"encryption_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "cloudwatch_encryption": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"cloudwatch_encryption_mode": {
							  "type": "string",
							  "optional": true
							},
							"kms_key_arn": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  },
					  "job_bookmarks_encryption": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"job_bookmarks_encryption_mode": {
							  "type": "string",
							  "optional": true
							},
							"kms_key_arn": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  },
					  "s3_encryption": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"kms_key_arn": {
							  "type": "string",
							  "optional": true
							},
							"s3_encryption_mode": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_glue_trigger": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"schedule": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"type": {
				  "type": "string",
				  "required": true
				},
				"workflow_name": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"actions": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "arguments": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  },
					  "crawler_name": {
						"type": "string",
						"optional": true
					  },
					  "job_name": {
						"type": "string",
						"optional": true
					  },
					  "timeout": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "min_items": 1
				},
				"predicate": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "logical": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "conditions": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"crawl_state": {
							  "type": "string",
							  "optional": true
							},
							"crawler_name": {
							  "type": "string",
							  "optional": true
							},
							"job_name": {
							  "type": "string",
							  "optional": true
							},
							"logical_operator": {
							  "type": "string",
							  "optional": true
							},
							"state": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"min_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_glue_workflow": {
			"version": 0,
			"block": {
			  "attributes": {
				"default_run_properties": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_guardduty_detector": {
			"version": 0,
			"block": {
			  "attributes": {
				"account_id": {
				  "type": "string",
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"enable": {
				  "type": "bool",
				  "optional": true
				},
				"finding_publishing_frequency": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_guardduty_invite_accepter": {
			"version": 0,
			"block": {
			  "attributes": {
				"detector_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"master_account_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_guardduty_ipset": {
			"version": 0,
			"block": {
			  "attributes": {
				"activate": {
				  "type": "bool",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"detector_id": {
				  "type": "string",
				  "required": true
				},
				"format": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"location": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_guardduty_member": {
			"version": 0,
			"block": {
			  "attributes": {
				"account_id": {
				  "type": "string",
				  "required": true
				},
				"detector_id": {
				  "type": "string",
				  "required": true
				},
				"disable_email_notification": {
				  "type": "bool",
				  "optional": true
				},
				"email": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"invitation_message": {
				  "type": "string",
				  "optional": true
				},
				"invite": {
				  "type": "bool",
				  "optional": true
				},
				"relationship_status": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_guardduty_organization_admin_account": {
			"version": 0,
			"block": {
			  "attributes": {
				"admin_account_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_guardduty_organization_configuration": {
			"version": 0,
			"block": {
			  "attributes": {
				"auto_enable": {
				  "type": "bool",
				  "required": true
				},
				"detector_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_guardduty_threatintelset": {
			"version": 0,
			"block": {
			  "attributes": {
				"activate": {
				  "type": "bool",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"detector_id": {
				  "type": "string",
				  "required": true
				},
				"format": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"location": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_iam_access_key": {
			"version": 0,
			"block": {
			  "attributes": {
				"encrypted_secret": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key_fingerprint": {
				  "type": "string",
				  "computed": true
				},
				"pgp_key": {
				  "type": "string",
				  "optional": true
				},
				"secret": {
				  "type": "string",
				  "computed": true,
				  "sensitive": true
				},
				"ses_smtp_password_v4": {
				  "type": "string",
				  "computed": true,
				  "sensitive": true
				},
				"status": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"user": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iam_account_alias": {
			"version": 0,
			"block": {
			  "attributes": {
				"account_alias": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_iam_account_password_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"allow_users_to_change_password": {
				  "type": "bool",
				  "optional": true
				},
				"expire_passwords": {
				  "type": "bool",
				  "computed": true
				},
				"hard_expiry": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"max_password_age": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"minimum_password_length": {
				  "type": "number",
				  "optional": true
				},
				"password_reuse_prevention": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"require_lowercase_characters": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"require_numbers": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"require_symbols": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"require_uppercase_characters": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_iam_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"path": {
				  "type": "string",
				  "optional": true
				},
				"unique_id": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_iam_group_membership": {
			"version": 0,
			"block": {
			  "attributes": {
				"group": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"users": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				}
			  }
			}
		  },
		  "aws_iam_group_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"group": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iam_group_policy_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"group": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iam_instance_profile": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"create_date": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"path": {
				  "type": "string",
				  "optional": true
				},
				"role": {
				  "type": "string",
				  "optional": true
				},
				"unique_id": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_iam_openid_connect_provider": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"client_id_list": {
				  "type": [
					"list",
					"string"
				  ],
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"thumbprint_list": {
				  "type": [
					"list",
					"string"
				  ],
				  "required": true
				},
				"url": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iam_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"path": {
				  "type": "string",
				  "optional": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iam_policy_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"policy_arn": {
				  "type": "string",
				  "required": true
				},
				"roles": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"users": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_iam_role": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"assume_role_policy": {
				  "type": "string",
				  "required": true
				},
				"create_date": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"force_detach_policies": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"max_session_duration": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"path": {
				  "type": "string",
				  "optional": true
				},
				"permissions_boundary": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"unique_id": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_iam_role_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				},
				"role": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iam_role_policy_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy_arn": {
				  "type": "string",
				  "required": true
				},
				"role": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iam_saml_provider": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"saml_metadata_document": {
				  "type": "string",
				  "required": true
				},
				"valid_until": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_iam_server_certificate": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"certificate_body": {
				  "type": "string",
				  "required": true
				},
				"certificate_chain": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"path": {
				  "type": "string",
				  "optional": true
				},
				"private_key": {
				  "type": "string",
				  "required": true,
				  "sensitive": true
				}
			  }
			}
		  },
		  "aws_iam_service_linked_role": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"aws_service_name": {
				  "type": "string",
				  "required": true
				},
				"create_date": {
				  "type": "string",
				  "computed": true
				},
				"custom_suffix": {
				  "type": "string",
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "computed": true
				},
				"path": {
				  "type": "string",
				  "computed": true
				},
				"unique_id": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_iam_user": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"force_destroy": {
				  "type": "bool",
				  "description": "Delete user even if it has non-Terraform-managed IAM access keys, login profile or MFA devices",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"path": {
				  "type": "string",
				  "optional": true
				},
				"permissions_boundary": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"unique_id": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_iam_user_group_membership": {
			"version": 0,
			"block": {
			  "attributes": {
				"groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"user": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iam_user_login_profile": {
			"version": 0,
			"block": {
			  "attributes": {
				"encrypted_password": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key_fingerprint": {
				  "type": "string",
				  "computed": true
				},
				"password_length": {
				  "type": "number",
				  "optional": true
				},
				"password_reset_required": {
				  "type": "bool",
				  "optional": true
				},
				"pgp_key": {
				  "type": "string",
				  "required": true
				},
				"user": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iam_user_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				},
				"user": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iam_user_policy_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy_arn": {
				  "type": "string",
				  "required": true
				},
				"user": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iam_user_ssh_key": {
			"version": 0,
			"block": {
			  "attributes": {
				"encoding": {
				  "type": "string",
				  "required": true
				},
				"fingerprint": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"public_key": {
				  "type": "string",
				  "required": true
				},
				"ssh_public_key_id": {
				  "type": "string",
				  "computed": true
				},
				"status": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"username": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_inspector_assessment_target": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"resource_group_arn": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_inspector_assessment_template": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"duration": {
				  "type": "number",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"rules_package_arns": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"target_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_inspector_resource_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "required": true
				}
			  }
			}
		  },
		  "aws_instance": {
			"version": 1,
			"block": {
			  "attributes": {
				"ami": {
				  "type": "string",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"associate_public_ip_address": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"availability_zone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cpu_core_count": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"cpu_threads_per_core": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"disable_api_termination": {
				  "type": "bool",
				  "optional": true
				},
				"ebs_optimized": {
				  "type": "bool",
				  "optional": true
				},
				"get_password_data": {
				  "type": "bool",
				  "optional": true
				},
				"hibernation": {
				  "type": "bool",
				  "optional": true
				},
				"host_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"iam_instance_profile": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_initiated_shutdown_behavior": {
				  "type": "string",
				  "optional": true
				},
				"instance_state": {
				  "type": "string",
				  "computed": true
				},
				"instance_type": {
				  "type": "string",
				  "required": true
				},
				"ipv6_address_count": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"ipv6_addresses": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"key_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"monitoring": {
				  "type": "bool",
				  "optional": true
				},
				"outpost_arn": {
				  "type": "string",
				  "computed": true
				},
				"password_data": {
				  "type": "string",
				  "computed": true
				},
				"placement_group": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"primary_network_interface_id": {
				  "type": "string",
				  "computed": true
				},
				"private_dns": {
				  "type": "string",
				  "computed": true
				},
				"private_ip": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"public_dns": {
				  "type": "string",
				  "computed": true
				},
				"public_ip": {
				  "type": "string",
				  "computed": true
				},
				"secondary_private_ips": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"security_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"source_dest_check": {
				  "type": "bool",
				  "optional": true
				},
				"subnet_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"tenancy": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"user_data": {
				  "type": "string",
				  "optional": true
				},
				"user_data_base64": {
				  "type": "string",
				  "optional": true
				},
				"volume_tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"vpc_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"credit_specification": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "cpu_credits": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"ebs_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "delete_on_termination": {
						"type": "bool",
						"optional": true
					  },
					  "device_name": {
						"type": "string",
						"required": true
					  },
					  "encrypted": {
						"type": "bool",
						"optional": true,
						"computed": true
					  },
					  "iops": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "kms_key_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "snapshot_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "volume_id": {
						"type": "string",
						"computed": true
					  },
					  "volume_size": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "volume_type": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  }
				},
				"ephemeral_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "device_name": {
						"type": "string",
						"required": true
					  },
					  "no_device": {
						"type": "bool",
						"optional": true
					  },
					  "virtual_name": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"metadata_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "http_endpoint": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "http_put_response_hop_limit": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "http_tokens": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				},
				"network_interface": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "delete_on_termination": {
						"type": "bool",
						"optional": true
					  },
					  "device_index": {
						"type": "number",
						"required": true
					  },
					  "network_interface_id": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"root_block_device": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "delete_on_termination": {
						"type": "bool",
						"optional": true
					  },
					  "device_name": {
						"type": "string",
						"computed": true
					  },
					  "encrypted": {
						"type": "bool",
						"optional": true,
						"computed": true
					  },
					  "iops": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "kms_key_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "volume_id": {
						"type": "string",
						"computed": true
					  },
					  "volume_size": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "volume_type": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_internet_gateway": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_iot_certificate": {
			"version": 0,
			"block": {
			  "attributes": {
				"active": {
				  "type": "bool",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"certificate_pem": {
				  "type": "string",
				  "computed": true,
				  "sensitive": true
				},
				"csr": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"private_key": {
				  "type": "string",
				  "computed": true,
				  "sensitive": true
				},
				"public_key": {
				  "type": "string",
				  "computed": true,
				  "sensitive": true
				}
			  }
			}
		  },
		  "aws_iot_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"default_version_id": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iot_policy_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				},
				"target": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iot_role_alias": {
			"version": 0,
			"block": {
			  "attributes": {
				"alias": {
				  "type": "string",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"credential_duration": {
				  "type": "number",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"role_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iot_thing": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"attributes": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"default_client_id": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"thing_type_name": {
				  "type": "string",
				  "optional": true
				},
				"version": {
				  "type": "number",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_iot_thing_principal_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"principal": {
				  "type": "string",
				  "required": true
				},
				"thing": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_iot_thing_type": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"deprecated": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"properties": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "description": {
						"type": "string",
						"optional": true
					  },
					  "searchable_attributes": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_iot_topic_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"enabled": {
				  "type": "bool",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"sql": {
				  "type": "string",
				  "required": true
				},
				"sql_version": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"cloudwatch_alarm": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "alarm_name": {
						"type": "string",
						"required": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "state_reason": {
						"type": "string",
						"required": true
					  },
					  "state_value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"cloudwatch_metric": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "metric_name": {
						"type": "string",
						"required": true
					  },
					  "metric_namespace": {
						"type": "string",
						"required": true
					  },
					  "metric_timestamp": {
						"type": "string",
						"optional": true
					  },
					  "metric_unit": {
						"type": "string",
						"required": true
					  },
					  "metric_value": {
						"type": "string",
						"required": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"dynamodb": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "hash_key_field": {
						"type": "string",
						"required": true
					  },
					  "hash_key_type": {
						"type": "string",
						"optional": true
					  },
					  "hash_key_value": {
						"type": "string",
						"required": true
					  },
					  "operation": {
						"type": "string",
						"optional": true
					  },
					  "payload_field": {
						"type": "string",
						"optional": true
					  },
					  "range_key_field": {
						"type": "string",
						"optional": true
					  },
					  "range_key_type": {
						"type": "string",
						"optional": true
					  },
					  "range_key_value": {
						"type": "string",
						"optional": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "table_name": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"dynamodbv2": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "role_arn": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "put_item": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"table_name": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  }
				},
				"elasticsearch": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "endpoint": {
						"type": "string",
						"required": true
					  },
					  "id": {
						"type": "string",
						"required": true
					  },
					  "index": {
						"type": "string",
						"required": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"error_action": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "cloudwatch_alarm": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"alarm_name": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							},
							"state_reason": {
							  "type": "string",
							  "required": true
							},
							"state_value": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "cloudwatch_metric": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"metric_name": {
							  "type": "string",
							  "required": true
							},
							"metric_namespace": {
							  "type": "string",
							  "required": true
							},
							"metric_timestamp": {
							  "type": "string",
							  "optional": true
							},
							"metric_unit": {
							  "type": "string",
							  "required": true
							},
							"metric_value": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "dynamodb": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"hash_key_field": {
							  "type": "string",
							  "required": true
							},
							"hash_key_type": {
							  "type": "string",
							  "optional": true
							},
							"hash_key_value": {
							  "type": "string",
							  "required": true
							},
							"operation": {
							  "type": "string",
							  "optional": true
							},
							"payload_field": {
							  "type": "string",
							  "optional": true
							},
							"range_key_field": {
							  "type": "string",
							  "optional": true
							},
							"range_key_type": {
							  "type": "string",
							  "optional": true
							},
							"range_key_value": {
							  "type": "string",
							  "optional": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							},
							"table_name": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "dynamodbv2": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"role_arn": {
							  "type": "string",
							  "required": true
							}
						  },
						  "block_types": {
							"put_item": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "table_name": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  },
					  "elasticsearch": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"endpoint": {
							  "type": "string",
							  "required": true
							},
							"id": {
							  "type": "string",
							  "required": true
							},
							"index": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "firehose": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"delivery_stream_name": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							},
							"separator": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "iot_analytics": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"channel_name": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "iot_events": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"input_name": {
							  "type": "string",
							  "required": true
							},
							"message_id": {
							  "type": "string",
							  "optional": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "kinesis": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"partition_key": {
							  "type": "string",
							  "optional": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							},
							"stream_name": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "lambda": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"function_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "republish": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"qos": {
							  "type": "number",
							  "optional": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							},
							"topic": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "s3": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"bucket_name": {
							  "type": "string",
							  "required": true
							},
							"key": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "sns": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"message_format": {
							  "type": "string",
							  "optional": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							},
							"target_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "sqs": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"queue_url": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							},
							"use_base64": {
							  "type": "bool",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "step_functions": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"execution_name_prefix": {
							  "type": "string",
							  "optional": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							},
							"state_machine_name": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"firehose": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "delivery_stream_name": {
						"type": "string",
						"required": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "separator": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"iot_analytics": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "channel_name": {
						"type": "string",
						"required": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"iot_events": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "input_name": {
						"type": "string",
						"required": true
					  },
					  "message_id": {
						"type": "string",
						"optional": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"kinesis": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "partition_key": {
						"type": "string",
						"optional": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "stream_name": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"lambda": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "function_arn": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"republish": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "qos": {
						"type": "number",
						"optional": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "topic": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"s3": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "bucket_name": {
						"type": "string",
						"required": true
					  },
					  "key": {
						"type": "string",
						"required": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"sns": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "message_format": {
						"type": "string",
						"optional": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "target_arn": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"sqs": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "queue_url": {
						"type": "string",
						"required": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "use_base64": {
						"type": "bool",
						"required": true
					  }
					}
				  }
				},
				"step_functions": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "execution_name_prefix": {
						"type": "string",
						"optional": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "state_machine_name": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_key_pair": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"fingerprint": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key_name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"key_pair_id": {
				  "type": "string",
				  "computed": true
				},
				"public_key": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_kinesis_analytics_application": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"code": {
				  "type": "string",
				  "optional": true
				},
				"create_timestamp": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"last_update_timestamp": {
				  "type": "string",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"version": {
				  "type": "number",
				  "computed": true
				}
			  },
			  "block_types": {
				"cloudwatch_logging_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "id": {
						"type": "string",
						"computed": true
					  },
					  "log_stream_arn": {
						"type": "string",
						"required": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"inputs": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "id": {
						"type": "string",
						"computed": true
					  },
					  "name_prefix": {
						"type": "string",
						"required": true
					  },
					  "starting_position_configuration": {
						"type": [
						  "list",
						  [
							"object",
							{
							  "starting_position": "string"
							}
						  ]
						],
						"computed": true
					  },
					  "stream_names": {
						"type": [
						  "set",
						  "string"
						],
						"computed": true
					  }
					},
					"block_types": {
					  "kinesis_firehose": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"resource_arn": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "kinesis_stream": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"resource_arn": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "parallelism": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"count": {
							  "type": "number",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "processing_configuration": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"lambda": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "resource_arn": {
									"type": "string",
									"required": true
								  },
								  "role_arn": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  },
					  "schema": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"record_encoding": {
							  "type": "string",
							  "optional": true
							}
						  },
						  "block_types": {
							"record_columns": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "mapping": {
									"type": "string",
									"optional": true
								  },
								  "name": {
									"type": "string",
									"required": true
								  },
								  "sql_type": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "min_items": 1
							},
							"record_format": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "record_format_type": {
									"type": "string",
									"computed": true
								  }
								},
								"block_types": {
								  "mapping_parameters": {
									"nesting_mode": "list",
									"block": {
									  "block_types": {
										"csv": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "record_column_delimiter": {
												"type": "string",
												"required": true
											  },
											  "record_row_delimiter": {
												"type": "string",
												"required": true
											  }
											}
										  },
										  "max_items": 1
										},
										"json": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "record_row_path": {
												"type": "string",
												"required": true
											  }
											}
										  },
										  "max_items": 1
										}
									  }
									},
									"max_items": 1
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"outputs": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "id": {
						"type": "string",
						"computed": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "kinesis_firehose": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"resource_arn": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "kinesis_stream": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"resource_arn": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "lambda": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"resource_arn": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "schema": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"record_format_type": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "max_items": 3
				},
				"reference_data_sources": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "id": {
						"type": "string",
						"computed": true
					  },
					  "table_name": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "s3": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"bucket_arn": {
							  "type": "string",
							  "required": true
							},
							"file_key": {
							  "type": "string",
							  "required": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  },
					  "schema": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"record_encoding": {
							  "type": "string",
							  "optional": true
							}
						  },
						  "block_types": {
							"record_columns": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "mapping": {
									"type": "string",
									"optional": true
								  },
								  "name": {
									"type": "string",
									"required": true
								  },
								  "sql_type": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "min_items": 1
							},
							"record_format": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "record_format_type": {
									"type": "string",
									"computed": true
								  }
								},
								"block_types": {
								  "mapping_parameters": {
									"nesting_mode": "list",
									"block": {
									  "block_types": {
										"csv": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "record_column_delimiter": {
												"type": "string",
												"required": true
											  },
											  "record_row_delimiter": {
												"type": "string",
												"required": true
											  }
											}
										  },
										  "max_items": 1
										},
										"json": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "record_row_path": {
												"type": "string",
												"required": true
											  }
											}
										  },
										  "max_items": 1
										}
									  }
									},
									"max_items": 1
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_kinesis_firehose_delivery_stream": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"destination": {
				  "type": "string",
				  "required": true
				},
				"destination_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"version_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"elasticsearch_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "buffering_interval": {
						"type": "number",
						"optional": true
					  },
					  "buffering_size": {
						"type": "number",
						"optional": true
					  },
					  "domain_arn": {
						"type": "string",
						"required": true
					  },
					  "index_name": {
						"type": "string",
						"required": true
					  },
					  "index_rotation_period": {
						"type": "string",
						"optional": true
					  },
					  "retry_duration": {
						"type": "number",
						"optional": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "s3_backup_mode": {
						"type": "string",
						"optional": true
					  },
					  "type_name": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "cloudwatch_logging_options": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"enabled": {
							  "type": "bool",
							  "optional": true
							},
							"log_group_name": {
							  "type": "string",
							  "optional": true
							},
							"log_stream_name": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "processing_configuration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"enabled": {
							  "type": "bool",
							  "optional": true
							}
						  },
						  "block_types": {
							"processors": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "type": {
									"type": "string",
									"required": true
								  }
								},
								"block_types": {
								  "parameters": {
									"nesting_mode": "list",
									"block": {
									  "attributes": {
										"parameter_name": {
										  "type": "string",
										  "required": true
										},
										"parameter_value": {
										  "type": "string",
										  "required": true
										}
									  }
									}
								  }
								}
							  }
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"extended_s3_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket_arn": {
						"type": "string",
						"required": true
					  },
					  "buffer_interval": {
						"type": "number",
						"optional": true
					  },
					  "buffer_size": {
						"type": "number",
						"optional": true
					  },
					  "compression_format": {
						"type": "string",
						"optional": true
					  },
					  "error_output_prefix": {
						"type": "string",
						"optional": true
					  },
					  "kms_key_arn": {
						"type": "string",
						"optional": true
					  },
					  "prefix": {
						"type": "string",
						"optional": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "s3_backup_mode": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "cloudwatch_logging_options": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"enabled": {
							  "type": "bool",
							  "optional": true
							},
							"log_group_name": {
							  "type": "string",
							  "optional": true
							},
							"log_stream_name": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "data_format_conversion_configuration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"enabled": {
							  "type": "bool",
							  "optional": true
							}
						  },
						  "block_types": {
							"input_format_configuration": {
							  "nesting_mode": "list",
							  "block": {
								"block_types": {
								  "deserializer": {
									"nesting_mode": "list",
									"block": {
									  "block_types": {
										"hive_json_ser_de": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "timestamp_formats": {
												"type": [
												  "list",
												  "string"
												],
												"optional": true
											  }
											}
										  },
										  "max_items": 1
										},
										"open_x_json_ser_de": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "case_insensitive": {
												"type": "bool",
												"optional": true
											  },
											  "column_to_json_key_mappings": {
												"type": [
												  "map",
												  "string"
												],
												"optional": true
											  },
											  "convert_dots_in_json_keys_to_underscores": {
												"type": "bool",
												"optional": true
											  }
											}
										  },
										  "max_items": 1
										}
									  }
									},
									"min_items": 1,
									"max_items": 1
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							},
							"output_format_configuration": {
							  "nesting_mode": "list",
							  "block": {
								"block_types": {
								  "serializer": {
									"nesting_mode": "list",
									"block": {
									  "block_types": {
										"orc_ser_de": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "block_size_bytes": {
												"type": "number",
												"optional": true
											  },
											  "bloom_filter_columns": {
												"type": [
												  "list",
												  "string"
												],
												"optional": true
											  },
											  "bloom_filter_false_positive_probability": {
												"type": "number",
												"optional": true
											  },
											  "compression": {
												"type": "string",
												"optional": true
											  },
											  "dictionary_key_threshold": {
												"type": "number",
												"optional": true
											  },
											  "enable_padding": {
												"type": "bool",
												"optional": true
											  },
											  "format_version": {
												"type": "string",
												"optional": true
											  },
											  "padding_tolerance": {
												"type": "number",
												"optional": true
											  },
											  "row_index_stride": {
												"type": "number",
												"optional": true
											  },
											  "stripe_size_bytes": {
												"type": "number",
												"optional": true
											  }
											}
										  },
										  "max_items": 1
										},
										"parquet_ser_de": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "block_size_bytes": {
												"type": "number",
												"optional": true
											  },
											  "compression": {
												"type": "string",
												"optional": true
											  },
											  "enable_dictionary_compression": {
												"type": "bool",
												"optional": true
											  },
											  "max_padding_bytes": {
												"type": "number",
												"optional": true
											  },
											  "page_size_bytes": {
												"type": "number",
												"optional": true
											  },
											  "writer_version": {
												"type": "string",
												"optional": true
											  }
											}
										  },
										  "max_items": 1
										}
									  }
									},
									"min_items": 1,
									"max_items": 1
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							},
							"schema_configuration": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "catalog_id": {
									"type": "string",
									"optional": true,
									"computed": true
								  },
								  "database_name": {
									"type": "string",
									"required": true
								  },
								  "region": {
									"type": "string",
									"optional": true,
									"computed": true
								  },
								  "role_arn": {
									"type": "string",
									"required": true
								  },
								  "table_name": {
									"type": "string",
									"required": true
								  },
								  "version_id": {
									"type": "string",
									"optional": true
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  },
					  "processing_configuration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"enabled": {
							  "type": "bool",
							  "optional": true
							}
						  },
						  "block_types": {
							"processors": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "type": {
									"type": "string",
									"required": true
								  }
								},
								"block_types": {
								  "parameters": {
									"nesting_mode": "list",
									"block": {
									  "attributes": {
										"parameter_name": {
										  "type": "string",
										  "required": true
										},
										"parameter_value": {
										  "type": "string",
										  "required": true
										}
									  }
									}
								  }
								}
							  }
							}
						  }
						},
						"max_items": 1
					  },
					  "s3_backup_configuration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"bucket_arn": {
							  "type": "string",
							  "required": true
							},
							"buffer_interval": {
							  "type": "number",
							  "optional": true
							},
							"buffer_size": {
							  "type": "number",
							  "optional": true
							},
							"compression_format": {
							  "type": "string",
							  "optional": true
							},
							"kms_key_arn": {
							  "type": "string",
							  "optional": true
							},
							"prefix": {
							  "type": "string",
							  "optional": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							}
						  },
						  "block_types": {
							"cloudwatch_logging_options": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "enabled": {
									"type": "bool",
									"optional": true
								  },
								  "log_group_name": {
									"type": "string",
									"optional": true
								  },
								  "log_stream_name": {
									"type": "string",
									"optional": true
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"kinesis_source_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "kinesis_stream_arn": {
						"type": "string",
						"required": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"redshift_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "cluster_jdbcurl": {
						"type": "string",
						"required": true
					  },
					  "copy_options": {
						"type": "string",
						"optional": true
					  },
					  "data_table_columns": {
						"type": "string",
						"optional": true
					  },
					  "data_table_name": {
						"type": "string",
						"required": true
					  },
					  "password": {
						"type": "string",
						"required": true,
						"sensitive": true
					  },
					  "retry_duration": {
						"type": "number",
						"optional": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "s3_backup_mode": {
						"type": "string",
						"optional": true
					  },
					  "username": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "cloudwatch_logging_options": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"enabled": {
							  "type": "bool",
							  "optional": true
							},
							"log_group_name": {
							  "type": "string",
							  "optional": true
							},
							"log_stream_name": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "processing_configuration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"enabled": {
							  "type": "bool",
							  "optional": true
							}
						  },
						  "block_types": {
							"processors": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "type": {
									"type": "string",
									"required": true
								  }
								},
								"block_types": {
								  "parameters": {
									"nesting_mode": "list",
									"block": {
									  "attributes": {
										"parameter_name": {
										  "type": "string",
										  "required": true
										},
										"parameter_value": {
										  "type": "string",
										  "required": true
										}
									  }
									}
								  }
								}
							  }
							}
						  }
						},
						"max_items": 1
					  },
					  "s3_backup_configuration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"bucket_arn": {
							  "type": "string",
							  "required": true
							},
							"buffer_interval": {
							  "type": "number",
							  "optional": true
							},
							"buffer_size": {
							  "type": "number",
							  "optional": true
							},
							"compression_format": {
							  "type": "string",
							  "optional": true
							},
							"kms_key_arn": {
							  "type": "string",
							  "optional": true
							},
							"prefix": {
							  "type": "string",
							  "optional": true
							},
							"role_arn": {
							  "type": "string",
							  "required": true
							}
						  },
						  "block_types": {
							"cloudwatch_logging_options": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "enabled": {
									"type": "bool",
									"optional": true
								  },
								  "log_group_name": {
									"type": "string",
									"optional": true
								  },
								  "log_stream_name": {
									"type": "string",
									"optional": true
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"s3_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket_arn": {
						"type": "string",
						"required": true
					  },
					  "buffer_interval": {
						"type": "number",
						"optional": true
					  },
					  "buffer_size": {
						"type": "number",
						"optional": true
					  },
					  "compression_format": {
						"type": "string",
						"optional": true
					  },
					  "kms_key_arn": {
						"type": "string",
						"optional": true
					  },
					  "prefix": {
						"type": "string",
						"optional": true
					  },
					  "role_arn": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "cloudwatch_logging_options": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"enabled": {
							  "type": "bool",
							  "optional": true
							},
							"log_group_name": {
							  "type": "string",
							  "optional": true
							},
							"log_stream_name": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"server_side_encryption": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"splunk_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "hec_acknowledgment_timeout": {
						"type": "number",
						"optional": true
					  },
					  "hec_endpoint": {
						"type": "string",
						"required": true
					  },
					  "hec_endpoint_type": {
						"type": "string",
						"optional": true
					  },
					  "hec_token": {
						"type": "string",
						"required": true
					  },
					  "retry_duration": {
						"type": "number",
						"optional": true
					  },
					  "s3_backup_mode": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "cloudwatch_logging_options": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"enabled": {
							  "type": "bool",
							  "optional": true
							},
							"log_group_name": {
							  "type": "string",
							  "optional": true
							},
							"log_stream_name": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "processing_configuration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"enabled": {
							  "type": "bool",
							  "optional": true
							}
						  },
						  "block_types": {
							"processors": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "type": {
									"type": "string",
									"required": true
								  }
								},
								"block_types": {
								  "parameters": {
									"nesting_mode": "list",
									"block": {
									  "attributes": {
										"parameter_name": {
										  "type": "string",
										  "required": true
										},
										"parameter_value": {
										  "type": "string",
										  "required": true
										}
									  }
									}
								  }
								}
							  }
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_kinesis_stream": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"encryption_type": {
				  "type": "string",
				  "optional": true
				},
				"enforce_consumer_deletion": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"retention_period": {
				  "type": "number",
				  "optional": true
				},
				"shard_count": {
				  "type": "number",
				  "required": true
				},
				"shard_level_metrics": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_kinesis_video_stream": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"creation_time": {
				  "type": "string",
				  "computed": true
				},
				"data_retention_in_hours": {
				  "type": "number",
				  "optional": true
				},
				"device_name": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"media_type": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"version": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_kms_alias": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"target_key_arn": {
				  "type": "string",
				  "computed": true
				},
				"target_key_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_kms_ciphertext": {
			"version": 0,
			"block": {
			  "attributes": {
				"ciphertext_blob": {
				  "type": "string",
				  "computed": true
				},
				"context": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key_id": {
				  "type": "string",
				  "required": true
				},
				"plaintext": {
				  "type": "string",
				  "required": true,
				  "sensitive": true
				}
			  }
			}
		  },
		  "aws_kms_external_key": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"deletion_window_in_days": {
				  "type": "number",
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"expiration_model": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key_material_base64": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"key_state": {
				  "type": "string",
				  "computed": true
				},
				"key_usage": {
				  "type": "string",
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"valid_to": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_kms_grant": {
			"version": 0,
			"block": {
			  "attributes": {
				"grant_creation_tokens": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"grant_id": {
				  "type": "string",
				  "computed": true
				},
				"grant_token": {
				  "type": "string",
				  "computed": true
				},
				"grantee_principal": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key_id": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"operations": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"retire_on_delete": {
				  "type": "bool",
				  "optional": true
				},
				"retiring_principal": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"constraints": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "encryption_context_equals": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  },
					  "encryption_context_subset": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_kms_key": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"customer_master_key_spec": {
				  "type": "string",
				  "optional": true
				},
				"deletion_window_in_days": {
				  "type": "number",
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"enable_key_rotation": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"is_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"key_id": {
				  "type": "string",
				  "computed": true
				},
				"key_usage": {
				  "type": "string",
				  "optional": true
				},
				"policy": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_lambda_alias": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"function_name": {
				  "type": "string",
				  "required": true
				},
				"function_version": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"invoke_arn": {
				  "type": "string",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"routing_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "additional_version_weights": {
						"type": [
						  "map",
						  "number"
						],
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_lambda_event_source_mapping": {
			"version": 0,
			"block": {
			  "attributes": {
				"batch_size": {
				  "type": "number",
				  "optional": true
				},
				"bisect_batch_on_function_error": {
				  "type": "bool",
				  "optional": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"event_source_arn": {
				  "type": "string",
				  "required": true
				},
				"function_arn": {
				  "type": "string",
				  "computed": true
				},
				"function_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"last_modified": {
				  "type": "string",
				  "computed": true
				},
				"last_processing_result": {
				  "type": "string",
				  "computed": true
				},
				"maximum_batching_window_in_seconds": {
				  "type": "number",
				  "optional": true
				},
				"maximum_record_age_in_seconds": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"maximum_retry_attempts": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"parallelization_factor": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"starting_position": {
				  "type": "string",
				  "optional": true
				},
				"starting_position_timestamp": {
				  "type": "string",
				  "optional": true
				},
				"state": {
				  "type": "string",
				  "computed": true
				},
				"state_transition_reason": {
				  "type": "string",
				  "computed": true
				},
				"uuid": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"destination_config": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "on_failure": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"destination_arn": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_lambda_function": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"filename": {
				  "type": "string",
				  "optional": true
				},
				"function_name": {
				  "type": "string",
				  "required": true
				},
				"handler": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"invoke_arn": {
				  "type": "string",
				  "computed": true
				},
				"kms_key_arn": {
				  "type": "string",
				  "optional": true
				},
				"last_modified": {
				  "type": "string",
				  "computed": true
				},
				"layers": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"memory_size": {
				  "type": "number",
				  "optional": true
				},
				"publish": {
				  "type": "bool",
				  "optional": true
				},
				"qualified_arn": {
				  "type": "string",
				  "computed": true
				},
				"reserved_concurrent_executions": {
				  "type": "number",
				  "optional": true
				},
				"role": {
				  "type": "string",
				  "required": true
				},
				"runtime": {
				  "type": "string",
				  "required": true
				},
				"s3_bucket": {
				  "type": "string",
				  "optional": true
				},
				"s3_key": {
				  "type": "string",
				  "optional": true
				},
				"s3_object_version": {
				  "type": "string",
				  "optional": true
				},
				"source_code_hash": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"source_code_size": {
				  "type": "number",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"timeout": {
				  "type": "number",
				  "optional": true
				},
				"version": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"dead_letter_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "target_arn": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"environment": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "variables": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"file_system_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "arn": {
						"type": "string",
						"required": true
					  },
					  "local_mount_path": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"tracing_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "mode": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"vpc_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "security_group_ids": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "subnet_ids": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "vpc_id": {
						"type": "string",
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_lambda_function_event_invoke_config": {
			"version": 0,
			"block": {
			  "attributes": {
				"function_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"maximum_event_age_in_seconds": {
				  "type": "number",
				  "optional": true
				},
				"maximum_retry_attempts": {
				  "type": "number",
				  "optional": true
				},
				"qualifier": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"destination_config": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "on_failure": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"destination": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "on_success": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"destination": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_lambda_layer_version": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"compatible_runtimes": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"created_date": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"filename": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"layer_arn": {
				  "type": "string",
				  "computed": true
				},
				"layer_name": {
				  "type": "string",
				  "required": true
				},
				"license_info": {
				  "type": "string",
				  "optional": true
				},
				"s3_bucket": {
				  "type": "string",
				  "optional": true
				},
				"s3_key": {
				  "type": "string",
				  "optional": true
				},
				"s3_object_version": {
				  "type": "string",
				  "optional": true
				},
				"source_code_hash": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"source_code_size": {
				  "type": "number",
				  "computed": true
				},
				"version": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_lambda_permission": {
			"version": 0,
			"block": {
			  "attributes": {
				"action": {
				  "type": "string",
				  "required": true
				},
				"event_source_token": {
				  "type": "string",
				  "optional": true
				},
				"function_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"principal": {
				  "type": "string",
				  "required": true
				},
				"qualifier": {
				  "type": "string",
				  "optional": true
				},
				"source_account": {
				  "type": "string",
				  "optional": true
				},
				"source_arn": {
				  "type": "string",
				  "optional": true
				},
				"statement_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"statement_id_prefix": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_lambda_provisioned_concurrency_config": {
			"version": 0,
			"block": {
			  "attributes": {
				"function_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"provisioned_concurrent_executions": {
				  "type": "number",
				  "required": true
				},
				"qualifier": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_launch_configuration": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"associate_public_ip_address": {
				  "type": "bool",
				  "optional": true
				},
				"ebs_optimized": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"enable_monitoring": {
				  "type": "bool",
				  "optional": true
				},
				"iam_instance_profile": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"image_id": {
				  "type": "string",
				  "required": true
				},
				"instance_type": {
				  "type": "string",
				  "required": true
				},
				"key_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"placement_tenancy": {
				  "type": "string",
				  "optional": true
				},
				"security_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"spot_price": {
				  "type": "string",
				  "optional": true
				},
				"user_data": {
				  "type": "string",
				  "optional": true
				},
				"user_data_base64": {
				  "type": "string",
				  "optional": true
				},
				"vpc_classic_link_id": {
				  "type": "string",
				  "optional": true
				},
				"vpc_classic_link_security_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"ebs_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "delete_on_termination": {
						"type": "bool",
						"optional": true
					  },
					  "device_name": {
						"type": "string",
						"required": true
					  },
					  "encrypted": {
						"type": "bool",
						"optional": true,
						"computed": true
					  },
					  "iops": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "no_device": {
						"type": "bool",
						"optional": true
					  },
					  "snapshot_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "volume_size": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "volume_type": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  }
				},
				"ephemeral_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "device_name": {
						"type": "string",
						"required": true
					  },
					  "virtual_name": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"root_block_device": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "delete_on_termination": {
						"type": "bool",
						"optional": true
					  },
					  "encrypted": {
						"type": "bool",
						"optional": true,
						"computed": true
					  },
					  "iops": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "volume_size": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "volume_type": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_launch_template": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"default_version": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"disable_api_termination": {
				  "type": "bool",
				  "optional": true
				},
				"ebs_optimized": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"image_id": {
				  "type": "string",
				  "optional": true
				},
				"instance_initiated_shutdown_behavior": {
				  "type": "string",
				  "optional": true
				},
				"instance_type": {
				  "type": "string",
				  "optional": true
				},
				"kernel_id": {
				  "type": "string",
				  "optional": true
				},
				"key_name": {
				  "type": "string",
				  "optional": true
				},
				"latest_version": {
				  "type": "number",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"ram_disk_id": {
				  "type": "string",
				  "optional": true
				},
				"security_group_names": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"update_default_version": {
				  "type": "bool",
				  "optional": true
				},
				"user_data": {
				  "type": "string",
				  "optional": true
				},
				"vpc_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"block_device_mappings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "device_name": {
						"type": "string",
						"optional": true
					  },
					  "no_device": {
						"type": "string",
						"optional": true
					  },
					  "virtual_name": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "ebs": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"delete_on_termination": {
							  "type": "string",
							  "optional": true
							},
							"encrypted": {
							  "type": "string",
							  "optional": true
							},
							"iops": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"kms_key_id": {
							  "type": "string",
							  "optional": true
							},
							"snapshot_id": {
							  "type": "string",
							  "optional": true
							},
							"volume_size": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"volume_type": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  }
				},
				"capacity_reservation_specification": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "capacity_reservation_preference": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "capacity_reservation_target": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"capacity_reservation_id": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"cpu_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "core_count": {
						"type": "number",
						"optional": true
					  },
					  "threads_per_core": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"credit_specification": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "cpu_credits": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"elastic_gpu_specifications": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"elastic_inference_accelerator": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"hibernation_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "configured": {
						"type": "bool",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"iam_instance_profile": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "arn": {
						"type": "string",
						"optional": true
					  },
					  "name": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"instance_market_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "market_type": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "spot_options": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"block_duration_minutes": {
							  "type": "number",
							  "optional": true
							},
							"instance_interruption_behavior": {
							  "type": "string",
							  "optional": true
							},
							"max_price": {
							  "type": "string",
							  "optional": true
							},
							"spot_instance_type": {
							  "type": "string",
							  "optional": true
							},
							"valid_until": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"license_specification": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "license_configuration_arn": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"metadata_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "http_endpoint": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "http_put_response_hop_limit": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "http_tokens": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				},
				"monitoring": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"network_interfaces": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "associate_public_ip_address": {
						"type": "string",
						"optional": true
					  },
					  "delete_on_termination": {
						"type": "string",
						"optional": true
					  },
					  "description": {
						"type": "string",
						"optional": true
					  },
					  "device_index": {
						"type": "number",
						"optional": true
					  },
					  "ipv4_address_count": {
						"type": "number",
						"optional": true
					  },
					  "ipv4_addresses": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "ipv6_address_count": {
						"type": "number",
						"optional": true
					  },
					  "ipv6_addresses": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "network_interface_id": {
						"type": "string",
						"optional": true
					  },
					  "private_ip_address": {
						"type": "string",
						"optional": true
					  },
					  "security_groups": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "subnet_id": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"placement": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "affinity": {
						"type": "string",
						"optional": true
					  },
					  "availability_zone": {
						"type": "string",
						"optional": true
					  },
					  "group_name": {
						"type": "string",
						"optional": true
					  },
					  "host_id": {
						"type": "string",
						"optional": true
					  },
					  "partition_number": {
						"type": "number",
						"optional": true
					  },
					  "spread_domain": {
						"type": "string",
						"optional": true
					  },
					  "tenancy": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"tag_specifications": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "resource_type": {
						"type": "string",
						"optional": true
					  },
					  "tags": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_lb": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"arn_suffix": {
				  "type": "string",
				  "computed": true
				},
				"dns_name": {
				  "type": "string",
				  "computed": true
				},
				"drop_invalid_header_fields": {
				  "type": "bool",
				  "optional": true
				},
				"enable_cross_zone_load_balancing": {
				  "type": "bool",
				  "optional": true
				},
				"enable_deletion_protection": {
				  "type": "bool",
				  "optional": true
				},
				"enable_http2": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"idle_timeout": {
				  "type": "number",
				  "optional": true
				},
				"internal": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"ip_address_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"load_balancer_type": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"security_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"subnets": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				},
				"zone_id": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"access_logs": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket": {
						"type": "string",
						"required": true
					  },
					  "enabled": {
						"type": "bool",
						"optional": true
					  },
					  "prefix": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"subnet_mapping": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "allocation_id": {
						"type": "string",
						"optional": true
					  },
					  "private_ipv4_address": {
						"type": "string",
						"optional": true
					  },
					  "subnet_id": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_lb_cookie_stickiness_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"cookie_expiration_period": {
				  "type": "number",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"lb_port": {
				  "type": "number",
				  "required": true
				},
				"load_balancer": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_lb_listener": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"certificate_arn": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"load_balancer_arn": {
				  "type": "string",
				  "required": true
				},
				"port": {
				  "type": "number",
				  "required": true
				},
				"protocol": {
				  "type": "string",
				  "optional": true
				},
				"ssl_policy": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"default_action": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "order": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "target_group_arn": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "authenticate_cognito": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"authentication_request_extra_params": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"on_unauthenticated_request": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"scope": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_cookie_name": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_timeout": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"user_pool_arn": {
							  "type": "string",
							  "required": true
							},
							"user_pool_client_id": {
							  "type": "string",
							  "required": true
							},
							"user_pool_domain": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "authenticate_oidc": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"authentication_request_extra_params": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"authorization_endpoint": {
							  "type": "string",
							  "required": true
							},
							"client_id": {
							  "type": "string",
							  "required": true
							},
							"client_secret": {
							  "type": "string",
							  "required": true,
							  "sensitive": true
							},
							"issuer": {
							  "type": "string",
							  "required": true
							},
							"on_unauthenticated_request": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"scope": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_cookie_name": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_timeout": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"token_endpoint": {
							  "type": "string",
							  "required": true
							},
							"user_info_endpoint": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "fixed_response": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"content_type": {
							  "type": "string",
							  "required": true
							},
							"message_body": {
							  "type": "string",
							  "optional": true
							},
							"status_code": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							}
						  }
						},
						"max_items": 1
					  },
					  "forward": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"stickiness": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "duration": {
									"type": "number",
									"required": true
								  },
								  "enabled": {
									"type": "bool",
									"optional": true
								  }
								}
							  },
							  "max_items": 1
							},
							"target_group": {
							  "nesting_mode": "set",
							  "block": {
								"attributes": {
								  "arn": {
									"type": "string",
									"required": true
								  },
								  "weight": {
									"type": "number",
									"optional": true
								  }
								}
							  },
							  "min_items": 2,
							  "max_items": 5
							}
						  }
						},
						"max_items": 1
					  },
					  "redirect": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"host": {
							  "type": "string",
							  "optional": true
							},
							"path": {
							  "type": "string",
							  "optional": true
							},
							"port": {
							  "type": "string",
							  "optional": true
							},
							"protocol": {
							  "type": "string",
							  "optional": true
							},
							"query": {
							  "type": "string",
							  "optional": true
							},
							"status_code": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "read": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_lb_listener_certificate": {
			"version": 0,
			"block": {
			  "attributes": {
				"certificate_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"listener_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_lb_listener_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"listener_arn": {
				  "type": "string",
				  "required": true
				},
				"priority": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"action": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "order": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "target_group_arn": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "authenticate_cognito": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"authentication_request_extra_params": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"on_unauthenticated_request": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"scope": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_cookie_name": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_timeout": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"user_pool_arn": {
							  "type": "string",
							  "required": true
							},
							"user_pool_client_id": {
							  "type": "string",
							  "required": true
							},
							"user_pool_domain": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "authenticate_oidc": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"authentication_request_extra_params": {
							  "type": [
								"map",
								"string"
							  ],
							  "optional": true
							},
							"authorization_endpoint": {
							  "type": "string",
							  "required": true
							},
							"client_id": {
							  "type": "string",
							  "required": true
							},
							"client_secret": {
							  "type": "string",
							  "required": true,
							  "sensitive": true
							},
							"issuer": {
							  "type": "string",
							  "required": true
							},
							"on_unauthenticated_request": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"scope": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_cookie_name": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"session_timeout": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"token_endpoint": {
							  "type": "string",
							  "required": true
							},
							"user_info_endpoint": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "fixed_response": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"content_type": {
							  "type": "string",
							  "required": true
							},
							"message_body": {
							  "type": "string",
							  "optional": true
							},
							"status_code": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							}
						  }
						},
						"max_items": 1
					  },
					  "forward": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"stickiness": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "duration": {
									"type": "number",
									"required": true
								  },
								  "enabled": {
									"type": "bool",
									"optional": true
								  }
								}
							  },
							  "max_items": 1
							},
							"target_group": {
							  "nesting_mode": "set",
							  "block": {
								"attributes": {
								  "arn": {
									"type": "string",
									"required": true
								  },
								  "weight": {
									"type": "number",
									"optional": true
								  }
								}
							  },
							  "min_items": 2,
							  "max_items": 5
							}
						  }
						},
						"max_items": 1
					  },
					  "redirect": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"host": {
							  "type": "string",
							  "optional": true
							},
							"path": {
							  "type": "string",
							  "optional": true
							},
							"port": {
							  "type": "string",
							  "optional": true
							},
							"protocol": {
							  "type": "string",
							  "optional": true
							},
							"query": {
							  "type": "string",
							  "optional": true
							},
							"status_code": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1
				},
				"condition": {
				  "nesting_mode": "set",
				  "block": {
					"block_types": {
					  "host_header": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"values": {
							  "type": [
								"set",
								"string"
							  ],
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "http_header": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"http_header_name": {
							  "type": "string",
							  "required": true
							},
							"values": {
							  "type": [
								"set",
								"string"
							  ],
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "http_request_method": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"values": {
							  "type": [
								"set",
								"string"
							  ],
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "path_pattern": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"values": {
							  "type": [
								"set",
								"string"
							  ],
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "query_string": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"key": {
							  "type": "string",
							  "optional": true
							},
							"value": {
							  "type": "string",
							  "required": true
							}
						  }
						}
					  },
					  "source_ip": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"values": {
							  "type": [
								"set",
								"string"
							  ],
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1
				}
			  }
			}
		  },
		  "aws_lb_ssl_negotiation_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"lb_port": {
				  "type": "number",
				  "required": true
				},
				"load_balancer": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"attribute": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "name": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_lb_target_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"arn_suffix": {
				  "type": "string",
				  "computed": true
				},
				"deregistration_delay": {
				  "type": "number",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"lambda_multi_value_headers_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"load_balancing_algorithm_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"port": {
				  "type": "number",
				  "optional": true
				},
				"protocol": {
				  "type": "string",
				  "optional": true
				},
				"proxy_protocol_v2": {
				  "type": "bool",
				  "optional": true
				},
				"slow_start": {
				  "type": "number",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"target_type": {
				  "type": "string",
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"health_check": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"optional": true
					  },
					  "healthy_threshold": {
						"type": "number",
						"optional": true
					  },
					  "interval": {
						"type": "number",
						"optional": true
					  },
					  "matcher": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "path": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "port": {
						"type": "string",
						"optional": true
					  },
					  "protocol": {
						"type": "string",
						"optional": true
					  },
					  "timeout": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "unhealthy_threshold": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"stickiness": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "cookie_duration": {
						"type": "number",
						"optional": true
					  },
					  "enabled": {
						"type": "bool",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_lb_target_group_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"availability_zone": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "optional": true
				},
				"target_group_arn": {
				  "type": "string",
				  "required": true
				},
				"target_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_licensemanager_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"license_configuration_arn": {
				  "type": "string",
				  "required": true
				},
				"resource_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_licensemanager_license_configuration": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"license_count": {
				  "type": "number",
				  "optional": true
				},
				"license_count_hard_limit": {
				  "type": "bool",
				  "optional": true
				},
				"license_counting_type": {
				  "type": "string",
				  "required": true
				},
				"license_rules": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_lightsail_domain": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"domain_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_lightsail_instance": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"availability_zone": {
				  "type": "string",
				  "required": true
				},
				"blueprint_id": {
				  "type": "string",
				  "required": true
				},
				"bundle_id": {
				  "type": "string",
				  "required": true
				},
				"cpu_count": {
				  "type": "number",
				  "computed": true
				},
				"created_at": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ipv6_address": {
				  "type": "string",
				  "computed": true
				},
				"is_static_ip": {
				  "type": "bool",
				  "computed": true
				},
				"key_pair_name": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"private_ip_address": {
				  "type": "string",
				  "computed": true
				},
				"public_ip_address": {
				  "type": "string",
				  "computed": true
				},
				"ram_size": {
				  "type": "number",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"user_data": {
				  "type": "string",
				  "optional": true
				},
				"username": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_lightsail_key_pair": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"encrypted_fingerprint": {
				  "type": "string",
				  "computed": true
				},
				"encrypted_private_key": {
				  "type": "string",
				  "computed": true
				},
				"fingerprint": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"pgp_key": {
				  "type": "string",
				  "optional": true
				},
				"private_key": {
				  "type": "string",
				  "computed": true
				},
				"public_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_lightsail_static_ip": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ip_address": {
				  "type": "string",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"support_code": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_lightsail_static_ip_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_name": {
				  "type": "string",
				  "required": true
				},
				"ip_address": {
				  "type": "string",
				  "computed": true
				},
				"static_ip_name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_load_balancer_backend_server_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_port": {
				  "type": "number",
				  "required": true
				},
				"load_balancer_name": {
				  "type": "string",
				  "required": true
				},
				"policy_names": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_load_balancer_listener_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"load_balancer_name": {
				  "type": "string",
				  "required": true
				},
				"load_balancer_port": {
				  "type": "number",
				  "required": true
				},
				"policy_names": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_load_balancer_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"load_balancer_name": {
				  "type": "string",
				  "required": true
				},
				"policy_name": {
				  "type": "string",
				  "required": true
				},
				"policy_type_name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"policy_attribute": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "name": {
						"type": "string",
						"optional": true
					  },
					  "value": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_macie_member_account_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"member_account_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_macie_s3_bucket_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"bucket_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"member_account_id": {
				  "type": "string",
				  "optional": true
				},
				"prefix": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"classification_type": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "continuous": {
						"type": "string",
						"optional": true
					  },
					  "one_time": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_main_route_table_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"original_route_table_id": {
				  "type": "string",
				  "computed": true
				},
				"route_table_id": {
				  "type": "string",
				  "required": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_media_convert_queue": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"pricing_plan": {
				  "type": "string",
				  "optional": true
				},
				"status": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"reservation_plan_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "commitment": {
						"type": "string",
						"required": true
					  },
					  "renewal_type": {
						"type": "string",
						"required": true
					  },
					  "reserved_slots": {
						"type": "number",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_media_package_channel": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"channel_id": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"hls_ingest": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"ingest_endpoints": [
						  "list",
						  [
							"object",
							{
							  "password": "string",
							  "url": "string",
							  "username": "string"
							}
						  ]
						]
					  }
					]
				  ],
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_media_store_container": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"endpoint": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_media_store_container_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"container_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_mq_broker": {
			"version": 0,
			"block": {
			  "attributes": {
				"apply_immediately": {
				  "type": "bool",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_minor_version_upgrade": {
				  "type": "bool",
				  "optional": true
				},
				"broker_name": {
				  "type": "string",
				  "required": true
				},
				"deployment_mode": {
				  "type": "string",
				  "optional": true
				},
				"engine_type": {
				  "type": "string",
				  "required": true
				},
				"engine_version": {
				  "type": "string",
				  "required": true
				},
				"host_instance_type": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instances": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"console_url": "string",
						"endpoints": [
						  "list",
						  "string"
						],
						"ip_address": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"publicly_accessible": {
				  "type": "bool",
				  "optional": true
				},
				"security_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "revision": {
						"type": "number",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				},
				"encryption_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "kms_key_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "use_aws_owned_key": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"logs": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "audit": {
						"type": "bool",
						"optional": true
					  },
					  "general": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"maintenance_window_start_time": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "day_of_week": {
						"type": "string",
						"required": true
					  },
					  "time_of_day": {
						"type": "string",
						"required": true
					  },
					  "time_zone": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"user": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "console_access": {
						"type": "bool",
						"optional": true
					  },
					  "groups": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true
					  },
					  "password": {
						"type": "string",
						"required": true,
						"sensitive": true
					  },
					  "username": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1
				}
			  }
			}
		  },
		  "aws_mq_configuration": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"data": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"engine_type": {
				  "type": "string",
				  "required": true
				},
				"engine_version": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"latest_revision": {
				  "type": "number",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_msk_cluster": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"bootstrap_brokers": {
				  "type": "string",
				  "computed": true
				},
				"bootstrap_brokers_tls": {
				  "type": "string",
				  "computed": true
				},
				"cluster_name": {
				  "type": "string",
				  "required": true
				},
				"current_version": {
				  "type": "string",
				  "computed": true
				},
				"enhanced_monitoring": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kafka_version": {
				  "type": "string",
				  "required": true
				},
				"number_of_broker_nodes": {
				  "type": "number",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"zookeeper_connect_string": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"broker_node_group_info": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "az_distribution": {
						"type": "string",
						"optional": true
					  },
					  "client_subnets": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  },
					  "ebs_volume_size": {
						"type": "number",
						"required": true
					  },
					  "instance_type": {
						"type": "string",
						"required": true
					  },
					  "security_groups": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"client_authentication": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "tls": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"certificate_authority_arns": {
							  "type": [
								"set",
								"string"
							  ],
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"configuration_info": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "arn": {
						"type": "string",
						"required": true
					  },
					  "revision": {
						"type": "number",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"encryption_info": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "encryption_at_rest_kms_key_arn": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					},
					"block_types": {
					  "encryption_in_transit": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"client_broker": {
							  "type": "string",
							  "optional": true
							},
							"in_cluster": {
							  "type": "bool",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"logging_info": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "broker_logs": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"cloudwatch_logs": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "enabled": {
									"type": "bool",
									"required": true
								  },
								  "log_group": {
									"type": "string",
									"optional": true
								  }
								}
							  },
							  "max_items": 1
							},
							"firehose": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "delivery_stream": {
									"type": "string",
									"optional": true
								  },
								  "enabled": {
									"type": "bool",
									"required": true
								  }
								}
							  },
							  "max_items": 1
							},
							"s3": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "bucket": {
									"type": "string",
									"optional": true
								  },
								  "enabled": {
									"type": "bool",
									"required": true
								  },
								  "prefix": {
									"type": "string",
									"optional": true
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"open_monitoring": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "prometheus": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"jmx_exporter": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "enabled_in_broker": {
									"type": "bool",
									"required": true
								  }
								}
							  },
							  "max_items": 1
							},
							"node_exporter": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "enabled_in_broker": {
									"type": "bool",
									"required": true
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_msk_configuration": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kafka_versions": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"latest_revision": {
				  "type": "number",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"server_properties": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_nat_gateway": {
			"version": 0,
			"block": {
			  "attributes": {
				"allocation_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"network_interface_id": {
				  "type": "string",
				  "computed": true
				},
				"private_ip": {
				  "type": "string",
				  "computed": true
				},
				"public_ip": {
				  "type": "string",
				  "computed": true
				},
				"subnet_id": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_neptune_cluster": {
			"version": 0,
			"block": {
			  "attributes": {
				"apply_immediately": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"availability_zones": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"backup_retention_period": {
				  "type": "number",
				  "optional": true
				},
				"cluster_identifier": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_identifier_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_members": {
				  "type": [
					"set",
					"string"
				  ],
				  "computed": true
				},
				"cluster_resource_id": {
				  "type": "string",
				  "computed": true
				},
				"deletion_protection": {
				  "type": "bool",
				  "optional": true
				},
				"enable_cloudwatch_logs_exports": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"endpoint": {
				  "type": "string",
				  "computed": true
				},
				"engine": {
				  "type": "string",
				  "optional": true
				},
				"engine_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"final_snapshot_identifier": {
				  "type": "string",
				  "optional": true
				},
				"hosted_zone_id": {
				  "type": "string",
				  "computed": true
				},
				"iam_database_authentication_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"iam_roles": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"neptune_cluster_parameter_group_name": {
				  "type": "string",
				  "optional": true
				},
				"neptune_subnet_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "optional": true
				},
				"preferred_backup_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"preferred_maintenance_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"reader_endpoint": {
				  "type": "string",
				  "computed": true
				},
				"replication_source_identifier": {
				  "type": "string",
				  "optional": true
				},
				"skip_final_snapshot": {
				  "type": "bool",
				  "optional": true
				},
				"snapshot_identifier": {
				  "type": "string",
				  "optional": true
				},
				"storage_encrypted": {
				  "type": "bool",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_neptune_cluster_instance": {
			"version": 0,
			"block": {
			  "attributes": {
				"address": {
				  "type": "string",
				  "computed": true
				},
				"apply_immediately": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_minor_version_upgrade": {
				  "type": "bool",
				  "optional": true
				},
				"availability_zone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_identifier": {
				  "type": "string",
				  "required": true
				},
				"dbi_resource_id": {
				  "type": "string",
				  "computed": true
				},
				"endpoint": {
				  "type": "string",
				  "computed": true
				},
				"engine": {
				  "type": "string",
				  "optional": true
				},
				"engine_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identifier": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identifier_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_class": {
				  "type": "string",
				  "required": true
				},
				"kms_key_arn": {
				  "type": "string",
				  "computed": true
				},
				"neptune_parameter_group_name": {
				  "type": "string",
				  "optional": true
				},
				"neptune_subnet_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "optional": true
				},
				"preferred_backup_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"preferred_maintenance_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"promotion_tier": {
				  "type": "number",
				  "optional": true
				},
				"publicly_accessible": {
				  "type": "bool",
				  "optional": true
				},
				"storage_encrypted": {
				  "type": "bool",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"writer": {
				  "type": "bool",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_neptune_cluster_parameter_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"family": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"parameter": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "apply_method": {
						"type": "string",
						"optional": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_neptune_cluster_snapshot": {
			"version": 0,
			"block": {
			  "attributes": {
				"allocated_storage": {
				  "type": "number",
				  "computed": true
				},
				"availability_zones": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"db_cluster_identifier": {
				  "type": "string",
				  "required": true
				},
				"db_cluster_snapshot_arn": {
				  "type": "string",
				  "computed": true
				},
				"db_cluster_snapshot_identifier": {
				  "type": "string",
				  "required": true
				},
				"engine": {
				  "type": "string",
				  "computed": true
				},
				"engine_version": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "computed": true
				},
				"license_model": {
				  "type": "string",
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "computed": true
				},
				"snapshot_type": {
				  "type": "string",
				  "computed": true
				},
				"source_db_cluster_snapshot_arn": {
				  "type": "string",
				  "computed": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"storage_encrypted": {
				  "type": "bool",
				  "computed": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_neptune_event_subscription": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"customer_aws_id": {
				  "type": "string",
				  "computed": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"event_categories": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"sns_topic_arn": {
				  "type": "string",
				  "required": true
				},
				"source_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"source_type": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_neptune_parameter_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"family": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"parameter": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "apply_method": {
						"type": "string",
						"optional": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_neptune_subnet_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_network_acl": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"egress": {
				  "type": [
					"set",
					[
					  "object",
					  {
						"action": "string",
						"cidr_block": "string",
						"from_port": "number",
						"icmp_code": "number",
						"icmp_type": "number",
						"ipv6_cidr_block": "string",
						"protocol": "string",
						"rule_no": "number",
						"to_port": "number"
					  }
					]
				  ],
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ingress": {
				  "type": [
					"set",
					[
					  "object",
					  {
						"action": "string",
						"cidr_block": "string",
						"from_port": "number",
						"icmp_code": "number",
						"icmp_type": "number",
						"ipv6_cidr_block": "string",
						"protocol": "string",
						"rule_no": "number",
						"to_port": "number"
					  }
					]
				  ],
				  "optional": true,
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_network_acl_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"cidr_block": {
				  "type": "string",
				  "optional": true
				},
				"egress": {
				  "type": "bool",
				  "optional": true
				},
				"from_port": {
				  "type": "number",
				  "optional": true
				},
				"icmp_code": {
				  "type": "string",
				  "optional": true
				},
				"icmp_type": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ipv6_cidr_block": {
				  "type": "string",
				  "optional": true
				},
				"network_acl_id": {
				  "type": "string",
				  "required": true
				},
				"protocol": {
				  "type": "string",
				  "required": true
				},
				"rule_action": {
				  "type": "string",
				  "required": true
				},
				"rule_number": {
				  "type": "number",
				  "required": true
				},
				"to_port": {
				  "type": "number",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_network_interface": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"mac_address": {
				  "type": "string",
				  "computed": true
				},
				"outpost_arn": {
				  "type": "string",
				  "computed": true
				},
				"private_dns_name": {
				  "type": "string",
				  "computed": true
				},
				"private_ip": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"private_ips": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"private_ips_count": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"security_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"source_dest_check": {
				  "type": "bool",
				  "optional": true
				},
				"subnet_id": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"attachment": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "attachment_id": {
						"type": "string",
						"computed": true
					  },
					  "device_index": {
						"type": "number",
						"required": true
					  },
					  "instance": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_network_interface_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"attachment_id": {
				  "type": "string",
				  "computed": true
				},
				"device_index": {
				  "type": "number",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_id": {
				  "type": "string",
				  "required": true
				},
				"network_interface_id": {
				  "type": "string",
				  "required": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_network_interface_sg_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"network_interface_id": {
				  "type": "string",
				  "required": true
				},
				"security_group_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_opsworks_application": {
			"version": 0,
			"block": {
			  "attributes": {
				"auto_bundle_on_deploy": {
				  "type": "string",
				  "optional": true
				},
				"aws_flow_ruby_settings": {
				  "type": "string",
				  "optional": true
				},
				"data_source_arn": {
				  "type": "string",
				  "optional": true
				},
				"data_source_database_name": {
				  "type": "string",
				  "optional": true
				},
				"data_source_type": {
				  "type": "string",
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"document_root": {
				  "type": "string",
				  "optional": true
				},
				"domains": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"enable_ssl": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"rails_env": {
				  "type": "string",
				  "optional": true
				},
				"short_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"stack_id": {
				  "type": "string",
				  "required": true
				},
				"type": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"app_source": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "password": {
						"type": "string",
						"optional": true,
						"sensitive": true
					  },
					  "revision": {
						"type": "string",
						"optional": true
					  },
					  "ssh_key": {
						"type": "string",
						"optional": true,
						"sensitive": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  },
					  "url": {
						"type": "string",
						"optional": true
					  },
					  "username": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"environment": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "key": {
						"type": "string",
						"required": true
					  },
					  "secure": {
						"type": "bool",
						"optional": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"ssl_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "certificate": {
						"type": "string",
						"required": true
					  },
					  "chain": {
						"type": "string",
						"optional": true
					  },
					  "private_key": {
						"type": "string",
						"required": true,
						"sensitive": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_opsworks_custom_layer": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_assign_elastic_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_assign_public_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_healing": {
				  "type": "bool",
				  "optional": true
				},
				"custom_configure_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_deploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_instance_profile_arn": {
				  "type": "string",
				  "optional": true
				},
				"custom_json": {
				  "type": "string",
				  "optional": true
				},
				"custom_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"custom_setup_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_shutdown_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_undeploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"drain_elb_on_shutdown": {
				  "type": "bool",
				  "optional": true
				},
				"elastic_load_balancer": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"install_updates_on_boot": {
				  "type": "bool",
				  "optional": true
				},
				"instance_shutdown_timeout": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"short_name": {
				  "type": "string",
				  "required": true
				},
				"stack_id": {
				  "type": "string",
				  "required": true
				},
				"system_packages": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"use_ebs_optimized_instances": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"ebs_volume": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "encrypted": {
						"type": "bool",
						"optional": true
					  },
					  "iops": {
						"type": "number",
						"optional": true
					  },
					  "mount_point": {
						"type": "string",
						"required": true
					  },
					  "number_of_disks": {
						"type": "number",
						"required": true
					  },
					  "raid_level": {
						"type": "string",
						"optional": true
					  },
					  "size": {
						"type": "number",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_opsworks_ganglia_layer": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_assign_elastic_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_assign_public_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_healing": {
				  "type": "bool",
				  "optional": true
				},
				"custom_configure_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_deploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_instance_profile_arn": {
				  "type": "string",
				  "optional": true
				},
				"custom_json": {
				  "type": "string",
				  "optional": true
				},
				"custom_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"custom_setup_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_shutdown_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_undeploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"drain_elb_on_shutdown": {
				  "type": "bool",
				  "optional": true
				},
				"elastic_load_balancer": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"install_updates_on_boot": {
				  "type": "bool",
				  "optional": true
				},
				"instance_shutdown_timeout": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"password": {
				  "type": "string",
				  "required": true
				},
				"stack_id": {
				  "type": "string",
				  "required": true
				},
				"system_packages": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"url": {
				  "type": "string",
				  "optional": true
				},
				"use_ebs_optimized_instances": {
				  "type": "bool",
				  "optional": true
				},
				"username": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"ebs_volume": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "encrypted": {
						"type": "bool",
						"optional": true
					  },
					  "iops": {
						"type": "number",
						"optional": true
					  },
					  "mount_point": {
						"type": "string",
						"required": true
					  },
					  "number_of_disks": {
						"type": "number",
						"required": true
					  },
					  "raid_level": {
						"type": "string",
						"optional": true
					  },
					  "size": {
						"type": "number",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_opsworks_haproxy_layer": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_assign_elastic_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_assign_public_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_healing": {
				  "type": "bool",
				  "optional": true
				},
				"custom_configure_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_deploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_instance_profile_arn": {
				  "type": "string",
				  "optional": true
				},
				"custom_json": {
				  "type": "string",
				  "optional": true
				},
				"custom_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"custom_setup_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_shutdown_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_undeploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"drain_elb_on_shutdown": {
				  "type": "bool",
				  "optional": true
				},
				"elastic_load_balancer": {
				  "type": "string",
				  "optional": true
				},
				"healthcheck_method": {
				  "type": "string",
				  "optional": true
				},
				"healthcheck_url": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"install_updates_on_boot": {
				  "type": "bool",
				  "optional": true
				},
				"instance_shutdown_timeout": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"stack_id": {
				  "type": "string",
				  "required": true
				},
				"stats_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"stats_password": {
				  "type": "string",
				  "required": true
				},
				"stats_url": {
				  "type": "string",
				  "optional": true
				},
				"stats_user": {
				  "type": "string",
				  "optional": true
				},
				"system_packages": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"use_ebs_optimized_instances": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"ebs_volume": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "encrypted": {
						"type": "bool",
						"optional": true
					  },
					  "iops": {
						"type": "number",
						"optional": true
					  },
					  "mount_point": {
						"type": "string",
						"required": true
					  },
					  "number_of_disks": {
						"type": "number",
						"required": true
					  },
					  "raid_level": {
						"type": "string",
						"optional": true
					  },
					  "size": {
						"type": "number",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_opsworks_instance": {
			"version": 0,
			"block": {
			  "attributes": {
				"agent_version": {
				  "type": "string",
				  "optional": true
				},
				"ami_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"architecture": {
				  "type": "string",
				  "optional": true
				},
				"auto_scaling_type": {
				  "type": "string",
				  "optional": true
				},
				"availability_zone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"created_at": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"delete_ebs": {
				  "type": "bool",
				  "optional": true
				},
				"delete_eip": {
				  "type": "bool",
				  "optional": true
				},
				"ebs_optimized": {
				  "type": "bool",
				  "optional": true
				},
				"ec2_instance_id": {
				  "type": "string",
				  "computed": true
				},
				"ecs_cluster_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"elastic_ip": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"hostname": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"infrastructure_class": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"install_updates_on_boot": {
				  "type": "bool",
				  "optional": true
				},
				"instance_profile_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_type": {
				  "type": "string",
				  "optional": true
				},
				"last_service_error_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"layer_ids": {
				  "type": [
					"list",
					"string"
				  ],
				  "required": true
				},
				"os": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"platform": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"private_dns": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"private_ip": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"public_dns": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"public_ip": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"registered_by": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"reported_agent_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"reported_os_family": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"reported_os_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"reported_os_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"root_device_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"root_device_volume_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"security_group_ids": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"ssh_host_dsa_key_fingerprint": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ssh_host_rsa_key_fingerprint": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ssh_key_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"stack_id": {
				  "type": "string",
				  "required": true
				},
				"state": {
				  "type": "string",
				  "optional": true
				},
				"status": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"subnet_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tenancy": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"virtualization_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"ebs_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "delete_on_termination": {
						"type": "bool",
						"optional": true
					  },
					  "device_name": {
						"type": "string",
						"required": true
					  },
					  "iops": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "snapshot_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "volume_size": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "volume_type": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  }
				},
				"ephemeral_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "device_name": {
						"type": "string",
						"required": true
					  },
					  "virtual_name": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"root_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "delete_on_termination": {
						"type": "bool",
						"optional": true
					  },
					  "iops": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "volume_size": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "volume_type": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  }
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_opsworks_java_app_layer": {
			"version": 0,
			"block": {
			  "attributes": {
				"app_server": {
				  "type": "string",
				  "optional": true
				},
				"app_server_version": {
				  "type": "string",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_assign_elastic_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_assign_public_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_healing": {
				  "type": "bool",
				  "optional": true
				},
				"custom_configure_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_deploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_instance_profile_arn": {
				  "type": "string",
				  "optional": true
				},
				"custom_json": {
				  "type": "string",
				  "optional": true
				},
				"custom_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"custom_setup_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_shutdown_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_undeploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"drain_elb_on_shutdown": {
				  "type": "bool",
				  "optional": true
				},
				"elastic_load_balancer": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"install_updates_on_boot": {
				  "type": "bool",
				  "optional": true
				},
				"instance_shutdown_timeout": {
				  "type": "number",
				  "optional": true
				},
				"jvm_options": {
				  "type": "string",
				  "optional": true
				},
				"jvm_type": {
				  "type": "string",
				  "optional": true
				},
				"jvm_version": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"stack_id": {
				  "type": "string",
				  "required": true
				},
				"system_packages": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"use_ebs_optimized_instances": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"ebs_volume": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "encrypted": {
						"type": "bool",
						"optional": true
					  },
					  "iops": {
						"type": "number",
						"optional": true
					  },
					  "mount_point": {
						"type": "string",
						"required": true
					  },
					  "number_of_disks": {
						"type": "number",
						"required": true
					  },
					  "raid_level": {
						"type": "string",
						"optional": true
					  },
					  "size": {
						"type": "number",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_opsworks_memcached_layer": {
			"version": 0,
			"block": {
			  "attributes": {
				"allocated_memory": {
				  "type": "number",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_assign_elastic_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_assign_public_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_healing": {
				  "type": "bool",
				  "optional": true
				},
				"custom_configure_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_deploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_instance_profile_arn": {
				  "type": "string",
				  "optional": true
				},
				"custom_json": {
				  "type": "string",
				  "optional": true
				},
				"custom_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"custom_setup_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_shutdown_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_undeploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"drain_elb_on_shutdown": {
				  "type": "bool",
				  "optional": true
				},
				"elastic_load_balancer": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"install_updates_on_boot": {
				  "type": "bool",
				  "optional": true
				},
				"instance_shutdown_timeout": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"stack_id": {
				  "type": "string",
				  "required": true
				},
				"system_packages": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"use_ebs_optimized_instances": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"ebs_volume": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "encrypted": {
						"type": "bool",
						"optional": true
					  },
					  "iops": {
						"type": "number",
						"optional": true
					  },
					  "mount_point": {
						"type": "string",
						"required": true
					  },
					  "number_of_disks": {
						"type": "number",
						"required": true
					  },
					  "raid_level": {
						"type": "string",
						"optional": true
					  },
					  "size": {
						"type": "number",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_opsworks_mysql_layer": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_assign_elastic_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_assign_public_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_healing": {
				  "type": "bool",
				  "optional": true
				},
				"custom_configure_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_deploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_instance_profile_arn": {
				  "type": "string",
				  "optional": true
				},
				"custom_json": {
				  "type": "string",
				  "optional": true
				},
				"custom_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"custom_setup_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_shutdown_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_undeploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"drain_elb_on_shutdown": {
				  "type": "bool",
				  "optional": true
				},
				"elastic_load_balancer": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"install_updates_on_boot": {
				  "type": "bool",
				  "optional": true
				},
				"instance_shutdown_timeout": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"root_password": {
				  "type": "string",
				  "optional": true
				},
				"root_password_on_all_instances": {
				  "type": "bool",
				  "optional": true
				},
				"stack_id": {
				  "type": "string",
				  "required": true
				},
				"system_packages": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"use_ebs_optimized_instances": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"ebs_volume": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "encrypted": {
						"type": "bool",
						"optional": true
					  },
					  "iops": {
						"type": "number",
						"optional": true
					  },
					  "mount_point": {
						"type": "string",
						"required": true
					  },
					  "number_of_disks": {
						"type": "number",
						"required": true
					  },
					  "raid_level": {
						"type": "string",
						"optional": true
					  },
					  "size": {
						"type": "number",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_opsworks_nodejs_app_layer": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_assign_elastic_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_assign_public_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_healing": {
				  "type": "bool",
				  "optional": true
				},
				"custom_configure_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_deploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_instance_profile_arn": {
				  "type": "string",
				  "optional": true
				},
				"custom_json": {
				  "type": "string",
				  "optional": true
				},
				"custom_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"custom_setup_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_shutdown_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_undeploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"drain_elb_on_shutdown": {
				  "type": "bool",
				  "optional": true
				},
				"elastic_load_balancer": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"install_updates_on_boot": {
				  "type": "bool",
				  "optional": true
				},
				"instance_shutdown_timeout": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"nodejs_version": {
				  "type": "string",
				  "optional": true
				},
				"stack_id": {
				  "type": "string",
				  "required": true
				},
				"system_packages": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"use_ebs_optimized_instances": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"ebs_volume": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "encrypted": {
						"type": "bool",
						"optional": true
					  },
					  "iops": {
						"type": "number",
						"optional": true
					  },
					  "mount_point": {
						"type": "string",
						"required": true
					  },
					  "number_of_disks": {
						"type": "number",
						"required": true
					  },
					  "raid_level": {
						"type": "string",
						"optional": true
					  },
					  "size": {
						"type": "number",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_opsworks_permission": {
			"version": 0,
			"block": {
			  "attributes": {
				"allow_ssh": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"allow_sudo": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"level": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"stack_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"user_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_opsworks_php_app_layer": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_assign_elastic_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_assign_public_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_healing": {
				  "type": "bool",
				  "optional": true
				},
				"custom_configure_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_deploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_instance_profile_arn": {
				  "type": "string",
				  "optional": true
				},
				"custom_json": {
				  "type": "string",
				  "optional": true
				},
				"custom_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"custom_setup_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_shutdown_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_undeploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"drain_elb_on_shutdown": {
				  "type": "bool",
				  "optional": true
				},
				"elastic_load_balancer": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"install_updates_on_boot": {
				  "type": "bool",
				  "optional": true
				},
				"instance_shutdown_timeout": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"stack_id": {
				  "type": "string",
				  "required": true
				},
				"system_packages": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"use_ebs_optimized_instances": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"ebs_volume": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "encrypted": {
						"type": "bool",
						"optional": true
					  },
					  "iops": {
						"type": "number",
						"optional": true
					  },
					  "mount_point": {
						"type": "string",
						"required": true
					  },
					  "number_of_disks": {
						"type": "number",
						"required": true
					  },
					  "raid_level": {
						"type": "string",
						"optional": true
					  },
					  "size": {
						"type": "number",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_opsworks_rails_app_layer": {
			"version": 0,
			"block": {
			  "attributes": {
				"app_server": {
				  "type": "string",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_assign_elastic_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_assign_public_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_healing": {
				  "type": "bool",
				  "optional": true
				},
				"bundler_version": {
				  "type": "string",
				  "optional": true
				},
				"custom_configure_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_deploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_instance_profile_arn": {
				  "type": "string",
				  "optional": true
				},
				"custom_json": {
				  "type": "string",
				  "optional": true
				},
				"custom_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"custom_setup_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_shutdown_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_undeploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"drain_elb_on_shutdown": {
				  "type": "bool",
				  "optional": true
				},
				"elastic_load_balancer": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"install_updates_on_boot": {
				  "type": "bool",
				  "optional": true
				},
				"instance_shutdown_timeout": {
				  "type": "number",
				  "optional": true
				},
				"manage_bundler": {
				  "type": "bool",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"passenger_version": {
				  "type": "string",
				  "optional": true
				},
				"ruby_version": {
				  "type": "string",
				  "optional": true
				},
				"rubygems_version": {
				  "type": "string",
				  "optional": true
				},
				"stack_id": {
				  "type": "string",
				  "required": true
				},
				"system_packages": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"use_ebs_optimized_instances": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"ebs_volume": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "encrypted": {
						"type": "bool",
						"optional": true
					  },
					  "iops": {
						"type": "number",
						"optional": true
					  },
					  "mount_point": {
						"type": "string",
						"required": true
					  },
					  "number_of_disks": {
						"type": "number",
						"required": true
					  },
					  "raid_level": {
						"type": "string",
						"optional": true
					  },
					  "size": {
						"type": "number",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_opsworks_rds_db_instance": {
			"version": 0,
			"block": {
			  "attributes": {
				"db_password": {
				  "type": "string",
				  "required": true,
				  "sensitive": true
				},
				"db_user": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"rds_db_instance_arn": {
				  "type": "string",
				  "required": true
				},
				"stack_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_opsworks_stack": {
			"version": 0,
			"block": {
			  "attributes": {
				"agent_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"berkshelf_version": {
				  "type": "string",
				  "optional": true
				},
				"color": {
				  "type": "string",
				  "optional": true
				},
				"configuration_manager_name": {
				  "type": "string",
				  "optional": true
				},
				"configuration_manager_version": {
				  "type": "string",
				  "optional": true
				},
				"custom_json": {
				  "type": "string",
				  "optional": true
				},
				"default_availability_zone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"default_instance_profile_arn": {
				  "type": "string",
				  "required": true
				},
				"default_os": {
				  "type": "string",
				  "optional": true
				},
				"default_root_device_type": {
				  "type": "string",
				  "optional": true
				},
				"default_ssh_key_name": {
				  "type": "string",
				  "optional": true
				},
				"default_subnet_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"hostname_theme": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"manage_berkshelf": {
				  "type": "bool",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"region": {
				  "type": "string",
				  "required": true
				},
				"service_role_arn": {
				  "type": "string",
				  "required": true
				},
				"stack_endpoint": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"use_custom_cookbooks": {
				  "type": "bool",
				  "optional": true
				},
				"use_opsworks_security_groups": {
				  "type": "bool",
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"custom_cookbooks_source": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "password": {
						"type": "string",
						"optional": true,
						"sensitive": true
					  },
					  "revision": {
						"type": "string",
						"optional": true
					  },
					  "ssh_key": {
						"type": "string",
						"optional": true,
						"sensitive": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  },
					  "url": {
						"type": "string",
						"required": true
					  },
					  "username": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_opsworks_static_web_layer": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_assign_elastic_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_assign_public_ips": {
				  "type": "bool",
				  "optional": true
				},
				"auto_healing": {
				  "type": "bool",
				  "optional": true
				},
				"custom_configure_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_deploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_instance_profile_arn": {
				  "type": "string",
				  "optional": true
				},
				"custom_json": {
				  "type": "string",
				  "optional": true
				},
				"custom_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"custom_setup_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_shutdown_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"custom_undeploy_recipes": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"drain_elb_on_shutdown": {
				  "type": "bool",
				  "optional": true
				},
				"elastic_load_balancer": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"install_updates_on_boot": {
				  "type": "bool",
				  "optional": true
				},
				"instance_shutdown_timeout": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"stack_id": {
				  "type": "string",
				  "required": true
				},
				"system_packages": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"use_ebs_optimized_instances": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"ebs_volume": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "encrypted": {
						"type": "bool",
						"optional": true
					  },
					  "iops": {
						"type": "number",
						"optional": true
					  },
					  "mount_point": {
						"type": "string",
						"required": true
					  },
					  "number_of_disks": {
						"type": "number",
						"required": true
					  },
					  "raid_level": {
						"type": "string",
						"optional": true
					  },
					  "size": {
						"type": "number",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_opsworks_user_profile": {
			"version": 0,
			"block": {
			  "attributes": {
				"allow_self_management": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ssh_public_key": {
				  "type": "string",
				  "optional": true
				},
				"ssh_username": {
				  "type": "string",
				  "required": true
				},
				"user_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_organizations_account": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"email": {
				  "type": "string",
				  "required": true
				},
				"iam_user_access_to_billing": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"joined_method": {
				  "type": "string",
				  "computed": true
				},
				"joined_timestamp": {
				  "type": "string",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"parent_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"role_name": {
				  "type": "string",
				  "optional": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_organizations_organization": {
			"version": 0,
			"block": {
			  "attributes": {
				"accounts": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"arn": "string",
						"email": "string",
						"id": "string",
						"name": "string",
						"status": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"aws_service_access_principals": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"enabled_policy_types": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"feature_set": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"master_account_arn": {
				  "type": "string",
				  "computed": true
				},
				"master_account_email": {
				  "type": "string",
				  "computed": true
				},
				"master_account_id": {
				  "type": "string",
				  "computed": true
				},
				"non_master_accounts": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"arn": "string",
						"email": "string",
						"id": "string",
						"name": "string",
						"status": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"roots": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"arn": "string",
						"id": "string",
						"name": "string",
						"policy_types": [
						  "list",
						  [
							"object",
							{
							  "status": "string",
							  "type": "string"
							}
						  ]
						]
					  }
					]
				  ],
				  "computed": true
				}
			  }
			}
		  },
		  "aws_organizations_organizational_unit": {
			"version": 0,
			"block": {
			  "attributes": {
				"accounts": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"arn": "string",
						"email": "string",
						"id": "string",
						"name": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"parent_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_organizations_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"content": {
				  "type": "string",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"type": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_organizations_policy_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy_id": {
				  "type": "string",
				  "required": true
				},
				"target_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_pinpoint_adm_channel": {
			"version": 0,
			"block": {
			  "attributes": {
				"application_id": {
				  "type": "string",
				  "required": true
				},
				"client_id": {
				  "type": "string",
				  "required": true,
				  "sensitive": true
				},
				"client_secret": {
				  "type": "string",
				  "required": true,
				  "sensitive": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_pinpoint_apns_channel": {
			"version": 0,
			"block": {
			  "attributes": {
				"application_id": {
				  "type": "string",
				  "required": true
				},
				"bundle_id": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"certificate": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"default_authentication_method": {
				  "type": "string",
				  "optional": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"private_key": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"team_id": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"token_key": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"token_key_id": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				}
			  }
			}
		  },
		  "aws_pinpoint_apns_sandbox_channel": {
			"version": 0,
			"block": {
			  "attributes": {
				"application_id": {
				  "type": "string",
				  "required": true
				},
				"bundle_id": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"certificate": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"default_authentication_method": {
				  "type": "string",
				  "optional": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"private_key": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"team_id": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"token_key": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"token_key_id": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				}
			  }
			}
		  },
		  "aws_pinpoint_apns_voip_channel": {
			"version": 0,
			"block": {
			  "attributes": {
				"application_id": {
				  "type": "string",
				  "required": true
				},
				"bundle_id": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"certificate": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"default_authentication_method": {
				  "type": "string",
				  "optional": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"private_key": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"team_id": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"token_key": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"token_key_id": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				}
			  }
			}
		  },
		  "aws_pinpoint_apns_voip_sandbox_channel": {
			"version": 0,
			"block": {
			  "attributes": {
				"application_id": {
				  "type": "string",
				  "required": true
				},
				"bundle_id": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"certificate": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"default_authentication_method": {
				  "type": "string",
				  "optional": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"private_key": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"team_id": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"token_key": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"token_key_id": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				}
			  }
			}
		  },
		  "aws_pinpoint_app": {
			"version": 0,
			"block": {
			  "attributes": {
				"application_id": {
				  "type": "string",
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"campaign_hook": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "lambda_function_name": {
						"type": "string",
						"optional": true
					  },
					  "mode": {
						"type": "string",
						"optional": true
					  },
					  "web_url": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"limits": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "daily": {
						"type": "number",
						"optional": true
					  },
					  "maximum_duration": {
						"type": "number",
						"optional": true
					  },
					  "messages_per_second": {
						"type": "number",
						"optional": true
					  },
					  "total": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"quiet_time": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "end": {
						"type": "string",
						"optional": true
					  },
					  "start": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_pinpoint_baidu_channel": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_key": {
				  "type": "string",
				  "required": true,
				  "sensitive": true
				},
				"application_id": {
				  "type": "string",
				  "required": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"secret_key": {
				  "type": "string",
				  "required": true,
				  "sensitive": true
				}
			  }
			}
		  },
		  "aws_pinpoint_email_channel": {
			"version": 0,
			"block": {
			  "attributes": {
				"application_id": {
				  "type": "string",
				  "required": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"from_address": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identity": {
				  "type": "string",
				  "required": true
				},
				"messages_per_second": {
				  "type": "number",
				  "computed": true
				},
				"role_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_pinpoint_event_stream": {
			"version": 0,
			"block": {
			  "attributes": {
				"application_id": {
				  "type": "string",
				  "required": true
				},
				"destination_stream_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"role_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_pinpoint_gcm_channel": {
			"version": 0,
			"block": {
			  "attributes": {
				"api_key": {
				  "type": "string",
				  "required": true,
				  "sensitive": true
				},
				"application_id": {
				  "type": "string",
				  "required": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_pinpoint_sms_channel": {
			"version": 0,
			"block": {
			  "attributes": {
				"application_id": {
				  "type": "string",
				  "required": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"promotional_messages_per_second": {
				  "type": "number",
				  "computed": true
				},
				"sender_id": {
				  "type": "string",
				  "optional": true
				},
				"short_code": {
				  "type": "string",
				  "optional": true
				},
				"transactional_messages_per_second": {
				  "type": "number",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_placement_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"placement_group_id": {
				  "type": "string",
				  "computed": true
				},
				"strategy": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_proxy_protocol_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_ports": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"load_balancer": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_qldb_ledger": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"deletion_protection": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_quicksight_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"aws_account_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"group_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"namespace": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_quicksight_user": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"aws_account_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"email": {
				  "type": "string",
				  "required": true
				},
				"iam_arn": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identity_type": {
				  "type": "string",
				  "required": true
				},
				"namespace": {
				  "type": "string",
				  "optional": true
				},
				"session_name": {
				  "type": "string",
				  "optional": true
				},
				"user_name": {
				  "type": "string",
				  "optional": true
				},
				"user_role": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ram_principal_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"principal": {
				  "type": "string",
				  "required": true
				},
				"resource_share_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ram_resource_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"resource_arn": {
				  "type": "string",
				  "required": true
				},
				"resource_share_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ram_resource_share": {
			"version": 0,
			"block": {
			  "attributes": {
				"allow_external_principals": {
				  "type": "bool",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_ram_resource_share_accepter": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"invitation_arn": {
				  "type": "string",
				  "computed": true
				},
				"receiver_account_id": {
				  "type": "string",
				  "computed": true
				},
				"resources": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"sender_account_id": {
				  "type": "string",
				  "computed": true
				},
				"share_arn": {
				  "type": "string",
				  "required": true
				},
				"share_id": {
				  "type": "string",
				  "computed": true
				},
				"share_name": {
				  "type": "string",
				  "computed": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_rds_cluster": {
			"version": 0,
			"block": {
			  "attributes": {
				"apply_immediately": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"availability_zones": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"backtrack_window": {
				  "type": "number",
				  "optional": true
				},
				"backup_retention_period": {
				  "type": "number",
				  "optional": true
				},
				"cluster_identifier": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_identifier_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_members": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"cluster_resource_id": {
				  "type": "string",
				  "computed": true
				},
				"copy_tags_to_snapshot": {
				  "type": "bool",
				  "optional": true
				},
				"database_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"db_cluster_parameter_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"db_subnet_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"deletion_protection": {
				  "type": "bool",
				  "optional": true
				},
				"enable_http_endpoint": {
				  "type": "bool",
				  "optional": true
				},
				"enabled_cloudwatch_logs_exports": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"endpoint": {
				  "type": "string",
				  "computed": true
				},
				"engine": {
				  "type": "string",
				  "optional": true
				},
				"engine_mode": {
				  "type": "string",
				  "optional": true
				},
				"engine_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"final_snapshot_identifier": {
				  "type": "string",
				  "optional": true
				},
				"global_cluster_identifier": {
				  "type": "string",
				  "optional": true
				},
				"hosted_zone_id": {
				  "type": "string",
				  "computed": true
				},
				"iam_database_authentication_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"iam_roles": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"master_password": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"master_username": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"preferred_backup_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"preferred_maintenance_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"reader_endpoint": {
				  "type": "string",
				  "computed": true
				},
				"replication_source_identifier": {
				  "type": "string",
				  "optional": true
				},
				"skip_final_snapshot": {
				  "type": "bool",
				  "optional": true
				},
				"snapshot_identifier": {
				  "type": "string",
				  "optional": true
				},
				"source_region": {
				  "type": "string",
				  "optional": true
				},
				"storage_encrypted": {
				  "type": "bool",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"s3_import": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket_name": {
						"type": "string",
						"required": true
					  },
					  "bucket_prefix": {
						"type": "string",
						"optional": true
					  },
					  "ingestion_role": {
						"type": "string",
						"required": true
					  },
					  "source_engine": {
						"type": "string",
						"required": true
					  },
					  "source_engine_version": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"scaling_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "auto_pause": {
						"type": "bool",
						"optional": true
					  },
					  "max_capacity": {
						"type": "number",
						"optional": true
					  },
					  "min_capacity": {
						"type": "number",
						"optional": true
					  },
					  "seconds_until_auto_pause": {
						"type": "number",
						"optional": true
					  },
					  "timeout_action": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_rds_cluster_endpoint": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"cluster_endpoint_identifier": {
				  "type": "string",
				  "required": true
				},
				"cluster_identifier": {
				  "type": "string",
				  "required": true
				},
				"custom_endpoint_type": {
				  "type": "string",
				  "required": true
				},
				"endpoint": {
				  "type": "string",
				  "computed": true
				},
				"excluded_members": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"static_members": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_rds_cluster_instance": {
			"version": 0,
			"block": {
			  "attributes": {
				"apply_immediately": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_minor_version_upgrade": {
				  "type": "bool",
				  "optional": true
				},
				"availability_zone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ca_cert_identifier": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_identifier": {
				  "type": "string",
				  "required": true
				},
				"copy_tags_to_snapshot": {
				  "type": "bool",
				  "optional": true
				},
				"db_parameter_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"db_subnet_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"dbi_resource_id": {
				  "type": "string",
				  "computed": true
				},
				"endpoint": {
				  "type": "string",
				  "computed": true
				},
				"engine": {
				  "type": "string",
				  "optional": true
				},
				"engine_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identifier": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identifier_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_class": {
				  "type": "string",
				  "required": true
				},
				"kms_key_id": {
				  "type": "string",
				  "computed": true
				},
				"monitoring_interval": {
				  "type": "number",
				  "optional": true
				},
				"monitoring_role_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"performance_insights_enabled": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"performance_insights_kms_key_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"port": {
				  "type": "number",
				  "computed": true
				},
				"preferred_backup_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"preferred_maintenance_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"promotion_tier": {
				  "type": "number",
				  "optional": true
				},
				"publicly_accessible": {
				  "type": "bool",
				  "optional": true
				},
				"storage_encrypted": {
				  "type": "bool",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"writer": {
				  "type": "bool",
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_rds_cluster_parameter_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"family": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"parameter": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "apply_method": {
						"type": "string",
						"optional": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_rds_global_cluster": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"database_name": {
				  "type": "string",
				  "optional": true
				},
				"deletion_protection": {
				  "type": "bool",
				  "optional": true
				},
				"engine": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"engine_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"force_destroy": {
				  "type": "bool",
				  "optional": true
				},
				"global_cluster_identifier": {
				  "type": "string",
				  "required": true
				},
				"global_cluster_members": {
				  "type": [
					"set",
					[
					  "object",
					  {
						"db_cluster_arn": "string",
						"is_writer": "bool"
					  }
					]
				  ],
				  "computed": true
				},
				"global_cluster_resource_id": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"source_db_cluster_identifier": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"storage_encrypted": {
				  "type": "bool",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_redshift_cluster": {
			"version": 0,
			"block": {
			  "attributes": {
				"allow_version_upgrade": {
				  "type": "bool",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"automated_snapshot_retention_period": {
				  "type": "number",
				  "optional": true
				},
				"availability_zone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_identifier": {
				  "type": "string",
				  "required": true
				},
				"cluster_parameter_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_public_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_revision_number": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_security_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"cluster_subnet_group_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cluster_version": {
				  "type": "string",
				  "optional": true
				},
				"database_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"dns_name": {
				  "type": "string",
				  "computed": true
				},
				"elastic_ip": {
				  "type": "string",
				  "optional": true
				},
				"encrypted": {
				  "type": "bool",
				  "optional": true
				},
				"endpoint": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"enhanced_vpc_routing": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"final_snapshot_identifier": {
				  "type": "string",
				  "optional": true
				},
				"iam_roles": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"master_password": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"master_username": {
				  "type": "string",
				  "optional": true
				},
				"node_type": {
				  "type": "string",
				  "required": true
				},
				"number_of_nodes": {
				  "type": "number",
				  "optional": true
				},
				"owner_account": {
				  "type": "string",
				  "optional": true
				},
				"port": {
				  "type": "number",
				  "optional": true
				},
				"preferred_maintenance_window": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"publicly_accessible": {
				  "type": "bool",
				  "optional": true
				},
				"skip_final_snapshot": {
				  "type": "bool",
				  "optional": true
				},
				"snapshot_cluster_identifier": {
				  "type": "string",
				  "optional": true
				},
				"snapshot_identifier": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"logging": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket_name": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "enable": {
						"type": "bool",
						"required": true
					  },
					  "s3_key_prefix": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				},
				"snapshot_copy": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "destination_region": {
						"type": "string",
						"required": true
					  },
					  "grant_name": {
						"type": "string",
						"optional": true
					  },
					  "retention_period": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_redshift_event_subscription": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"customer_aws_id": {
				  "type": "string",
				  "computed": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"event_categories": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"severity": {
				  "type": "string",
				  "optional": true
				},
				"sns_topic_arn": {
				  "type": "string",
				  "required": true
				},
				"source_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"source_type": {
				  "type": "string",
				  "optional": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_redshift_parameter_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"family": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"parameter": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "name": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_redshift_security_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"ingress": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "cidr": {
						"type": "string",
						"optional": true
					  },
					  "security_group_name": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "security_group_owner_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "min_items": 1
				}
			  }
			}
		  },
		  "aws_redshift_snapshot_copy_grant": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"snapshot_copy_grant_name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_redshift_snapshot_schedule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"definitions": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"force_destroy": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identifier": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identifier_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_redshift_snapshot_schedule_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"cluster_identifier": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"schedule_identifier": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_redshift_subnet_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_resourcegroups_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"resource_query": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "query": {
						"type": "string",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_route": {
			"version": 0,
			"block": {
			  "attributes": {
				"destination_cidr_block": {
				  "type": "string",
				  "optional": true
				},
				"destination_ipv6_cidr_block": {
				  "type": "string",
				  "optional": true
				},
				"destination_prefix_list_id": {
				  "type": "string",
				  "computed": true
				},
				"egress_only_gateway_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"gateway_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_owner_id": {
				  "type": "string",
				  "computed": true
				},
				"nat_gateway_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"network_interface_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"origin": {
				  "type": "string",
				  "computed": true
				},
				"route_table_id": {
				  "type": "string",
				  "required": true
				},
				"state": {
				  "type": "string",
				  "computed": true
				},
				"transit_gateway_id": {
				  "type": "string",
				  "optional": true
				},
				"vpc_peering_connection_id": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_route53_delegation_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_servers": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"reference_name": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_route53_health_check": {
			"version": 0,
			"block": {
			  "attributes": {
				"child_health_threshold": {
				  "type": "number",
				  "optional": true
				},
				"child_healthchecks": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"cloudwatch_alarm_name": {
				  "type": "string",
				  "optional": true
				},
				"cloudwatch_alarm_region": {
				  "type": "string",
				  "optional": true
				},
				"disabled": {
				  "type": "bool",
				  "optional": true
				},
				"enable_sni": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"failure_threshold": {
				  "type": "number",
				  "optional": true
				},
				"fqdn": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"insufficient_data_health_status": {
				  "type": "string",
				  "optional": true
				},
				"invert_healthcheck": {
				  "type": "bool",
				  "optional": true
				},
				"ip_address": {
				  "type": "string",
				  "optional": true
				},
				"measure_latency": {
				  "type": "bool",
				  "optional": true
				},
				"port": {
				  "type": "number",
				  "optional": true
				},
				"reference_name": {
				  "type": "string",
				  "optional": true
				},
				"regions": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"request_interval": {
				  "type": "number",
				  "optional": true
				},
				"resource_path": {
				  "type": "string",
				  "optional": true
				},
				"search_string": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"type": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_route53_query_log": {
			"version": 0,
			"block": {
			  "attributes": {
				"cloudwatch_log_group_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"zone_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_route53_record": {
			"version": 2,
			"block": {
			  "attributes": {
				"allow_overwrite": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"fqdn": {
				  "type": "string",
				  "computed": true
				},
				"health_check_id": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"multivalue_answer_routing_policy": {
				  "type": "bool",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"records": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"set_identifier": {
				  "type": "string",
				  "optional": true
				},
				"ttl": {
				  "type": "number",
				  "optional": true
				},
				"type": {
				  "type": "string",
				  "required": true
				},
				"zone_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"alias": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "evaluate_target_health": {
						"type": "bool",
						"required": true
					  },
					  "name": {
						"type": "string",
						"required": true
					  },
					  "zone_id": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"failover_routing_policy": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"geolocation_routing_policy": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "continent": {
						"type": "string",
						"optional": true
					  },
					  "country": {
						"type": "string",
						"optional": true
					  },
					  "subdivision": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"latency_routing_policy": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "region": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"weighted_routing_policy": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "weight": {
						"type": "number",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_route53_resolver_endpoint": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"direction": {
				  "type": "string",
				  "required": true
				},
				"host_vpc_id": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"ip_address": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "ip": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "ip_id": {
						"type": "string",
						"computed": true
					  },
					  "subnet_id": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 2,
				  "max_items": 10
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_route53_resolver_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"domain_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"resolver_endpoint_id": {
				  "type": "string",
				  "optional": true
				},
				"rule_type": {
				  "type": "string",
				  "required": true
				},
				"share_status": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"target_ip": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "ip": {
						"type": "string",
						"required": true
					  },
					  "port": {
						"type": "number",
						"optional": true
					  }
					}
				  }
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_route53_resolver_rule_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"resolver_rule_id": {
				  "type": "string",
				  "required": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_route53_vpc_association_authorization": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				},
				"vpc_region": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"zone_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_route53_zone": {
			"version": 0,
			"block": {
			  "attributes": {
				"comment": {
				  "type": "string",
				  "optional": true
				},
				"delegation_set_id": {
				  "type": "string",
				  "optional": true
				},
				"force_destroy": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"name_servers": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"zone_id": {
				  "type": "string",
				  "computed": true
				}
			  },
			  "block_types": {
				"vpc": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "vpc_id": {
						"type": "string",
						"required": true
					  },
					  "vpc_region": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_route53_zone_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"owning_account": {
				  "type": "string",
				  "computed": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				},
				"vpc_region": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"zone_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_route_table": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"propagating_vgws": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"route": {
				  "type": [
					"set",
					[
					  "object",
					  {
						"cidr_block": "string",
						"egress_only_gateway_id": "string",
						"gateway_id": "string",
						"instance_id": "string",
						"ipv6_cidr_block": "string",
						"nat_gateway_id": "string",
						"network_interface_id": "string",
						"transit_gateway_id": "string",
						"vpc_peering_connection_id": "string"
					  }
					]
				  ],
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_route_table_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"gateway_id": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"route_table_id": {
				  "type": "string",
				  "required": true
				},
				"subnet_id": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_s3_access_point": {
			"version": 0,
			"block": {
			  "attributes": {
				"account_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"bucket": {
				  "type": "string",
				  "required": true
				},
				"domain_name": {
				  "type": "string",
				  "computed": true
				},
				"has_public_access_policy": {
				  "type": "bool",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"network_origin": {
				  "type": "string",
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"public_access_block_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "block_public_acls": {
						"type": "bool",
						"optional": true
					  },
					  "block_public_policy": {
						"type": "bool",
						"optional": true
					  },
					  "ignore_public_acls": {
						"type": "bool",
						"optional": true
					  },
					  "restrict_public_buckets": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"vpc_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "vpc_id": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_s3_account_public_access_block": {
			"version": 0,
			"block": {
			  "attributes": {
				"account_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"block_public_acls": {
				  "type": "bool",
				  "optional": true
				},
				"block_public_policy": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ignore_public_acls": {
				  "type": "bool",
				  "optional": true
				},
				"restrict_public_buckets": {
				  "type": "bool",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_s3_bucket": {
			"version": 0,
			"block": {
			  "attributes": {
				"acceleration_status": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"acl": {
				  "type": "string",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"bucket": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"bucket_domain_name": {
				  "type": "string",
				  "computed": true
				},
				"bucket_prefix": {
				  "type": "string",
				  "optional": true
				},
				"bucket_regional_domain_name": {
				  "type": "string",
				  "computed": true
				},
				"force_destroy": {
				  "type": "bool",
				  "optional": true
				},
				"hosted_zone_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "optional": true
				},
				"region": {
				  "type": "string",
				  "computed": true
				},
				"request_payer": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"website_domain": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"website_endpoint": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"cors_rule": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allowed_headers": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "allowed_methods": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  },
					  "allowed_origins": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  },
					  "expose_headers": {
						"type": [
						  "list",
						  "string"
						],
						"optional": true
					  },
					  "max_age_seconds": {
						"type": "number",
						"optional": true
					  }
					}
				  }
				},
				"grant": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "id": {
						"type": "string",
						"optional": true
					  },
					  "permissions": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  },
					  "uri": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"lifecycle_rule": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "abort_incomplete_multipart_upload_days": {
						"type": "number",
						"optional": true
					  },
					  "enabled": {
						"type": "bool",
						"required": true
					  },
					  "id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "prefix": {
						"type": "string",
						"optional": true
					  },
					  "tags": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  }
					},
					"block_types": {
					  "expiration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"date": {
							  "type": "string",
							  "optional": true
							},
							"days": {
							  "type": "number",
							  "optional": true
							},
							"expired_object_delete_marker": {
							  "type": "bool",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "noncurrent_version_expiration": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"days": {
							  "type": "number",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "noncurrent_version_transition": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"days": {
							  "type": "number",
							  "optional": true
							},
							"storage_class": {
							  "type": "string",
							  "required": true
							}
						  }
						}
					  },
					  "transition": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"date": {
							  "type": "string",
							  "optional": true
							},
							"days": {
							  "type": "number",
							  "optional": true
							},
							"storage_class": {
							  "type": "string",
							  "required": true
							}
						  }
						}
					  }
					}
				  }
				},
				"logging": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "target_bucket": {
						"type": "string",
						"required": true
					  },
					  "target_prefix": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"object_lock_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "object_lock_enabled": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "rule": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"default_retention": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "days": {
									"type": "number",
									"optional": true
								  },
								  "mode": {
									"type": "string",
									"required": true
								  },
								  "years": {
									"type": "number",
									"optional": true
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"replication_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "role": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "rules": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"id": {
							  "type": "string",
							  "optional": true
							},
							"prefix": {
							  "type": "string",
							  "optional": true
							},
							"priority": {
							  "type": "number",
							  "optional": true
							},
							"status": {
							  "type": "string",
							  "required": true
							}
						  },
						  "block_types": {
							"destination": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "account_id": {
									"type": "string",
									"optional": true
								  },
								  "bucket": {
									"type": "string",
									"required": true
								  },
								  "replica_kms_key_id": {
									"type": "string",
									"optional": true
								  },
								  "storage_class": {
									"type": "string",
									"optional": true
								  }
								},
								"block_types": {
								  "access_control_translation": {
									"nesting_mode": "list",
									"block": {
									  "attributes": {
										"owner": {
										  "type": "string",
										  "required": true
										}
									  }
									},
									"max_items": 1
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							},
							"filter": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "prefix": {
									"type": "string",
									"optional": true
								  },
								  "tags": {
									"type": [
									  "map",
									  "string"
									],
									"optional": true
								  }
								}
							  },
							  "max_items": 1
							},
							"source_selection_criteria": {
							  "nesting_mode": "list",
							  "block": {
								"block_types": {
								  "sse_kms_encrypted_objects": {
									"nesting_mode": "list",
									"block": {
									  "attributes": {
										"enabled": {
										  "type": "bool",
										  "required": true
										}
									  }
									},
									"max_items": 1
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"min_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"server_side_encryption_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "rule": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"apply_server_side_encryption_by_default": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "kms_master_key_id": {
									"type": "string",
									"optional": true
								  },
								  "sse_algorithm": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"versioning": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "enabled": {
						"type": "bool",
						"optional": true
					  },
					  "mfa_delete": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"website": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "error_document": {
						"type": "string",
						"optional": true
					  },
					  "index_document": {
						"type": "string",
						"optional": true
					  },
					  "redirect_all_requests_to": {
						"type": "string",
						"optional": true
					  },
					  "routing_rules": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_s3_bucket_analytics_configuration": {
			"version": 0,
			"block": {
			  "attributes": {
				"bucket": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"filter": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "prefix": {
						"type": "string",
						"optional": true
					  },
					  "tags": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"storage_class_analysis": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "data_export": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"output_schema_version": {
							  "type": "string",
							  "optional": true
							}
						  },
						  "block_types": {
							"destination": {
							  "nesting_mode": "list",
							  "block": {
								"block_types": {
								  "s3_bucket_destination": {
									"nesting_mode": "list",
									"block": {
									  "attributes": {
										"bucket_account_id": {
										  "type": "string",
										  "optional": true
										},
										"bucket_arn": {
										  "type": "string",
										  "required": true
										},
										"format": {
										  "type": "string",
										  "optional": true
										},
										"prefix": {
										  "type": "string",
										  "optional": true
										}
									  }
									},
									"min_items": 1,
									"max_items": 1
								  }
								}
							  },
							  "min_items": 1,
							  "max_items": 1
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_s3_bucket_inventory": {
			"version": 0,
			"block": {
			  "attributes": {
				"bucket": {
				  "type": "string",
				  "required": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"included_object_versions": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"optional_fields": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"destination": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "bucket": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"account_id": {
							  "type": "string",
							  "optional": true
							},
							"bucket_arn": {
							  "type": "string",
							  "required": true
							},
							"format": {
							  "type": "string",
							  "required": true
							},
							"prefix": {
							  "type": "string",
							  "optional": true
							}
						  },
						  "block_types": {
							"encryption": {
							  "nesting_mode": "list",
							  "block": {
								"block_types": {
								  "sse_kms": {
									"nesting_mode": "list",
									"block": {
									  "attributes": {
										"key_id": {
										  "type": "string",
										  "required": true
										}
									  }
									},
									"max_items": 1
								  },
								  "sse_s3": {
									"nesting_mode": "list",
									"block": {},
									"max_items": 1
								  }
								}
							  },
							  "max_items": 1
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"filter": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "prefix": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"schedule": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "frequency": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_s3_bucket_metric": {
			"version": 0,
			"block": {
			  "attributes": {
				"bucket": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"filter": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "prefix": {
						"type": "string",
						"optional": true
					  },
					  "tags": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_s3_bucket_notification": {
			"version": 0,
			"block": {
			  "attributes": {
				"bucket": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"lambda_function": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "events": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "filter_prefix": {
						"type": "string",
						"optional": true
					  },
					  "filter_suffix": {
						"type": "string",
						"optional": true
					  },
					  "id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "lambda_function_arn": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"queue": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "events": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "filter_prefix": {
						"type": "string",
						"optional": true
					  },
					  "filter_suffix": {
						"type": "string",
						"optional": true
					  },
					  "id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "queue_arn": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"topic": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "events": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "filter_prefix": {
						"type": "string",
						"optional": true
					  },
					  "filter_suffix": {
						"type": "string",
						"optional": true
					  },
					  "id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "topic_arn": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_s3_bucket_object": {
			"version": 0,
			"block": {
			  "attributes": {
				"acl": {
				  "type": "string",
				  "optional": true
				},
				"bucket": {
				  "type": "string",
				  "required": true
				},
				"cache_control": {
				  "type": "string",
				  "optional": true
				},
				"content": {
				  "type": "string",
				  "optional": true
				},
				"content_base64": {
				  "type": "string",
				  "optional": true
				},
				"content_disposition": {
				  "type": "string",
				  "optional": true
				},
				"content_encoding": {
				  "type": "string",
				  "optional": true
				},
				"content_language": {
				  "type": "string",
				  "optional": true
				},
				"content_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"etag": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"force_destroy": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key": {
				  "type": "string",
				  "required": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true
				},
				"metadata": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"object_lock_legal_hold_status": {
				  "type": "string",
				  "optional": true
				},
				"object_lock_mode": {
				  "type": "string",
				  "optional": true
				},
				"object_lock_retain_until_date": {
				  "type": "string",
				  "optional": true
				},
				"server_side_encryption": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"source": {
				  "type": "string",
				  "optional": true
				},
				"storage_class": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"version_id": {
				  "type": "string",
				  "computed": true
				},
				"website_redirect": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_s3_bucket_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"bucket": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_s3_bucket_public_access_block": {
			"version": 0,
			"block": {
			  "attributes": {
				"block_public_acls": {
				  "type": "bool",
				  "optional": true
				},
				"block_public_policy": {
				  "type": "bool",
				  "optional": true
				},
				"bucket": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ignore_public_acls": {
				  "type": "bool",
				  "optional": true
				},
				"restrict_public_buckets": {
				  "type": "bool",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_sagemaker_endpoint": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"endpoint_config_name": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_sagemaker_endpoint_configuration": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_arn": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"production_variants": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "accelerator_type": {
						"type": "string",
						"optional": true
					  },
					  "initial_instance_count": {
						"type": "number",
						"required": true
					  },
					  "initial_variant_weight": {
						"type": "number",
						"optional": true
					  },
					  "instance_type": {
						"type": "string",
						"required": true
					  },
					  "model_name": {
						"type": "string",
						"required": true
					  },
					  "variant_name": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "min_items": 1
				}
			  }
			}
		  },
		  "aws_sagemaker_model": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"enable_network_isolation": {
				  "type": "bool",
				  "optional": true
				},
				"execution_role_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"container": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "container_hostname": {
						"type": "string",
						"optional": true
					  },
					  "environment": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  },
					  "image": {
						"type": "string",
						"required": true
					  },
					  "model_data_url": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"primary_container": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "container_hostname": {
						"type": "string",
						"optional": true
					  },
					  "environment": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  },
					  "image": {
						"type": "string",
						"required": true
					  },
					  "model_data_url": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"vpc_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "security_group_ids": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  },
					  "subnets": {
						"type": [
						  "set",
						  "string"
						],
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_sagemaker_notebook_instance": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"direct_internet_access": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_type": {
				  "type": "string",
				  "required": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true
				},
				"lifecycle_config_name": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"role_arn": {
				  "type": "string",
				  "required": true
				},
				"security_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"subnet_id": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_sagemaker_notebook_instance_lifecycle_configuration": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"on_create": {
				  "type": "string",
				  "optional": true
				},
				"on_start": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_secretsmanager_secret": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_key_id": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "optional": true
				},
				"recovery_window_in_days": {
				  "type": "number",
				  "optional": true
				},
				"rotation_enabled": {
				  "type": "bool",
				  "computed": true
				},
				"rotation_lambda_arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"rotation_rules": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "automatically_after_days": {
						"type": "number",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_secretsmanager_secret_rotation": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"rotation_enabled": {
				  "type": "bool",
				  "computed": true
				},
				"rotation_lambda_arn": {
				  "type": "string",
				  "required": true
				},
				"secret_id": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"rotation_rules": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "automatically_after_days": {
						"type": "number",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_secretsmanager_secret_version": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"secret_binary": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"secret_id": {
				  "type": "string",
				  "required": true
				},
				"secret_string": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"version_id": {
				  "type": "string",
				  "computed": true
				},
				"version_stages": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_security_group": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"egress": {
				  "type": [
					"set",
					[
					  "object",
					  {
						"cidr_blocks": [
						  "list",
						  "string"
						],
						"description": "string",
						"from_port": "number",
						"ipv6_cidr_blocks": [
						  "list",
						  "string"
						],
						"prefix_list_ids": [
						  "list",
						  "string"
						],
						"protocol": "string",
						"security_groups": [
						  "set",
						  "string"
						],
						"self": "bool",
						"to_port": "number"
					  }
					]
				  ],
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ingress": {
				  "type": [
					"set",
					[
					  "object",
					  {
						"cidr_blocks": [
						  "list",
						  "string"
						],
						"description": "string",
						"from_port": "number",
						"ipv6_cidr_blocks": [
						  "list",
						  "string"
						],
						"prefix_list_ids": [
						  "list",
						  "string"
						],
						"protocol": "string",
						"security_groups": [
						  "set",
						  "string"
						],
						"self": "bool",
						"to_port": "number"
					  }
					]
				  ],
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"revoke_rules_on_delete": {
				  "type": "bool",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_security_group_rule": {
			"version": 2,
			"block": {
			  "attributes": {
				"cidr_blocks": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"from_port": {
				  "type": "number",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ipv6_cidr_blocks": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"prefix_list_ids": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"protocol": {
				  "type": "string",
				  "required": true
				},
				"security_group_id": {
				  "type": "string",
				  "required": true
				},
				"self": {
				  "type": "bool",
				  "optional": true
				},
				"source_security_group_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"to_port": {
				  "type": "number",
				  "required": true
				},
				"type": {
				  "type": "string",
				  "description": "Type of rule, ingress (inbound) or egress (outbound).",
				  "required": true
				}
			  }
			}
		  },
		  "aws_securityhub_account": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_securityhub_member": {
			"version": 0,
			"block": {
			  "attributes": {
				"account_id": {
				  "type": "string",
				  "required": true
				},
				"email": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"invite": {
				  "type": "bool",
				  "optional": true
				},
				"master_id": {
				  "type": "string",
				  "computed": true
				},
				"member_status": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_securityhub_product_subscription": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"product_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_securityhub_standards_subscription": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"standards_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_service_discovery_http_namespace": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_service_discovery_private_dns_namespace": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"hosted_zone": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_service_discovery_public_dns_namespace": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"hosted_zone": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_service_discovery_service": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"namespace_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"dns_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "namespace_id": {
						"type": "string",
						"required": true
					  },
					  "routing_policy": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "dns_records": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"ttl": {
							  "type": "number",
							  "required": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"health_check_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "failure_threshold": {
						"type": "number",
						"optional": true
					  },
					  "resource_path": {
						"type": "string",
						"optional": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"health_check_custom_config": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "failure_threshold": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_servicecatalog_portfolio": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"created_time": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"provider_name": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_servicequotas_service_quota": {
			"version": 0,
			"block": {
			  "attributes": {
				"adjustable": {
				  "type": "bool",
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"default_value": {
				  "type": "number",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"quota_code": {
				  "type": "string",
				  "required": true
				},
				"quota_name": {
				  "type": "string",
				  "computed": true
				},
				"request_id": {
				  "type": "string",
				  "computed": true
				},
				"request_status": {
				  "type": "string",
				  "computed": true
				},
				"service_code": {
				  "type": "string",
				  "required": true
				},
				"service_name": {
				  "type": "string",
				  "computed": true
				},
				"value": {
				  "type": "number",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ses_active_receipt_rule_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"rule_set_name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ses_configuration_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ses_domain_dkim": {
			"version": 0,
			"block": {
			  "attributes": {
				"dkim_tokens": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"domain": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_ses_domain_identity": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"domain": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"verification_token": {
				  "type": "string",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_ses_domain_identity_verification": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"domain": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_ses_domain_mail_from": {
			"version": 0,
			"block": {
			  "attributes": {
				"behavior_on_mx_failure": {
				  "type": "string",
				  "optional": true
				},
				"domain": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"mail_from_domain": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ses_email_identity": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"email": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_ses_event_destination": {
			"version": 0,
			"block": {
			  "attributes": {
				"configuration_set_name": {
				  "type": "string",
				  "required": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"matching_types": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"cloudwatch_destination": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "default_value": {
						"type": "string",
						"required": true
					  },
					  "dimension_name": {
						"type": "string",
						"required": true
					  },
					  "value_source": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"kinesis_destination": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "role_arn": {
						"type": "string",
						"required": true
					  },
					  "stream_arn": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"sns_destination": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "topic_arn": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_ses_identity_notification_topic": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identity": {
				  "type": "string",
				  "required": true
				},
				"include_original_headers": {
				  "type": "bool",
				  "optional": true
				},
				"notification_type": {
				  "type": "string",
				  "required": true
				},
				"topic_arn": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_ses_identity_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identity": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ses_receipt_filter": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"cidr": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ses_receipt_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"after": {
				  "type": "string",
				  "optional": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"recipients": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"rule_set_name": {
				  "type": "string",
				  "required": true
				},
				"scan_enabled": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"tls_policy": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  },
			  "block_types": {
				"add_header_action": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "header_name": {
						"type": "string",
						"required": true
					  },
					  "header_value": {
						"type": "string",
						"required": true
					  },
					  "position": {
						"type": "number",
						"required": true
					  }
					}
				  }
				},
				"bounce_action": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "message": {
						"type": "string",
						"required": true
					  },
					  "position": {
						"type": "number",
						"required": true
					  },
					  "sender": {
						"type": "string",
						"required": true
					  },
					  "smtp_reply_code": {
						"type": "string",
						"required": true
					  },
					  "status_code": {
						"type": "string",
						"optional": true
					  },
					  "topic_arn": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"lambda_action": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "function_arn": {
						"type": "string",
						"required": true
					  },
					  "invocation_type": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "position": {
						"type": "number",
						"required": true
					  },
					  "topic_arn": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"s3_action": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "bucket_name": {
						"type": "string",
						"required": true
					  },
					  "kms_key_arn": {
						"type": "string",
						"optional": true
					  },
					  "object_key_prefix": {
						"type": "string",
						"optional": true
					  },
					  "position": {
						"type": "number",
						"required": true
					  },
					  "topic_arn": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"sns_action": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "position": {
						"type": "number",
						"required": true
					  },
					  "topic_arn": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"stop_action": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "position": {
						"type": "number",
						"required": true
					  },
					  "scope": {
						"type": "string",
						"required": true
					  },
					  "topic_arn": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"workmail_action": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "organization_arn": {
						"type": "string",
						"required": true
					  },
					  "position": {
						"type": "number",
						"required": true
					  },
					  "topic_arn": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_ses_receipt_rule_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"rule_set_name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ses_template": {
			"version": 0,
			"block": {
			  "attributes": {
				"html": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"subject": {
				  "type": "string",
				  "optional": true
				},
				"text": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_sfn_activity": {
			"version": 0,
			"block": {
			  "attributes": {
				"creation_date": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_sfn_state_machine": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"creation_date": {
				  "type": "string",
				  "computed": true
				},
				"definition": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"role_arn": {
				  "type": "string",
				  "required": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_shield_protection": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"resource_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_simpledb_domain": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_snapshot_create_volume_permission": {
			"version": 0,
			"block": {
			  "attributes": {
				"account_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"snapshot_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_sns_platform_application": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"event_delivery_failure_topic_arn": {
				  "type": "string",
				  "optional": true
				},
				"event_endpoint_created_topic_arn": {
				  "type": "string",
				  "optional": true
				},
				"event_endpoint_deleted_topic_arn": {
				  "type": "string",
				  "optional": true
				},
				"event_endpoint_updated_topic_arn": {
				  "type": "string",
				  "optional": true
				},
				"failure_feedback_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"platform": {
				  "type": "string",
				  "required": true
				},
				"platform_credential": {
				  "type": "string",
				  "required": true,
				  "sensitive": true
				},
				"platform_principal": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"success_feedback_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"success_feedback_sample_rate": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_sns_sms_preferences": {
			"version": 0,
			"block": {
			  "attributes": {
				"default_sender_id": {
				  "type": "string",
				  "optional": true
				},
				"default_sms_type": {
				  "type": "string",
				  "optional": true
				},
				"delivery_status_iam_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"delivery_status_success_sampling_rate": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"monthly_spend_limit": {
				  "type": "string",
				  "optional": true
				},
				"usage_report_s3_bucket": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_sns_topic": {
			"version": 0,
			"block": {
			  "attributes": {
				"application_failure_feedback_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"application_success_feedback_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"application_success_feedback_sample_rate": {
				  "type": "number",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"delivery_policy": {
				  "type": "string",
				  "optional": true
				},
				"display_name": {
				  "type": "string",
				  "optional": true
				},
				"http_failure_feedback_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"http_success_feedback_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"http_success_feedback_sample_rate": {
				  "type": "number",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_master_key_id": {
				  "type": "string",
				  "optional": true
				},
				"lambda_failure_feedback_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"lambda_success_feedback_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"lambda_success_feedback_sample_rate": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"policy": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"sqs_failure_feedback_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"sqs_success_feedback_role_arn": {
				  "type": "string",
				  "optional": true
				},
				"sqs_success_feedback_sample_rate": {
				  "type": "number",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_sns_topic_policy": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_sns_topic_subscription": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"confirmation_timeout_in_minutes": {
				  "type": "number",
				  "optional": true
				},
				"delivery_policy": {
				  "type": "string",
				  "optional": true
				},
				"endpoint": {
				  "type": "string",
				  "required": true
				},
				"endpoint_auto_confirms": {
				  "type": "bool",
				  "optional": true
				},
				"filter_policy": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"protocol": {
				  "type": "string",
				  "required": true
				},
				"raw_message_delivery": {
				  "type": "bool",
				  "optional": true
				},
				"topic_arn": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_spot_datafeed_subscription": {
			"version": 0,
			"block": {
			  "attributes": {
				"bucket": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"prefix": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_spot_fleet_request": {
			"version": 1,
			"block": {
			  "attributes": {
				"allocation_strategy": {
				  "type": "string",
				  "optional": true
				},
				"client_token": {
				  "type": "string",
				  "computed": true
				},
				"excess_capacity_termination_policy": {
				  "type": "string",
				  "optional": true
				},
				"fleet_type": {
				  "type": "string",
				  "optional": true
				},
				"iam_fleet_role": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_interruption_behaviour": {
				  "type": "string",
				  "optional": true
				},
				"instance_pools_to_use_count": {
				  "type": "number",
				  "optional": true
				},
				"load_balancers": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"replace_unhealthy_instances": {
				  "type": "bool",
				  "optional": true
				},
				"spot_price": {
				  "type": "string",
				  "optional": true
				},
				"spot_request_state": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"target_capacity": {
				  "type": "number",
				  "required": true
				},
				"target_group_arns": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"terminate_instances_with_expiration": {
				  "type": "bool",
				  "optional": true
				},
				"valid_from": {
				  "type": "string",
				  "optional": true
				},
				"valid_until": {
				  "type": "string",
				  "optional": true
				},
				"wait_for_fulfillment": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"launch_specification": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "ami": {
						"type": "string",
						"required": true
					  },
					  "associate_public_ip_address": {
						"type": "bool",
						"optional": true
					  },
					  "availability_zone": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "ebs_optimized": {
						"type": "bool",
						"optional": true
					  },
					  "iam_instance_profile": {
						"type": "string",
						"optional": true
					  },
					  "iam_instance_profile_arn": {
						"type": "string",
						"optional": true
					  },
					  "instance_type": {
						"type": "string",
						"required": true
					  },
					  "key_name": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "monitoring": {
						"type": "bool",
						"optional": true
					  },
					  "placement_group": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "placement_tenancy": {
						"type": "string",
						"optional": true
					  },
					  "spot_price": {
						"type": "string",
						"optional": true
					  },
					  "subnet_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "tags": {
						"type": [
						  "map",
						  "string"
						],
						"optional": true
					  },
					  "user_data": {
						"type": "string",
						"optional": true
					  },
					  "vpc_security_group_ids": {
						"type": [
						  "set",
						  "string"
						],
						"optional": true,
						"computed": true
					  },
					  "weighted_capacity": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "ebs_block_device": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"delete_on_termination": {
							  "type": "bool",
							  "optional": true
							},
							"device_name": {
							  "type": "string",
							  "required": true
							},
							"encrypted": {
							  "type": "bool",
							  "optional": true,
							  "computed": true
							},
							"iops": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"kms_key_id": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"snapshot_id": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"volume_size": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"volume_type": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							}
						  }
						}
					  },
					  "ephemeral_block_device": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"device_name": {
							  "type": "string",
							  "required": true
							},
							"virtual_name": {
							  "type": "string",
							  "required": true
							}
						  }
						}
					  },
					  "root_block_device": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"delete_on_termination": {
							  "type": "bool",
							  "optional": true
							},
							"encrypted": {
							  "type": "bool",
							  "optional": true,
							  "computed": true
							},
							"iops": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"kms_key_id": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"volume_size": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"volume_type": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							}
						  }
						}
					  }
					}
				  }
				},
				"launch_template_config": {
				  "nesting_mode": "set",
				  "block": {
					"block_types": {
					  "launch_template_specification": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"id": {
							  "type": "string",
							  "optional": true
							},
							"name": {
							  "type": "string",
							  "optional": true
							},
							"version": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  },
					  "overrides": {
						"nesting_mode": "set",
						"block": {
						  "attributes": {
							"availability_zone": {
							  "type": "string",
							  "optional": true
							},
							"instance_type": {
							  "type": "string",
							  "optional": true
							},
							"priority": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							},
							"spot_price": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"subnet_id": {
							  "type": "string",
							  "optional": true,
							  "computed": true
							},
							"weighted_capacity": {
							  "type": "number",
							  "optional": true,
							  "computed": true
							}
						  }
						}
					  }
					}
				  }
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_spot_instance_request": {
			"version": 0,
			"block": {
			  "attributes": {
				"ami": {
				  "type": "string",
				  "required": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"associate_public_ip_address": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"availability_zone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"block_duration_minutes": {
				  "type": "number",
				  "optional": true
				},
				"cpu_core_count": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"cpu_threads_per_core": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"disable_api_termination": {
				  "type": "bool",
				  "optional": true
				},
				"ebs_optimized": {
				  "type": "bool",
				  "optional": true
				},
				"get_password_data": {
				  "type": "bool",
				  "optional": true
				},
				"hibernation": {
				  "type": "bool",
				  "optional": true
				},
				"host_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"iam_instance_profile": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_initiated_shutdown_behavior": {
				  "type": "string",
				  "optional": true
				},
				"instance_interruption_behaviour": {
				  "type": "string",
				  "optional": true
				},
				"instance_state": {
				  "type": "string",
				  "computed": true
				},
				"instance_type": {
				  "type": "string",
				  "required": true
				},
				"ipv6_address_count": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"ipv6_addresses": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"key_name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"launch_group": {
				  "type": "string",
				  "optional": true
				},
				"monitoring": {
				  "type": "bool",
				  "optional": true
				},
				"outpost_arn": {
				  "type": "string",
				  "computed": true
				},
				"password_data": {
				  "type": "string",
				  "computed": true
				},
				"placement_group": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"primary_network_interface_id": {
				  "type": "string",
				  "computed": true
				},
				"private_dns": {
				  "type": "string",
				  "computed": true
				},
				"private_ip": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"public_dns": {
				  "type": "string",
				  "computed": true
				},
				"public_ip": {
				  "type": "string",
				  "computed": true
				},
				"secondary_private_ips": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"security_groups": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"source_dest_check": {
				  "type": "bool",
				  "optional": true
				},
				"spot_bid_status": {
				  "type": "string",
				  "computed": true
				},
				"spot_instance_id": {
				  "type": "string",
				  "computed": true
				},
				"spot_price": {
				  "type": "string",
				  "optional": true
				},
				"spot_request_state": {
				  "type": "string",
				  "computed": true
				},
				"spot_type": {
				  "type": "string",
				  "optional": true
				},
				"subnet_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"tenancy": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"user_data": {
				  "type": "string",
				  "optional": true
				},
				"user_data_base64": {
				  "type": "string",
				  "optional": true
				},
				"valid_from": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"valid_until": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"volume_tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"wait_for_fulfillment": {
				  "type": "bool",
				  "optional": true
				}
			  },
			  "block_types": {
				"credit_specification": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "cpu_credits": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"ebs_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "delete_on_termination": {
						"type": "bool",
						"optional": true
					  },
					  "device_name": {
						"type": "string",
						"required": true
					  },
					  "encrypted": {
						"type": "bool",
						"optional": true,
						"computed": true
					  },
					  "iops": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "kms_key_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "snapshot_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "volume_id": {
						"type": "string",
						"computed": true
					  },
					  "volume_size": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "volume_type": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  }
				},
				"ephemeral_block_device": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "device_name": {
						"type": "string",
						"required": true
					  },
					  "no_device": {
						"type": "bool",
						"optional": true
					  },
					  "virtual_name": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				},
				"metadata_options": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "http_endpoint": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "http_put_response_hop_limit": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "http_tokens": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				},
				"network_interface": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "delete_on_termination": {
						"type": "bool",
						"optional": true
					  },
					  "device_index": {
						"type": "number",
						"required": true
					  },
					  "network_interface_id": {
						"type": "string",
						"required": true
					  }
					}
				  }
				},
				"root_block_device": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "delete_on_termination": {
						"type": "bool",
						"optional": true
					  },
					  "device_name": {
						"type": "string",
						"computed": true
					  },
					  "encrypted": {
						"type": "bool",
						"optional": true,
						"computed": true
					  },
					  "iops": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "kms_key_id": {
						"type": "string",
						"optional": true,
						"computed": true
					  },
					  "volume_id": {
						"type": "string",
						"computed": true
					  },
					  "volume_size": {
						"type": "number",
						"optional": true,
						"computed": true
					  },
					  "volume_type": {
						"type": "string",
						"optional": true,
						"computed": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_sqs_queue": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"content_based_deduplication": {
				  "type": "bool",
				  "optional": true
				},
				"delay_seconds": {
				  "type": "number",
				  "optional": true
				},
				"fifo_queue": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_data_key_reuse_period_seconds": {
				  "type": "number",
				  "optional": true,
				  "computed": true
				},
				"kms_master_key_id": {
				  "type": "string",
				  "optional": true
				},
				"max_message_size": {
				  "type": "number",
				  "optional": true
				},
				"message_retention_seconds": {
				  "type": "number",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"policy": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"receive_wait_time_seconds": {
				  "type": "number",
				  "optional": true
				},
				"redrive_policy": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"visibility_timeout_seconds": {
				  "type": "number",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_sqs_queue_policy": {
			"version": 1,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "required": true
				},
				"queue_url": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ssm_activation": {
			"version": 0,
			"block": {
			  "attributes": {
				"activation_code": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"expiration_date": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"expired": {
				  "type": "bool",
				  "computed": true
				},
				"iam_role": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"registration_count": {
				  "type": "number",
				  "computed": true
				},
				"registration_limit": {
				  "type": "number",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_ssm_association": {
			"version": 1,
			"block": {
			  "attributes": {
				"association_id": {
				  "type": "string",
				  "computed": true
				},
				"association_name": {
				  "type": "string",
				  "optional": true
				},
				"automation_target_parameter_name": {
				  "type": "string",
				  "optional": true
				},
				"compliance_severity": {
				  "type": "string",
				  "optional": true
				},
				"document_version": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_id": {
				  "type": "string",
				  "optional": true
				},
				"max_concurrency": {
				  "type": "string",
				  "optional": true
				},
				"max_errors": {
				  "type": "string",
				  "optional": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"parameters": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"schedule_expression": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"output_location": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "s3_bucket_name": {
						"type": "string",
						"required": true
					  },
					  "s3_key_prefix": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"targets": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "key": {
						"type": "string",
						"required": true
					  },
					  "values": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  }
					}
				  },
				  "max_items": 5
				}
			  }
			}
		  },
		  "aws_ssm_document": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"content": {
				  "type": "string",
				  "required": true
				},
				"created_date": {
				  "type": "string",
				  "computed": true
				},
				"default_version": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "computed": true
				},
				"document_format": {
				  "type": "string",
				  "optional": true
				},
				"document_type": {
				  "type": "string",
				  "required": true
				},
				"document_version": {
				  "type": "string",
				  "computed": true
				},
				"hash": {
				  "type": "string",
				  "computed": true
				},
				"hash_type": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"latest_version": {
				  "type": "string",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"owner": {
				  "type": "string",
				  "computed": true
				},
				"parameter": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"default_value": "string",
						"description": "string",
						"name": "string",
						"type": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"permissions": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"platform_types": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"schema_version": {
				  "type": "string",
				  "computed": true
				},
				"status": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"target_type": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"attachments_source": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "key": {
						"type": "string",
						"required": true
					  },
					  "name": {
						"type": "string",
						"optional": true
					  },
					  "values": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_ssm_maintenance_window": {
			"version": 0,
			"block": {
			  "attributes": {
				"allow_unassociated_targets": {
				  "type": "bool",
				  "optional": true
				},
				"cutoff": {
				  "type": "number",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"duration": {
				  "type": "number",
				  "required": true
				},
				"enabled": {
				  "type": "bool",
				  "optional": true
				},
				"end_date": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"schedule": {
				  "type": "string",
				  "required": true
				},
				"schedule_timezone": {
				  "type": "string",
				  "optional": true
				},
				"start_date": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_ssm_maintenance_window_target": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"owner_information": {
				  "type": "string",
				  "optional": true
				},
				"resource_type": {
				  "type": "string",
				  "required": true
				},
				"window_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"targets": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "key": {
						"type": "string",
						"required": true
					  },
					  "values": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 5
				}
			  }
			}
		  },
		  "aws_ssm_maintenance_window_task": {
			"version": 0,
			"block": {
			  "attributes": {
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"max_concurrency": {
				  "type": "string",
				  "required": true
				},
				"max_errors": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "optional": true
				},
				"priority": {
				  "type": "number",
				  "optional": true
				},
				"service_role_arn": {
				  "type": "string",
				  "required": true
				},
				"task_arn": {
				  "type": "string",
				  "required": true
				},
				"task_type": {
				  "type": "string",
				  "required": true
				},
				"window_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"targets": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "key": {
						"type": "string",
						"required": true
					  },
					  "values": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  }
					}
				  },
				  "min_items": 1
				},
				"task_invocation_parameters": {
				  "nesting_mode": "list",
				  "block": {
					"block_types": {
					  "automation_parameters": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"document_version": {
							  "type": "string",
							  "optional": true
							}
						  },
						  "block_types": {
							"parameter": {
							  "nesting_mode": "set",
							  "block": {
								"attributes": {
								  "name": {
									"type": "string",
									"required": true
								  },
								  "values": {
									"type": [
									  "list",
									  "string"
									],
									"required": true
								  }
								}
							  }
							}
						  }
						},
						"max_items": 1
					  },
					  "lambda_parameters": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"client_context": {
							  "type": "string",
							  "optional": true
							},
							"payload": {
							  "type": "string",
							  "optional": true,
							  "sensitive": true
							},
							"qualifier": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  },
					  "run_command_parameters": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"comment": {
							  "type": "string",
							  "optional": true
							},
							"document_hash": {
							  "type": "string",
							  "optional": true
							},
							"document_hash_type": {
							  "type": "string",
							  "optional": true
							},
							"output_s3_bucket": {
							  "type": "string",
							  "optional": true
							},
							"output_s3_key_prefix": {
							  "type": "string",
							  "optional": true
							},
							"service_role_arn": {
							  "type": "string",
							  "optional": true
							},
							"timeout_seconds": {
							  "type": "number",
							  "optional": true
							}
						  },
						  "block_types": {
							"notification_config": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "notification_arn": {
									"type": "string",
									"optional": true
								  },
								  "notification_events": {
									"type": [
									  "list",
									  "string"
									],
									"optional": true
								  },
								  "notification_type": {
									"type": "string",
									"optional": true
								  }
								}
							  },
							  "max_items": 1
							},
							"parameter": {
							  "nesting_mode": "set",
							  "block": {
								"attributes": {
								  "name": {
									"type": "string",
									"required": true
								  },
								  "values": {
									"type": [
									  "list",
									  "string"
									],
									"required": true
								  }
								}
							  }
							}
						  }
						},
						"max_items": 1
					  },
					  "step_functions_parameters": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"input": {
							  "type": "string",
							  "optional": true,
							  "sensitive": true
							},
							"name": {
							  "type": "string",
							  "optional": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_ssm_parameter": {
			"version": 0,
			"block": {
			  "attributes": {
				"allowed_pattern": {
				  "type": "string",
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"data_type": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"key_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"overwrite": {
				  "type": "bool",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"tier": {
				  "type": "string",
				  "optional": true
				},
				"type": {
				  "type": "string",
				  "required": true
				},
				"value": {
				  "type": "string",
				  "required": true,
				  "sensitive": true
				},
				"version": {
				  "type": "number",
				  "computed": true
				}
			  }
			}
		  },
		  "aws_ssm_patch_baseline": {
			"version": 0,
			"block": {
			  "attributes": {
				"approved_patches": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"approved_patches_compliance_level": {
				  "type": "string",
				  "optional": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"operating_system": {
				  "type": "string",
				  "optional": true
				},
				"rejected_patches": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"approval_rule": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "approve_after_days": {
						"type": "number",
						"required": true
					  },
					  "compliance_level": {
						"type": "string",
						"optional": true
					  },
					  "enable_non_security": {
						"type": "bool",
						"optional": true
					  }
					},
					"block_types": {
					  "patch_filter": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"key": {
							  "type": "string",
							  "required": true
							},
							"values": {
							  "type": [
								"list",
								"string"
							  ],
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 10
					  }
					}
				  }
				},
				"global_filter": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "key": {
						"type": "string",
						"required": true
					  },
					  "values": {
						"type": [
						  "list",
						  "string"
						],
						"required": true
					  }
					}
				  },
				  "max_items": 4
				}
			  }
			}
		  },
		  "aws_ssm_patch_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"baseline_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"patch_group": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_ssm_resource_data_sync": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"s3_destination": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "bucket_name": {
						"type": "string",
						"required": true
					  },
					  "kms_key_arn": {
						"type": "string",
						"optional": true
					  },
					  "prefix": {
						"type": "string",
						"optional": true
					  },
					  "region": {
						"type": "string",
						"required": true
					  },
					  "sync_format": {
						"type": "string",
						"optional": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_storagegateway_cache": {
			"version": 0,
			"block": {
			  "attributes": {
				"disk_id": {
				  "type": "string",
				  "required": true
				},
				"gateway_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_storagegateway_cached_iscsi_volume": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"chap_enabled": {
				  "type": "bool",
				  "computed": true
				},
				"gateway_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"lun_number": {
				  "type": "number",
				  "computed": true
				},
				"network_interface_id": {
				  "type": "string",
				  "required": true
				},
				"network_interface_port": {
				  "type": "number",
				  "computed": true
				},
				"snapshot_id": {
				  "type": "string",
				  "optional": true
				},
				"source_volume_arn": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"target_arn": {
				  "type": "string",
				  "computed": true
				},
				"target_name": {
				  "type": "string",
				  "required": true
				},
				"volume_arn": {
				  "type": "string",
				  "computed": true
				},
				"volume_id": {
				  "type": "string",
				  "computed": true
				},
				"volume_size_in_bytes": {
				  "type": "number",
				  "required": true
				}
			  }
			}
		  },
		  "aws_storagegateway_gateway": {
			"version": 0,
			"block": {
			  "attributes": {
				"activation_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"cloudwatch_log_group_arn": {
				  "type": "string",
				  "optional": true
				},
				"gateway_id": {
				  "type": "string",
				  "computed": true
				},
				"gateway_ip_address": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"gateway_name": {
				  "type": "string",
				  "required": true
				},
				"gateway_timezone": {
				  "type": "string",
				  "required": true
				},
				"gateway_type": {
				  "type": "string",
				  "optional": true
				},
				"gateway_vpc_endpoint": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"medium_changer_type": {
				  "type": "string",
				  "optional": true
				},
				"smb_guest_password": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"tape_drive_type": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"smb_active_directory_settings": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "domain_name": {
						"type": "string",
						"required": true
					  },
					  "password": {
						"type": "string",
						"required": true,
						"sensitive": true
					  },
					  "username": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_storagegateway_nfs_file_share": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"client_list": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"default_storage_class": {
				  "type": "string",
				  "optional": true
				},
				"fileshare_id": {
				  "type": "string",
				  "computed": true
				},
				"gateway_arn": {
				  "type": "string",
				  "required": true
				},
				"guess_mime_type_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"kms_encrypted": {
				  "type": "bool",
				  "optional": true
				},
				"kms_key_arn": {
				  "type": "string",
				  "optional": true
				},
				"location_arn": {
				  "type": "string",
				  "required": true
				},
				"object_acl": {
				  "type": "string",
				  "optional": true
				},
				"path": {
				  "type": "string",
				  "computed": true
				},
				"read_only": {
				  "type": "bool",
				  "optional": true
				},
				"requester_pays": {
				  "type": "bool",
				  "optional": true
				},
				"role_arn": {
				  "type": "string",
				  "required": true
				},
				"squash": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"nfs_file_share_defaults": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "directory_mode": {
						"type": "string",
						"optional": true
					  },
					  "file_mode": {
						"type": "string",
						"optional": true
					  },
					  "group_id": {
						"type": "number",
						"optional": true
					  },
					  "owner_id": {
						"type": "number",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_storagegateway_smb_file_share": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"authentication": {
				  "type": "string",
				  "optional": true
				},
				"default_storage_class": {
				  "type": "string",
				  "optional": true
				},
				"fileshare_id": {
				  "type": "string",
				  "computed": true
				},
				"gateway_arn": {
				  "type": "string",
				  "required": true
				},
				"guess_mime_type_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"invalid_user_list": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"kms_encrypted": {
				  "type": "bool",
				  "optional": true
				},
				"kms_key_arn": {
				  "type": "string",
				  "optional": true
				},
				"location_arn": {
				  "type": "string",
				  "required": true
				},
				"object_acl": {
				  "type": "string",
				  "optional": true
				},
				"path": {
				  "type": "string",
				  "computed": true
				},
				"read_only": {
				  "type": "bool",
				  "optional": true
				},
				"requester_pays": {
				  "type": "bool",
				  "optional": true
				},
				"role_arn": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"valid_user_list": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_storagegateway_upload_buffer": {
			"version": 0,
			"block": {
			  "attributes": {
				"disk_id": {
				  "type": "string",
				  "required": true
				},
				"gateway_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_storagegateway_working_storage": {
			"version": 0,
			"block": {
			  "attributes": {
				"disk_id": {
				  "type": "string",
				  "required": true
				},
				"gateway_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_subnet": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"assign_ipv6_address_on_creation": {
				  "type": "bool",
				  "optional": true
				},
				"availability_zone": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"availability_zone_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"cidr_block": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ipv6_cidr_block": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ipv6_cidr_block_association_id": {
				  "type": "string",
				  "computed": true
				},
				"map_public_ip_on_launch": {
				  "type": "bool",
				  "optional": true
				},
				"outpost_arn": {
				  "type": "string",
				  "optional": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_swf_domain": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name_prefix": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"workflow_execution_retention_period_in_days": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_transfer_server": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"endpoint": {
				  "type": "string",
				  "computed": true
				},
				"endpoint_type": {
				  "type": "string",
				  "optional": true
				},
				"force_destroy": {
				  "type": "bool",
				  "optional": true
				},
				"host_key": {
				  "type": "string",
				  "optional": true,
				  "sensitive": true
				},
				"host_key_fingerprint": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"identity_provider_type": {
				  "type": "string",
				  "optional": true
				},
				"invocation_role": {
				  "type": "string",
				  "optional": true
				},
				"logging_role": {
				  "type": "string",
				  "optional": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"url": {
				  "type": "string",
				  "optional": true
				}
			  },
			  "block_types": {
				"endpoint_details": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "vpc_endpoint_id": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_transfer_ssh_key": {
			"version": 0,
			"block": {
			  "attributes": {
				"body": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"server_id": {
				  "type": "string",
				  "required": true
				},
				"user_name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_transfer_user": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"home_directory": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "optional": true
				},
				"role": {
				  "type": "string",
				  "required": true
				},
				"server_id": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"user_name": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_volume_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"device_name": {
				  "type": "string",
				  "required": true
				},
				"force_detach": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_id": {
				  "type": "string",
				  "required": true
				},
				"skip_destroy": {
				  "type": "bool",
				  "optional": true
				},
				"volume_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_vpc": {
			"version": 1,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"assign_generated_ipv6_cidr_block": {
				  "type": "bool",
				  "optional": true
				},
				"cidr_block": {
				  "type": "string",
				  "required": true
				},
				"default_network_acl_id": {
				  "type": "string",
				  "computed": true
				},
				"default_route_table_id": {
				  "type": "string",
				  "computed": true
				},
				"default_security_group_id": {
				  "type": "string",
				  "computed": true
				},
				"dhcp_options_id": {
				  "type": "string",
				  "computed": true
				},
				"enable_classiclink": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"enable_classiclink_dns_support": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"enable_dns_hostnames": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"enable_dns_support": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"instance_tenancy": {
				  "type": "string",
				  "optional": true
				},
				"ipv6_association_id": {
				  "type": "string",
				  "computed": true
				},
				"ipv6_cidr_block": {
				  "type": "string",
				  "computed": true
				},
				"main_route_table_id": {
				  "type": "string",
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_vpc_dhcp_options": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"domain_name": {
				  "type": "string",
				  "optional": true
				},
				"domain_name_servers": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"netbios_name_servers": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"netbios_node_type": {
				  "type": "string",
				  "optional": true
				},
				"ntp_servers": {
				  "type": [
					"list",
					"string"
				  ],
				  "optional": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_vpc_dhcp_options_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"dhcp_options_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_vpc_endpoint": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"auto_accept": {
				  "type": "bool",
				  "optional": true
				},
				"cidr_blocks": {
				  "type": [
					"list",
					"string"
				  ],
				  "computed": true
				},
				"dns_entry": {
				  "type": [
					"list",
					[
					  "object",
					  {
						"dns_name": "string",
						"hosted_zone_id": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"network_interface_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "computed": true
				},
				"owner_id": {
				  "type": "string",
				  "computed": true
				},
				"policy": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"prefix_list_id": {
				  "type": "string",
				  "computed": true
				},
				"private_dns_enabled": {
				  "type": "bool",
				  "optional": true
				},
				"requester_managed": {
				  "type": "bool",
				  "computed": true
				},
				"route_table_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"security_group_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"service_name": {
				  "type": "string",
				  "required": true
				},
				"state": {
				  "type": "string",
				  "computed": true
				},
				"subnet_ids": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_endpoint_type": {
				  "type": "string",
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_vpc_endpoint_connection_notification": {
			"version": 0,
			"block": {
			  "attributes": {
				"connection_events": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"connection_notification_arn": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"notification_type": {
				  "type": "string",
				  "computed": true
				},
				"state": {
				  "type": "string",
				  "computed": true
				},
				"vpc_endpoint_id": {
				  "type": "string",
				  "optional": true
				},
				"vpc_endpoint_service_id": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_vpc_endpoint_route_table_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"route_table_id": {
				  "type": "string",
				  "required": true
				},
				"vpc_endpoint_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_vpc_endpoint_service": {
			"version": 0,
			"block": {
			  "attributes": {
				"acceptance_required": {
				  "type": "bool",
				  "required": true
				},
				"allowed_principals": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"availability_zones": {
				  "type": [
					"set",
					"string"
				  ],
				  "computed": true
				},
				"base_endpoint_dns_names": {
				  "type": [
					"set",
					"string"
				  ],
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"manages_vpc_endpoints": {
				  "type": "bool",
				  "computed": true
				},
				"network_load_balancer_arns": {
				  "type": [
					"set",
					"string"
				  ],
				  "required": true
				},
				"private_dns_name": {
				  "type": "string",
				  "computed": true
				},
				"service_name": {
				  "type": "string",
				  "computed": true
				},
				"service_type": {
				  "type": "string",
				  "computed": true
				},
				"state": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_vpc_endpoint_service_allowed_principal": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"principal_arn": {
				  "type": "string",
				  "required": true
				},
				"vpc_endpoint_service_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_vpc_endpoint_subnet_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"subnet_id": {
				  "type": "string",
				  "required": true
				},
				"vpc_endpoint_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_vpc_ipv4_cidr_block_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"cidr_block": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_vpc_peering_connection": {
			"version": 0,
			"block": {
			  "attributes": {
				"accept_status": {
				  "type": "string",
				  "computed": true
				},
				"auto_accept": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"peer_owner_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"peer_region": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"peer_vpc_id": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"accepter": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allow_classic_link_to_remote_vpc": {
						"type": "bool",
						"optional": true
					  },
					  "allow_remote_vpc_dns_resolution": {
						"type": "bool",
						"optional": true
					  },
					  "allow_vpc_to_remote_classic_link": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"requester": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allow_classic_link_to_remote_vpc": {
						"type": "bool",
						"optional": true
					  },
					  "allow_remote_vpc_dns_resolution": {
						"type": "bool",
						"optional": true
					  },
					  "allow_vpc_to_remote_classic_link": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"timeouts": {
				  "nesting_mode": "single",
				  "block": {
					"attributes": {
					  "create": {
						"type": "string",
						"optional": true
					  },
					  "delete": {
						"type": "string",
						"optional": true
					  },
					  "update": {
						"type": "string",
						"optional": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_vpc_peering_connection_accepter": {
			"version": 0,
			"block": {
			  "attributes": {
				"accept_status": {
				  "type": "string",
				  "computed": true
				},
				"auto_accept": {
				  "type": "bool",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"peer_owner_id": {
				  "type": "string",
				  "computed": true
				},
				"peer_region": {
				  "type": "string",
				  "computed": true
				},
				"peer_vpc_id": {
				  "type": "string",
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "computed": true
				},
				"vpc_peering_connection_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"accepter": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allow_classic_link_to_remote_vpc": {
						"type": "bool",
						"optional": true
					  },
					  "allow_remote_vpc_dns_resolution": {
						"type": "bool",
						"optional": true
					  },
					  "allow_vpc_to_remote_classic_link": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"requester": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allow_classic_link_to_remote_vpc": {
						"type": "bool",
						"optional": true
					  },
					  "allow_remote_vpc_dns_resolution": {
						"type": "bool",
						"optional": true
					  },
					  "allow_vpc_to_remote_classic_link": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_vpc_peering_connection_options": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"vpc_peering_connection_id": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"accepter": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allow_classic_link_to_remote_vpc": {
						"type": "bool",
						"optional": true
					  },
					  "allow_remote_vpc_dns_resolution": {
						"type": "bool",
						"optional": true
					  },
					  "allow_vpc_to_remote_classic_link": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				},
				"requester": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "allow_classic_link_to_remote_vpc": {
						"type": "bool",
						"optional": true
					  },
					  "allow_remote_vpc_dns_resolution": {
						"type": "bool",
						"optional": true
					  },
					  "allow_vpc_to_remote_classic_link": {
						"type": "bool",
						"optional": true
					  }
					}
				  },
				  "max_items": 1
				}
			  }
			}
		  },
		  "aws_vpn_connection": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"customer_gateway_configuration": {
				  "type": "string",
				  "computed": true
				},
				"customer_gateway_id": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"routes": {
				  "type": [
					"set",
					[
					  "object",
					  {
						"destination_cidr_block": "string",
						"source": "string",
						"state": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"static_routes_only": {
				  "type": "bool",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"transit_gateway_attachment_id": {
				  "type": "string",
				  "computed": true
				},
				"transit_gateway_id": {
				  "type": "string",
				  "optional": true
				},
				"tunnel1_address": {
				  "type": "string",
				  "computed": true
				},
				"tunnel1_bgp_asn": {
				  "type": "string",
				  "computed": true
				},
				"tunnel1_bgp_holdtime": {
				  "type": "number",
				  "computed": true
				},
				"tunnel1_cgw_inside_address": {
				  "type": "string",
				  "computed": true
				},
				"tunnel1_inside_cidr": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tunnel1_preshared_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true,
				  "sensitive": true
				},
				"tunnel1_vgw_inside_address": {
				  "type": "string",
				  "computed": true
				},
				"tunnel2_address": {
				  "type": "string",
				  "computed": true
				},
				"tunnel2_bgp_asn": {
				  "type": "string",
				  "computed": true
				},
				"tunnel2_bgp_holdtime": {
				  "type": "number",
				  "computed": true
				},
				"tunnel2_cgw_inside_address": {
				  "type": "string",
				  "computed": true
				},
				"tunnel2_inside_cidr": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tunnel2_preshared_key": {
				  "type": "string",
				  "optional": true,
				  "computed": true,
				  "sensitive": true
				},
				"tunnel2_vgw_inside_address": {
				  "type": "string",
				  "computed": true
				},
				"type": {
				  "type": "string",
				  "required": true
				},
				"vgw_telemetry": {
				  "type": [
					"set",
					[
					  "object",
					  {
						"accepted_route_count": "number",
						"last_status_change": "string",
						"outside_ip_address": "string",
						"status": "string",
						"status_message": "string"
					  }
					]
				  ],
				  "computed": true
				},
				"vpn_gateway_id": {
				  "type": "string",
				  "optional": true
				}
			  }
			}
		  },
		  "aws_vpn_connection_route": {
			"version": 0,
			"block": {
			  "attributes": {
				"destination_cidr_block": {
				  "type": "string",
				  "required": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"vpn_connection_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_vpn_gateway": {
			"version": 0,
			"block": {
			  "attributes": {
				"amazon_side_asn": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"availability_zone": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				},
				"vpc_id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				}
			  }
			}
		  },
		  "aws_vpn_gateway_attachment": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"vpc_id": {
				  "type": "string",
				  "required": true
				},
				"vpn_gateway_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_vpn_gateway_route_propagation": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"route_table_id": {
				  "type": "string",
				  "required": true
				},
				"vpn_gateway_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_waf_byte_match_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"byte_match_tuples": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "positional_constraint": {
						"type": "string",
						"required": true
					  },
					  "target_string": {
						"type": "string",
						"optional": true
					  },
					  "text_transformation": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "field_to_match": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"data": {
							  "type": "string",
							  "optional": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_waf_geo_match_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"geo_match_constraint": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "type": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_waf_ipset": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"ip_set_descriptors": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "type": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_waf_rate_based_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"metric_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"rate_key": {
				  "type": "string",
				  "required": true
				},
				"rate_limit": {
				  "type": "number",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"predicates": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "data_id": {
						"type": "string",
						"required": true
					  },
					  "negated": {
						"type": "bool",
						"required": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_waf_regex_match_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"regex_match_tuple": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "regex_pattern_set_id": {
						"type": "string",
						"required": true
					  },
					  "text_transformation": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "field_to_match": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"data": {
							  "type": "string",
							  "optional": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_waf_regex_pattern_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"regex_pattern_strings": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_waf_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"metric_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"predicates": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "data_id": {
						"type": "string",
						"required": true
					  },
					  "negated": {
						"type": "bool",
						"required": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_waf_rule_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"metric_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"activated_rule": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "priority": {
						"type": "number",
						"required": true
					  },
					  "rule_id": {
						"type": "string",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "action": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_waf_size_constraint_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"size_constraints": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "comparison_operator": {
						"type": "string",
						"required": true
					  },
					  "size": {
						"type": "number",
						"required": true
					  },
					  "text_transformation": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "field_to_match": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"data": {
							  "type": "string",
							  "optional": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_waf_sql_injection_match_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"sql_injection_match_tuples": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "text_transformation": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "field_to_match": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"data": {
							  "type": "string",
							  "optional": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_waf_web_acl": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"metric_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"default_action": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"logging_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "log_destination": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "redacted_fields": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"field_to_match": {
							  "nesting_mode": "set",
							  "block": {
								"attributes": {
								  "data": {
									"type": "string",
									"optional": true
								  },
								  "type": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "min_items": 1
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"rules": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "priority": {
						"type": "number",
						"required": true
					  },
					  "rule_id": {
						"type": "string",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "action": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "override_action": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_waf_xss_match_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"xss_match_tuples": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "text_transformation": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "field_to_match": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"data": {
							  "type": "string",
							  "optional": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_wafregional_byte_match_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"byte_match_tuples": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "positional_constraint": {
						"type": "string",
						"required": true
					  },
					  "target_string": {
						"type": "string",
						"optional": true
					  },
					  "text_transformation": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "field_to_match": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"data": {
							  "type": "string",
							  "optional": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_wafregional_geo_match_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"geo_match_constraint": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "type": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_wafregional_ipset": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"ip_set_descriptor": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "type": {
						"type": "string",
						"required": true
					  },
					  "value": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_wafregional_rate_based_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"metric_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"rate_key": {
				  "type": "string",
				  "required": true
				},
				"rate_limit": {
				  "type": "number",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"predicate": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "data_id": {
						"type": "string",
						"required": true
					  },
					  "negated": {
						"type": "bool",
						"required": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_wafregional_regex_match_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"regex_match_tuple": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "regex_pattern_set_id": {
						"type": "string",
						"required": true
					  },
					  "text_transformation": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "field_to_match": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"data": {
							  "type": "string",
							  "optional": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_wafregional_regex_pattern_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"regex_pattern_strings": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_wafregional_rule": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"metric_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"predicate": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "data_id": {
						"type": "string",
						"required": true
					  },
					  "negated": {
						"type": "bool",
						"required": true
					  },
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_wafregional_rule_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"metric_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"activated_rule": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "priority": {
						"type": "number",
						"required": true
					  },
					  "rule_id": {
						"type": "string",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "action": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_wafregional_size_constraint_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"size_constraints": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "comparison_operator": {
						"type": "string",
						"required": true
					  },
					  "size": {
						"type": "number",
						"required": true
					  },
					  "text_transformation": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "field_to_match": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"data": {
							  "type": "string",
							  "optional": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_wafregional_sql_injection_match_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"sql_injection_match_tuple": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "text_transformation": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "field_to_match": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"data": {
							  "type": "string",
							  "optional": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_wafregional_web_acl": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"metric_name": {
				  "type": "string",
				  "required": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"default_action": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "type": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "min_items": 1,
				  "max_items": 1
				},
				"logging_configuration": {
				  "nesting_mode": "list",
				  "block": {
					"attributes": {
					  "log_destination": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "redacted_fields": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"field_to_match": {
							  "nesting_mode": "set",
							  "block": {
								"attributes": {
								  "data": {
									"type": "string",
									"optional": true
								  },
								  "type": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "min_items": 1
							}
						  }
						},
						"max_items": 1
					  }
					}
				  },
				  "max_items": 1
				},
				"rule": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "priority": {
						"type": "number",
						"required": true
					  },
					  "rule_id": {
						"type": "string",
						"required": true
					  },
					  "type": {
						"type": "string",
						"optional": true
					  }
					},
					"block_types": {
					  "action": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  },
					  "override_action": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_wafregional_web_acl_association": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"resource_arn": {
				  "type": "string",
				  "required": true
				},
				"web_acl_id": {
				  "type": "string",
				  "required": true
				}
			  }
			}
		  },
		  "aws_wafregional_xss_match_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				}
			  },
			  "block_types": {
				"xss_match_tuple": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "text_transformation": {
						"type": "string",
						"required": true
					  }
					},
					"block_types": {
					  "field_to_match": {
						"nesting_mode": "list",
						"block": {
						  "attributes": {
							"data": {
							  "type": "string",
							  "optional": true
							},
							"type": {
							  "type": "string",
							  "required": true
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  }
					}
				  }
				}
			  }
			}
		  },
		  "aws_wafv2_ip_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"addresses": {
				  "type": [
					"set",
					"string"
				  ],
				  "optional": true
				},
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"ip_address_version": {
				  "type": "string",
				  "required": true
				},
				"lock_token": {
				  "type": "string",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"scope": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  }
			}
		  },
		  "aws_wafv2_regex_pattern_set": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"lock_token": {
				  "type": "string",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"scope": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"regular_expression": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "regex_string": {
						"type": "string",
						"required": true
					  }
					}
				  },
				  "max_items": 10
				}
			  }
			}
		  },
		  "aws_wafv2_rule_group": {
			"version": 0,
			"block": {
			  "attributes": {
				"arn": {
				  "type": "string",
				  "computed": true
				},
				"capacity": {
				  "type": "number",
				  "required": true
				},
				"description": {
				  "type": "string",
				  "optional": true
				},
				"id": {
				  "type": "string",
				  "optional": true,
				  "computed": true
				},
				"lock_token": {
				  "type": "string",
				  "computed": true
				},
				"name": {
				  "type": "string",
				  "required": true
				},
				"scope": {
				  "type": "string",
				  "required": true
				},
				"tags": {
				  "type": [
					"map",
					"string"
				  ],
				  "optional": true
				}
			  },
			  "block_types": {
				"rule": {
				  "nesting_mode": "set",
				  "block": {
					"attributes": {
					  "name": {
						"type": "string",
						"required": true
					  },
					  "priority": {
						"type": "number",
						"required": true
					  }
					},
					"block_types": {
					  "action": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"allow": {
							  "nesting_mode": "list",
							  "block": {},
							  "max_items": 1
							},
							"block": {
							  "nesting_mode": "list",
							  "block": {},
							  "max_items": 1
							},
							"count": {
							  "nesting_mode": "list",
							  "block": {},
							  "max_items": 1
							}
						  }
						},
						"min_items": 1,
						"max_items": 1
					  },
					  "statement": {
						"nesting_mode": "list",
						"block": {
						  "block_types": {
							"and_statement": {
							  "nesting_mode": "list",
							  "block": {
								"block_types": {
								  "statement": {
									"nesting_mode": "list",
									"block": {
									  "block_types": {
										"and_statement": {
										  "nesting_mode": "list",
										  "block": {
											"block_types": {
											  "statement": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"byte_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "positional_constraint": {
															"type": "string",
															"required": true
														  },
														  "search_string": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"geo_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "country_codes": {
															"type": [
															  "list",
															  "string"
															],
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"ip_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"regex_pattern_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"size_constraint_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "comparison_operator": {
															"type": "string",
															"required": true
														  },
														  "size": {
															"type": "number",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"sqli_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"xss_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"byte_match_statement": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "positional_constraint": {
												"type": "string",
												"required": true
											  },
											  "search_string": {
												"type": "string",
												"required": true
											  }
											},
											"block_types": {
											  "field_to_match": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"all_query_arguments": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"body": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"method": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"query_string": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"single_header": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"single_query_argument": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"uri_path": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													}
												  }
												},
												"max_items": 1
											  },
											  "text_transformation": {
												"nesting_mode": "set",
												"block": {
												  "attributes": {
													"priority": {
													  "type": "number",
													  "required": true
													},
													"type": {
													  "type": "string",
													  "required": true
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"geo_match_statement": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "country_codes": {
												"type": [
												  "list",
												  "string"
												],
												"required": true
											  }
											}
										  },
										  "max_items": 1
										},
										"ip_set_reference_statement": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "arn": {
												"type": "string",
												"required": true
											  }
											}
										  },
										  "max_items": 1
										},
										"not_statement": {
										  "nesting_mode": "list",
										  "block": {
											"block_types": {
											  "statement": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"byte_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "positional_constraint": {
															"type": "string",
															"required": true
														  },
														  "search_string": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"geo_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "country_codes": {
															"type": [
															  "list",
															  "string"
															],
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"ip_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"regex_pattern_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"size_constraint_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "comparison_operator": {
															"type": "string",
															"required": true
														  },
														  "size": {
															"type": "number",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"sqli_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"xss_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"or_statement": {
										  "nesting_mode": "list",
										  "block": {
											"block_types": {
											  "statement": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"byte_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "positional_constraint": {
															"type": "string",
															"required": true
														  },
														  "search_string": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"geo_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "country_codes": {
															"type": [
															  "list",
															  "string"
															],
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"ip_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"regex_pattern_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"size_constraint_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "comparison_operator": {
															"type": "string",
															"required": true
														  },
														  "size": {
															"type": "number",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"sqli_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"xss_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"regex_pattern_set_reference_statement": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "arn": {
												"type": "string",
												"required": true
											  }
											},
											"block_types": {
											  "field_to_match": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"all_query_arguments": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"body": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"method": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"query_string": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"single_header": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"single_query_argument": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"uri_path": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													}
												  }
												},
												"max_items": 1
											  },
											  "text_transformation": {
												"nesting_mode": "set",
												"block": {
												  "attributes": {
													"priority": {
													  "type": "number",
													  "required": true
													},
													"type": {
													  "type": "string",
													  "required": true
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"size_constraint_statement": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "comparison_operator": {
												"type": "string",
												"required": true
											  },
											  "size": {
												"type": "number",
												"required": true
											  }
											},
											"block_types": {
											  "field_to_match": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"all_query_arguments": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"body": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"method": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"query_string": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"single_header": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"single_query_argument": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"uri_path": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													}
												  }
												},
												"max_items": 1
											  },
											  "text_transformation": {
												"nesting_mode": "set",
												"block": {
												  "attributes": {
													"priority": {
													  "type": "number",
													  "required": true
													},
													"type": {
													  "type": "string",
													  "required": true
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"sqli_match_statement": {
										  "nesting_mode": "list",
										  "block": {
											"block_types": {
											  "field_to_match": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"all_query_arguments": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"body": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"method": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"query_string": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"single_header": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"single_query_argument": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"uri_path": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													}
												  }
												},
												"max_items": 1
											  },
											  "text_transformation": {
												"nesting_mode": "set",
												"block": {
												  "attributes": {
													"priority": {
													  "type": "number",
													  "required": true
													},
													"type": {
													  "type": "string",
													  "required": true
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"xss_match_statement": {
										  "nesting_mode": "list",
										  "block": {
											"block_types": {
											  "field_to_match": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"all_query_arguments": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"body": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"method": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"query_string": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"single_header": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"single_query_argument": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"uri_path": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													}
												  }
												},
												"max_items": 1
											  },
											  "text_transformation": {
												"nesting_mode": "set",
												"block": {
												  "attributes": {
													"priority": {
													  "type": "number",
													  "required": true
													},
													"type": {
													  "type": "string",
													  "required": true
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										}
									  }
									},
									"min_items": 1
								  }
								}
							  },
							  "max_items": 1
							},
							"byte_match_statement": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "positional_constraint": {
									"type": "string",
									"required": true
								  },
								  "search_string": {
									"type": "string",
									"required": true
								  }
								},
								"block_types": {
								  "field_to_match": {
									"nesting_mode": "list",
									"block": {
									  "block_types": {
										"all_query_arguments": {
										  "nesting_mode": "list",
										  "block": {},
										  "max_items": 1
										},
										"body": {
										  "nesting_mode": "list",
										  "block": {},
										  "max_items": 1
										},
										"method": {
										  "nesting_mode": "list",
										  "block": {},
										  "max_items": 1
										},
										"query_string": {
										  "nesting_mode": "list",
										  "block": {},
										  "max_items": 1
										},
										"single_header": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "name": {
												"type": "string",
												"required": true
											  }
											}
										  },
										  "max_items": 1
										},
										"single_query_argument": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "name": {
												"type": "string",
												"required": true
											  }
											}
										  },
										  "max_items": 1
										},
										"uri_path": {
										  "nesting_mode": "list",
										  "block": {},
										  "max_items": 1
										}
									  }
									},
									"max_items": 1
								  },
								  "text_transformation": {
									"nesting_mode": "set",
									"block": {
									  "attributes": {
										"priority": {
										  "type": "number",
										  "required": true
										},
										"type": {
										  "type": "string",
										  "required": true
										}
									  }
									},
									"min_items": 1
								  }
								}
							  },
							  "max_items": 1
							},
							"geo_match_statement": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "country_codes": {
									"type": [
									  "list",
									  "string"
									],
									"required": true
								  }
								}
							  },
							  "max_items": 1
							},
							"ip_set_reference_statement": {
							  "nesting_mode": "list",
							  "block": {
								"attributes": {
								  "arn": {
									"type": "string",
									"required": true
								  }
								}
							  },
							  "max_items": 1
							},
							"not_statement": {
							  "nesting_mode": "list",
							  "block": {
								"block_types": {
								  "statement": {
									"nesting_mode": "list",
									"block": {
									  "block_types": {
										"and_statement": {
										  "nesting_mode": "list",
										  "block": {
											"block_types": {
											  "statement": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"byte_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "positional_constraint": {
															"type": "string",
															"required": true
														  },
														  "search_string": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"geo_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "country_codes": {
															"type": [
															  "list",
															  "string"
															],
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"ip_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"regex_pattern_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"size_constraint_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "comparison_operator": {
															"type": "string",
															"required": true
														  },
														  "size": {
															"type": "number",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"sqli_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"xss_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"byte_match_statement": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "positional_constraint": {
												"type": "string",
												"required": true
											  },
											  "search_string": {
												"type": "string",
												"required": true
											  }
											},
											"block_types": {
											  "field_to_match": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"all_query_arguments": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"body": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"method": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"query_string": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"single_header": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"single_query_argument": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"uri_path": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													}
												  }
												},
												"max_items": 1
											  },
											  "text_transformation": {
												"nesting_mode": "set",
												"block": {
												  "attributes": {
													"priority": {
													  "type": "number",
													  "required": true
													},
													"type": {
													  "type": "string",
													  "required": true
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"geo_match_statement": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "country_codes": {
												"type": [
												  "list",
												  "string"
												],
												"required": true
											  }
											}
										  },
										  "max_items": 1
										},
										"ip_set_reference_statement": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "arn": {
												"type": "string",
												"required": true
											  }
											}
										  },
										  "max_items": 1
										},
										"not_statement": {
										  "nesting_mode": "list",
										  "block": {
											"block_types": {
											  "statement": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"byte_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "positional_constraint": {
															"type": "string",
															"required": true
														  },
														  "search_string": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"geo_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "country_codes": {
															"type": [
															  "list",
															  "string"
															],
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"ip_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"regex_pattern_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"size_constraint_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "comparison_operator": {
															"type": "string",
															"required": true
														  },
														  "size": {
															"type": "number",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"sqli_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"xss_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"or_statement": {
										  "nesting_mode": "list",
										  "block": {
											"block_types": {
											  "statement": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"byte_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "positional_constraint": {
															"type": "string",
															"required": true
														  },
														  "search_string": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"geo_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "country_codes": {
															"type": [
															  "list",
															  "string"
															],
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"ip_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"regex_pattern_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"size_constraint_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "comparison_operator": {
															"type": "string",
															"required": true
														  },
														  "size": {
															"type": "number",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"sqli_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"xss_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"regex_pattern_set_reference_statement": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "arn": {
												"type": "string",
												"required": true
											  }
											},
											"block_types": {
											  "field_to_match": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"all_query_arguments": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"body": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"method": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"query_string": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"single_header": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"single_query_argument": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"uri_path": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													}
												  }
												},
												"max_items": 1
											  },
											  "text_transformation": {
												"nesting_mode": "set",
												"block": {
												  "attributes": {
													"priority": {
													  "type": "number",
													  "required": true
													},
													"type": {
													  "type": "string",
													  "required": true
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"size_constraint_statement": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "comparison_operator": {
												"type": "string",
												"required": true
											  },
											  "size": {
												"type": "number",
												"required": true
											  }
											},
											"block_types": {
											  "field_to_match": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"all_query_arguments": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"body": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"method": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"query_string": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"single_header": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"single_query_argument": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"uri_path": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													}
												  }
												},
												"max_items": 1
											  },
											  "text_transformation": {
												"nesting_mode": "set",
												"block": {
												  "attributes": {
													"priority": {
													  "type": "number",
													  "required": true
													},
													"type": {
													  "type": "string",
													  "required": true
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"sqli_match_statement": {
										  "nesting_mode": "list",
										  "block": {
											"block_types": {
											  "field_to_match": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"all_query_arguments": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"body": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"method": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"query_string": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"single_header": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"single_query_argument": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"uri_path": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													}
												  }
												},
												"max_items": 1
											  },
											  "text_transformation": {
												"nesting_mode": "set",
												"block": {
												  "attributes": {
													"priority": {
													  "type": "number",
													  "required": true
													},
													"type": {
													  "type": "string",
													  "required": true
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"xss_match_statement": {
										  "nesting_mode": "list",
										  "block": {
											"block_types": {
											  "field_to_match": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"all_query_arguments": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"body": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"method": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"query_string": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"single_header": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"single_query_argument": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"uri_path": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													}
												  }
												},
												"max_items": 1
											  },
											  "text_transformation": {
												"nesting_mode": "set",
												"block": {
												  "attributes": {
													"priority": {
													  "type": "number",
													  "required": true
													},
													"type": {
													  "type": "string",
													  "required": true
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										}
									  }
									},
									"min_items": 1
								  }
								}
							  },
							  "max_items": 1
							},
							"or_statement": {
							  "nesting_mode": "list",
							  "block": {
								"block_types": {
								  "statement": {
									"nesting_mode": "list",
									"block": {
									  "block_types": {
										"and_statement": {
										  "nesting_mode": "list",
										  "block": {
											"block_types": {
											  "statement": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"byte_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "positional_constraint": {
															"type": "string",
															"required": true
														  },
														  "search_string": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"geo_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "country_codes": {
															"type": [
															  "list",
															  "string"
															],
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"ip_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"regex_pattern_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"size_constraint_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "comparison_operator": {
															"type": "string",
															"required": true
														  },
														  "size": {
															"type": "number",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"sqli_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"xss_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"byte_match_statement": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "positional_constraint": {
												"type": "string",
												"required": true
											  },
											  "search_string": {
												"type": "string",
												"required": true
											  }
											},
											"block_types": {
											  "field_to_match": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"all_query_arguments": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"body": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"method": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"query_string": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													},
													"single_header": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"single_query_argument": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "name": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"uri_path": {
													  "nesting_mode": "list",
													  "block": {},
													  "max_items": 1
													}
												  }
												},
												"max_items": 1
											  },
											  "text_transformation": {
												"nesting_mode": "set",
												"block": {
												  "attributes": {
													"priority": {
													  "type": "number",
													  "required": true
													},
													"type": {
													  "type": "string",
													  "required": true
													}
												  }
												},
												"min_items": 1
											  }
											}
										  },
										  "max_items": 1
										},
										"geo_match_statement": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "country_codes": {
												"type": [
												  "list",
												  "string"
												],
												"required": true
											  }
											}
										  },
										  "max_items": 1
										},
										"ip_set_reference_statement": {
										  "nesting_mode": "list",
										  "block": {
											"attributes": {
											  "arn": {
												"type": "string",
												"required": true
											  }
											}
										  },
										  "max_items": 1
										},
										"not_statement": {
										  "nesting_mode": "list",
										  "block": {
											"block_types": {
											  "statement": {
												"nesting_mode": "list",
												"block": {
												  "block_types": {
													"byte_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "positional_constraint": {
															"type": "string",
															"required": true
														  },
														  "search_string": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"geo_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "country_codes": {
															"type": [
															  "list",
															  "string"
															],
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"ip_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														}
													  },
													  "max_items": 1
													},
													"regex_pattern_set_reference_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "arn": {
															"type": "string",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"size_constraint_statement": {
													  "nesting_mode": "list",
													  "block": {
														"attributes": {
														  "comparison_operator": {
															"type": "string",
															"required": true
														  },
														  "size": {
															"type": "number",
															"required": true
														  }
														},
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
															},
															"min_items": 1
														  }
														}
													  },
													  "max_items": 1
													},
													"sqli_match_statement": {
													  "nesting_mode": "list",
													  "block": {
														"block_types": {
														  "field_to_match": {
															"nesting_mode": "list",
															"block": {
															  "block_types": {
																"all_query_arguments": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"body": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"method": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"query_string": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																},
																"single_header": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"single_query_argument": {
																  "nesting_mode": "list",
																  "block": {
																	"attributes": {
																	  "name": {
																		"type": "string",
																		"required": true
																	  }
																	}
																  },
																  "max_items": 1
																},
																"uri_path": {
																  "nesting_mode": "list",
																  "block": {},
																  "max_items": 1
																}
															  }
															},
															"max_items": 1
														  },
														  "text_transformation": {
															"nesting_mode": "set",
															"block": {
															  "attributes": {
																"priority": {
																  "type": "number",
																  "required": true
																},
																"type": {
																  "type": "string",
																  "required": true
																}
															  }
								}