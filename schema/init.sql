USER aurora;

DROP TABLE if EXISTS `alert_metrics`;
CREATE TABLE `alert_metrics`
(
    `id`          varchar(40)  NOT NULL,
    `name`        varchar(200) NOT NULL COMMENT '指标名称',
    `type`        varchar(50)  NOT NULL COMMENT '指标类型',
    `expression`  varchar(500) NOT NULL COMMENT '表达式',
    `unit`        varchar(100) NOT NULL COMMENT '单位',
    `operator`    varchar(22)  NOT NULL COMMENT '操作符',
    `description` varchar(500) COMMENT '备注',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='告警指标表' row_format=dynamic;


DROP TABLE if EXISTS `alert_rules`;
CREATE TABLE `alert_rules`
(
    `id`                varchar(40)  NOT NULL,
    `name`              varchar(200) NOT NULL COMMENT '规则名称',
    `alert_object`      text         NOT NULL COMMENT '告警对象,json{key:value}',
    `rules_status`      int          DEFAULT 1 COMMENT '0:禁用,1:启用',
    `severity`          varchar(64)  DEFAULT NULL COMMENT '告警等级,hint,minor,importance,urgency',
    `webhook`           varchar(200) DEFAULT NULL COMMENT '回调接口',
    `alert_silences_id` varchar(40)  DEFAULT NULL COMMENT '告警静默id',
    `persistent`        varchar(64)  DEFAULT NULL COMMENT '持续时间（默认为s）,单位:ms(毫秒),s（秒）,m(分),h(时),d(天),如果为null表示不告警',
    `alert_interval`    varchar(64)  DEFAULT NULL COMMENT '告警间隔时间（默认为s）,单位:ms(毫秒),s（秒）,m(分),h(时),d(天),如果为null表示不告警',
    `store_interval`    varchar(64)  DEFAULT NULL COMMENT '存储间隔时间（默认为s）,单位:ms(毫秒),s（秒）,m(分),h(时),d(天),如果为null表示不存储',
    `description`       varchar(500) COMMENT '备注',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='告警规则表' row_format=dynamic;

DROP TABLE if EXISTS `rule_metric_relation`;
CREATE TABLE `rule_metric_relation`
(
    `id`          int         NOT NULL AUTO_INCREMENT,
    `rule_id`     varchar(40) NOT NULL COMMENT '规则id',
    `metric_id`   varchar(40) NOT NULL COMMENT '指标id',
    `statistics`  varchar(40) NOT NULL COMMENT '统计时间',
    `operator`    varchar(40) NOT NULL COMMENT '操作符',
    `alert_value` varchar(64) NOT NULL COMMENT '告警值',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='告警规则指标关系表' row_format=dynamic;

DROP TABLE if EXISTS `alert_silences`;
CREATE TABLE `alert_silences`
(
    `id`          varchar(40)  NOT NULL,
    `name`        varchar(200) NOT NULL COMMENT '静默名称',
    `type`        varchar(40)  NOT NULL COMMENT '静默类型,everyday,block,offday',
    `start_time`  datetime DEFAULT NULL COMMENT '静默开始时间',
    `end_time`    datetime DEFAULT NULL COMMENT '静默结束时间',
    `description` varchar(500) COMMENT '备注',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='告警静默表' row_format=dynamic;

DROP TABLE if EXISTS `alert_records`;
CREATE TABLE `alert_records`
(
    `id`          varchar(40)  NOT NULL,
    `alert_name`  varchar(200) NOT NULL COMMENT '告警名称',
    `rule_name`   varchar(200) NOT NULL COMMENT '规则名称',
    `rule_id`     varchar(40)  NOT NULL COMMENT '规则id',
    `severity`    varchar(64) DEFAULT NULL COMMENT '告警等级,hint,minor,importance,urgency',
    `summary`     text        DEFAULT NULL COMMENT '概述',
    `value`       varchar(64) DEFAULT NULL COMMENT '值',
    `attribute`   text        DEFAULT NULL COMMENT '属性',
    `create_time` datetime    DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='告警记录表' row_format=dynamic;

DROP TABLE if EXISTS `alert_users`;
CREATE TABLE `alert_users`
(
    `id`          varchar(40)  NOT NULL,
    `name`        varchar(200) NOT NULL COMMENT '姓名',
    `department`  varchar(200) NOT NULL COMMENT '部门',
    `email`       varchar(200) NOT NULL COMMENT '邮箱',
    `phone`       varchar(200) DEFAULT NULL COMMENT '手机',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='告警用户表' row_format=dynamic;