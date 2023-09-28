package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type SyncResourceInfo struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	SubjectId        int64  `bson:"subject_id,omitempty" json:"subjectId,omitempty"`
	SubjectType      string `bson:"subject_type,omitempty" json:"subjectType,omitempty"`
	ResourceUniqcode string `bson:"resource_uniqcode,omitempty" json:"resourceUniqcode,omitempty"`
	ServiceType      string `bson:"service_type,omitempty" json:"serviceType,omitempty"`
	WorkspaceUuid    string `bson:"workspace_uuid,omitempty" json:"workspaceUuid,omitempty"`
	CreatedDate      int64  `bson:"created_date,omitempty" json:"createdDate,omitempty"`
}
