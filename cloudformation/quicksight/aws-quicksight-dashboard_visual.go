// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Dashboard_Visual AWS CloudFormation Resource (AWS::QuickSight::Dashboard.Visual)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html
type Dashboard_Visual struct {

	// BarChartVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-barchartvisual
	BarChartVisual *Dashboard_BarChartVisual `json:"BarChartVisual,omitempty"`

	// BoxPlotVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-boxplotvisual
	BoxPlotVisual *Dashboard_BoxPlotVisual `json:"BoxPlotVisual,omitempty"`

	// ComboChartVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-combochartvisual
	ComboChartVisual *Dashboard_ComboChartVisual `json:"ComboChartVisual,omitempty"`

	// CustomContentVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-customcontentvisual
	CustomContentVisual *Dashboard_CustomContentVisual `json:"CustomContentVisual,omitempty"`

	// EmptyVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-emptyvisual
	EmptyVisual *Dashboard_EmptyVisual `json:"EmptyVisual,omitempty"`

	// FilledMapVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-filledmapvisual
	FilledMapVisual *Dashboard_FilledMapVisual `json:"FilledMapVisual,omitempty"`

	// FunnelChartVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-funnelchartvisual
	FunnelChartVisual *Dashboard_FunnelChartVisual `json:"FunnelChartVisual,omitempty"`

	// GaugeChartVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-gaugechartvisual
	GaugeChartVisual *Dashboard_GaugeChartVisual `json:"GaugeChartVisual,omitempty"`

	// GeospatialMapVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-geospatialmapvisual
	GeospatialMapVisual *Dashboard_GeospatialMapVisual `json:"GeospatialMapVisual,omitempty"`

	// HeatMapVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-heatmapvisual
	HeatMapVisual *Dashboard_HeatMapVisual `json:"HeatMapVisual,omitempty"`

	// HistogramVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-histogramvisual
	HistogramVisual *Dashboard_HistogramVisual `json:"HistogramVisual,omitempty"`

	// InsightVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-insightvisual
	InsightVisual *Dashboard_InsightVisual `json:"InsightVisual,omitempty"`

	// KPIVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-kpivisual
	KPIVisual *Dashboard_KPIVisual `json:"KPIVisual,omitempty"`

	// LineChartVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-linechartvisual
	LineChartVisual *Dashboard_LineChartVisual `json:"LineChartVisual,omitempty"`

	// PieChartVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-piechartvisual
	PieChartVisual *Dashboard_PieChartVisual `json:"PieChartVisual,omitempty"`

	// PivotTableVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-pivottablevisual
	PivotTableVisual *Dashboard_PivotTableVisual `json:"PivotTableVisual,omitempty"`

	// RadarChartVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-radarchartvisual
	RadarChartVisual *Dashboard_RadarChartVisual `json:"RadarChartVisual,omitempty"`

	// SankeyDiagramVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-sankeydiagramvisual
	SankeyDiagramVisual *Dashboard_SankeyDiagramVisual `json:"SankeyDiagramVisual,omitempty"`

	// ScatterPlotVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-scatterplotvisual
	ScatterPlotVisual *Dashboard_ScatterPlotVisual `json:"ScatterPlotVisual,omitempty"`

	// TableVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-tablevisual
	TableVisual *Dashboard_TableVisual `json:"TableVisual,omitempty"`

	// TreeMapVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-treemapvisual
	TreeMapVisual *Dashboard_TreeMapVisual `json:"TreeMapVisual,omitempty"`

	// WaterfallVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-waterfallvisual
	WaterfallVisual *Dashboard_WaterfallVisual `json:"WaterfallVisual,omitempty"`

	// WordCloudVisual AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visual.html#cfn-quicksight-dashboard-visual-wordcloudvisual
	WordCloudVisual *Dashboard_WordCloudVisual `json:"WordCloudVisual,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Dashboard_Visual) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.Visual"
}
