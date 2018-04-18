package cloudformation

// AWSSESReceiptFilter_IpFilter AWS CloudFormation Resource (AWS::SES::ReceiptFilter.IpFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptfilter-ipfilter.html
type AWSSESReceiptFilter_IpFilter struct {

	// Cidr AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptfilter-ipfilter.html#cfn-ses-receiptfilter-ipfilter-cidr
	Cidr string `json:"Cidr,omitempty"`

	// Policy AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptfilter-ipfilter.html#cfn-ses-receiptfilter-ipfilter-policy
	Policy string `json:"Policy,omitempty"`

	DeletionPolicy *string                 `json:"-"`
	DependsOn      *[]string               `json:"-"`
	Metadata       *map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESReceiptFilter_IpFilter) AWSCloudFormationType() string {
	return "AWS::SES::ReceiptFilter.IpFilter"
}
