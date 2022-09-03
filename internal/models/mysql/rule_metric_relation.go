package mysql

import (
	"github.com/DWHengr/aurora/internal/models"
	"gorm.io/gorm"
)

type ruleMetricRelationRepo struct{}

func NewRuleMetricRelationRepo() models.RuleMetricRelationRepo {
	return &ruleMetricRelationRepo{}
}

func (r *ruleMetricRelationRepo) TableName() string {
	return RuleMetricRelation
}

func (r *ruleMetricRelationRepo) GetRuleMetricByRuleId(db *gorm.DB, ruleId string) ([]*models.RuleMetricRelation, error) {

	entity := make([]*models.RuleMetricRelation, 0)
	err := db.Table(r.TableName()).
		Select("rule_metric_relation.*,alert_metrics.expression").
		Joins("left join alert_metrics on alert_metrics.id = rule_metric_relation.metric_id").
		Where("rule_id=?", ruleId).
		Scan(&entity).Error

	if err != nil {
		return nil, err
	}

	return entity, nil
}
