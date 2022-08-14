USER aurora;

DROP TABLE if EXISTS `alert_metrics`;
CREATE TABLE `alert_metrics`
(
    `id`          varchar(40) NOT NULL,
    `name`        varchar(200) NOT NULL COMMENT '指标名称',
    `type`        varchar(50) NOT NULL COMMENT '指标类型',
    `unit`        varchar(22) NOT NULL COMMENT '单位',
    `operator`    varchar(22) NOT NULL COMMENT '操作符',
    `description` varchar(22) COMMENT '备注',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='告警指标表' row_format=dynamic;


DROP TABLE if EXISTS `alert_rules`;
CREATE TABLE `alert_rules`
(
    `id`             varchar(40)  NOT NULL,
    `name`           varchar(200)  NOT NULL COMMENT '规则名称',
    `alert_object`   text  NOT NULL COMMENT '告警对象,json{key:value}',
    `rules`          text  NOT NULL COMMENT '告警规则,json{metric:metricId,operator:operator,value:value}',
    `rules_status`   int         DEFAULT 1 COMMENT '0:禁用,1:启用',
    `webhook`        varchar(200) NOT NULL COMMENT '回调接口',
    `persistent`     varchar(64) DEFAULT NULL COMMENT '持续时间（默认为s）,单位:ms(毫秒),s（秒）,m(分),h(时),d(天),如果为null表示不告警',
    `alert_interval` varchar(64) DEFAULT NULL COMMENT '告警间隔时间（默认为s）,单位:ms(毫秒),s（秒）,m(分),h(时),d(天),如果为null表示不告警',
    `store_interval` varchar(64) DEFAULT NULL COMMENT '存储间隔时间（默认为s）,单位:ms(毫秒),s（秒）,m(分),h(时),d(天),如果为null表示不存储',
    `description`    varchar(200) COMMENT '备注',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='告警规则表' row_format=dynamic;