package cloudformation

// AWSDataPipelinePipeline_PipelineTag AWS CloudFormation Resource (AWS::DataPipeline::Pipeline.PipelineTag)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelinetags.html
type AWSDataPipelinePipeline_PipelineTag struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelinetags.html#cfn-datapipeline-pipeline-pipelinetags-key
	Key string `json:"Key,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelinetags.html#cfn-datapipeline-pipeline-pipelinetags-value
	Value string `json:"Value,omitempty"`

	DeletionPolicy *string                 `json:"-"`
	DependsOn      *[]string               `json:"-"`
	Metadata       *map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDataPipelinePipeline_PipelineTag) AWSCloudFormationType() string {
	return "AWS::DataPipeline::Pipeline.PipelineTag"
}
