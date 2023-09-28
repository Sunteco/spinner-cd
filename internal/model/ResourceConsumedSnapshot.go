package model

type ResourceConsumedSnapshot struct {
	baseEntity
	// TODO: Fill your own fields
	WorkspaceName string `bson:"workspace_name,omitempty" json:"workspaceName,omitempty"`
	ClusterName   string `bson:"cluster_name,omitempty" json:"ClusterName,omitempty"`
	ResourceType  string `bson:"resource_type,omitempty" json:"resourceType,omitempty"`
	ResourceUsage uint64 `bson:"resource_usage,omitempty" json:"resourceUsage,omitempty"`
	TimeStart     uint64 `bson:"time_start,omitempty" json:"timeStart,omitempty"`
	TimeEnd       uint64 `bson:"time_end,omitempty" json:"timeEnd,omitempty"`
	TimeStep      string `bson:"time_step,omitempty" json:"timeStep,omitempty"`
	SubjectId     int64  `bson:"subject_id,omitempty" json:"subjectId,omitempty"`
	SubjectType   string `bson:"subject_type,omitempty" json:"subjectType,omitempty"`
}
